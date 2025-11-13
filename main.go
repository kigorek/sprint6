package main

import (
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

var text = ".--. .-. .. .-- . -"

//var text = "Привет"

func main() {
	conv, _ := service.MorseOrText(text)
	fmt.Println(conv)
}
