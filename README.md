# api-go-test

Boilerplate REST API sederhana dengan Go, PostgreSQL, `httprouter`, dan struktur layer `controller/service/repository`.

## Fitur Dasar
- Public route untuk read endpoint
- Protected route untuk write endpoint dengan `X-API-Key`
- Health check: `/healthz` dan `/readyz`
- Request ID di header `X-Request-ID`
- Logging request dasar
- Graceful shutdown
- Konfigurasi via environment variable
- Database migration command

## Environment Variable
- `APP_ENV`
- `PORT`
- `DB_DSN`
- `API_KEY`
- `READ_TIMEOUT`
- `WRITE_TIMEOUT`
- `IDLE_TIMEOUT`
- `SHUTDOWN_TIMEOUT`

## Keterangan Route
- `/healthz`
  - Endpoint liveness check
  - Dipakai untuk mengecek apakah service hidup

- `/readyz`
  - Endpoint readiness check
  - Dipakai untuk mengecek apakah service siap menerima request, termasuk koneksi database

- `/api/`
  - Prefix utama untuk endpoint API aplikasi
  - Dipakai untuk endpoint resource public maupun protected

- `/api/data`
  - Contoh resource bawaan di boilerplate ini
  - Bisa diganti nanti menjadi resource utama seperti `/api/hijri`, `/api/calendar`, `/api/prayer-times`, dan lain-lain

## Default Endpoint
- `GET /healthz`
  - Mengecek apakah aplikasi hidup

- `GET /readyz`
  - Mengecek apakah aplikasi siap melayani request

- `GET /api/data`
  - Public endpoint untuk mengambil semua data

- `GET /api/data/:dataId`
  - Public endpoint untuk mengambil detail data

- `POST /api/data`
  - Protected endpoint untuk membuat data baru

- `PUT /api/data/:dataId`
  - Protected endpoint untuk mengubah data

- `DELETE /api/data/:dataId`
  - Protected endpoint untuk menghapus data

## Menjalankan Project
```powershell
$env:DB_DSN="postgres://user:password@127.0.0.1:5432/dbname?sslmode=disable"
$env:API_KEY="change-me"
go run ./cmd/migrate up
go run .
```

Atau jika memakai `make`:

```bash
make migrate-up
make run
```

## Testing
Jalankan seluruh unit test:

```powershell
go test ./...
```

Atau jika memakai `make`:

```bash
make test
```

Unit test saat ini mencakup:
- helper
- request ID helper
- auth middleware
- request/logging middleware
- exception handler
- service layer dengan `sqlmock`

## Database Migration
Migration file disimpan di folder [migrations](c:/Users/VICTUS/OneDrive/Documents/TNOD/api/api-go-test/migrations).

Command yang tersedia:

```powershell
go run ./cmd/migrate create add_users_table
go run ./cmd/migrate up
go run ./cmd/migrate down
go run ./cmd/migrate version
go run ./cmd/migrate force 1
```

Versi `make`:

```bash
make migrate-create NAME=add_users_table
make migrate-up
make migrate-down
make migrate-version
make migrate-force VERSION=1
```

Keterangan:
- `create <name>` membuat file migration `.up.sql` dan `.down.sql` baru
- `up` menjalankan semua migration yang belum diaplikasikan
- `down` rollback 1 step migration terakhir
- `version` melihat versi migration aktif
- `force <version>` dipakai jika status migration dirty dan perlu dipaksa ke versi tertentu

Migration pertama yang tersedia:
- `000001_create_data_table.up.sql`
  - membuat tabel `data`
  - membuat index `idx_data_status`

## Auth
Untuk endpoint write, kirim header:

```text
X-API-Key: change-me
```
