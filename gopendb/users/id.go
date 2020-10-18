package users

import (
	"github.com/mazrean/gopendb-generator-proto/gopendb/types"
)

var ID = id{}

type id struct{}

func (id id) column(user) string {
	return "id"
}

type idBoolean struct {
	Column
	types.IntBoolean
}

func (id idBoolean) sql(user) (string, []interface{}, error) {
	query, args, err := id.Sql(id.column(userVal))
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

func (id id) Eq(v int64) Boolean {
	return idBoolean{
		Column: ID,
		IntBoolean: types.IntBoolean{
			ComparisonOperator: types.Eq,
			Values:             []int64{v},
		},
	}
}
