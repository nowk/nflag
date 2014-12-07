package nflag

import (
	"flag"
	"os"
	"time"
)

// def looks up env and returns that or the default depending on whether it was
// set
func def(n string, d interface{}) interface{} {
	s := os.Getenv(n)
	if s == "" {
		return d
	}

	return s
}

// BoolVar flag
func BoolVar(p *bool, n, fl string, d bool, usage string) {
	flag.BoolVar(p, fl, def(n, d).(bool), usage)
}

// DurationVar flag
func DurationVar(p *time.Duration, n, fl string, d time.Duration, usage string) {
	flag.DurationVar(p, fl, def(n, d).(time.Duration), usage)
}

// Float64Var flag
func Float64Var(p *float64, n, fl string, d float64, usage string) {
	flag.Float64Var(p, fl, def(n, d).(float64), usage)
}

// IntVar flag
func IntVar(p *int, n, fl string, d int, usage string) {
	flag.IntVar(p, fl, def(n, d).(int), usage)
}

// Int64Var flag
func Int64Var(p *int64, n, fl string, d int64, usage string) {
	flag.Int64Var(p, fl, def(n, d).(int64), usage)
}

// StringVar flag
func StringVar(p *string, n, fl string, d string, usage string) {
	flag.StringVar(p, fl, def(n, d).(string), usage)
}

// UintVar flag
func UintVar(p *uint, n, fl string, d uint, usage string) {
	flag.UintVar(p, fl, def(n, d).(uint), usage)
}

// Uint64Var flag
func Uint64Var(p *uint64, n, fl string, d uint64, usage string) {
	flag.Uint64Var(p, fl, def(n, d).(uint64), usage)
}

// Bool flag
func Bool(n, fl string, d bool, usage string) *bool {
	var v bool
	BoolVar(&v, n, fl, d, usage)

	return &v
}

// Duration flag
func Duration(n, fl string, d time.Duration, usage string) *time.Duration {
	var v time.Duration
	DurationVar(&v, n, fl, d, usage)

	return &v
}

// Float64 flag
func Float64(n, fl string, d float64, usage string) *float64 {
	var v float64
	Float64Var(&v, n, fl, d, usage)

	return &v
}

// Int flag
func Int(n, fl string, d int, usage string) *int {
	var v int
	IntVar(&v, n, fl, d, usage)

	return &v
}

// Int64 flag
func Int64(n, fl string, d int64, usage string) *int64 {
	var v int64
	Int64Var(&v, n, fl, d, usage)

	return &v
}

// String flag
func String(n, fl string, d string, usage string) *string {
	var v string
	StringVar(&v, n, fl, d, usage)

	return &v
}

// Uint flag
func Uint(n, fl string, d uint, usage string) *uint {
	var v uint
	UintVar(&v, n, fl, d, usage)

	return &v
}

// Uint64 flag
func Uint64(n, fl string, d uint64, usage string) *uint64 {
	var v uint64
	Uint64Var(&v, n, fl, d, usage)

	return &v
}
