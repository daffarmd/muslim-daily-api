package service

import (
	"context"
	"errors"
	"testing"

	"api-go-test/exception"
	"api-go-test/model/domain"
	"api-go-test/model/web"
)

type stubMuslimRepository struct {
	findDuasFn       func(ctx context.Context, category string) []domain.Dua
	findRandomDuaFn  func(ctx context.Context) (domain.Dua, error)
	findPrayerTimeFn func(ctx context.Context, city string, date string) (domain.PrayerTime, error)
}

func (stub *stubMuslimRepository) FindDuas(ctx context.Context, category string) []domain.Dua {
	return stub.findDuasFn(ctx, category)
}

func (stub *stubMuslimRepository) FindRandomDua(ctx context.Context) (domain.Dua, error) {
	return stub.findRandomDuaFn(ctx)
}

func (stub *stubMuslimRepository) FindPrayerTime(ctx context.Context, city string, date string) (domain.PrayerTime, error) {
	return stub.findPrayerTimeFn(ctx, city, date)
}

func TestMuslimServiceFindDuas(t *testing.T) {
	service := NewMuslimService(&stubMuslimRepository{
		findDuasFn: func(ctx context.Context, category string) []domain.Dua {
			if category != "daily" {
				t.Fatalf("expected category %q, got %q", "daily", category)
			}

			return []domain.Dua{
				{
					ID:          1,
					Title:       "Doa Sebelum Makan",
					Arabic:      "اللهم بارك لنا فيما رزقتنا",
					Latin:       "Allahumma barik lana fima razaqtana",
					Translation: "Ya Allah, berkahilah rezeki yang Engkau berikan kepada kami.",
					Category:    "daily",
				},
			}
		},
		findRandomDuaFn: func(ctx context.Context) (domain.Dua, error) {
			return domain.Dua{}, nil
		},
		findPrayerTimeFn: func(ctx context.Context, city string, date string) (domain.PrayerTime, error) {
			return domain.PrayerTime{}, nil
		},
	})

	response := service.FindDuas(t.Context(), "daily")

	if len(response) != 1 {
		t.Fatalf("expected 1 dua, got %d", len(response))
	}

	if response[0].Title != "Doa Sebelum Makan" {
		t.Fatalf("unexpected dua title: %#v", response[0])
	}
}

func TestMuslimServiceFindRandomDuaPanicsOnMissingData(t *testing.T) {
	service := NewMuslimService(&stubMuslimRepository{
		findDuasFn: func(ctx context.Context, category string) []domain.Dua {
			return nil
		},
		findRandomDuaFn: func(ctx context.Context) (domain.Dua, error) {
			return domain.Dua{}, errors.New("dua data is not available")
		},
		findPrayerTimeFn: func(ctx context.Context, city string, date string) (domain.PrayerTime, error) {
			return domain.PrayerTime{}, nil
		},
	})

	defer func() {
		rec := recover()
		if rec == nil {
			t.Fatal("expected panic on missing dua data")
		}

		if _, ok := rec.(exception.NotFoundError); !ok {
			t.Fatalf("expected NotFoundError panic, got %#v", rec)
		}
	}()

	service.FindRandomDua(t.Context())
}

func TestMuslimServiceFindPrayerTime(t *testing.T) {
	service := NewMuslimService(&stubMuslimRepository{
		findDuasFn: func(ctx context.Context, category string) []domain.Dua {
			return nil
		},
		findRandomDuaFn: func(ctx context.Context) (domain.Dua, error) {
			return domain.Dua{}, nil
		},
		findPrayerTimeFn: func(ctx context.Context, city string, date string) (domain.PrayerTime, error) {
			if city != "jakarta" || date != "2026-03-23" {
				t.Fatalf("unexpected lookup arguments city=%q date=%q", city, date)
			}

			return domain.PrayerTime{
				City:    "Jakarta",
				Date:    "2026-03-23",
				Fajr:    "04:32",
				Dhuhr:   "12:01",
				Asr:     "15:15",
				Maghrib: "18:05",
				Isha:    "19:15",
			}, nil
		},
	})

	response := service.FindPrayerTime(t.Context(), web.PrayerTimeRequest{
		City: "jakarta",
		Date: "2026-03-23",
	})

	if response.City != "Jakarta" || response.Times.Maghrib != "18:05" {
		t.Fatalf("unexpected prayer time response: %#v", response)
	}
}

func TestMuslimServiceFindPrayerTimePanicsOnInvalidDate(t *testing.T) {
	service := NewMuslimService(&stubMuslimRepository{
		findDuasFn: func(ctx context.Context, category string) []domain.Dua {
			return nil
		},
		findRandomDuaFn: func(ctx context.Context) (domain.Dua, error) {
			return domain.Dua{}, nil
		},
		findPrayerTimeFn: func(ctx context.Context, city string, date string) (domain.PrayerTime, error) {
			t.Fatal("repository should not be called on invalid request")
			return domain.PrayerTime{}, nil
		},
	})

	defer func() {
		rec := recover()
		if rec == nil {
			t.Fatal("expected panic on invalid prayer time request")
		}

		badRequestError, ok := rec.(exception.BadRequestError)
		if !ok {
			t.Fatalf("expected BadRequestError panic, got %#v", rec)
		}

		if badRequestError.Errors["date"] != "must use format YYYY-MM-DD" {
			t.Fatalf("unexpected validation errors: %#v", badRequestError.Errors)
		}
	}()

	service.FindPrayerTime(t.Context(), web.PrayerTimeRequest{
		City: "jakarta",
		Date: "23-03-2026",
	})
}
