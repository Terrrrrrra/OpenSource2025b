package greeting

import "fmt"

//func hello() 함수 소문자면 외부노출안됨
func Hello() {
	fmt.Printf("Hello!\n")
}

func Hi() {
	fmt.Printf("Hi~\n")
}
