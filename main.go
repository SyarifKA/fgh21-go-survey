package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// func main(){
// 	fmt.Println("Siapa nama Anda? ")
// 	var name string
// 	fmt.Scan(&name)
// 	fmt.Println("Berapa usia Anda? ")
// 	var age int
// 	fmt.Scan(&age)
// 	fmt.Println("Apa jenis kelamin Anda? ")
// 	var gender string
// 	fmt.Scan(&gender)
// 	fmt.Println("Apakah Anda merokok? \n", "1. Ya 2. Tidak")
// 	var isSmoker bool
// 	var smoker string
// 	fmt.Scan(&smoker)
// 	// fmt.Scan(&isSmoker)
// 	if smoker == "1" {
// 		isSmoker = true
// 	}else{
// 		isSmoker = false
// 	}
// 	// fmt.Println("Rokok apa yang pernah Anda coba? ")
// 	// var cigarVariant []string
// 	// fmt.Scanln(&cigarVariant)
// 	fmt.Printf( "Nama Anda : %s\n", name)
// 	fmt.Printf( "Usia Anda : %d\n", age)
// 	fmt.Printf( "Gender Anda : %s\n", gender)
// 	fmt.Printf( "Merokok : %t\n", isSmoker)
// }

type response struct{
	name string
	age string
	gender string
	isSmoker string
	cigarVariant []string
}

func ask(question string)string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s ", question)
	answer, _ := reader.ReadString('\n')
	result := ""
	if(runtime.GOOS == "windows"){ // untuk menjaga agar program dapat dijalankan disemua sistem operasi
		result = strings.TrimSuffix(answer, "\r\n")
	}else{
		result = strings.TrimSuffix(answer, "\n")
	}
	return result
}

func doSurvey()response{
	name := ask("siapa nama anda?")
	age := ask("berapa umur anda?")
	gender := ask("apa jenis kelamin anda?")
	isSmoker := ask("apakah anda merokok?")
	cigarVariant := []string{}
	cont := true
	for cont {
		variant := ask("rokok apa yg pernah anda coba?")
		if variant == "0"{
			cont = false
		}else{
			cigarVariant = append(cigarVariant, variant)
		}
	}
	answer := response{
		name,
		age,
		gender,
		isSmoker,
		cigarVariant,
	}
	return answer
}

func main(){
	
	inputResponse := doSurvey()
	// fmt.Println()
	// fmt.Println("")
	// fmt.Println("")
	resp := ask("1. input lagi\n2. lihat hasil\n3. akhiri program\n")
	if resp == "2"{
		fmt.Println(inputResponse)
	}
}