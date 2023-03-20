package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	studentNumber, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Oh no, there is something error. The input must be a number")
		return
	}

	if studentNumber < 1 {
		fmt.Println("The input number is invalid")
		return
	}

	if studentNumber > 10 {
		fmt.Println("Sorry, there is only 10 students currently")
		return
	}

	displayStudentData(studentNumber)
}

type Student struct {
	name    string
	address string
	job     string
	reason  string
}

func displayStudentData(studentNumber int) {
	students := []Student{
		{name: "Alex", address: "Jakarta", job: "Java Developer", reason: "Get a new job as a Golang developer"},
		{name: "Bennet", address: "Bandung", job: "Frontend Developer", reason: "Get a new job as a Golang developer"},
		{name: "Cintya", address: "Surabaya", job: "Data Analyst", reason: "Get a new job as a Golang developer"},
		{name: "Dean", address: "Subang", job: "Student", reason: "Get a new job as a Golang developer"},
		{name: "Elijah", address: "Medan", job: "Teacher", reason: "Get a new job as a Golang developer"},
		{name: "Franklin", address: "Pontianak", job: "Student", reason: "Get a new job as a Golang developer"},
		{name: "Gustavo", address: "Malang", job: "C# Developer", reason: "Get a new job as a Golang developer"},
		{name: "Hunter", address: "Pekanbaru", job: "Java Developer", reason: "Get a new job as a Golang developer"},
		{name: "Isaac", address: "Lombok", job: "Quality Assurance", reason: "Get a new job as a Golang developer"},
		{name: "Jackson", address: "Bali", job: "Student", reason: "Get a new job as a Golang developer"},
	}

	fmt.Println()
	fmt.Println("STUDENT DATA")
	fmt.Println("Name:", students[studentNumber-1].name)
	fmt.Println("Adress:", students[studentNumber-1].address)
	fmt.Println("Current Job:", students[studentNumber-1].job)
	fmt.Println("Reason to Study Golang:", students[studentNumber-1].reason)
}
