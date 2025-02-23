package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main(){
	fileName := flag.String("f", "", "file containing the http request")
	flag.Parse()
	
	if *fileName == "" {
		log.Fatal("Please provide a file name\nUsage: gocurl -f <filename>")
	}
	// Read the file
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	request, err := buildRequest(file)
	if err != nil {
		log.Fatal(err)
	}
	defer request.Body.Close()

	response, err := sendRequest(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	
	status := response.StatusCode
	colorStr := "\033[36m"
	if status >= 400 {
		colorStr = "\033[31m"
	}
	cyan := "\033[36m"
	// Print the response
	fmt.Printf("%s================**HEADER**====================\033[0m\n", cyan)
	fmt.Printf("Response Status: %s%s\033[0m\n",colorStr, response.Status)
	fmt.Println("Response Headers:", response.Header)

	fmt.Printf("%s==================**BODY**====================\033[0m\n", cyan)
	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	fmt.Printf("%s==============================================\033[0m\n", cyan)

}

func sendRequest(request *http.Request) (*http.Response,error) {
	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func buildRequest(file *os.File) (*http.Request, error){

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err	
	}
	fields := strings.Split(line, " ")
	if len(fields) < 2 {
		return nil, errors.New("invalid request")	
	}
	
	method := strings.Trim(fields[0], " \n")
	url := strings.Trim(fields[1], " \n")

	fmt.Println(method, url)
	request, err := http.NewRequest(method, url, nil)
	if err != nil {	
		return nil, err
	}

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		fields := strings.Split(line, ":")
		if len(fields) < 2 {
			break
		}
		request.Header.Add(fields[0], fields[1])
	}
	
	body := ""
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		body += line
	}
	
	request.Body = io.NopCloser(strings.NewReader(body)) 
	return request, nil

}











