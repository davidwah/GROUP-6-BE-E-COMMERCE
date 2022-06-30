package presentation

import (
	"errors"
	"net/http"
	"strconv"

	"construct-week1/features/helper"
	"construct-week1/features/products"
	"construct-week1/features/products/presentation/request"
	"construct-week1/features/products/presentation/response"
	"construct-week1/middlewares"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productBusiness products.Business
}

func NewProductHandler(business products.Business) *ProductHandler {
	return &ProductHandler{
		productBusiness: business,
	}
}

//	e.POST("/products")
func (handle *ProductHandler) AddProduct(c echo.Context) error {
	//	Merupakan validasi user id yaitu menggunakan token
	idUser := middlewares.ExtractToken(c)
	if idUser == 0 {
		return errors.New("failed")
	}

	//	Proses pembuatan data product
	var newProduct request.Product
	newProduct.IDUser = uint(idUser)
	errCreate := c.Bind(&newProduct)
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed("Failed to bind your data product, check your input"))
	}

	dataProduct := request.ToCore(newProduct)
	row, err := handle.productBusiness.InsertProduct(dataProduct)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, helper.ResponseFailed(err.Error()))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success to insert your data product"))
}

//	e.GET("/products")
func (handle *ProductHandler) FindProduct(c echo.Context) error {
	//	Proses pengambilan semua data product
	limit, offset := c.QueryParam("limit"), c.QueryParam("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	res, err := handle.productBusiness.GetProduct(limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get all data product"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to get all data product", response.FromCoreList(res)))
}

//	e.GET("/users/products/:id")
func (handler *ProductHandler) FindProductbyID(c echo.Context) error {
	//	Merupakan validasi user id yaitu menggunakan token
	idToken := middlewares.ExtractToken(c)

	id := c.Param("id")
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("ID not recognized"))
	}

	if idToken != idProduct {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
	}

	//	Proses pengambilan data product yang dimiliki oleh user
	res, err := handler.productBusiness.GetProductbyID(uint(idProduct))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get your data product"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to get your data product", response.FromCoreListGetID(res)))
}

//	e.PUT("/products/:id")
func (handle *ProductHandler) UpdateProduct(c echo.Context) error {
	//	Merupakan validasi user id yaitu menggunakan token
	idToken := middlewares.ExtractToken(c)

	id := c.Param("id")
	idUser, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("ID not recognized"))
	}

	if idToken != idUser {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unauthorized"))
	}

	data := map[string]interface{}{
		"id":          c.FormValue("id"),
		"name":        c.FormValue("name"),
		"price":       c.FormValue("price"),
		"quantity":    c.FormValue("qty"),
		"description": c.FormValue("desc"),
		// "UpdatedAt":   time.Now(),
	}

	idproducts, _ := data["id"].(string)
	res := handle.productBusiness.UpdateProduct(idproducts, data)
	if res != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to your update data product"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessNoData("Success update your data product"))
}

func (handle *ProductHandler) FindProductByID(c echo.Context) error {
	err := middlewares.ExtractToken(c)
	if err == 0 {
		return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("Unathorized"))
	}

	id := c.Param("id")
	idProduct, errId := strconv.Atoi(id)
	if errId != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("ID not recognized"))
	}

	productId, err2 := handle.productBusiness.GetProductID(idProduct)
	if err2 != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed get your data"))
	}
	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to get your data", productId))
}

// func (handle *ProductHandler) DeleteProduct(id uint) error {

// }
