package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"api-go-test/model/web"

	"github.com/julienschmidt/httprouter"
)

type stubMuslimService struct {
	findDuasFn       func(ctx context.Context, category string) []web.DuaResponse
	findRandomDuaFn  func(ctx context.Context) web.DuaResponse
	findPrayerTimeFn func(ctx context.Context, request web.PrayerTimeRequest) web.PrayerTimeResponse
}

func (stub *stubMuslimService) FindDuas(ctx context.Context, category string) []web.DuaResponse {
	return stub.findDuasFn(ctx, category)
}

func (stub *stubMuslimService) FindRandomDua(ctx context.Context) web.DuaResponse {
	return stub.findRandomDuaFn(ctx)
}

func (stub *stubMuslimService) FindPrayerTime(ctx context.Context, request web.PrayerTimeRequest) web.PrayerTimeResponse {
	return stub.findPrayerTimeFn(ctx, request)
}

func TestMuslimControllerFindDuas(t *testing.T) {
	controller := NewMuslimController(&stubMuslimService{
		findDuasFn: func(ctx context.Context, category string) []web.DuaResponse {
			if category != "daily" {
				t.Fatalf("expected category %q, got %q", "daily", category)
			}

			return []web.DuaResponse{
				{
					ID:       1,
					Title:    "Doa Sebelum Makan",
					Category: "daily",
				},
			}
		},
		findRandomDuaFn: func(ctx context.Context) web.DuaResponse {
			return web.DuaResponse{}
		},
		findPrayerTimeFn: func(ctx context.Context, request web.PrayerTimeRequest) web.PrayerTimeResponse {
			return web.PrayerTimeResponse{}
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/duas?category=daily", nil)
	recorder := httptest.NewRecorder()

	controller.FindDuas(recorder, req, httprouter.Params{})

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	var response web.WebResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !response.Success || response.Message != "Duas fetched successfully" {
		t.Fatalf("unexpected response metadata: %#v", response)
	}
}

func TestMuslimControllerFindPrayerTime(t *testing.T) {
	controller := NewMuslimController(&stubMuslimService{
		findDuasFn: func(ctx context.Context, category string) []web.DuaResponse {
			return nil
		},
		findRandomDuaFn: func(ctx context.Context) web.DuaResponse {
			return web.DuaResponse{}
		},
		findPrayerTimeFn: func(ctx context.Context, request web.PrayerTimeRequest) web.PrayerTimeResponse {
			if request.City != "jakarta" || request.Date != "2026-03-23" {
				t.Fatalf("unexpected prayer time request: %#v", request)
			}

			return web.PrayerTimeResponse{
				City: "Jakarta",
				Date: "2026-03-23",
				Times: web.PrayerTimeSlots{
					Fajr:    "04:32",
					Dhuhr:   "12:01",
					Asr:     "15:15",
					Maghrib: "18:05",
					Isha:    "19:15",
				},
			}
		},
	})

	req := httptest.NewRequest(http.MethodGet, "/prayer-times?city=jakarta&date=2026-03-23", nil)
	recorder := httptest.NewRecorder()

	controller.FindPrayerTime(recorder, req, httprouter.Params{})

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	var response web.WebResponse
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !response.Success || response.Message != "Prayer times fetched successfully" {
		t.Fatalf("unexpected response metadata: %#v", response)
	}
}
