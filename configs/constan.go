package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetPathAPI() string {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}
	result := os.Getenv("PATH_API")
	fmt.Println(result)
	return result
}
