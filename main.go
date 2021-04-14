package main

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
)

func main() {
	var productsUploaded []shopifyProducts

	wb, _ := xlsx.OpenFile("./mafo.xlsx")
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
}
