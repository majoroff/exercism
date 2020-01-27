package accumulate

// Accumulate : given a collection and an
// operation to perform on each element of the collection, returns a new
// collection containing the result of applying that operation to each element of
// the input collection.
func Accumulate(collection []string, fn func(string) string) []string {
	var result []string
	for _, s := range collection {
		result = append(result, fn(s))
	}
	return result
}
