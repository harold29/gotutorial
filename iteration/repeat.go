package iteration

func Repeat(character string, iternumber int) string {
	if iternumber < 0 {
		iternumber = 5
	}
	var repeated string
	for i := 0; i < iternumber; i++ {
		repeated += character
	}
	return repeated
}
