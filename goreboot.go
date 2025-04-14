package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func login() {
	fmt.Println("checklogin.xml...")
	postURL := "http://192.168.1.1:8000/cgi-bin/cbCheckLogin.xml"

	// Form data
	form := url.Values{}
	// rqUsername=~%7Brvq&rqPasswd=HTGN%2CIpk&rqTimeout=600
	form.Add("rqUsername", "~{rvq")
	form.Add("rqPasswd", "HTGN,Ipk")
	form.Add("rqTimeout", "600")

	// Create a new POST request
	req, err := http.NewRequest("POST", postURL, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make the HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response status:", resp.Status)
	fmt.Println("---------------------------")
	fmt.Println("Response body:", string(body))
	fmt.Println("=======================a")
}

func getsessionkey() (sessionKey string) {
	// Define the URL you want to post to
	fmt.Println("get sessionkey...")
	getURL := "http://192.168.1.1:8000/reboot.asp"
	// Create a new GET request
	resp, err := http.Get(getURL)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	// defer resp.Body.Close()
	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response status:", resp.Status)
	fmt.Println("---------------------------")
	fmt.Println("Response body:", string(body))

	// Define regex to extract sessionKey value
	re := regexp.MustCompile(`var\s+sessionKey\s*=\s*'(\d+)';`)
	matches := re.FindStringSubmatch(string(body))
	sessionKey = ""
	if len(matches) > 1 {
		sessionKey = matches[1]
		fmt.Println("Extracted sessionKey:", sessionKey)
	} else {
		fmt.Println("sessionKey not found")
	}

	fmt.Println("Session Key:", sessionKey)

	fmt.Println("=======================a")
	return sessionKey
}

func reboot(sessionKey string) {
	fmt.Println("reboot...")
	postURL := "http://192.168.1.1:8000/cgi-bin/cbReboot.xml"
	// Create a new POST request
	formreboot := url.Values{}
	// rqUsername=~%7Brvq&rqPasswd=HTGN%2CIpk&rqTimeout=600
	formreboot.Add("sessionKey", sessionKey)
	req, err := http.NewRequest("POST", postURL, strings.NewReader(""))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	// defer resp.Body.Close()
	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response status:", resp.Status)
	fmt.Println("---------------------------")
	fmt.Println("Response body:", string(body))
	fmt.Println("=======================a")
}

func main() {
	login()
	skey := getsessionkey()
	// reboot(skey)
}
