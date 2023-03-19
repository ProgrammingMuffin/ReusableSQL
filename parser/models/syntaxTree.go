package parserModels

type SQLTree interface {
}

type IdentTree struct {
	SQLTree
	Name string
}

type WhereTree struct {
	SQLTree
	Left  SQLTree
	Right SQLTree
	Op    string
}

type SelectTree struct {
	SQLTree
	Fields []IdentTree
	From   SQLTree
	Where  WhereTree
}

const (
	IDENT_TREE  = "IDENT_TREE"
	WHERE_TREE  = "WHERE_TREE"
	SELECT_TREE = "SELECT_TREE"
)
