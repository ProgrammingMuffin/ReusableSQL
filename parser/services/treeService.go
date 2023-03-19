package services

import parserModels "github.com/ProgrammingMuffin/ReusableSQL/parser/models"

func IdentTree(name string) parserModels.SQLTree {
	return parserModels.IdentTree{Name: name}
}

func WhereTree(left parserModels.SQLTree, right parserModels.SQLTree, op string) parserModels.SQLTree {
	return parserModels.WhereTree{Left: left, Right: right, Op: op}
}

func SelectTree(fields []parserModels.IdentTree, from parserModels.SQLTree, where parserModels.WhereTree) parserModels.SQLTree {
	return parserModels.SelectTree{Fields: fields, From: from, Where: where}
}
