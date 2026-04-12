package main

import (
	"fmt"

	"github.com/adhitamafikri/go-simple-pms/services/users/configs"
)

func main() {
	fmt.Println("This is the users service")
	configs.Bootstrap()
}
