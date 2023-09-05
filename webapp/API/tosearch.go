package gt

import (
	// handlers "gt/webapp/handlers"
	"strings"
	// "fmt"
)

func Tosearch(typedData string, APIcall []Artists) []Artists {
	typedDataInt := Atoi(typedData)
	typedData = strings.ToUpper(typedData)
	var dataToReturn []Artists
	var toDisplay string
	ifMatching := false // boolean used so that whenever the data matches the artist details are appended
	for i, oneArtist := range APIcall {
		if strings.ToUpper(oneArtist.Name) == typedData {
			toDisplay = "Search matches Name " + oneArtist.Name
			ifMatching = true
		}
		for _, oneMember := range oneArtist.Member {
			if strings.ToUpper(oneMember) == typedData {
				toDisplay = "Search matches "  + oneArtist.Name + " Member name " + oneMember
				ifMatching = true
			}
		}
		if typedDataInt == oneArtist.Creationdate {
			toDisplay = "Search matches " + oneArtist.Name +" Creation Date " + typedData
			ifMatching = true
		}
		firstAlbum := Atoi(oneArtist.FirstAlbum)
		if typedDataInt == firstAlbum {
			toDisplay = "Search matches "  + oneArtist.Name + "\nFirst Album Date " + typedData
			ifMatching = true
		}
		for _, oneLocation := range oneArtist.Locations {
			typedData = strings.ReplaceAll(typedData, "-", ", ")
			typedData = strings.ReplaceAll(typedData, "_", " ")
			if strings.ToUpper(oneLocation) == typedData {
				toDisplay = "Search matches " + oneArtist.Name + " Location " + oneLocation
				ifMatching = true
			}
		}
		for _, oneDate := range oneArtist.Dates {
			oneDateInt := Atoi(oneDate)
			if oneDateInt == typedDataInt {
				toDisplay = "Search matches Date " + oneDate
				ifMatching = true
			}
		}
		if ifMatching {
			ArtistDetails := APIcall[i] // the data is appended according to the artist that matches from the index
			ArtistDetails.SearchResult = toDisplay // data added to the artist struct according to the matching result
			dataToReturn = append(dataToReturn, ArtistDetails) // all the artists that match are appended and returned as an []Artists
		}
		ifMatching = false
	}
	return dataToReturn
}
