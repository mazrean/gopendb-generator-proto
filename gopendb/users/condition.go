package users

import "github.com/mazrean/gopendb-generator-proto/gopendb/types"

type Condition struct {
	conditions []*condition
	err        error
}

func NewCondition(b Boolean) *Condition {
	return &Condition{
		conditions: []*condition{{
			logicalOperator: types.Init,
			Boolean:         b,
		}},
		err: nil,
	}
}

type condition struct {
	logicalOperator types.LogicalOperator
	Boolean
}
