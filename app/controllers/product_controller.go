package controllers

import (
	"net/http"

	"github.com/ssr0016/gotoko/app/models"
	"github.com/unrolled/render"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	productModal := models.Product{}
	products, err := productModal.GetProducts(server.DB)
	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products": products,
	})
}
