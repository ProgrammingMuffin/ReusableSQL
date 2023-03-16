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

type Token struct {
	tokenType string
	value     string
}
