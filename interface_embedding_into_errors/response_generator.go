package interface_embedding_into_errors

type errorResponseGenerator interface {
	GetMessage() string
}

type BadAddressResponseGenerator struct {
	msg string
}

func (c *BadAddressResponseGenerator) GetMessage() string {
	return `{"bad address":"` + c.msg + `"}`

}
