package testPackageGoLD

import "fmt"

func runTest(paramA string) {
	if paramA == "" {
		fmt.Println("paramA is empty")
	}
	paramA = "Hello"
	fmt.Println("%s from test ", paramA)
}
