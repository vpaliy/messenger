package store

import "time"

type Options struct {
	from  time.Time
	to    time.Time
	Limit int16
}

type BetweenRange struct {
	To   time.Time
	From time.Time
}

func (o *Options) TimeRange() *BetweenRange {
	if !o.from.IsZero() {
		tr := new(BetweenRange)
		tr.From = o.from
		if !o.to.IsZero() {
			tr.To = o.to
		} else {
			tr.To = time.Now()
		}
		return tr
	}
	return nil
}

type Option func(*Options)

func From(t time.Time) Option {
	return func(args *Options) {
		args.from = t
	}
}

func To(t time.Time) Option {
	return func(args *Options) {
		args.to = t
	}
}

func Limit(limit int16) Option {
	return func(args *Options) {
		if limit <= 0 {
			args.Limit = -1
		} else {
			args.Limit = limit
		}
	}
}

func NewOptions(args ...Option) *Options {
	options := &Options{Limit: -1}
	for _, set := range args {
		set(options)
	}
	return options
}
