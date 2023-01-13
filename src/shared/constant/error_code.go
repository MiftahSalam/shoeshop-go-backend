package constant

import (
	"net/http"
	"shoeshop-backend/src/infrastructure/errors"
)

var (
	ErrorDataNotFound   = errors.New(http.StatusNotFound, "4043000", "Error Data Not Found")
	ErrorInternalServer = errors.New(http.StatusInternalServerError, "5003001", "Internal Server Error")
)
