package app

import (
	"image_storage/internal/services"
	"image_storage/internal/transport"
	"log"
	"net/http"
)

func Run() {
	mux := http.NewServeMux()

	service := services.NewImgService()
	handler := transport.NewImgHandler(service)

	mux.HandleFunc("POST /case/{id}/img", handler.UploadCaseImg)
	mux.HandleFunc("GET /case/{id}/img", handler.GetCaseImg)
	mux.HandleFunc("GET /test", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("OK!"))
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
