package testPackageGoLD

import "fmt"

func RunTest(paramA string) {
	if paramA == "" {
		fmt.Println("paramA is empty")
	}
	paramA = "Hello"
	// use paramA in printLn along with some static content like use paramA inside the string
	fmt.Println("paramA is: ", paramA)

}
