package store

type (
	Query struct {
		params []*Param
	}

	Param struct {
		key   string
		value interface{}
	}
)

func CreateQuery(m map[string]interface{}) *Query {
	query := &Query{make([]*Param, len(m))}
	for k, v := range m {
		query.Append(k, v)
	}
	return query
}

func (q *Query) Append(key string, value interface{}) {
	q.params = append(q.params, &Param{key, value})
}
