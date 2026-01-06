package main

import (
	"fmt"

	"github.com/lolorine16/super-succotash/project01/utils" // importation d'un package local
)

func main() {
	message := utils.DireSalut("Laureen")
	fmt.Println(message)
}
