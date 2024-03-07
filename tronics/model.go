package tronics

import "gopkg.in/go-playground/validator.v9"

type body struct {
	Type string `json="type" validate="required,min=4`
}

type pros struct {
	Products []body `json="products"`
}

var products = []body{
	{Type: "mobile"},
	{Type: "tv"},
	{Type: "laundry machine"},
	{Type: "printer"},
}

type ProductValidator struct {
	validator *validator.Validate
}
