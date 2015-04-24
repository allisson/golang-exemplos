package main

import (
	"github.com/disintegration/imaging"
	"io"
	"log"
	"net/http"
	"strconv"
)

type ImageFetchResult struct {
	Error  error
	Reader io.Reader
}

func fetchImage(url string, channel chan ImageFetchResult) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		channel <- ImageFetchResult{Error: err}
		return
	}
	channel <- ImageFetchResult{Reader: res.Body}
}

func index(w http.ResponseWriter, r *http.Request) {
	imageUrl := r.URL.Query().Get("url")
	imageWidth := r.URL.Query().Get("width")
	imageHeight := r.URL.Query().Get("height")

	if imageUrl == "" || imageWidth == "" || imageHeight == "" {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	channel := make(chan ImageFetchResult)
	go fetchImage(imageUrl, channel)
	result := <-channel
	if result.Error != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	img, err := imaging.Decode(result.Reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	widthFinal, _ := strconv.Atoi(imageWidth)
	heightFinal, _ := strconv.Atoi(imageHeight)
	thumb := imaging.Thumbnail(img, widthFinal, heightFinal, imaging.CatmullRom)
	imaging.Encode(w, thumb, imaging.JPEG)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
