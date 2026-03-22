package service

import (
	"context"

	"api-go-test/model/web"
)

type MuslimService interface {
	FindDuas(ctx context.Context, category string) []web.DuaResponse
	FindRandomDua(ctx context.Context) web.DuaResponse
	FindPrayerTime(ctx context.Context, request web.PrayerTimeRequest) web.PrayerTimeResponse
}
