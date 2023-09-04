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
	var data []string
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
	if matchingLetters+2 >= len(typedData) && len(typedData)>4 {
		data = append(data, oneArtist.Name)
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
		if matchingLetters+2 >= len(typedData) && len(typedData)>4{
		data = append(data, oneMember)
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
				if matchingLetters+2 >= len(typedData)&& len(typedData)>4 {
					data = append(data, oneLocation)
				}
				matchingLetters = 0
				nonMatchingLetters = 0
			}
	}
if data != nil{
for i:=0;i<len(data);i++{
	count := 0	
	for j:= len(data)-1;j>i;j--{
	if data[i] == data[j] {
		count++
		break
	}	
}
if count == 0 {
dataToReturn = append(dataToReturn, data[i])}
}	
}
return dataToReturn
}