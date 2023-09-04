package gt

import ("strings"
// "fmt"
)

func NoMatchFound(typedData string, APIcall []Artists) []string{
	typedData = strings.ToUpper(typedData)
	typedDataRune := []rune{}
	for _, letter := range typedData {
		typedDataRune = append(typedDataRune, letter)
	}
	var dataToReturn []string
	matchingLetters := 0
	nonMatchingLetters := 0
	for _, oneArtist := range APIcall {
		for i:=0;i<len(oneArtist.Name);i++{
		if i < len(typedDataRune) {
		if rune(strings.ToUpper(oneArtist.Name)[i]) == typedDataRune[i] {
			matchingLetters++
		} else {
			nonMatchingLetters++
		}
		if nonMatchingLetters == 3 {
			break
		}
	}
}
	if matchingLetters+1 >= len(typedData) {
		dataToReturn = append(dataToReturn, oneArtist.Name)
	}
	matchingLetters = 0
	nonMatchingLetters = 0
	for _, oneMember := range oneArtist.Member {
		for i:=0;i<len(oneMember);i++{
			if i < len(typedDataRune) {
			if rune(strings.ToUpper(oneMember)[i]) == typedDataRune[i] {
				matchingLetters++
			} else {
				nonMatchingLetters++
			}
			if nonMatchingLetters == 3 {
				break
			}
		}}
		if matchingLetters+1 >= len(typedData) {
		dataToReturn = append(dataToReturn, oneMember)
	}
	matchingLetters = 0
	nonMatchingLetters = 0
}
		for _, oneLocation := range oneArtist.Locations {
			typedData = strings.ReplaceAll(typedData, "-", ", ")
			typedData = strings.ReplaceAll(typedData, "_", " ")
				for i:=0;i<len(oneLocation);i++{
					if i < len(typedDataRune) {
					if rune(strings.ToUpper(oneLocation)[i]) == typedDataRune[i] {
						matchingLetters++
					} else {
						nonMatchingLetters++
					}
					if nonMatchingLetters == 3 {
						break
					}
				}
			}
				if matchingLetters+1 >= len(typedData) {
					dataToReturn = append(dataToReturn, oneLocation)
				}
				matchingLetters = 0
				nonMatchingLetters = 0
			}
	}
	return dataToReturn
}