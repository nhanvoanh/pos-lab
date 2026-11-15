package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// version được gán lúc build bằng -ldflags, giống cách CI đóng dấu phiên bản.
var version = "dev"

func main() {
	hostname, _ := os.Hostname()

	mux := http.NewServeMux()

	// Kubernetes gọi đường dẫn này để biết pod còn sống và sẵn sàng nhận request.
	mux.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"service": "pos-gateway-lab",
			"version": version,
			"env":     os.Getenv("APP_ENV"),
			"pod":     hostname,
			// Chỉ báo có nhận được secret hay không, không bao giờ in giá trị ra.
			"apiKey": map[bool]string{true: "đã nhận", false: "chưa có"}[os.Getenv("API_KEY") != ""],
		})
	})

	log.Printf("pos-gateway %s listening on :3000", version)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
