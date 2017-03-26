package view

import (
//"html/template"
)

type View struct {
	view interface{}
}

func New() *View {
	return &View{}
}
