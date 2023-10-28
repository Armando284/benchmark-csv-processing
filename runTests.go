package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("node", "js/test1.js") // Replace "ls -l" with your desired Bash command

	combined, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(combined))

	tests := [4]string{"node js/test1.js", "python python/test1.py", "php php/test1.php", "go run go/test1.go"}

	for _, test := range tests {
		for i := 0; i < 10; i++ {
			arr := strings.Fields(test)
			fmt.Println("Test: ", arr)
			command := arr[0]
			args := arr[1:]
			cmd := exec.Command(command, args...)

			_, err := cmd.CombinedOutput()
			if err != nil {
				panic(err)
			}

		}
	}

	fmt.Println("All tests done")
}
