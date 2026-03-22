package exception

import (
	"log"
	"net/http"
	"strings"

	"api-go-test/helper"

	"github.com/go-playground/validator"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	if notFoundError(w, r, err) {
		return
	}

	if badRequest(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func notFoundError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		helper.WriteError(w, r, http.StatusNotFound, "resource not found", map[string]string{
			"resource": exception.Error,
		})
		return true
	} else {
		return false
	}

}

func badRequest(w http.ResponseWriter, r *http.Request, err any) bool {
	badRequestError, ok := err.(BadRequestError)
	if ok {
		helper.WriteError(w, r, http.StatusBadRequest, badRequestError.Message, badRequestError.Errors)
		return true
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		helper.WriteError(w, r, http.StatusBadRequest, "validation failed", toValidationErrors(validationErrors))
		return true
	} else {
		return false
	}

}

func internalServerError(w http.ResponseWriter, r *http.Request, err any) {
	log.Printf(
		"request_id=%s method=%s path=%s panic=%v",
		helper.RequestIDFromContext(r.Context()),
		r.Method,
		r.URL.Path,
		err,
	)

	helper.WriteError(w, r, http.StatusInternalServerError, "internal server error", nil)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	helper.WriteError(w, r, http.StatusNotFound, "route not found", nil)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	helper.WriteError(w, r, http.StatusMethodNotAllowed, "method not allowed", nil)
}

func toValidationErrors(validationErrors validator.ValidationErrors) map[string]string {
	errors := make(map[string]string, len(validationErrors))
	for _, fieldError := range validationErrors {
		field := strings.ToLower(fieldError.Field())
		errors[field] = validationMessage(fieldError)
	}

	return errors
}

func validationMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "field is required"
	case "min":
		return "minimum length is " + fieldError.Param()
	case "max":
		return "maximum length is " + fieldError.Param()
	default:
		return "validation failed on rule " + fieldError.Tag()
	}
}
