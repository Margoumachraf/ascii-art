package art

import "os"

func Write(c string, fileclise **string) {
	file, _ := os.Create(**fileclise)
	defer file.Close()
	file.Write([]byte(c))
}