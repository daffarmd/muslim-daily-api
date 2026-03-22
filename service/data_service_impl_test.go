package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"api-go-test/exception"
	"api-go-test/model/domain"
	"api-go-test/model/web"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator"
)

type stubDataRepository struct {
	createFn   func(ctx context.Context, tx *sql.Tx, data domain.Data) domain.Data
	updateFn   func(ctx context.Context, tx *sql.Tx, data domain.Data) domain.Data
	deleteFn   func(ctx context.Context, tx *sql.Tx, data domain.Data)
	findByIDFn func(ctx context.Context, tx *sql.Tx, dataID int) (domain.Data, error)
	findAllFn  func(ctx context.Context, tx *sql.Tx) []domain.Data
	findAscFn  func(ctx context.Context, tx *sql.Tx) []domain.Data
}

func (s *stubDataRepository) Create(ctx context.Context, tx *sql.Tx, data domain.Data) domain.Data {
	return s.createFn(ctx, tx, data)
}

func (s *stubDataRepository) Update(ctx context.Context, tx *sql.Tx, data domain.Data) domain.Data {
	return s.updateFn(ctx, tx, data)
}

func (s *stubDataRepository) Delete(ctx context.Context, tx *sql.Tx, data domain.Data) {
	s.deleteFn(ctx, tx, data)
}

func (s *stubDataRepository) FindById(ctx context.Context, tx *sql.Tx, dataID int) (domain.Data, error) {
	return s.findByIDFn(ctx, tx, dataID)
}

func (s *stubDataRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.Data {
	return s.findAllFn(ctx, tx)
}

func (s *stubDataRepository) FindAllAsc(ctx context.Context, tx *sql.Tx) []domain.Data {
	return s.findAscFn(ctx, tx)
}

func TestDataServiceCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectCommit()

	repo := &stubDataRepository{
		createFn: func(ctx context.Context, tx *sql.Tx, data domain.Data) domain.Data {
			data.Id = 1
			data.Status = 1
			return data
		},
	}

	service := NewDataService(repo, db, validator.New())

	response := service.Create(t.Context(), web.DataCreateRequest{Name: "Ramadan"})

	if response.Id != 1 || response.Name != "Ramadan" || response.Status != 1 {
		t.Fatalf("unexpected response: %#v", response)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestDataServiceCreatePanicsOnValidationError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	service := NewDataService(&stubDataRepository{}, db, validator.New())

	defer func() {
		if recover() == nil {
			t.Fatal("expected panic on invalid create request")
		}
	}()

	service.Create(t.Context(), web.DataCreateRequest{})
}

func TestDataServiceFindByIDPanicsOnNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectRollback()

	repo := &stubDataRepository{
		findByIDFn: func(ctx context.Context, tx *sql.Tx, dataID int) (domain.Data, error) {
			return domain.Data{}, errors.New("data is not found")
		},
	}

	service := NewDataService(repo, db, validator.New())

	defer func() {
		rec := recover()
		if rec == nil {
			t.Fatal("expected panic on missing data")
		}

		if _, ok := rec.(exception.NotFoundError); !ok {
			t.Fatalf("expected NotFoundError panic, got %#v", rec)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet sql expectations: %v", err)
		}
	}()

	service.FindById(t.Context(), 99)
}

func TestDataServiceFindAllAsc(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectCommit()

	repo := &stubDataRepository{
		findAscFn: func(ctx context.Context, tx *sql.Tx) []domain.Data {
			return []domain.Data{
				{Id: 1, Name: "Muharram", Status: 1},
				{Id: 2, Name: "Safar", Status: 1},
			}
		},
	}

	service := NewDataService(repo, db, validator.New())

	response := service.FindAllAsc(t.Context())

	if len(response) != 2 {
		t.Fatalf("expected 2 responses, got %d", len(response))
	}

	if response[0].Name != "Muharram" || response[1].Name != "Safar" {
		t.Fatalf("unexpected response ordering: %#v", response)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}
