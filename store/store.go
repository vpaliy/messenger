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

func NewQuery(m map[string]interface{}) Query {
	query := Query{make([]*Param, len(m))}
	for k, v := range m {
		query.Append(k, v)
	}
	return query
}

func (q *Query) Append(key string, value interface{}) {
	q.params = append(q.params, &Param{key, value})
}

func (q *Query) ToMap() map[string]interface{} {
	mapping := make(map[string]interface{})
	for _, param := range q.params {
		mapping[param.key] = param.value
	}
	return mapping
}
