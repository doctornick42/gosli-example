package types

func (r *A) equal(another *A) bool {
	if r == nil && another == nil {
		return true
	}
	if (r == nil && another != nil) || (r != nil && another == nil) {
		return false
	}
	// `equal` method has to be implemented manually

	return r.FieldInt == another.FieldInt
}
