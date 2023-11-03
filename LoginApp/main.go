package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	var username, password string
	selectRole := 0
	isLogin := false

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Please select login role\n------------------")
	fmt.Println("0 - Admin")
	fmt.Println("1 - Student")
	fmt.Scanln(&selectRole)
	fmt.Println("\n------------------")
	for i := 5; i > 0; i-- {
		currentTime := time.Now()

		fmt.Println("Enter Username: ")
		fmt.Scanln(&username)
		fmt.Println("Enter Password: ")
		fmt.Scanln(&password)

		if selectRole == 0 && username == "admin" && password == "admin" {
			io.WriteString(file, "Username: "+username+"\nPassword: "+password+"\nStatus: \033[32mSUCCESS\033[0m \nTime: "+currentTime.Format("01-02-2006 15:04:05")+"\n\n")
			fmt.Println("\033[32mLogin Successful\033[0m \n------------------")
			isLogin = true
			break
		} else if selectRole == 1 && username == "student" && password == "student" {
			io.WriteString(file, "Username: "+username+"\nPassword: "+password+"\nStatus: \033[32mSUCCESS\033[0m \nTime: "+currentTime.Format("01-02-2006 15:04:05")+"\n\n")
			fmt.Println("\033[32mLogin Successful\033[0m \n------------------")
			isLogin = true
			break
		} else {
			io.WriteString(file, "Username: "+username+"\nPassword: "+password+"\nStatus: \033[31mFAILURE\033[0m \nTime: "+currentTime.Format("01-02-2006 15:04:05")+"\n\n")
			fmt.Println("Login Failed. Left Attempts: ", i-1)

		}
	}
	adminSelect := 0

	if isLogin && username == "admin" {
		fmt.Println("Welcome Admin\n------------------")
		fmt.Println("1. Read Logs")
		fmt.Println("2. Log out")
		fmt.Scan(&adminSelect)
		if adminSelect == 1 {
			fmt.Println("Reading Logs... \n------------------")
			readFile, err := os.Open("logs.txt")
			if err != nil {
				fmt.Println(err)
			}
			fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)
			var fileLines []string

			for fileScanner.Scan() {
				fileLines = append(fileLines, fileScanner.Text())
			}

			readFile.Close()

			for _, line := range fileLines {
				fmt.Println(line)
			}

			//fmt.Println(fileLines)
		} else if adminSelect == 2 {
			fmt.Println("Logging out...")
		} else {
			fmt.Println("Invalid Selection")
		}
	} else if isLogin && username == "student" {
		fmt.Println("Welcome student\n------------------")
	} else {
		fmt.Println("Login Failed")
	}

}
