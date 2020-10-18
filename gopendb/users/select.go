package users

import (
	"fmt"

	"github.com/mazrean/gopendb-generator-proto/gopendb/types"
)

func (c *Chain) Error() []error {
	return c.errs
}

func (c *Chain) Where(condition *Condition) *Chain {
	if condition.err != nil {
		c.errs = append(c.errs, fmt.Errorf("where condition err: %w", condition.err))
		return c
	}

	c.whereCondition = condition

	return c
}

func (c *Chain) Select(columns ...Column) *SelectQuery {
	return &SelectQuery{
		table:     userVal,
		queryType: types.Select,
		columns:   columns,
		chain:     c,
	}
}
