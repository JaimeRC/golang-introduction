package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	readFile()
	readFile2()
	writeFile()
	writeFile2()
}

func readFile(){
	fileTest, err := ioutil.ReadFile("./file_test.txt")
	if err != nil {
		fmt.Println("Error readFile")
	}else {
		fmt.Println(string(fileTest))
	}
}

func readFile2(){
	fileTest, err := os.Open("./fileManager/file_test.txt")

	if err != nil {
		fmt.Println("Error readFile2")
	}else {
		scanner := bufio.NewScanner(fileTest)
		for scanner.Scan(){
			registro := scanner.Text()
			fmt.Println("Line > " + registro + "\n")
		}
	}
	fileTest.Close()
}

func writeFile(){
	fileTest, err := os.Create("./fileManager/new_file_test.txt")

	if err != nil {
		fmt.Println("Error writeFile")
		return
	}

	fmt.Fprintln(fileTest,"This is a new line")
	fileTest.Close()
}

func writeFile2(){
	fileName := "./fileManager/new_file_test.txt"

	if Append(fileName, "\nThis is a second line.") == false {
		fmt.Println("Error in the second line. writeFile2")
		return
	}


}

func Append(file string, text string) bool{
	fileTest, err := os.OpenFile(file,os.O_WRONLY | os.O_APPEND,0644)

	if err != nil {
		fmt.Println("Error Append")
		return false
	}
	_, err = fileTest.WriteString(text)

	if err != nil {
		fmt.Println("Error2 Append")
		return false
	}

	return true
}