package main

import (
	"fmt"
	splitCsv "github.com/tolik505/split-csv"
	"strings"

	"github.com/tealeg/xlsx/v3"
)

func main() {
	var productsUploaded []shopifyProducts

	wb, _ := xlsx.OpenFile("./omoo.xlsx")
	wbSheet := wb.Sheets[0]
	fmt.Printf("the max row is %v: ", wbSheet.MaxRow)

	var firstRow *xlsx.Row
	firstRow, err := wbSheet.Row(0)
	if err != nil {
		fmt.Printf("error reading first row of file")
	}

	firstRowColIndex := 0
	headersMap := map[string]int{}
	for {
		cell := firstRow.GetCell(firstRowColIndex)
		if len(cell.Value) == 0 {
			break
		} else {
			headersMap[strings.ToLower(cell.Value)] = firstRowColIndex
			firstRowColIndex++
		}
	}

	otherRowsIndex := 1
	shopifyExcelObject := NewShopifyExcel()
	for {
		var otherRow *xlsx.Row
		product := shopifyProducts{
			Handle:       "",
			Title:        "",
			Body:         "",
			Category:     "",
			Option1Name:  "",
			Option1Value: "",
			Option2Name:  "",
			Option2Value: "",
			Image:        "",
			Price:        "",
		}
		otherRow, err = wbSheet.Row(otherRowsIndex)
		if err != nil || otherRow == nil {
			fmt.Printf("error reading row(%d)", otherRowsIndex)
			break
		}
		for k, v := range headersMap {
			otherRow.GetCell(v)
			product = shopifyExcelObject.ParseShopifyProducts(product, k, otherRow.GetCell(v))
		}

		productsUploaded = append(productsUploaded, product)
		otherRowsIndex++

		if otherRowsIndex == wbSheet.MaxRow {
			break
		}
	}

	fmt.Printf("the length is %v: \n", len(productsUploaded))

	convertToWoocommerceSyntax(productsUploaded)

	// after the conversion, split the csv file into smaller ones.
	fmt.Printf("start calling this function.")
	splitter := splitCsv.New()
	splitter.FileChunkSize = 500000 //in bytes (500KB)
	result, _ := splitter.Split("./finally.csv", "./result")
	fmt.Println(result)
}
