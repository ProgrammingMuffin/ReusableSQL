package main

import (
	"fmt"

	"github.com/ProgrammingMuffin/ReusableSQL/lexer/usecases"
)

func main() {
	for _, val := range usecases.TokenizeSingleLineQuery("SELECT * FROM something WHERE something = something") {
		fmt.Println(val)
	}
	return
}
