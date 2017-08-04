package main

import (
	"os"
	"strings"
	"bufio"
	"fmt"
	"strconv"
)

type caseSuite struct {
	credit int
	items int
	priceS []string
	price []int
}

func main(){
	inputFileName, _ := os.Open("A-large-practice.in")
	scanner := bufio.NewScanner(inputFileName)
	scanner.Split(bufio.ScanLines)
	firstLine := true
	var caseNum, evCount,caseCount  int =0,1,0
	var caseNumS,creditS,itemS, priceS string
	var testCase []caseSuite
	for scanner.Scan(){
		if firstLine {
			caseNumS = scanner.Text()
			caseNum, _ = strconv.Atoi(caseNumS)
			firstLine = false
			testCase = make([]caseSuite, caseNum)
		} else{
			switch evCount {
			case 1:
				creditS = scanner.Text()
				evCount++
			case 2:
				itemS = scanner.Text()
				evCount++
			case 3:
				priceS = scanner.Text()
				evCount++
			}
			if evCount==4 {
				evCount = 1
				testCase[caseCount].credit, _ = strconv.Atoi(creditS)
				testCase[caseCount].items, _ = strconv.Atoi(itemS)
				testCase[caseCount].priceS = strings.Fields(priceS)
				testCase[caseCount].price = make([]int,len(testCase[caseCount].priceS))
				for i:=0; i < len(testCase[caseCount].priceS); i++{
					testCase[caseCount].price[i], _ = strconv.Atoi(testCase[caseCount].priceS[i])
				}
				caseCount++
			} 
		}
	}
	/*
	for i:=0; i<caseNum; i++{
		solveSmall(testCase[i], i)
	}
	*/
	for i:=0; i<caseNum; i++{
		solveLarge(testCase[i], i)
	}
}

func solveSmall(data caseSuite, cases int){
	for i:=0; i < data.items; i++ {
		for j:=1; j < data.items-1; j++{
			if data.credit == data.price[i] + data.price[j] {
				if data.price[i] < data.price[j]{
					fmt.Println("Case#", cases+1, ":", data.price[i], data.price[j])
				} else{
					fmt.Println("Case#", cases+1, ":", data.price[j], data.price[i])
				}
				return
			}	
		}
	} 
}

/*
	solveLarge
	by filtering impossible value before actually find solution.
*/

func solveLarge(data caseSuite, cases int){
	filterPrice := make([]int,data.items)
	priceCount := 0
	for _,v := range data.price{
		if v < data.credit {
			filterPrice[priceCount] = v
			priceCount++
		}
	}
	for i:=0; i < priceCount-1; i++ {
		for j:= i+1; j < priceCount; j++{
			if filterPrice[i]+filterPrice[j]==data.credit{
				if filterPrice[i] < filterPrice[j]{
					fmt.Println("Case #", cases+1, ":", filterPrice[i], filterPrice[j])
				} else{
					fmt.Println("Case #", cases+1, ":", filterPrice[j], filterPrice[i])
				}
				return
			}
		}
	}
}
