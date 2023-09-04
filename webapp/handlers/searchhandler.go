package gt

import (
	// "fmt"
	API "gt/webapp/API"
	"html/template"
	"net/http"
	// "strings"
	// "strconv"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Pattern (Path)
	if r.URL.Path != "/search/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	typedData := r.FormValue("search")
	typedData = r.URL.Query().Get("search")
	if typedData == "" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	matchingArtists := API.Tosearch(typedData, APIcall)
		if matchingArtists == nil && typedData != "" {
		var examples API.NoMatch
		suggestion := API.NoMatchFound(typedData, APIcall)
		examples = API.NoMatch{Thereis: false, Suggestion: suggestion}
		t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "nomatch.html", examples)
	} else {
		t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "search.html", matchingArtists)
	}
}	

