package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	data := "x < y"

	raw, _ := json.Marshal(data)
	fmt.Println(string(raw))
	//prints: "x \u003c y" <- probably not what you expected

	var b1 bytes.Buffer
	json.NewEncoder(&b1).Encode(data)
	fmt.Println(b1.String())
	//prints: "x \u003c y" <- probably not what you expected

	var b2 bytes.Buffer
	enc := json.NewEncoder(&b2)
	enc.SetEscapeHTML(false)
	enc.Encode(data)
	fmt.Println(b2.String())
	//prints: "x < y" <- looks better
}
