package messages

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mazrean/gopendb-generator-proto/gopendb/types"
	"gopkg.in/guregu/null.v4"
)

type user struct{}

func (user) table() string {
	return "users"
}

var userVal = user(struct{}{})

type order struct {
	columns []Column
	types.Order
}

type Table struct {
	ID        int64
	UserID    int64
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

func (t *Table) table() string {
	return "users"
}

type Nullable struct {
	ID        null.Int
	UserID      null.Int
	Body  null.String
	CreatedAt null.Time
	UpdatedAt null.Time
	DeletedAt null.Time
}

type Chain struct {
	table           user
	db              *sql.DB
	whereCondition  *Condition
	havingCondition *Condition
	orderBy         *order
	groupBy         *order
	limit           int64
	offset          int64
	errs            []error
}

func NewChain(db *sql.DB) *Chain {
	return &Chain{
		table: userVal,
		db:    db,
	}
}

type SelectQuery struct {
	table     user
	queryType types.QueryType
	columns   []Column
	chain     *Chain
}

func (sq *SelectQuery) buildQuery() (string, []interface{}, error) {
	query := ""
	args := []interface{}{}
	switch sq.queryType {
	case types.Select:
		columnQuery := ""
		for i, column := range sq.columns {
			columnQuery += column.column(userVal)
			if i != len(sq.columns)-1 {
				columnQuery += ", "
			}
		}
		if len(sq.columns) == 0 {
			columnQuery = "*"
		}
		query = fmt.Sprintf("SELECT %s FROM %s", columnQuery, sq.table.table())
	default:
		return "", nil, fmt.Errorf("invalid query type(%d)", sq.queryType)
	}

	return query, args, nil
}

func (sq *SelectQuery) Do() ([]*Table, []error) {
	query, args, err := sq.buildQuery()
	if err != nil {
		sq.chain.errs = append(sq.chain.errs, fmt.Errorf("failed to build query: %w", err))
	}
	if len(sq.chain.errs) != 0 {
		return nil, sq.chain.errs
	}

	rows, err := sq.chain.db.Query(query, args...)
	if err != nil {
		return nil, []error{fmt.Errorf("failed to execute query: %w", err)}
	}

	ans := []*Table{}
	for rows.Next() {
		table := Table{}
		dests := []interface{}{}
		for _, column := range sq.columns {
			switch column {
			case UserID:
				dests = append(dests, &table.ID)
			default:
				return nil, []error{fmt.Errorf("invalid column(%s)", column.column(userVal))}
			}
		}
		err := rows.Scan(dests...)
		if err != nil {
			return nil, []error{fmt.Errorf("failed to scan: %w", err)}
		}

		ans = append(ans, &table)
	}

	return ans, nil
}
