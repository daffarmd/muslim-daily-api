package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MuslimController interface {
	FindDuas(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindRandomDua(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindPrayerTime(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
