package repository

import (
	"context"
	"database/sql"
	"errors"

	"api-go-test/helper"
	"api-go-test/model/domain"
)

type DataRepositoryImpl struct {
}

func NewDataRepository() DataRepository {
	return &DataRepositoryImpl{}
}

func (t *DataRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, data domain.Data) domain.Data {
	sqlInsert := "INSERT INTO data (name) VALUES ($1) RETURNING id, status"

	err := tx.QueryRowContext(ctx, sqlInsert, data.Name).Scan(&data.Id, &data.Status)
	helper.PanicIfErr(err)

	return data
}

func (t *DataRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, data domain.Data) domain.Data {
	sqlUpdate := "UPDATE data SET name = $1 where id = $2"
	_, err := tx.ExecContext(ctx, sqlUpdate, data.Name, data.Id)
	helper.PanicIfErr(err)

	return data
}

func (t *DataRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, data domain.Data) {
	sqlDelete := "UPDATE data SET status = $1 where id = $2"
	_, err := tx.ExecContext(ctx, sqlDelete, -2, data.Id)
	helper.PanicIfErr(err)
}

func (t *DataRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, dataId int) (domain.Data, error) {
	sqlq := "SELECT id, name, status FROM data WHERE id = $1"
	rows, err := tx.QueryContext(ctx, sqlq, dataId)
	helper.PanicIfErr(err)
	defer rows.Close()

	data := domain.Data{}
	if rows.Next() {
		err := rows.Scan(&data.Id, &data.Name, &data.Status)
		helper.PanicIfErr(err)
		return data, nil
	}
	return data, errors.New("data is not found")
}

func (t *DataRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Data {
	sqlq := "SELECT id, name, status FROM data"
	rows, err := tx.QueryContext(ctx, sqlq)
	helper.PanicIfErr(err)
	defer rows.Close()

	var datas []domain.Data
	for rows.Next() {
		data := domain.Data{}
		err := rows.Scan(&data.Id, &data.Name, &data.Status)
		helper.PanicIfErr(err)
		datas = append(datas, data)
	}
	helper.PanicIfErr(rows.Err())

	return datas
}

func (t *DataRepositoryImpl) FindAllAsc(ctx context.Context, tx *sql.Tx) []domain.Data {
	sqlq := "SELECT id, name, status FROM data ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, sqlq)
	helper.PanicIfErr(err)
	defer rows.Close()

	var datas []domain.Data
	for rows.Next() {
		data := domain.Data{}
		err := rows.Scan(&data.Id, &data.Name, &data.Status)
		helper.PanicIfErr(err)
		datas = append(datas, data)
	}
	helper.PanicIfErr(rows.Err())

	return datas
}
