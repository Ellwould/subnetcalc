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
	"log"
	"net/http"
	"subnetcalcresource"
)

func main() {

	var startHTML string
	startHTML = subnetcalcresource.StartHTML()

	var endHTML string
	endHTML = subnetcalcresource.EndHTML()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprint(w, startHTML)
		fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "&nbsp &nbsp &nbsp")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th><h2>&nbsp &nbsp RFC 1918 Subnet Calculater &nbsp &nbsp</h2></th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th><h3><a href=\"https://datatracker.ietf.org/doc/html/rfc1918\" class=\"externalButton tableButton>IETF RFC 1918 Document</a></h3></th>")
		fmt.Fprintf(w, "    <th><h3><a href=\"https://ell.today\" class=\"externalButton tableButton\">Written by Elliot Keavney (Website)</a></h3></th>")
		fmt.Fprintf(w, "    <th><h3><a href=\"https://github.com/Ellwould/subnetcalc\" class=\"externalButton tableButton\">Subnetcalc Source Code (GitHub)</a></h3></th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<hr class=\"roundedbar\">")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td><h2>&nbsp &nbsp 10.0.0.0/8 &nbsp &nbsp</h2>")
		fmt.Fprintf(w, "    <h3>&nbsp &nbsp (10.0.0.0 to 10.255.255.255) &nbsp &nbsp</h3>")
		fmt.Fprintf(w, "    <br>")
		fmt.Fprintf(w, "    <h2>&nbsp &nbsp 192.168.0.0/16 &nbsp &nbsp</h2>")
		fmt.Fprintf(w, "    <h3>&nbsp &nbsp (192.168.0.0 to 192.168.255.255) &nbsp &nbsp<h3></td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
                fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<div class=ipbox>")
		fmt.Fprintf(w, "<form method=\"POST\" action=\"/subnet-result\">")
		fmt.Fprintf(w, "  <label for=\"ip_address\"><b>IP Address:</b>")
		fmt.Fprintf(w, "  </label>")
                fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "  <select id=\"ip_address\" name=\"ip_address\">")
		fmt.Fprintf(w, "    <option value=\"10.0.0.0\">10.0.0.0</option>")
		fmt.Fprintf(w, "    <option value=\"192.168.0.0\">192.168.0.0</option>")
		fmt.Fprintf(w, "  </select>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "  <label for=\"cidr_notation\"><b>CIDR Notation:</b>")
		fmt.Fprintf(w, "  </label>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "  <select id=\"cidr_notation\" name=\"cidr_notation\">")
		fmt.Fprintf(w, "    <option value=\"30\">/30 (2 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"29\">/29 (6 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"28\">/28 (14 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"27\">/27 (30 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"26\">/26 (62 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"25\">/25 (126 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"24\">/24 (254 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"23\">/23 (510 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"22\">/22 (1,022 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"21\">/21 (2,046 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"20\">/20 (4,094 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"19\">/19 (8,190 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"18\">/18 (16,382 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"17\">/17 (32,766 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"16\">/16 (65,534 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"15\">/15 (131,070 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"14\">/14 (262,142 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"13\">/13 (524,286 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"12\">/12 (1,048,574 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"11\">/11 (2,097,150 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"10\">/10 (4,194,302 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"9\">/9 (8,388,606 Hosts)</option>")
		fmt.Fprintf(w, "    <option value=\"8\">/8 (16,777,214 Hosts)</option>")
		fmt.Fprintf(w, "  </select>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "  <input type=\"submit\" value=\"submit\" />")
		fmt.Fprintf(w, "</form>")
		fmt.Fprintf(w, "</div>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<hr class=\"roundedbar\">")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<table>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <th>1st Octet</th>")
		fmt.Fprintf(w, "    <th>2nd Octet</th>")
		fmt.Fprintf(w, "    <th>3rd Octet</th>")
		fmt.Fprintf(w, "    <th>4th Octet</th>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>256^3 = 16777216</td>")
		fmt.Fprintf(w, "    <td>256^2 = 65536</td>")
		fmt.Fprintf(w, "    <td>256^1 = 256</td>")
		fmt.Fprintf(w, "    <td>256^0 = 1</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "  <tr>")
		fmt.Fprintf(w, "    <td>/0 (4294967296)")
		fmt.Fprintf(w, "    <br>/1 (2147483648)")
		fmt.Fprintf(w, "    <br>/2 (1073741824)")
		fmt.Fprintf(w, "    <br>/3 (536870912) ")
		fmt.Fprintf(w, "    <br>/4 (268435456)")
		fmt.Fprintf(w, "    <br>/5 (134217728)")
		fmt.Fprintf(w, "    <br>/6 (67108864)")
		fmt.Fprintf(w, "    <br>/7 (33554432)")
		fmt.Fprintf(w, "    <br><br></td>")
		fmt.Fprintf(w, "    <td>/8 (16777216)")
		fmt.Fprintf(w, "    <br>/9 (8388608)")
		fmt.Fprintf(w, "    <br>/10 (4194304)")
		fmt.Fprintf(w, "    <br>/11 (2097152)")
		fmt.Fprintf(w, "    <br>/12 (1048576)")
		fmt.Fprintf(w, "    <br>/13 (524288)")
		fmt.Fprintf(w, "    <br>/14 (262144)")
		fmt.Fprintf(w, "    <br>/15 (131072)")
		fmt.Fprintf(w, "    <br><br></td>")
		fmt.Fprintf(w, "    <td>/16 (65536)")
		fmt.Fprintf(w, "    <br>/17 (32768)")
		fmt.Fprintf(w, "    <br>/18 (16384)")
		fmt.Fprintf(w, "    <br>/19 (8192)")
		fmt.Fprintf(w, "    <br>/20 (4096)")
		fmt.Fprintf(w, "    <br>/21 (2048)")
		fmt.Fprintf(w, "    <br>/22 (1024)")
		fmt.Fprintf(w, "    <br>/23 (512)")
		fmt.Fprintf(w, "    <br><br></td>")
		fmt.Fprintf(w, "    <td>/24 (256)")
		fmt.Fprintf(w, "    <br>/25 (128)")
		fmt.Fprintf(w, "    <br>/26 (64)")
		fmt.Fprintf(w, "    <br>/27 (32)")
		fmt.Fprintf(w, "    <br>/28 (16)")
		fmt.Fprintf(w, "    <br>/29 (8)")
		fmt.Fprintf(w, "    <br>/30 (4)")
		fmt.Fprintf(w, "    <br>/31 (2)")
		fmt.Fprintf(w, "    <br>/32 (1)</td>")
		fmt.Fprintf(w, "  </tr>")
		fmt.Fprintf(w, "</table>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprintf(w, "<br>")
		fmt.Fprint(w, endHTML)
	})
	port := "localhost:8000"
	fmt.Println("subnethome is running on localhost and port " + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}

// Contributor(s):
// Elliot Michael Keavney
