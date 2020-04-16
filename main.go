package main

import (
	"html/template"
	"encoding/json"
)

type Story struct {
	Title string
	Story []string
	Options MyOptions
}

type MyOptions struct {
	text string
	arc string
}
