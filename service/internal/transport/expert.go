package transport

import (
	"TrafficPolice/internal/services"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

const (
	expertsDir            = "experts"
	expertContentImageKey = "image"
	expertIDPathValue     = "id"
)

type ExpertHandler struct {
	imgService    services.ImgService
	expertService services.ExpertService
}

func NewExpertHandler(imgService services.ImgService, expertService services.ExpertService) *ExpertHandler {
	return &ExpertHandler{
		imgService:    imgService,
		expertService: expertService,
	}
}

func (h *ExpertHandler) UploadExpertImg(w http.ResponseWriter, r *http.Request) {
	expertID := r.PathValue(expertIDPathValue)
	if expertID == "" {
		http.Error(w, "id is empty", http.StatusBadRequest)
		return
	}

	file, header, err := parseMultipartForm(r, expertContentImageKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contentType := header.Header.Get(contentTypeKey)
	extension, err := getImgExtension(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	imgFilePath := fmt.Sprintf("%s/%s.%s", expertsDir, expertID, extension)
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error while reading fileBytes: %v\n", fileBytes)
		return
	}

	err = h.imgService.SaveImg(fileBytes, imgFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Succesfully uploaded image")
}

func (h *ExpertHandler) GetExpertImg(w http.ResponseWriter, r *http.Request) {
	expertID := r.PathValue(expertIDPathValue)
	if expertID == "" {
		http.Error(w, "bad expert id", http.StatusBadRequest)
		return
	}

	pattern := fmt.Sprintf("%s/%s.*", expertsDir, expertID)
	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	http.ServeFile(w, r, files[0])
}
