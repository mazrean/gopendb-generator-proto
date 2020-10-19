package messages

type Column interface {
	column(user) string
}

type Boolean interface {
	sql(user) (query string, args []interface{}, err error)
}
