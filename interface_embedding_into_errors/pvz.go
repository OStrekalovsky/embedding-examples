package interface_embedding_into_errors

// implements errorResponseGenerator
type BadPVZAddressError struct {
	*BadAddressResponseGenerator
}

func (b *BadPVZAddressError) Error() string {
	return "Bad PVZ address"
}
