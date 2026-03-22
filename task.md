# 🌙 Muslim API MVP — Task Plan

Dokumen ini adalah panduan untuk membangun API sederhana yang bermanfaat bagi umat Muslim, dengan fokus pada fitur inti yang realistis untuk MVP (Minimum Viable Product).

---

## 🎯 Tujuan

Membangun API yang:

* Menyediakan jadwal sholat berdasarkan kota dan tanggal
* Menyediakan doa harian
* Memiliki struktur endpoint yang rapi dan mudah digunakan oleh frontend

---

## 📦 Scope MVP

### Fitur Wajib

1. `GET /prayer-times?city=jakarta&date=2026-03-23`
2. `GET /duas`
3. `GET /duas/random`

### Fitur Opsional (Setelah MVP)

* `GET /cities`
* `GET /health`
* `GET /hijri-date`
* `GET /ayah/random`

---

## 🛠 Tech Stack

* Golang
* PostgreSQL
* Postman / Thunder Client (untuk testing)
* dotenv (untuk environment variables)

---

## 🧱 Struktur Data

### Prayer Times

Field minimal:

* city
* date
* fajr
* dhuhr
* asr
* maghrib
* isha

### Dua

Field minimal:

* id
* title
* arabic
* latin
* translation
* category

---

## 📋 Task Breakdown

### 1. Planning

* [ ] Tentukan nama project API
* [ ] Tentukan scope MVP
* [ ] Tentukan format response JSON
* [ ] Tentukan daftar endpoint
* [ ] Tentukan sumber data awal

---

### 2. Setup Project

* [ ] Inisialisasi project Golang
* [ ] Setup routing (gunakan framework seperti Gin / Fiber jika perlu)
* [ ] Setup koneksi database PostgreSQL
* [ ] Setup environment variables (.env)
* [ ] Struktur folder project

Contoh struktur:

```txt
muslim-api/
  handlers/
  routes/
  models/
  database/
  utils/
  middlewares/
  main.go
  .env
  go.mod
```

---

### 3. Endpoint: Prayer Times

* [ ] Buat endpoint `GET /prayer-times`
* [ ] Tambahkan query param `city`
* [ ] Tambahkan query param `date`
* [ ] Validasi input
* [ ] Handle jika data tidak ditemukan
* [ ] Pastikan response konsisten

Contoh response:

```json
{
  "success": true,
  "message": "Prayer times fetched successfully",
  "data": {
    "city": "Jakarta",
    "date": "2026-03-23",
    "times": {
      "fajr": "04:32",
      "dhuhr": "12:01",
      "asr": "15:15",
      "maghrib": "18:05",
      "isha": "19:15"
    }
  }
}
```

---

### 4. Endpoint: Duas

* [ ] Buat endpoint `GET /duas`
* [ ] Tambahkan filter kategori (opsional)
* [ ] Buat endpoint `GET /duas/random`
* [ ] Ambil satu data secara acak

Contoh response:

```json
{
  "success": true,
  "message": "Random dua fetched successfully",
  "data": {
    "id": 1,
    "title": "Doa Sebelum Makan",
    "arabic": "اللَّهُمَّ بَارِكْ لَنَا فِيمَا رَزَقْتَنَا",
    "latin": "Allahumma barik lana fima razaqtana",
    "translation": "Ya Allah, berkahilah rezeki yang Engkau berikan kepada kami.",
    "category": "daily"
  }
}
```

---

### 5. Utility & Error Handling

* [ ] Buat helper untuk response sukses
* [ ] Buat helper untuk response error
* [ ] Middleware untuk 404
* [ ] Middleware error handler
* [ ] Pastikan semua endpoint konsisten

---

### 6. Data Management

* [ ] Siapkan data doa (JSON atau database)
* [ ] Siapkan data jadwal sholat
* [ ] Tentukan apakah data statis atau dari API eksternal
* [ ] Pastikan konsistensi naming field

---

### 7. Testing

* [ ] Test endpoint di Postman
* [ ] Test success case
* [ ] Test error case
* [ ] Test input tidak valid
* [ ] Test data tidak ditemukan

---

### 8. Documentation

* [ ] Tulis README
* [ ] Cara install project
* [ ] Cara menjalankan server
* [ ] Dokumentasi endpoint
* [ ] Contoh request & response

---

### 9. Bonus Improvement

* [ ] Tambahkan endpoint `GET /health`
* [ ] Tambahkan endpoint `GET /cities`
* [ ] Tambahkan pagination untuk `/duas`
* [ ] Tambahkan Swagger / OpenAPI
* [ ] Deploy ke Railway / Render

---

## 🚀 Urutan Pengerjaan

1. Setup project
2. Implement `GET /duas`
3. Implement `GET /duas/random`
4. Implement `GET /prayer-times`
5. Tambahkan error handling
6. Buat dokumentasi
7. Testing
8. Deploy

---

## 💡 Nama Project

* Sahabat Ibadah API
* Muslim Daily API
* Barakah API
* Deen API
* Nur API

---

## 🎯 Target MVP

Fokus pada fitur berikut:

* [ ] `GET /duas`
* [ ] `GET /duas/random`
* [ ] `GET /prayer-times`
* [ ] README sederhana
* [ ] Testing endpoint

---

## 📝 Catatan Penting

* Fokus pada kesederhanaan terlebih dahulu
* Pastikan struktur response konsisten
* Gunakan error message yang jelas
* Jangan menambahkan terlalu banyak fitur di awal
* Prioritaskan penyelesaian MVP

---

✨ Dengan menyelesaikan MVP ini, kamu sudah punya fondasi API yang bisa dikembangkan menjadi produk yang lebih besar.