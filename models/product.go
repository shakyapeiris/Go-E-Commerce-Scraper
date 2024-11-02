package models

type Product struct {
	id           string
	name         string
	variantCount string
}

type ProductVariant struct {
	properties map[string]interface{}
	price      float32
}
