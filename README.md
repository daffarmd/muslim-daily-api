# Muslim Daily API

API MVP sederhana untuk kebutuhan frontend Muslim Daily. Fokus utama saat ini:

- `GET /prayer-times?city=jakarta&date=2026-03-23`
- `GET /duas`
- `GET /duas/random`
- `GET /health`, `GET /healthz`, dan `GET /readyz`

Implementasi sekarang memakai PostgreSQL sebagai source of truth untuk data doa dan jadwal sholat.

## Stack

- Go
- `httprouter`
- dotenv loader sederhana dari file `.env`
- PostgreSQL
- golang-migrate

## Response Format

Semua endpoint mengembalikan JSON konsisten:

```json
{
  "success": true,
  "message": "Prayer times fetched successfully",
  "data": {}
}
```

Contoh error:

```json
{
  "success": false,
  "message": "validation failed",
  "requestId": "abc123",
  "errors": {
    "date": "must use format YYYY-MM-DD"
  }
}
```

## Endpoint

### `GET /duas`

Mengambil semua doa. Mendukung filter kategori opsional:

```text
/duas?category=daily
```

Contoh response:

```json
{
  "success": true,
  "message": "Duas fetched successfully",
  "data": [
    {
      "id": 1,
      "title": "Doa Sebelum Makan",
      "arabic": "اللهم بارك لنا فيما رزقتنا",
      "latin": "Allahumma barik lana fima razaqtana",
      "translation": "Ya Allah, berkahilah rezeki yang Engkau berikan kepada kami.",
      "category": "daily"
    }
  ]
}
```

### `GET /duas/random`

Mengambil satu doa secara acak.

### `GET /prayer-times`

Query wajib:

- `city`
- `date` dengan format `YYYY-MM-DD`

Contoh:

```text
/prayer-times?city=jakarta&date=2026-03-23
```

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

### Health Check

- `GET /health`
- `GET /healthz`
- `GET /readyz`

## Menjalankan Project

1. Salin `.env.example` menjadi `.env`.
2. Isi `DB_DSN`.
3. Jalankan migration:

```powershell
go run ./cmd/migrate up
```

4. Jalankan server:

```powershell
go run .
```

Contoh `DB_DSN`:

```powershell
$env:DB_DSN="postgres://user:password@127.0.0.1:5432/muslim_daily_api?sslmode=disable"
```

## Migration

Schema baru yang ditambahkan:

- tabel `duas`
- tabel `prayer_times`

Seed awal juga disiapkan melalui migration agar endpoint langsung punya data setelah `migrate up`.

Command yang tersedia:

```powershell
go run ./cmd/migrate create add_table_name
go run ./cmd/migrate up
go run ./cmd/migrate down
go run ./cmd/migrate version
go run ./cmd/migrate force 1
```

## Testing

Jalankan seluruh test:

```powershell
go test ./...
```

Test saat ini mencakup:

- helper dan request ID
- middleware request/auth
- exception handler
- Muslim service
- Muslim controller
- Muslim repository dengan `sqlmock`

## OpenAPI

Spesifikasi endpoint tersedia di `apispec.yaml`.
