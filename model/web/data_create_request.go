package web

type DataCreateRequest struct {
	Name string `validate:"required,max=200,min=1"`
}
