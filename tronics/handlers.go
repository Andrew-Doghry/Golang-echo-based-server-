package tronics

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetProductByBody(ctx echo.Context) error {
	var reqProduct mBody
	if err := ctx.Bind(&reqProduct); err != nil {
		return err
	}
	if err := ctx.Validate(reqProduct); err != nil {
		return err
	}
	for _, pro := range products {
		if pro.Type == reqProduct.Product {
			return ctx.JSON(http.StatusOK, pro)
		}
	}
	return ctx.JSON(http.StatusNotFound, nil)
}
func ListProducts(ctx echo.Context) error {
	fmt.Println("inside list products func")
	ctx.Response().Header().Set("Content-Type", "application/json")
	ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return ctx.JSON(http.StatusOK, products)
}
func PostProduct(ctx echo.Context) error {
	ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
	var productSlice pros
	if err := ctx.Bind(&productSlice); err != nil {
		e.Logger.Fatalf("there is and error %v", err)
		return err
	}
	for _, pro := range productSlice.Products {
		products = append(products, pro)
	}
	return ctx.JSON(http.StatusOK, productSlice.Products)
}
func UpdateProduct(ctx echo.Context) error {
	id := ctx.Param("id")
	var prod body
	if err := ctx.Bind(&prod); err != nil {
		return err
	}
	if err := ctx.Validate(prod); err != nil {
		return err
	}

	dId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if dId > len(products) && dId <= 0 {
		return errors.New("the product not found ")
	}
	products[dId-1] = prod
	return ctx.JSON(http.StatusAccepted, products[dId-1])
}
func DeleteProduct(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	id = id - 1
	if err != nil {
		return errors.New("error converting the id ")
	}
	if id > len(products) && id <= 0 && len(products) == 0 {
		return errors.New("out of range")

	}

	// for i := id; i < len(products)-1; i++ {
	// 	products[i] = products[i+1]
	// }
	// products = products[:len(products)-1]
	products = splice(products, id)
	return ctx.JSON(http.StatusOK, products)
}
func (pv ProductValidator) Validate(i interface{}) error {
	return pv.validator.Struct(i)
}
