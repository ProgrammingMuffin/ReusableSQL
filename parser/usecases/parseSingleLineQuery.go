package usecases

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/ProgrammingMuffin/ReusableSQL/lexer/models"
	parserModels "github.com/ProgrammingMuffin/ReusableSQL/parser/models"
	"github.com/ProgrammingMuffin/ReusableSQL/shared"
)

var tokenIndex int = 0

var tokens []models.Token

func ParseSingleLineQuery(tokenStream []models.Token) {
	tokens = tokenStream
	parseStatement()
}

func parseStatement() {
	currentToken := tokens[tokenIndex]
	switch currentToken.Value {
	case models.SELECT:
		val, _ := json.Marshal(parseSelect())
		fmt.Println(string(val))
		fmt.Println(tokenIndex)
	}
}

func parseSelect() shared.Result[parserModels.SQLTree] {
	accept(models.TokenMap[models.SELECT])
	tree := parserModels.SelectTree{}
	parseFields(tree)
	accept(models.TokenMap[models.FROM])
	res := parseIdent()
	if res.Errored {
		return res
	}
	tree.From = res.Value
	return shared.Result[parserModels.SQLTree]{Value: tree}
}

func parseFields(tree parserModels.SelectTree) shared.Result[parserModels.SQLTree] {
	if tokens[tokenIndex].Value == models.STAR {
		accept(models.TokenMap[models.STAR])
	} else {
		if tokens[tokenIndex].Value == models.COMMA {
			next()
			res := parseFields(tree)
			if res.Errored {
				return res
			}
		} else {
			result := parseIdent()
			if result.Errored {
				return result
			}
			tree.Fields = append(tree.Fields, parserModels.IdentTree{SQLTree: result.Value})
		}
	}
	return shared.Result[parserModels.SQLTree]{Value: tree.Fields}
}

func parseIdent() shared.Result[parserModels.SQLTree] {
	_, contains := models.TokenMap[strings.ToUpper(tokens[tokenIndex].Value)]
	if contains {
		return shared.Result[parserModels.SQLTree]{Value: nil, Err: shared.Error{Message: "keyword used as identifier, pos: "}, Errored: true}
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
