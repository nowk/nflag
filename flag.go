package nflag

import (
	"flag"
	"github.com/segmentio/go-env"
	"time"
)

// Var looks up the env value first, then defaults to flag
func Var(p interface{}, en, fl string, def interface{}, usage string) {
	var d interface{}

	e, err := env.Get(en)
	if err != nil {
		d = def
	} else {
		d = e
	}

	// var method is determined by the default value (def)
	switch def.(type) {
	case bool:
		flag.BoolVar(p.(*bool), fl, d.(bool), usage)

	case time.Duration:
		flag.DurationVar(p.(*time.Duration), fl, d.(time.Duration), usage)

	case float64:
		flag.Float64Var(p.(*float64), fl, d.(float64), usage)

	case int:
		flag.IntVar(p.(*int), fl, d.(int), usage)

	case int64:
		flag.Int64Var(p.(*int64), fl, d.(int64), usage)

	case string:
		flag.StringVar(p.(*string), fl, d.(string), usage)

	case uint:
		flag.UintVar(p.(*uint), fl, d.(uint), usage)

	case uint64:
		flag.Uint64Var(p.(*uint64), fl, d.(uint64), usage)
	}
}
