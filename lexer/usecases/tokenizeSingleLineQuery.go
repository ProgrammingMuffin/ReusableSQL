package usecases

import (
	"github.com/ProgrammingMuffin/ReusableSQL/lexer/models"
	"github.com/ProgrammingMuffin/ReusableSQL/lexer/services"
)

func TokenizeSingleLineQuery(request string) []models.Token {
	return services.ConvertStringToTokenStream(request)
}
