package presentation

import (
	"construct-week1/features/helper"
	"construct-week1/features/products"
	"construct-week1/features/products/presentation/request"
	"construct-week1/features/products/presentation/response"
	"construct-week1/middlewares"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

type ProductHandler struct {
	productBusiness products.Business
}

func NewProductHandler(business products.Business) *ProductHandler {
	return &ProductHandler{
		productBusiness: business,
	}
}

func (handle *ProductHandler) AddProduct(c echo.Context) error {
	var newProduct request.Product
	errCreate := c.Bind(*&newProduct)
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

func (handle *ProductHandler) GetAllProduct(c echo.Context) error {
	limit, offset := c.QueryParam("limit"), c.QueryParam("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	res, err := handle.productBusiness.GetProduct(limitInt, offsetInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Faile to get all data product"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to get all data product", response.FromCoreList(res)))
}

func (handler *ProductHandler) GetProductbyID(c echo.Context) error {
	id := c.Param("id")
	idproduct, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("ID not recognized"))
	}

	res, err := handler.productBusiness.GetProductbyID(uint(idproduct))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to get your data product"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success to get your data product", response.FromCoreListGetID(res)))
}

func (handle *ProductHandler) UpdateProduct(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)

	id := c.Param("id")
	iduser, errid := strconv.Atoi(id)
	if errid != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "ID not recognized",
		})
	}

	if idToken != iduser {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Unauthorized",
		})
	}

	data := map[string]interface{}{
		"ID":        c.FormValue("ID"),
		"name":      c.FormValue("name"),
		"price":     c.FormValue("price"),
		"qty":       c.FormValue("qty"),
		"desc":      c.FormValue("desc"),
		"UpdatedAt": time.Now(),
	}

	idproducts, _ := data["ID"].(string)
	res := handle.productBusiness.UpdateProduct(idproducts, data)
	if res != nil {
		return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("Failed to update your data product"))
	}

	return c.JSON(http.StatusOK, helper.ResponseSuccessWithData("Success update your data product", data))
}
