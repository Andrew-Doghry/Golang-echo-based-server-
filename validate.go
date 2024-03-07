// package main

// import (
// 	"net/http"
// 	"os"

// 	"github.com/labstack/echo"
// 	"gopkg.in/go-playground/validator.v9"
// )

// type UserValidator struct {
// 	validator *validator.Validate
// }

// func (u *UserValidator) Validate(i interface{}) error {
// 	return u.validator.Struct(i)
// }
// func main() {
// 	e := echo.New()
// 	port := os.Getenv("my-app-key")
// 	if port == "" {
// 		port = "8800"
// 	}
// 	v := validator.New()
// 	e.Validator = &UserValidator{validator: v}
// 	type Body struct {
// 		Product    string `json:"product" validate:"required,min=4"`
// 		Vendor     string `json:"vendor" validate:"required,min=4,max=10"`
// 		Email      string `json:"email" validate:"required_with=Vendor,email"`
// 		Url        string `json:"url" validate:"required,url"`
// 		Default_id string `json:"default_device_ip" validate:"ip"`
// 	}
// 	e.GET("/product", func(ctx echo.Context) error {
// 		var prod Body
// 		if err := ctx.Bind(&prod); err != nil {
// 			return err
// 		}
// 		if err := ctx.Validate(prod); err != nil {

// 			return err
// 		}
// 		return ctx.JSON(http.StatusOK, prod)
// 	})

// 	e.Start(":" + port)
// }

// // echo validator a pointer to struct implements the validte func
package main
