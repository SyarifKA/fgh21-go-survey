package main

import "fmt"

func main(){
	fmt.Println("Siapa nama Anda? ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Berapa usia Anda? ")
	var age int
	fmt.Scan(&age)
	fmt.Println("Apa jenis kelamin Anda? ")
	var gender string
	fmt.Scan(&gender)
	fmt.Println("Apakah Anda merokok? \n", "1. Ya 2. Tidak")
	var isSmoker bool
	var smoker string
	fmt.Scan(&smoker)
	// fmt.Scan(&isSmoker)
	if smoker == "1" {
		isSmoker = true
	}else{
		isSmoker = false
	}
	// var smoker string
	// fmt.Println("Rokok apa yang pernah Anda coba? ")
	// var cigarVariant []string
	// fmt.Scanln(&cigarVariant)
	fmt.Printf( "Nama Anda : %s\n", name)
	fmt.Printf( "Usia Anda : %d\n", age)
	fmt.Printf( "Gender Anda : %s\n", gender)
	fmt.Printf( "Merokok : %t\n", isSmoker)
}