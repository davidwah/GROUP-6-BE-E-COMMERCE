package presentation

import (
	"errors"
	"net/http"
	"strconv"
	"time"

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

//	POST /products
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

//	GET /products
func (handle *ProductHandler) GetAllProduct(c echo.Context) error {
	//	Merupakan validasi user id yaitu menggunakan token
	id := middlewares.ExtractToken(c)
	if id == 0 {
		return errors.New("failed")
	}

	//	Proses pengambilan semua data product
	limit, offset := c.QueryParam("limit"), c.QueryParam("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	res, err := handle.productBusiness.GetProduct(limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Faile to get all data product"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to get all data product", response.FromCoreList(res)))
}

//	GET /products:id
func (handler *ProductHandler) GetProductbyID(c echo.Context) error {
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

//	PUT /products:id
func (handle *ProductHandler) UpdateProduct(c echo.Context) error {
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

	//	Proses pengupdatan data product yang dimiliki oleh user
	data := map[string]interface{}{
		"ID":        c.FormValue("ID"),
		"name":      c.FormValue("name"),
		"price":     c.FormValue("price"),
		"qty":       c.FormValue("qty"),
		"desc":      c.FormValue("desc"),
		"UpdatedAt": time.Now(),
	}

	idProducts, _ := data["ID"].(string)
	res := handle.productBusiness.UpdateProduct(idProducts, data)
	if res != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update your data product"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success update your data product", data))
}

// func (handle *ProductHandler) DeleteProduct(id uint) error {

// }