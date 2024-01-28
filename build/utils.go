package main

import "text/template"

var CustomFns = template.FuncMap{
	"neq": func(x, y interface{}) bool {
		return x != y
	},
	"eq": func(x, y interface{}) bool {
		return x == y
	},
	"sub": func(x, y int) int {
		return x - y
	},
	"add": func(x, y int) int {
		return x + y
	},
}
