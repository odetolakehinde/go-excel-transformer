package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func convertToWoocommerceSyntax(shopifyProducts []shopifyProducts) {
	var wPs []wooCommerceProducts
	var sku string
	productsNo := 0
	lastParentIndex := 0
	attributes1 := ""
	attributes2 := ""

	for i, sP := range shopifyProducts {
		isParent := make(map[int]bool)

		if len(shopifyProducts[i].Title) > 5 {
			isParent[0] = true
			productsNo++
			s := strings.Split(sP.Handle, "-")
			sku = s[len(s)-1]

			// update the last parent index and all
			if lastParentIndex > 0 {
				fmt.Printf("lastParentIndex: %v \n", lastParentIndex)
				wPs[lastParentIndex].Attribute1Value = attributes1
				wPs[lastParentIndex].Attribute2Value = attributes2

				// reinitialize attributes
				attributes1 = ""
				attributes2 = ""
			}
			// assign the new parent index
			lastParentIndex = i
		} else {
			isParent[0] = false
			// start concatenating the attribute values
			if sP.Option1Value != "" {
				var b bytes.Buffer
				if !strings.Contains(attributes1, sP.Option1Value) {
					b.WriteString(attributes1)
					if len(attributes1) != 0 {
						b.WriteString(",")
					}
					b.WriteString(sP.Option1Value)
					attributes1 = b.String()
				}
			}
			if sP.Option2Value != "" {
				var c bytes.Buffer
				if !strings.Contains(attributes2, sP.Option2Value) {
					c.WriteString(attributes2)
					if len(attributes2) != 0 {
						c.WriteString(",")
					}
					c.WriteString(sP.Option2Value)
					attributes2 = c.String()
				}
			}
		}

		wP := wooCommerceProducts{
			Published:            "1",
			IsFeatured:           "0",
			Visibility:           "visible",
			TaxStatus:            "taxable",
			InStock:              "1",
			BackOrder:            "0",
			SoldIndividually:     "0",
			AllowCustomerReviews: "",
			Name:                 sP.Title,
			RegularPrice:         sP.Price,
			Description:          sP.Body,
			Image:                sP.Image,
			Attribute1Name:       "Size",
			Attribute1Visibility: "0",
			Attribute1Global:     "1",
			Attribute2Name:       "Color",
			Attribute2Visibility: "0",
			Attribute2Global:     "1",
		}

		if isParent[0] {
			wP.Type = "variable"
			wP.Parent = ""
			wP.SKU = sku
		} else {
			wP.Type = "variation"
			wP.SKU = ""
			wP.Parent = sku
			wP.Attribute1Value = sP.Option1Value
			wP.Attribute2Value = sP.Option2Value
		}
		wPs = append(wPs, wP)
	}

	createWoocommerceCSV(wPs)
}

func createWoocommerceCSV(products []wooCommerceProducts) {
	f, err := os.Create("products3.csv")
	defer f.Close()

	if err != nil {
		fmt.Printf("failed to open file err: %v", err)
	}

	//var buf bytes.Buffer
	writer := csv.NewWriter(f)

	//csv header
	if writer.Write([]string{"Type", "Published", "Is Featured", "Visibility", "Tax Status", "In Stock", "Backorder", "Sold Individually",
		"Allow Customer Reviews", "Parent", "Name", "Regular Price", "Description", "SKU", "Image", "Attribute 1 Name", "Attribute 1 Value",
		"Attribute 1 Visibility", "Attribute 1 Global", "Attribute 2 Name", "Attribute 2 Value", "Attribute 1 Visibility", "Attribute 1 Global"}) != nil {
		fmt.Printf("Write report to csv error: %v", err)
	}

	for _, row := range products {
		err := writer.Write([]string{
			row.Type,
			row.Published,
			row.IsFeatured,
			row.Visibility,
			row.TaxStatus,
			row.InStock,
			row.BackOrder,
			row.SoldIndividually,
			row.AllowCustomerReviews,
			row.Parent,
			row.Name,
			row.RegularPrice,
			row.Description,
			row.SKU,
			row.Image,
			row.Attribute1Name,
			row.Attribute1Value,
			row.Attribute1Visibility,
			row.Attribute1Global,
			row.Attribute2Name,
			row.Attribute2Value,
			row.Attribute2Visibility,
			row.Attribute2Global,
		})
		if err != nil {
			fmt.Printf("Write report to csv error: %v", err)
		}
	}

	defer writer.Flush()
}
