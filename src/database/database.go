package database

import (
	"fmt"
	"os"
)

func init() {
	println("Init database")

	// Check if base directory exists
	if info, err := os.Stat("data"); err != nil || !info.IsDir() {
		// Create base directory
		err := os.Mkdir("data", 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
}
