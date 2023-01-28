package main

import (
	"fmt"
	"os/exec"
)

func main() {
	lambdaNames := []string{
		"lambdaName",
	}
	for _, lambda := range lambdaNames {
		cmd := exec.Command("./"+lambda, "GOOS=linux GOARCH=amd64 go build -o main main.go")
		res, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
		cmd = exec.Command("./"+lambda, "zip main.zip main")
		res, err = cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(res))
	}

	fmt.Println("lambdas created successfully")

}
