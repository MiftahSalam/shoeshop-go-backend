package main

import (
	"shoeshop-backend/src/di"
	"shoeshop-backend/src/interfaces"
)

func main() {
	interfaces.Start(di.Setup())
}
