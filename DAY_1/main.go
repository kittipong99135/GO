package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Company struct {
	RegisNumber       string
	Company           string
	RegisteredCapital int
	Status            string
}

func main() {
	var numX int = 0
	fmt.Println("EX1-----")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d,", i)
			numX++
		}
	}
	fmt.Println(numX)
	fmt.Println("----------------------------------------------------------------------")

	//EX2
	fmt.Println("EX2-----")
	xEX2 := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	minMax(xEX2)
	fmt.Println("----------------------------------------------------------------------")

	fmt.Println("EX3, EX3.1-----")
	var count int = 0
	for i := 0; i < 1000; i++ {
		st1 := strconv.Itoa(i)
		for _, val := range st1 {
			if string(val) == "9" {
				count++
			}
		}
	}
	fmt.Println(count)
	fmt.Println("------")
	// EX3.1
	someFunc(10000)
	fmt.Println("----------------------------------------------------------------------")

	// EX4
	fmt.Println("EX4-----")
	x1 := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17, 49, 199, 289, 2, 15, 67, 125, 112}
	for i := 0; i < len(x1); i++ {
		for j := 0; j < len(x1)-i-1; j++ {
			if x1[j] > x1[j+1] {
				t := x1[j+1]
				x1[j+1] = x1[j]
				x1[j] = t
			}
		}
	}
	fmt.Println(x1)
	fmt.Println("----------------------------------------------------------------------")

	fmt.Println("EX5-----")
	peoples := map[string]map[string]string{
		"emp_01": {
			"fname": "Marry",
			"lname": "Jane",
		},
		"emp_02": {
			"fname": "Gwenn",
			"lname": "Steframe",
		},
		"emp_03": {
			"fname": "Gwenn",
			"lname": "Steframe",
		},
	}
	fmt.Println("----------------------------------------------------------------------")

	for peoplesKey, peopleElement := range peoples {
		fmt.Println(peoplesKey)
		for elementKey, elementDetail := range peopleElement {
			fmt.Printf("\t %s : %s\n", elementKey, elementDetail)
		}
	}
	fmt.Println("----------------------------------------------------------------------")

	fmt.Println("EX6-----")
	company := Company{
		"01253356548415",
		"Tan Company",
		7000000000,
		"Operating",
	}
	fmt.Printf("Company :\t %s,\n Regisnumber :\t %s,\n RegisteredCapital :\t %d฿,\n Status :\t %s,\n",
		company.Company, company.RegisNumber, company.RegisteredCapital, company.Status)
	fmt.Println("----------------------------------------------------------------------")

	fmt.Println("EX Spacails-----**")
	for i := 0; i < 6; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------------------------------------------")
}

func minMax(num []int) {
	var max int = math.MinInt
	var min int = math.MaxInt
	for _, val := range num {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
	}
	fmt.Printf("Max : %d && Min : %d \n", max, min)
}

func someFunc(num int) {
	// EX3
	var count int = 0
	for i := 0; i < num; i++ {
		st1 := strconv.Itoa(i)
		for _, val := range st1 {
			if strings.Contains(string(val), "9") {
				count++
			}
		}
	}
	fmt.Println(count)
}
