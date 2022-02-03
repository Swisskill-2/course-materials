// bhg-scanner/scanner.go modified from Black Hat Go > CH2 > tcp-scanner-final > main.go
// Code : https://github.com/blackhat-go/bhg/blob/c27347f6f9019c8911547d6fc912aa1171e6c362/ch-2/tcp-scanner-final/main.go
// License: {$RepoRoot}/materials/BHG-LICENSE
// Useage: This is a port scanner. The steps for running are: cd workspace/course-materials/materials/lab/2/bhg-scanner/main
//															  go build
//															  time ./main
// You can also test it by: cd workspace/course-materials/materials/lab/2/bhg-scanner/scanner
//							go test
// Will Brant, 2/3/22
// {TODO 1: FILL IN}

package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

//TODO 3 : ADD closed ports; currently code only tracks open ports
var openports []int // notice the capitalization here. access limited!
var closports []int 

func worker(ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second) // TODO 2 : REPLACE THIS WITH DialTimeout (before testing!)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

// for Part 5 - consider
// easy: taking in a variable for the ports to scan (int? slice? ); a target address (string?)?
// med: easy + return  complex data structure(s?) (maps or slices) containing the ports.
// hard: restructuring code - consider modification to class/object
// No matter what you do, modify scanner_test.go to align; note the single test currently fails
//does this change this

//let's take in a parameter for what ports we are scanning for
//add a data structure that tells what each port does
//
func PortScanner() (string, int) {

	ports := make(chan int, 100) // TODO 4: TUNE THIS FOR CODEANYWHERE / LOCAL MACHINE (am i doing this right?)
	results := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		} else {
			closports = append(closports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	
	//TODO 5 : Enhance the output for easier consumption, include closed ports

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	
	}
	return fmt.Sprintf("There are %d ports open and %d ports closed", len(openports), closports), len(openports)
	//return len(openports) 
	//return len(closports)
	
	// TODO 6 : Return total number of ports scanned (number open, number closed);
	//you'll have to modify the function parameter list in the defintion and the values in the scanner_test
}
