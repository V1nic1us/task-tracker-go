package function

import (
	"fmt"
	"os"
)

func VerifyFile() []byte {
	content, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("tasks.json does not exist, creating it...")
		e := os.WriteFile("tasks.json", []byte(""), 0666)
		if e != nil {
			panic(e)
		}
		return []byte("")
	}

	return content
}
