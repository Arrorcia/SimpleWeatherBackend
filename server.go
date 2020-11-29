package main

import (
	"fmt"
	"github.com/ame-lm/SimpleWeather/handler"
	"net/http"
)

func main() {
	http.Handle("/demo", new(handler.DemoHandler))

	fmt.Println(http.ListenAndServe("localhost:12345", nil))
}
