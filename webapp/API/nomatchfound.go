package gt

import "strings"

func NoMatchFound(typedData string, APIcall []Artists) []string{
	if typedData :=
	typedData = strings.ToUpper(typedData)
	typedDataRune := []rune{}
	for i, letter := range typedData {
		typedDataRune[i] = letter
	}
	var dataToReturn []string
	matchingLetters := 0
	nonMatchingLetters := 0
	for _, oneArtist := range APIcall {
		for i:=0;i<len(oneArtist.Name);i++{
		if rune(strings.ToUpper(oneArtist.Name)[i]) == typedDataRune[i] {
			matchingLetters++
		} else {
			nonMatchingLetters++
		}
		if nonMatchingLetters == 3 {
			break
		}
	}
	if matchingLetters+2 >= len(typedData) {
		dataToReturn = append(dataToReturn, oneArtist.Name)
	}
	matchingLetters = 0
	nonMatchingLetters = 0
	for _, oneMember := range oneArtist.Member {
		for i:=0;i<len(oneMember);i++{
			if rune(strings.ToUpper(oneMember)[i]) == typedDataRune[i] {
				matchingLetters++
			} else {
				nonMatchingLetters++
			}
			if nonMatchingLetters == 3 {
				break
			}
		}
		if matchingLetters+2 >= len(typedData) {
		dataToReturn = append(dataToReturn, oneMember)
	}
	matchingLetters = 0
	nonMatchingLetters = 0
}
		for _, oneLocation := range oneArtist.Locations {
			typedData = strings.ReplaceAll(typedData, " ", ", ")
			typedData = strings.ReplaceAll(typedData, "-", ", ")
			typedData = strings.ReplaceAll(typedData, "_", " ")
				for i:=0;i<len(oneLocation);i++{
					if rune(strings.ToUpper(oneLocation)[i]) == typedDataRune[i] {
						matchingLetters++
					} else {
						nonMatchingLetters++
					}
					if nonMatchingLetters == 3 {
						break
					}
				}
				if matchingLetters+2 >= len(typedData) {
					dataToReturn = append(dataToReturn, oneLocation)
				}
				matchingLetters = 0
				nonMatchingLetters = 0
			}
	}
	return dataToReturn
}