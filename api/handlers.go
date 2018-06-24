package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

func CreaterUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Creater User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}