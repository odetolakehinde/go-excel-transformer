package main

import "github.com/tealeg/xlsx/v3"

func NewShopifyExcel() *ExcelObject {
	cells := map[string]ExcelCell{
		FieldHandle: {
			Name:        FieldHandle,
			Description: "Full name",
		},
		FieldTitle: {
			Name:        FieldTitle,
			Description: "Telephone Number",
		},
		FieldBody: {
			Name:        FieldBody,
			Description: "Email address",
		},
		FieldImage: {
			Name:        FieldImage,
			Description: "User password",
		},
		FieldOption1Name: {
			Name:        FieldOption1Name,
			Description: "Region",
		},
		FieldOption1Value: {
			Name:        FieldOption1Value,
			Description: "Designation",
		},
		FieldOption2Name: {
			Name:        FieldOption2Name,
			Description: "Region",
		},
		FieldOption2Value: {
			Name:        FieldOption2Value,
			Description: "Designation",
		},
		FieldPrice: {
			Name:        FieldPrice,
			Description: "Price",
		},
		FieldCategory: {
			Name:        FieldCategory,
			Description: "Category",
		},
	}
	return &ExcelObject{Cells: cells}
}

func (eO *ExcelObject) ParseShopifyProducts(product shopifyProducts, key string, cell *xlsx.Cell) shopifyProducts {
	if cell == nil {
		return product
	}

	switch key {
	case FieldHandle:
		product.Handle = cell.String()
	case FieldTitle:
		product.Title = cell.String()
	case FieldBody:
		product.Body = cell.String()
	case FieldOption1Name:
		product.Option1Name = cell.String()
	case FieldOption1Value:
		product.Option1Value = cell.String()
	case FieldOption2Name:
		product.Option2Name = cell.String()
	case FieldOption2Value:
		product.Option2Value = cell.String()
	case FieldImage:
		product.Image = cell.String()
	case FieldPrice:
		product.Price = cell.String()
	case FieldCategory:
		product.Category = cell.String()
	}
	return product
}
