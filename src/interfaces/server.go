package interfaces

import (
	"fmt"

	"golang.org/x/sync/errgroup"

	"shoeshop-backend/src/di"
	"shoeshop-backend/src/interfaces/http"
)

func Start(di *di.DI) {
	errs := errgroup.Group{}

	errs.Go(func() error {
		http.Start(di)
		return nil
	})

	if err := errs.Wait(); err != nil {
		fmt.Println(err)
	}
}
