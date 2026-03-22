package service

import (
	"context"

	"api-go-test/model/web"
)

type DataService interface {
	Create(ctx context.Context, request web.DataCreateRequest) web.DataResponse
	Update(ctx context.Context, request web.DataUpdateRequest) web.DataResponse
	Delete(ctx context.Context, dataId int)
	FindById(ctx context.Context, dataId int) web.DataResponse
	FindAll(ctx context.Context) []web.DataResponse
	FindAllAsc(ctx context.Context) []web.DataResponse
}
