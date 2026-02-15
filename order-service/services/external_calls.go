package services

import (
	"net/http"
)

func ValidateUser(userID uint) bool {
	resp, err := http.Get("http://user-service:8001/users")
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}

func ValidateProduct(productID uint) bool {
	resp, err := http.Get("http://product-service:8002/products")
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	return true
}
