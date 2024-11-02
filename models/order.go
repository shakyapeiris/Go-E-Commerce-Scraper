package models

type Order struct {
	id              string
	user            string
	shippingDetails map[string]string
}
