package gt

import (
	"fmt"
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
	matchingArtists := API.Tosearch(typedData, APIcall)
	fmt.Println(matchingArtists)
	if matchingArtists == nil && typedData != "" {
		var examples API.NoMatch
		suggestion := API.NoMatchFound(typedData, APIcall)
		examples = API.NoMatch{Thereis: false, Suggestion: suggestion}
		fmt.Println(examples)
		t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "search.html", examples)
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

