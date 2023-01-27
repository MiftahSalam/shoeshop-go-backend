package interfaces

import (
	"shoeshop-backend/src/di"
	"shoeshop-backend/src/interfaces/http"
)

func Start(di *di.DI) {
	http.Start(di)
}
