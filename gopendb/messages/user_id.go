package messages

import (
	"github.com/mazrean/gopendb-generator-proto/gopendb/types"
)

var UserID = userID{}
type userID struct{}

func (id userID) column(user) string {
	return "id"
}

func (id userID) Eq(v int64) Boolean {
	return userIDBoolean{
		Column: UserID,
		IntBoolean: types.IntBoolean{
			ComparisonOperator: types.Eq,
			Values: []int64{v},
		},
	}
}

type userIDBoolean struct {
	Column
	types.IntBoolean
}

func (id userIDBoolean) sql(user) (string,[]interface{}, error) {
	query, args, err := id.Sql(id.column(userVal))
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}
