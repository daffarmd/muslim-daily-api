package web

type PrayerTimeSlots struct {
	Fajr    string `json:"fajr"`
	Dhuhr   string `json:"dhuhr"`
	Asr     string `json:"asr"`
	Maghrib string `json:"maghrib"`
	Isha    string `json:"isha"`
}

type PrayerTimeResponse struct {
	City  string          `json:"city"`
	Date  string          `json:"date"`
	Times PrayerTimeSlots `json:"times"`
}
