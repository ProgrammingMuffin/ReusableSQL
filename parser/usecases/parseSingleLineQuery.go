package usecases

import (
	"reflect"
	"strings"

	"github.com/ProgrammingMuffin/ReusableSQL/lexer/models"
	parserModels "github.com/ProgrammingMuffin/ReusableSQL/parser/models"
	"github.com/ProgrammingMuffin/ReusableSQL/shared"
)

var tokenIndex int = 0

var tokens []models.Token

func ParseSingleLineQuery(tokenStream []models.Token) parserModels.SQLTree {
	tokens = tokenStream
}

func parseStatement() {
	currentToken := tokens[tokenIndex]
	switch currentToken.TokenType {
	case models.SELECT:

	}
}

func parseSelect() shared.Result[parserModels.SQLTree] {
	fields := parseFields()
	return shared.Result[parserModels.SQLTree]{Value: fields}
}

func parseFields() shared.Result[parserModels.SQLTree] {
	fields := make([]parserModels.IdentTree, 0)
	if tokens[tokenIndex].Value == models.STAR {
		accept(models.TokenMap[models.STAR])
	} else {
		if tokens[tokenIndex].Value == models.COMMA {
			next()
			return parseFields()
		}
		result := parseIdent()
		if result.Errored {
			return result
		}
	}
}

func parseIdent() shared.Result[parserModels.SQLTree] {
	_, contains := models.TokenMap[strings.ToUpper(tokens[tokenIndex].Value)]
	if contains {
		return shared.Result[parserModels.SQLTree]{Value: nil, Err: shared.Error{Message: "keyword used as identifier"}, Errored: true}
	}
	result := shared.Result[parserModels.SQLTree]{Value: parserModels.IdentTree{Name: tokens[tokenIndex].Value}}
	next()
	return result
}

func accept(token models.Token) {
	if reflect.DeepEqual(token, tokens[tokenIndex]) {
		next()
	}
}

func next() {
	tokenIndex++
}
