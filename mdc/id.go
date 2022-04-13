package mdc

var idCount int

func allocID() int {
	defer func() {
		idCount++
	}()

	return idCount
}
