package main

import (
	"DataAnalysis/analysis"
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Starting a test")
	file := "C:\\code\\go\\src\\github.com\\DataAnalysis\\data\\Ethereum_24h_ActiveAccount.xlsx"
	sheet := "ActiveAccounts"
	aes := analysis.GetAnalysisData(file, sheet)
	sort.Sort(aes)
	//for i, ae := range aes {
	//	fmt.Println("row",i, "ae.address", ae.Address, "ae.count",ae.Weight)
	//}
	h := analysis.GetSelfSimilarH(aes)
	fmt.Println("h",h)
}
