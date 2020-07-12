package main

import (
	"net/http"

	"github.com/Ylazerson/chitchat/data"
)

func junk(writer http.ResponseWriter, request *http.Request) {

	gizmos, err := data.Gizmos()

	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {

		generateHTML(writer, gizmos, "layout", "public.navbar", "junk")

	}
}
