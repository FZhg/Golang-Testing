//https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

package hello_world

import "fmt"

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}
