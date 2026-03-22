package repository

import (
	"context"

	"api-go-test/model/domain"
)

type MuslimRepository interface {
	FindDuas(ctx context.Context, category string) []domain.Dua
	FindRandomDua(ctx context.Context) (domain.Dua, error)
	FindPrayerTime(ctx context.Context, city string, date string) (domain.PrayerTime, error)
}
