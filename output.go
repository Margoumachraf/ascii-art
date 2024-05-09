package art

import (
	// "fmt"
	"os"
)

func Write(c string, fileclise **string) error {
	file, err := os.Create(**fileclise)
	// fmt.Println(c)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = file.WriteString(c)
    return err
}