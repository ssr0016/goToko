package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/ssr0016/gotoko/app/models"
)

func GetShoppingCartID(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, "shopping-cart-session")
	if session.Values["cart-id"] == nil {
		session.Values["cart-id"] = uuid.New().String()
		session.Save(r, w)
	}

	return fmt.Sprintf("%v", session.Values["cart-id"])
}

func (server *Server) GetCart(w http.ResponseWriter, r *http.Request) {
	panic("cart items")
}
func (server *Server) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	productID := r.FormValue("product_id")
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	productModel := models.Product{}
	product, err := productModel.FindByID(server.DB, productID)
	if err != nil {
		http.Redirect(w, r, "/products/"+product.Slug, http.StatusSeeOther)
	}

	if qty > product.Stock {
		http.Redirect(w, r, "/products/"+product.Slug, http.StatusSeeOther)
	}

	var cart *models.Cart

	cartID := GetShoppingCartID(w, r)
	http.Redirect(w, r, "/carts/", http.StatusSeeOther)

}
