package gogo

/*
All returns true if all elements of the iterable are true (or if the iterable is empty).
*/
func All(iterable []bool) bool {
	for _, element := range iterable {
		if !element {
			return false
		}
	}
	return true
}
