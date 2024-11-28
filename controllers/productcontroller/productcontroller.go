package productcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/izzuddinm/go-restapi-mux/helper"
	"github.com/izzuddinm/go-restapi-mux/models"
	"gorm.io/gorm"
)

var ResponseJson = helper.ResponseJson
var ResponseSuccess = helper.ResponseSuccess
var ResponseSuccessWithMessage = helper.ResponseSuccessWithMessage
var ResponseError = helper.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	// Implementation for listing products
	var products []models.Product
	if err := models.DB.Find(&products).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseSuccess(w, http.StatusOK, products)
}

func Show(w http.ResponseWriter, r *http.Request) {
	// Implementation for showing a single product
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Product was not found.")
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if product.Id == 0 {
		ResponseError(w, http.StatusNotFound, "Product was not found.")
		return
	}

	ResponseSuccess(w, http.StatusOK, product)
}

func Create(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a new product
	var product models.Product

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Save the new product to the database
	if err := models.DB.Create(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseSuccess(w, http.StatusCreated, product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating an existing product
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Product not found")
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Decode the updated product data
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Update the product in the database
	if err := models.DB.Save(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseSuccess(w, http.StatusOK, product)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the URL parameters
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	// Find the product by ID
	var product models.Product
	if err := models.DB.First(&product, id).Error; err != nil {
		// Handle the case where the product is not found
		if err == gorm.ErrRecordNotFound {
			ResponseError(w, http.StatusNotFound, "Product not found")
		} else {
			// Handle any other errors
			ResponseError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	// Delete the product
	if err := models.DB.Delete(&product).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, "Failed to delete product: "+err.Error())
		return
	}

	// Return a success message after deletion
	ResponseSuccessWithMessage(w, http.StatusOK, "Product successfully deleted")
}
