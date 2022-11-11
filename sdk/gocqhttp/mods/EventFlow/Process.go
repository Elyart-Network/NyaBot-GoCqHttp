package EventFlow

import "fmt"

func Process(data interface{}) {
	fmt.Println(string(data.([]byte)))
}
