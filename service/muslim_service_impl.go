package service

import (
	"context"
	"strings"
	"time"

	"api-go-test/exception"
	"api-go-test/helper"
	"api-go-test/model/web"
	"api-go-test/repository"
)

const prayerTimeDateLayout = "2006-01-02"

type MuslimServiceImpl struct {
	MuslimRepository repository.MuslimRepository
}

func NewMuslimService(muslimRepository repository.MuslimRepository) MuslimService {
	return &MuslimServiceImpl{
		MuslimRepository: muslimRepository,
	}
}

func (service *MuslimServiceImpl) FindDuas(ctx context.Context, category string) []web.DuaResponse {
	duas := service.MuslimRepository.FindDuas(ctx, category)
	return helper.ToDuaResponses(duas)
}

func (service *MuslimServiceImpl) FindRandomDua(ctx context.Context) web.DuaResponse {
	dua, err := service.MuslimRepository.FindRandomDua(ctx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToDuaResponse(dua)
}

func (service *MuslimServiceImpl) FindPrayerTime(ctx context.Context, request web.PrayerTimeRequest) web.PrayerTimeResponse {
	validationErrors := make(map[string]string)

	if strings.TrimSpace(request.City) == "" {
		validationErrors["city"] = "field is required"
	}

	if strings.TrimSpace(request.Date) == "" {
		validationErrors["date"] = "field is required"
	} else if _, err := time.Parse(prayerTimeDateLayout, request.Date); err != nil {
		validationErrors["date"] = "must use format YYYY-MM-DD"
	}

	if len(validationErrors) > 0 {
		panic(exception.NewBadRequestError("validation failed", validationErrors))
	}

	prayerTime, err := service.MuslimRepository.FindPrayerTime(ctx, request.City, request.Date)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPrayerTimeResponse(prayerTime)
}
