package gorm

import (
	"fmt"
	"github.com/vpaliy/telex/store"
)

func multipleMatch(a store.Arg) string {
	return "%" + fmt.Sprint(a) + "%"
}
