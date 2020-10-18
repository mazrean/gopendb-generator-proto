package types

import (
	"errors"
	"fmt"
)

type IntBoolean struct {
	ComparisonOperator
	Values []int64
}

func (b IntBoolean) Sql(column string) (string, []interface{}, error) {
	switch b.ComparisonOperator {
	case Eq:
		return fmt.Sprintf("(%s = ?)", column), []interface{}{b.Values[0]}, nil
	default:
		return "", []interface{}{}, errors.New("invalid comparison operator")
	}
}
