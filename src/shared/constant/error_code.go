package constant

import (
	"net/http"
	"shoeshop-backend/src/infrastructure/errors"
)

var (
	ErrorDataNotFound     = errors.New(http.StatusNotFound, "4043000", "Error Data Not Found")
	ErrorInvalidToken     = errors.New(http.StatusForbidden, "4013000", "Error Token Invalid")
	ErrorInvalidTokenData = errors.New(http.StatusForbidden, "4013001", "Error Token Data Invalid")
	ErrorInvalidPassword  = errors.New(http.StatusInternalServerError, "4003001", "Invalid Password")
	ErrorInternalServer   = errors.New(http.StatusInternalServerError, "5003001", "Internal Server Error")
)
