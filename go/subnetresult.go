/*
MIT License

Copyright (c) 2023 Elliot Michael Keavney

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"math"
	"net/http"
	"subnetcalcresource"
	"strconv"
)

// Function to convert CIDR notation into total IPv4 Addresses and minus one
func cidrFormula(var1 string) int {
	var var2 float64
	var2, _ = strconv.ParseFloat(var1, 0)
	var var3 float64
	var3 = math.Pow(2, (32-var2)) - 1
	return int(var3)
}

// Function to give total IPv4 Addresses
func totalIp(var1 http.ResponseWriter, var2 int) {
	fmt.Fprint(var1, "<br>")
	fmt.Fprint(var1, "<h1>Total IPv4 Addresses: ", var2+1, "</h1>")
	fmt.Fprint(var1, "<h1>Total Usable IPv4 Host Addresses: ", var2-1, "</h1>")
}

// Function to provide HTML button to home page
func homeButton(var1 http.ResponseWriter, var2 string) {
	fmt.Fprint(var1, "<br>")
	fmt.Fprint(var1, "<br>")
	fmt.Fprint(var1, "<a href=\"https://"+var2+"\" class=\"tableButton\"><h2>Home</h2></a>")
}

func main() {

	//Get HTML and CSS from file
	var startHTML string
	startHTML = subnetcalcresource.StartHTML()

	//Get HTML from file
	var endHTML string
	endHTML = subnetcalcresource.EndHTML()

	//Get FQDN from file
	var domainName string
	domainName = subnetcalcresource.FQDN()

	http.HandleFunc("/subnet-result", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}

		//Get IP Address and validate
		f1 := r.FormValue("ip_address")
		var ipAddress string
		ipAddress = f1
		validateIpAddress := validator.New()
		validateIpAddressErr := validateIpAddress.Var(ipAddress, "oneof=10.0.0.0 192.168.0.0")

		//Get CIDR notation and validate
		f2 := r.FormValue("cidr_notation")
		var cidrNotation string
		cidrNotation = f2
		validateCidrNotation := validator.New()
		validateCidrNotationErr := validateCidrNotation.Var(cidrNotation , "oneof=30 29 28 27 26 25 24 23 22 21 20 19 18 17 16 15 14 13 12 11 10 9 8")

		//Conditional statment that tests the user input has correct IPv4's and CIDR notation
		if validateIpAddressErr != nil || validateCidrNotationErr != nil {
			fmt.Fprint(w, startHTML)
			fmt.Fprint(w, "<table>")
			fmt.Fprint(w, "<tr>")
			fmt.Fprint(w, "<th>")
			fmt.Fprint(w, "<h1>Incorrect IPv4 and/or CIDR notation</h1>")
			fmt.Fprint(w, "</th>")
			fmt.Fprint(w, "</tr>")
			fmt.Fprint(w, "</table>"
			homeButton(w, domainName)
			fmt.Fprint(w, endHTML)
		} else if ipAddress == "10.0.0.0" && validateCidrNotationErr == nil {
			var cidr int
			cidr = cidrFormula(cidrNotation)
			const octet1 = int(10)
			var octet2, octet3, octet4 int
			octet2 = cidr / 65536
			octet3 = cidr / 256
			octet4 = cidr / 1
			fmt.Fprint(w, startHTML)
			fmt.Fprint(w, "<table class=\"resultTable\">")
			fmt.Fprint(w, "<tr>")
			fmt.Fprint(w, "<th>")
			fmt.Fprint(w, "<p>IPv4 Network ID: 10.0.0.0</p>")
			fmt.Fprint(w, "<br>")
			fmt.Fprint(w, "<p>First Usable IPv4 Host Address: 10.0.0.1</p>")
			if octet3 > 255 && octet4 > 255 {
				fmt.Fprint(w, "<p>Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", 255, ".", 255-1, "</p>")
				fmt.Fprint(w, "<p>IPv4 Broadcast Address: ", octet1, ".", octet2, ".", 255, ".", 255, "</p>")
			} else if octet3 > 255 {
				fmt.Fprint(w, "<p>Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", 255, ".", octet4-1, "</p>")
				fmt.Fprint(w, "<p>IPv4 Broadcast Address: ", octet1, ".", octet2, ".", 255, ".", octet4, "</p>")
			} else if octet4 > 255 {
				fmt.Fprint(w, "<p>Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", octet3, ".", 255-1, "</p>")
				fmt.Fprint(w, "<p>IPv4 Broadcast Address: ", octet1, ".", octet2, ".", octet3, ".", 255, "</p>")
			} else {
				fmt.Fprint(w, "<p>Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", octet3, ".", octet4-1, "</p>")
				fmt.Fprint(w, "<p>IPv4 Broadcast Address: ", octet1, ".", octet2, ".", octet3, ".", octet4, "</p>")
			}
			totalIp(w, cidr)
			homeButton(w, domainName)
			fmt.Fprint(w, "</th>")
			fmt.Fprint(w, "</tr>")
			fmt.Fprint(w, "</table>")
			fmt.Fprint(w, endHTML)
		} else if ipAddress == "192.168.0.0" && validateCidrNotationErr == nil {
			var cidr int
		        cidr = cidrFormula(cidrNotation)
			const octet1 = int(192)
			const octet2 = int(168)
			var octet3, octet4 int
			octet3 = cidr / 256
			octet4 = cidr / 1
			fmt.Fprint(w, startHTML)
			fmt.Fprint(w, "<table class=\"resultTable\">")
			fmt.Fprint(w, "<tr>")
			fmt.Fprint(w, "<th>")
			if octet4 > 65535 {
				fmt.Fprint(w, "<p>192.168.0.0/16 can only have</p>")
				fmt.Fprint(w, "<p>CIDR Notation between</p>")
				fmt.Fprint(w, "<p>/16 to /30</p>")
			} else if octet4 < 256 {
				fmt.Fprint(w, "<p>IPv4 Network ID: 192.168.0.0</h1")
				fmt.Fprint(w, "<br>")
				fmt.Fprint(w, "<p>First Usable IPv4 Host Address: 192.168.0.1</p>")
				fmt.Fprint(w, "<p>Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", 0, ".", octet4-1, "</p>")
				fmt.Fprint(w, "<p>IPv4 Broadcast Address: ", octet1, ".", octet2, ".", 0, ".", octet4, "</p>")
				totalIp(w, cidr)
			} else {
				fmt.Fprint(w, "<p>IPv4 Network ID: 192.168.0.0</p>")
				fmt.Fprint(w, "<p>First Usable IPv4 Host Address: 192.168.0.1</p>")
				fmt.Fprint(w, "<p>Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", octet3, ".", 255-1, "</p>")
				fmt.Fprint(w, "<p>IPv4 Broadcast Address: ", octet1, ".", octet2, ".", octet3, ".", 255, "</p>")
				totalIp(w, cidr)
			}
			homeButton(w, domainName)
			fmt.Fprint(w, "</th>")
			fmt.Fprint(w, "</tr>")
			fmt.Fprint(w, "</table>")
			fmt.Fprint(w, endHTML)
		} else {
			fmt.Fprint(w, startHTML)
			fmt.Fprint(w, "<table>")
			fmt.Fprint(w, "<tr>")
			fmt.Fprint(w, "<th>")
			fmt.Fprint(w, "<h1>Incorrect IPv4 and/or CIDR Notation")
			homeButton(w, domainName)
			fmt.Fprint(w, "</th>")
			fmt.Fprint(w, "</tr>")
			fmt.Fprint(w, "</table>")
			fmt.Fprint(w, endHTML)
		}
	})

	port := "localhost:8001"
	fmt.Println("subnet result is running on localhost and port " + port)

	//Log error
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// Contributor(s):
// Elliot Michael Keavney
