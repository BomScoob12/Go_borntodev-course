package main

import "fmt"

var products = make(map[string]float64)

func main() {
	fmt.Println("Products : ", products)

	//add data
	// var-name["KEY"] = value
	products["IPhone"] = 32000
	products["Mackbook"] = 40000
	products["Console Game"] = 20000
	fmt.Println("Products : ", products)

	//delete data
	delete(products, "Console Game")
	fmt.Println("Products after delete : ", products)

	//update data
	products["Iphone"] = 2000
	products["IPhone"] = 10000
	//case sensitive
	//if it not exist in the map it will be adding the new one.
	fmt.Println("Products after update : ", products)

	iphonePrice := products["Iphone"]
	ipadPrice := products["ipad"]
	fmt.Println("IphonePrice : ", iphonePrice)
	fmt.Println("Ipad price :", ipadPrice) //default float = 0

	//default value (it sort in asc key)
	courseName := map[string]string{"int101":"Java fund", "gen210":"How to be a good person.", "aot100":"How to fly"}
	fmt.Println("Course name : ", courseName)
}
