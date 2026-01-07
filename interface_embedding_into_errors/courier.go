package interface_embedding_into_errors

// implements errorResponseGenerator
type BadCourierAddressError struct {
	errorResponseGenerator
}

func (b *BadCourierAddressError) Error() string {
	return "Bad Courier address"
}
