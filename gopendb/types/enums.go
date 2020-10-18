package types

type LogicalOperator int

const (
	Init LogicalOperator = iota
	And
	Or
	Xor
	Not
)

type ComparisonOperator int

const (
	Eq ComparisonOperator = iota
	NullEq
	Neq
	NullNeq
	Lt
	Leq
	Gt
	Geq
	Like
	Between
	In
)

type Order int

const (
	Asc Order = iota
	Desc
)

type QueryType int

const (
	Select QueryType = iota
	Insert
	Update
	Delete
)
