package messages

import (
	"github.com/mazrean/gopendb-generator-proto/gopendb/users"
)

type JoinUsersCondition struct {
	messagesColumn Column
	usersColumn    users.Column
}

func (c *Chain) JoinUsers(conds ...JoinUsersCondition) *users.Chain
