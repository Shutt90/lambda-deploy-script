package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	lambdaNames := []string{}
	lambdaNames = recursiveAdd(lambdaNames)

	for _, lambda := range lambdaNames {
		cmd := exec.Command("go", "build", "-o", "main", "main.go")
		cmd.Dir = "./" + lambda
		_, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		cmd = exec.Command("zip", "main.zip", "main")
		cmd.Dir = "./" + lambda
		_, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}

		e := os.Remove("./" + lambda + "/main")
		if e != nil {
			log.Fatal(e)
		}
	}

	fmt.Println("lambdas created successfully")
}

func recursiveAdd(arr []string) []string {
	var lambdaName string
	fmt.Printf("Name of lambda: \n(Type exit to finish)\n")
	fmt.Scan(&lambdaName)

	if lambdaName == "exit" {
		return arr
	}

	arr = append(arr, lambdaName)

	return recursiveAdd(arr)
}
