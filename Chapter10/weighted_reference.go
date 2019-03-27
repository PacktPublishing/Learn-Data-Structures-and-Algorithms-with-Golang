

//Reference Counter
type ReferenceCounter struct {
	num     *uint32
	pool    *sync.Pool
	removed *uint32
	weight  int
}

//WeightedReference method
func WeightedReference() int {

	var references []ReferenceCounter

	references = GetReferences(root)

	var reference ReferenceCounter

	var sum int
	for _, reference = range references {

		sum = sum + reference.weight

	}

	return sum

}
