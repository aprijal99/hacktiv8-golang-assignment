package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		ticker := time.NewTicker(15 * time.Second)

		for {
			<-ticker.C

			task1()
			task2()
		}
	}()

	wg.Wait()
}

// Task 1
func task1() {
	url := "https://jsonplaceholder.typicode.com/posts"
	client := &http.Client{}

	reqData := map[string]interface{}{
		"userId": genRandomNumber(),
	}

	reqJson, errToJson := json.Marshal(reqData)
	if errToJson != nil {
		panic(errToJson)
	}

	req, errToReq := http.NewRequest("POST", url, bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")
	if errToReq != nil {
		panic(errToReq)
	}

	res, errClientReq := client.Do(req)
	if errClientReq != nil {
		panic(errClientReq)
	}
	defer res.Body.Close()

	resBody, errReadBody := ioutil.ReadAll(res.Body)
	if errReadBody != nil {
		panic(errReadBody)
	}

	fmt.Printf("=== TASK 1 ===\n\n")
	fmt.Printf("REQUEST:\n")
	fmt.Println(req)
	fmt.Printf("\n")
	fmt.Printf("RESPONSE:\n")
	fmt.Println(string(resBody))
	fmt.Printf("\n")
}

// Task 2
func task2() {
	waterRand := genRandomNumber()
	waterStatus := getWaterStatus(waterRand)

	windRand := genRandomNumber()
	windStatus := getWindStatus(windRand)

	waterWind := map[string]interface{}{
		"water": waterRand,
		"wind":  windRand,
	}

	waterWindJson, errWaterWind := json.MarshalIndent(waterWind, "", "  ")
	if errWaterWind != nil {
		panic(errWaterWind)
	}

	fmt.Printf("=== TASK 2 ===\n\n")
	fmt.Println(string(waterWindJson))
	fmt.Println("status water :", waterStatus)
	fmt.Println("status wind :", windStatus)
	fmt.Printf("\n")
}

func genRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100-1+1) + 1
}

func getWaterStatus(waterRand int) string {
	if waterRand < 5 {
		return "aman"
	} else if waterRand >= 6 && waterRand <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func getWindStatus(windRand int) string {
	if windRand < 6 {
		return "aman"
	} else if windRand >= 7 && windRand <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
