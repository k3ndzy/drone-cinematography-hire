package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	addr := ":8151"
	fmt.Printf("🚁 🚁 SkyLens - 4K Drone Filming & Pilot Hire Server đang chạy tại http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
