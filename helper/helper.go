package helper

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"api-go-test/model/domain"
	"api-go-test/model/web"
)

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFromReqBody(r *http.Request, result any) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfErr(err)
}

func WriteResponseBody(w http.ResponseWriter, response any) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfErr(err)
}

func WriteSuccess(w http.ResponseWriter, statusCode int, message string, data any) {
	w.WriteHeader(statusCode)
	WriteResponseBody(w, web.WebResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func WriteError(w http.ResponseWriter, r *http.Request, statusCode int, message string, errors any) {
	w.WriteHeader(statusCode)
	WriteResponseBody(w, web.WebResponse{
		Success:   false,
		Message:   message,
		RequestID: RequestIDFromContext(r.Context()),
		Errors:    errors,
	})
}

func CommitOrRollback(tx *sql.Tx) {
	if rec := recover(); rec != nil {
		_ = tx.Rollback()
		panic(rec)
	}

	err := tx.Commit()
	PanicIfErr(err)
}

func ToDataResponse(data domain.Data) web.DataResponse {
	return web.DataResponse{
		Id:     data.Id,
		Name:   data.Name,
		Status: data.Status,
	}
}

func ToDataResponses(datas []domain.Data) []web.DataResponse {
	responses := make([]web.DataResponse, 0, len(datas))
	for _, data := range datas {
		responses = append(responses, ToDataResponse(data))
	}

	return responses
}

func ToDuaResponse(dua domain.Dua) web.DuaResponse {
	return web.DuaResponse{
		ID:          dua.ID,
		Title:       dua.Title,
		Arabic:      dua.Arabic,
		Latin:       dua.Latin,
		Translation: dua.Translation,
		Category:    dua.Category,
	}
}

func ToDuaResponses(duas []domain.Dua) []web.DuaResponse {
	responses := make([]web.DuaResponse, 0, len(duas))
	for _, dua := range duas {
		responses = append(responses, ToDuaResponse(dua))
	}

	return responses
}

func ToPrayerTimeResponse(prayerTime domain.PrayerTime) web.PrayerTimeResponse {
	return web.PrayerTimeResponse{
		City: prayerTime.City,
		Date: prayerTime.Date,
		Times: web.PrayerTimeSlots{
			Fajr:    prayerTime.Fajr,
			Dhuhr:   prayerTime.Dhuhr,
			Asr:     prayerTime.Asr,
			Maghrib: prayerTime.Maghrib,
			Isha:    prayerTime.Isha,
		},
	}
}
