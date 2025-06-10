package controller

import (
	"backend/src/domain"
	"net/http"
)

func statusCode(err error) int {
	switch err.Error() {
	case domain.BadRequest:
		return http.StatusBadRequest
	case domain.Unauthorized:
		return http.StatusUnauthorized
	case domain.Forbidden:
		return http.StatusForbidden
	case domain.NotFound:
		return http.StatusNotFound
	case domain.Conflict:
		return http.StatusConflict
	case domain.InternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusNotImplemented
	}
}
