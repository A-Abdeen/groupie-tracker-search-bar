package gt

import (
	// "fmt"
	API "gt/webapp/API"
	"html/template"
	"net/http"
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
	typedData := r.FormValue("search") // data entered in the searchbar
	if typedData == "" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	t, err := template.ParseFiles(HtmlTmpl...)  // parsing of all html files
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	matchingArtists := API.Tosearch(typedData, APIcall) // function used to create an []Artists full of the data matching the search
	if matchingArtists == nil && typedData != "" {
		var examples API.NoMatch
		suggestion := API.NoMatchFound(typedData, APIcall) // if no data is found nomatch is used to find data that is similar
		examples = API.NoMatch{Thereis: false, Suggestion: suggestion}
		w.WriteHeader(http.StatusOK)
		t.ExecuteTemplate(w, "nomatch.html", examples) 
	} else {
		w.WriteHeader(http.StatusOK)
		t.ExecuteTemplate(w, "search.html", matchingArtists)
	}
}
