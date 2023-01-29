package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	lambdaNames := []string{
		"user_authentication",
	}
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
