package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func envStr(name string, defaultValue string) string {
	value, isDefined := os.LookupEnv(name)
	if value == "" && !isDefined {
		return defaultValue
	}

	return value
}

func main() {
	publicDir := envStr("PUBLIC_DIR", "public")
	port := fmt.Sprintf(":%v", envStr("PORT", "8000"))

	publicPath := filepath.Join(".", publicDir)
	if _, err := os.Stat(publicPath); os.IsNotExist(err) {
		err := os.Mkdir(publicPath, 0777)
		if err != nil {
			fmt.Printf("Unable to create directory %v\n", publicPath)
		}
	}

	router := gin.Default()
	router.Static("/", publicPath)

	router.Run(port)
}
