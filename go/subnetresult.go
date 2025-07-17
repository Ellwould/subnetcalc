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
	"github.com/ellwould/csvcell"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Constant for subnetcalc.env absolute path
const subnetCalcEnv string = "/etc/subnetcalc/subnetcalc.env"

// Constant for directory path that contains the files subnetcalc-start.html and subnetcalc-end.html
const dirHTML string = "/etc/subnetcalc/html-css"

// Constant for fileStartHTML file
const fileStartHTML string = "subnetcalc-start.html"

// Constant for fileEndHTML file
const fileEndHTML string = "subnetcalc-end.html"

// Constant for American National Standards Institute (ANSI) reset colour code
const resetColour string = "\033[0m"

// Constant for American National Standards Institute (ANSI) text colour codes
const textBoldWhite string = "\033[1;37m"

// Constant for American National Standards Institute (ANSI) background colour codes
const bgRed string = "\033[41m"

// Clear screen function for GNU/Linux OS's
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// Function to draw box with squares around message, must have a message with characters that total a odd number
func messageBox(bgColour string, messageColour string, message string) {
	topBottomSquare := strings.Repeat(" □", (len(message)/2)+6)
	inbetweenSpace := strings.Repeat(" ", len(message)+8)
	fmt.Println(bgColour + messageColour)
	fmt.Println(topBottomSquare + " ")
	fmt.Println(" □" + inbetweenSpace + "□ ")
	fmt.Println(" □    " + message + "    □ ")
	fmt.Println(" □" + inbetweenSpace + "□ ")
	fmt.Println(topBottomSquare + " ")
	fmt.Print(resetColour)
}

// Function to display message on CLI informing the user the configuration file has a wrong value
func invalidEnv(message string) {
	clearScreen()
	messageBox(bgRed, textBoldWhite, message)
	fmt.Println("")
	os.Exit(0)
}

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
	fmt.Fprint(var1, "<p>&nbsp &nbsp &nbsp &nbsp Total IPv4 Addresses: ", var2+1, " &nbsp &nbsp &nbsp &nbsp</p>")
	fmt.Fprint(var1, "<p>&nbsp &nbsp &nbsp &nbsp Total Usable IPv4 Host Addresses: ", var2-1, " &nbsp &nbsp &nbsp &nbsp</p>")
}

// Function to provide HTML button to home page
func homeButton(var1 http.ResponseWriter, var2 string) {
	fmt.Fprint(var1, "<br>")
	fmt.Fprint(var1, "<br>")
	fmt.Fprint(var1, "<a href=\""+var2+"\" class=\"tableButton\"><h2>Home</h2></a>")
	fmt.Fprint(var1, "<br>")
	fmt.Fprint(var1, "<br>")
}

func main() {

	startHTML := csvcell.FileData(dirHTML, fileStartHTML)
	endHTML := csvcell.FileData(dirHTML, fileEndHTML)

	err := godotenv.Load(subnetCalcEnv)
	if err != nil {
		panic("Error loading subnetcalc.env file")
	}

	envAddress := os.Getenv("address")
	envPort := os.Getenv("subnet_result_port")
	envHomePort := os.Getenv("subnet_home_port")
	envURL := os.Getenv("URL")

	validateEnvIP := validator.New()
	validateEnvIPErr := validateEnvIP.Var(envAddress, "required,ip_addr")

	validateEnvURL := validator.New()

	validateEnvURLErr := validateEnvURL.Var(envURL, "required,url")

	envPortInt, err := strconv.Atoi(envPort)
	if err != nil {
		invalidEnv("Port must be a number in " + subnetCalcEnv)
	}

	if envPortInt <= 0 || envPortInt >= 65536 {
		invalidEnv("Port number in " + subnetCalcEnv + " must be between 1 and 65535")
	} else if envPort == envHomePort {
		invalidEnv("Home web page and result web page port numbers cannot be the same in " + subnetCalcEnv)
	} else if validateEnvIPErr != nil && envAddress != "localhost" {
		invalidEnv("Address in " + subnetCalcEnv + " must be a valid Internet Protocol (IP) address or localhost")
	} else if validateEnvURLErr != nil {
		invalidEnv("URL in " + subnetCalcEnv + " must be a vaild URL")
	} else {

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
			validateCidrNotationErr := validateCidrNotation.Var(cidrNotation, "oneof=30 29 28 27 26 25 24 23 22 21 20 19 18 17 16 15 14 13 12 11 10 9 8")

			//Conditional statment that tests the user input has correct IPv4's and CIDR notation
			if validateIpAddressErr != nil || validateCidrNotationErr != nil {
				fmt.Fprint(w, startHTML)
				fmt.Fprint(w, "&nbsp &nbsp &nbsp &nbsp")
				fmt.Fprint(w, "<table>")
				fmt.Fprint(w, "<tr>")
				fmt.Fprint(w, "<th>")
				fmt.Fprint(w, "<h1>&nbsp &nbsp &nbsp &nbsp Incorrect IPv4 and/or CIDR notation &nbsp &nbsp &nbsp &nbsp</h1>")
				fmt.Fprint(w, "</th>")
				fmt.Fprint(w, "</tr>")
				fmt.Fprint(w, "</table>")
				homeButton(w, envURL)
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
				fmt.Fprint(w, "&nbsp &nbsp &nbsp &nbsp")
				fmt.Fprint(w, "<table class=\"resultTable\">")
				fmt.Fprint(w, "<tr>")
				fmt.Fprint(w, "<th>")
				fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp IPv4 Network ID: 10.0.0.0 &nbsp &nbsp &nbsp &nbsp</p>")
				fmt.Fprint(w, "<br>")
				fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp First Usable IPv4 Host Address: 10.0.0.1 &nbsp &nbsp &nbsp &nbsp</p>")
				if octet3 > 255 && octet4 > 255 {
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", 255, ".", 255-1, " &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp IPv4 Broadcast Address: ", octet1, ".", octet2, ".", 255, ".", 255, " &nbsp &nbsp &nbsp &nbsp</p>")
				} else if octet3 > 255 {
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", 255, ".", octet4-1, " &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp IPv4 Broadcast Address: ", octet1, ".", octet2, ".", 255, ".", octet4, " &nbsp &nbsp &nbsp &nbsp</p>")
				} else if octet4 > 255 {
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", octet3, ".", 255-1, " &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp IPv4 Broadcast Address: ", octet1, ".", octet2, ".", octet3, ".", 255, " &nbsp &nbsp &nbsp &nbsp</p>")
				} else {
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", octet3, ".", octet4-1, " &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp IPv4 Broadcast Address: ", octet1, ".", octet2, ".", octet3, ".", octet4, " &nbsp &nbsp &nbsp &nbsp</p>")
				}
				totalIp(w, cidr)
				homeButton(w, envURL)
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
				fmt.Fprint(w, "&nbsp &nbsp")
				fmt.Fprint(w, "<table class=\"resultTable\">")
				fmt.Fprint(w, "<tr>")
				fmt.Fprint(w, "<th>")
				if octet4 > 65535 {
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp 192.168.0.0/16 can only have &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp CIDR Notation between &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>/16 to /30</p>")
				} else if octet4 < 256 {
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp IPv4 Network ID: 192.168.0.0 &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<br>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp First Usable IPv4 Host Address: 192.168.0.1 &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", 0, ".", octet4-1, " &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp IPv4 Broadcast Address: ", octet1, ".", octet2, ".", 0, ".", octet4, " &nbsp &nbsp &nbsp &nbsp</p>")
					totalIp(w, cidr)
				} else {
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp IPv4 Network ID: 192.168.0.0 &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp First Usable IPv4 Host Address: 192.168.0.1 &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp Last Usable IPv4 Host Address: ", octet1, ".", octet2, ".", octet3, ".", 255-1, " &nbsp &nbsp &nbsp &nbsp</p>")
					fmt.Fprint(w, "<p>&nbsp &nbsp &nbsp &nbsp IPv4 Broadcast Address: ", octet1, ".", octet2, ".", octet3, ".", 255, " &nbsp &nbsp &nbsp &nbsp</p>")
					totalIp(w, cidr)
				}
				homeButton(w, envURL)
				fmt.Fprint(w, "</th>")
				fmt.Fprint(w, "</tr>")
				fmt.Fprint(w, "</table>")
				fmt.Fprint(w, endHTML)
			} else {
				fmt.Fprint(w, startHTML)
				fmt.Fprint(w, "&nbsp &nbsp &nbsp &nbsp")
				fmt.Fprint(w, "<table>")
				fmt.Fprint(w, "<tr>")
				fmt.Fprint(w, "<th>")
				fmt.Fprint(w, "<h1>&nbsp &nbsp &nbsp &nbsp Incorrect IPv4 and/or CIDR Notation &nbsp &nbsp &nbsp &nbsp</h>")
				homeButton(w, envURL)
				fmt.Fprint(w, "</th>")
				fmt.Fprint(w, "</tr>")
				fmt.Fprint(w, "</table>")
				fmt.Fprint(w, endHTML)
			}
		})

		socket := envAddress + ":" + envPort
		fmt.Println("subnet result is running on " + socket)

		//Log error
		if err := http.ListenAndServe(socket, nil); err != nil {
			log.Fatal(err)
		}
	}
}

// Contributor(s):
// Elliot Michael Keavney
