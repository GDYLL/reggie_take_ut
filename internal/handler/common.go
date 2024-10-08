package handler

import (
	"net/http"
	"path/filepath"
)

type CommonController struct{}

var path = "static/images/"

func (c CommonController) Download(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Query().Get("name")
	imgUrl := filepath.Join(path, imageName)
	w.Header().Add("Content-Disposition", "attachment;filename="+imageName)
	http.ServeFile(w, r, imgUrl)
}
