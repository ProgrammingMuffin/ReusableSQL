package models

/*
Lists all the keywords and reserved characters supported by the compiler frontend
*/
const (
	SELECT = "SELECT"
	STAR   = "*"
	FROM   = "FROM"
	COMMA  = ","
	WHERE  = "WHERE"
	EQ     = "="
	IN     = "IN"
)

/*
Lists all the token types
*/
const (
	KEYWORD    = "KEYWORD"
	OPERATOR   = "OPERATOR"
	IDENTIFIER = "IDENTIFIER"
)

type Token struct {
	TokenType string
	Value     string
}

var TokenMap map[string]Token = make(map[string]Token)

var TokenMapInitialized bool = false
