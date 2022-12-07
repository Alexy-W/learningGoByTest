package iteration

func Repeat(character string, numberOfIteration int) string {
	var finalString string
	for i := 0; i < numberOfIteration; i++ {
		finalString += character
	}
	return finalString
}
