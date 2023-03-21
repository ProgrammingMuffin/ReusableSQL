package services

import (
	"regexp"
	"strings"

	"github.com/ProgrammingMuffin/ReusableSQL/lexer/models"
)

var initTokenMaps = func() {
	if models.TokenMapInitialized {
		return
	}
	models.TokenMap[models.SELECT] = models.Token{TokenType: models.KEYWORD, Value: models.SELECT}
	models.TokenMap[models.STAR] = models.Token{TokenType: models.OPERATOR, Value: models.STAR}
	models.TokenMap[models.FROM] = models.Token{TokenType: models.KEYWORD, Value: models.FROM}
	models.TokenMap[models.COMMA] = models.Token{TokenType: models.OPERATOR, Value: models.COMMA}
	models.TokenMap[models.WHERE] = models.Token{TokenType: models.KEYWORD, Value: models.WHERE}
	models.TokenMap[models.EQ] = models.Token{TokenType: models.OPERATOR, Value: models.EQ}
	models.TokenMap[models.IN] = models.Token{TokenType: models.KEYWORD, Value: models.IN}
	models.TokenMapInitialized = true
}

func ConvertStringToTokenStream(str string) []models.Token {
	initTokenMaps()
	var tokenStream []models.Token
	re := regexp.MustCompile(`\s+`)
	str = re.ReplaceAllString(str, " ")
	strList := strings.Split(str, " ")
	for _, val := range strList {
		if val[len(val)-1] == ',' {
			tokenStream = append(tokenStream, getToken(val[0:len(val)-1]), getToken(models.COMMA))
		} else {
			tokenStream = append(tokenStream, getToken(val))
		}
	}
	return tokenStream
}

func getToken(val string) models.Token {
	mapVal, exists := models.TokenMap[strings.ToUpper(val)]
	if exists == false {
		return models.Token{TokenType: models.IDENTIFIER, Value: strings.ToUpper(val)}
	} else {
		return mapVal
	}
}
