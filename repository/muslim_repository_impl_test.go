package repository

import (
	"database/sql"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestMuslimRepositoryFindDuas(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	repository := NewMuslimRepository(db)
	query := regexp.QuoteMeta(`
		SELECT id, title, arabic, latin, translation, category
		FROM duas
		WHERE ($1 = '' OR LOWER(category) = LOWER($1))
		ORDER BY id ASC
	`)

	rows := sqlmock.NewRows([]string{"id", "title", "arabic", "latin", "translation", "category"}).
		AddRow(1, "Doa Sebelum Makan", "arabic-1", "latin-1", "translation-1", "daily").
		AddRow(2, "Doa Sesudah Makan", "arabic-2", "latin-2", "translation-2", "daily")

	mock.ExpectQuery(query).
		WithArgs("daily").
		WillReturnRows(rows)

	duas := repository.FindDuas(t.Context(), "daily")

	if len(duas) != 2 {
		t.Fatalf("expected 2 duas, got %d", len(duas))
	}

	if duas[0].Title != "Doa Sebelum Makan" || duas[1].Title != "Doa Sesudah Makan" {
		t.Fatalf("unexpected duas: %#v", duas)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMuslimRepositoryFindRandomDuaReturnsNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	repository := NewMuslimRepository(db)
	query := regexp.QuoteMeta(`
		SELECT id, title, arabic, latin, translation, category
		FROM duas
		ORDER BY RANDOM()
		LIMIT 1
	`)

	mock.ExpectQuery(query).
		WillReturnError(sql.ErrNoRows)

	_, err = repository.FindRandomDua(t.Context())
	if err == nil {
		t.Fatal("expected not found error")
	}

	if err.Error() != "dua data is not available" {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMuslimRepositoryFindPrayerTime(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	repository := NewMuslimRepository(db)
	query := regexp.QuoteMeta(`
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
	`)

	rows := sqlmock.NewRows([]string{"city_name", "prayer_date", "fajr", "dhuhr", "asr", "maghrib", "isha"}).
		AddRow("Jakarta", "2026-03-23", "04:32", "12:01", "15:15", "18:05", "19:15")

	mock.ExpectQuery(query).
		WithArgs("jakarta", "2026-03-23").
		WillReturnRows(rows)

	prayerTime, err := repository.FindPrayerTime(t.Context(), "  jakarta  ", "2026-03-23")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if prayerTime.City != "Jakarta" || prayerTime.Maghrib != "18:05" {
		t.Fatalf("unexpected prayer time: %#v", prayerTime)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}
