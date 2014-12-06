package nflag

import (
	"flag"
	"os"
	"time"
)

// Var type checks the default value to call the proper <T>Var function
func Var(p interface{}, n, fl string, d interface{}, usage string) {
	switch d.(type) {
	case bool:
		BoolVar(p.(*bool), n, fl, d.(bool), usage)

	case time.Duration:
		DurationVar(p.(*time.Duration), n, fl, d.(time.Duration), usage)

	case float64:
		Float64Var(p.(*float64), n, fl, d.(float64), usage)

	case int:
		IntVar(p.(*int), n, fl, d.(int), usage)

	case int64:
		Int64Var(p.(*int64), n, fl, d.(int64), usage)

	case string:
		StringVar(p.(*string), n, fl, d.(string), usage)

	case uint:
		UintVar(p.(*uint), n, fl, d.(uint), usage)

	case uint64:
		Uint64Var(p.(*uint64), n, fl, d.(uint64), usage)
	}
}

// def looks up env and returns that or the default depending on whether it was
// set
func def(n string, d interface{}) interface{} {
	s := os.Getenv(n)
	if s == "" {
		return d
	}

	return s
}

// BoolVar calls flag.BoolVar
func BoolVar(p *bool, n, fl string, d bool, usage string) {
	flag.BoolVar(p, fl, def(n, d).(bool), usage)
}

// DurationVar calls flag.DurationVar
func DurationVar(p *time.Duration, n, fl string, d time.Duration, usage string) {
	flag.DurationVar(p, fl, def(n, d).(time.Duration), usage)
}

// Float64Var calls flag.Float64Var
func Float64Var(p *float64, n, fl string, d float64, usage string) {
	flag.Float64Var(p, fl, def(n, d).(float64), usage)
}

// IntVar calls flag.IntVar
func IntVar(p *int, n, fl string, d int, usage string) {
	flag.IntVar(p, fl, def(n, d).(int), usage)
}

// Int64Var calls flag.Int64Var
func Int64Var(p *int64, n, fl string, d int64, usage string) {
	flag.Int64Var(p, fl, def(n, d).(int64), usage)
}

// StringVar calls flag.StringVar
func StringVar(p *string, n, fl string, d string, usage string) {
	flag.StringVar(p, fl, def(n, d).(string), usage)
}

// UintVar calls flag.UintVar
func UintVar(p *uint, n, fl string, d uint, usage string) {
	flag.UintVar(p, fl, def(n, d).(uint), usage)
}

// Uint64Var calls flag.Uint64Var
func Uint64Var(p *uint64, n, fl string, d uint64, usage string) {
	flag.Uint64Var(p, fl, def(n, d).(uint64), usage)
}

// Bool calls flag.Bool
func Bool(p *bool, n, fl string, d bool, usage string) *bool {
	return flag.Bool(fl, def(n, d).(bool), usage)
}

// Duration calls flag.Duration
func Duration(p *time.Duration, n, fl string, d time.Duration, usage string) *time.Duration {
	return flag.Duration(fl, def(n, d).(time.Duration), usage)
}

// Float64 calls flag.Float64
func Float64(n, fl string, d float64, usage string) *float64 {
	return flag.Float64(fl, def(n, d).(float64), usage)
}

// Int calls flag.Int
func Int(n, fl string, d int, usage string) *int {
	return flag.Int(fl, def(n, d).(int), usage)
}

// Int64 calls flag.Int64
func Int64(n, fl string, d int64, usage string) *int64 {
	return flag.Int64(fl, def(n, d).(int64), usage)
}

// String calls flag.String
func String(n, fl string, d string, usage string) *string {
	return flag.String(fl, def(n, d).(string), usage)
}

// Uint calls flag.Uint
func Uint(n, fl string, d uint, usage string) *uint {
	return flag.Uint(fl, def(n, d).(uint), usage)
}

// Uint64 calls flag.Uint64
func Uint64(n, fl string, d uint64, usage string) *uint64 {
	return flag.Uint64(fl, def(n, d).(uint64), usage)
}
