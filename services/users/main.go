package main

import (
	"fmt"

	"github.com/adhitamafikri/go-simple-pms/services/users/configs"
)

func main() {
	fmt.Println("Booting up 'users' service...")
	configs.Bootstrap()
	fmt.Println("'users' service running!!!...")
}
