package main

// error handling
func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}
