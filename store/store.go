package store

import "time"

type Query interface {
	Selection() map[string]interface{}
	TimeRange() *TimeRange
	Preloads() []string
	Limit() int
	AddPreload(string) Query
	Append(key string, value interface{}) Query
	SetLimit(limit int) Query
	SetTimeRange(time.Time, time.Time) Query
}

type TimeRange struct {
	From time.Time
	To   time.Time
}

type query struct {
	selection map[string]interface{}
	limit     int
	tmrange   *TimeRange
	preloads  []string
}

func NewQuery() Query {
	return &query{
		selection: make(map[string]interface{}),
		limit:     -1,
	}
}

func (q *query) SetTimeRange(oldest, latest time.Time) Query {
	q.tmrange = &TimeRange{oldest, latest}
	return q
}
func (q *query) SetLimit(limit int) Query {
	q.limit = limit
	return q
}

func (q *query) Append(key string, value interface{}) Query {
	q.selection[key] = value
	return q
}

func (q *query) Limit() int {
	return q.limit
}

func (q *query) TimeRange() *TimeRange {
	return q.tmrange
}

func (q *query) AddPreload(preload string) Query {
	q.preloads = append(q.preloads, preload)
	return q
}

func (q *query) Preloads() []string {
	return q.preloads
}

func (q *query) Selection() map[string]interface{} {
	return q.selection
}
