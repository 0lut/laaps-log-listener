// Go offers built-in support for JSON encoding and
// decoding, including to and from built-in and custom
// data types.

package main

import "fmt"

func main() {
	msg := "Hello ugur ocu"
	s := encrypt([]byte(msg), "kral")
	fmt.Println(string(s))
	sDecrypted := decrypt(s, "kral")
	fmt.Println(string(sDecrypted))

}
