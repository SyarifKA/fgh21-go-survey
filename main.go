package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

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
	if isSmoker != "tidak" && isSmoker != "" {
		cont := true
		for cont {
			variant := ask("rokok apa yg pernah anda coba?")
			if variant == "0"{
				cont = false
			}else{
				cigarVariant = append(cigarVariant, variant)
			}
		}
	}
	answer := response{
			name:         name,
			age:          age,
			gender:       gender,
			isSmoker:     isSmoker,
			cigarVariant: cigarVariant,
	}
	return answer
}

func main() {
	allResponse := []response{}
	writeSurvey := true

	for writeSurvey {
		resp := ask("1. input lagi\n2. lihat hasil\n3. akhiri program\n")
		if resp == "1" {
			inputResponse := doSurvey()
			allResponse = append(allResponse, inputResponse)
		} else if resp == "2" {
			for _, item := range allResponse {
				fmt.Println(item)
			}
			// fmt.Println(allResponse)
		} else if resp == "3" {
			writeSurvey = false
		}
	}
}
