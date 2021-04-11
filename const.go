package main

import "reflect"

type (
	ExcelObject struct {
		Cells map[string]ExcelCell
	}

	// ExcelCell holding identifiable cells supported for application to parse xlsx file uploaded
	ExcelCell struct {
		Name        string
		Description string
		Type        reflect.Kind
		Required    bool
	}

	shopifyProducts struct {
		Handle       string
		Title        string
		Body         string
		Option1Name  string
		Option1Value string
		Option2Name  string
		Option2Value string
		Image        string
		Price        string
	}

	wooCommerceProducts struct {
		Type                 string
		Published            string
		IsFeatured           string
		Visibility           string
		TaxStatus            string
		InStock              string
		BackOrder            string
		SoldIndividually     string
		AllowCustomerReviews string
		Parent               string
		Name                 string
		RegularPrice         string
		Description          string
		SKU                  string
		Image                string
		Attribute1Name       string
		Attribute1Value      string
		Attribute1Visibility string
		Attribute1Global     string
		Attribute2Name       string
		Attribute2Value      string
		Attribute2Visibility string
		Attribute2Global     string
	}
)

const (
	// HEADERS FOR SHOPIFY PRODUCTS
	FieldHandle       string = "handle"
	FieldTitle        string = "title"
	FieldBody         string = "body"
	FieldOption1Name  string = "option_1_name"
	FieldOption1Value string = "option_1_value"
	FieldOption2Name  string = "option_2_name"
	FieldOption2Value string = "option_2_value"
	FieldImage        string = "image"
	FieldPrice        string = "price"

	// HEADERS FOR WOOCOMMERCE PRODUCTS
	FieldType string = "type"
	Published string = "published"
)
