package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func App() error {
	var (
		err error
		input string
		mainPass string = ""
	)
	reader := bufio.NewReader(os.Stdin)

	for {
		input, err = reader.ReadString('\n')
		if err != nil{
			break
		}

		input = strings.TrimSpace(input)

		if mainPass != "" {
			
		} else if strings.HasPrefix(input, "login") {
			
		} else if strings.HasPrefix(input, "create") {
			
		} else {
			fmt.Println("you need to login\n please write `login [YOUR_MAINPASS]`")
		}
	}

	return err
}