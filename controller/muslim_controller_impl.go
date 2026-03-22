package controller

import (
	"net/http"

	"api-go-test/helper"
	"api-go-test/model/web"
	"api-go-test/service"

	"github.com/julienschmidt/httprouter"
)

type MuslimControllerImpl struct {
	MuslimService service.MuslimService
}

func NewMuslimController(muslimService service.MuslimService) MuslimController {
	return &MuslimControllerImpl{
		MuslimService: muslimService,
	}
}

func (controller *MuslimControllerImpl) FindDuas(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	category := r.URL.Query().Get("category")
	data := controller.MuslimService.FindDuas(r.Context(), category)

	helper.WriteSuccess(w, http.StatusOK, "Duas fetched successfully", data)
}

func (controller *MuslimControllerImpl) FindRandomDua(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := controller.MuslimService.FindRandomDua(r.Context())

	helper.WriteSuccess(w, http.StatusOK, "Random dua fetched successfully", data)
}

func (controller *MuslimControllerImpl) FindPrayerTime(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	request := web.PrayerTimeRequest{
		City: r.URL.Query().Get("city"),
		Date: r.URL.Query().Get("date"),
	}

	data := controller.MuslimService.FindPrayerTime(r.Context(), request)

	helper.WriteSuccess(w, http.StatusOK, "Prayer times fetched successfully", data)
}
