package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"api-go-test/helper"
	"api-go-test/model/domain"
)

type MuslimRepositoryImpl struct {
	DB *sql.DB
}

func NewMuslimRepository(db *sql.DB) MuslimRepository {
	return &MuslimRepositoryImpl{
		DB: db,
	}
}

func (repository *MuslimRepositoryImpl) FindDuas(ctx context.Context, category string) []domain.Dua {
	query := `
		SELECT id, title, arabic, latin, translation, category
		FROM duas
		WHERE ($1 = '' OR LOWER(category) = LOWER($1))
		ORDER BY id ASC
	`

	rows, err := repository.DB.QueryContext(ctx, query, strings.TrimSpace(category))
	helper.PanicIfErr(err)
	defer rows.Close()

	duas := make([]domain.Dua, 0)
	for rows.Next() {
		dua := domain.Dua{}
		err := rows.Scan(
			&dua.ID,
			&dua.Title,
			&dua.Arabic,
			&dua.Latin,
			&dua.Translation,
			&dua.Category,
		)
		helper.PanicIfErr(err)
		duas = append(duas, dua)
	}

	helper.PanicIfErr(rows.Err())

	return duas
}

func (repository *MuslimRepositoryImpl) FindRandomDua(ctx context.Context) (domain.Dua, error) {
	query := `
		SELECT id, title, arabic, latin, translation, category
		FROM duas
		ORDER BY RANDOM()
		LIMIT 1
	`

	dua := domain.Dua{}
	err := repository.DB.QueryRowContext(ctx, query).Scan(
		&dua.ID,
		&dua.Title,
		&dua.Arabic,
		&dua.Latin,
		&dua.Translation,
		&dua.Category,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Dua{}, errors.New("dua data is not available")
	}
	if err != nil {
		return domain.Dua{}, err
	}

	return dua, nil
}

func (repository *MuslimRepositoryImpl) FindPrayerTime(ctx context.Context, city string, date string) (domain.PrayerTime, error) {
	query := `
		SELECT
			city_name,
			TO_CHAR(prayer_date, 'YYYY-MM-DD') AS prayer_date,
			fajr,
			dhuhr,
			asr,
			maghrib,
			isha
		FROM prayer_times
		WHERE city_slug = $1 AND prayer_date = $2::date
	`

	prayerTime := domain.PrayerTime{}
	err := repository.DB.QueryRowContext(ctx, query, normalizeLookupKey(city), strings.TrimSpace(date)).Scan(
		&prayerTime.City,
		&prayerTime.Date,
		&prayerTime.Fajr,
		&prayerTime.Dhuhr,
		&prayerTime.Asr,
		&prayerTime.Maghrib,
		&prayerTime.Isha,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.PrayerTime{}, errors.New("prayer times not found")
	}
	if err != nil {
		return domain.PrayerTime{}, err
	}

	return prayerTime, nil
}

func normalizeLookupKey(value string) string {
	normalized := strings.ToLower(strings.TrimSpace(value))
	return strings.Join(strings.Fields(normalized), "-")
}
