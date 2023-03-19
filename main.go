package main

import (
	"github.com/ProgrammingMuffin/ReusableSQL/lexer/lexUsecases"
	"github.com/ProgrammingMuffin/ReusableSQL/parser/usecases"
)

func main() {
	tokenStream := lexUsecases.TokenizeSingleLineQuery("SELECT * FROM something")
	usecases.ParseSingleLineQuery(tokenStream)
}
