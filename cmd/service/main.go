package main

import (
	"errors"
	"fmt"
	"frisboo-bank/customer-service/internal/service"
	"os"
)

var ErrCustomerService = errors.New("customer service: failed unexpectedly")

func main() {
	err := service.Start()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
