package sdk

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	products, err := GetProducts(ProductQueryParams{Ecological: true})
	rt := reflect.TypeOf(products)

	assert.Nil(t, err)
	assert.Equal(t, rt.Kind(), reflect.Slice)
}

func TestGetProduct(t *testing.T) {
	product, err := GetProduct("8617101")
	rt := reflect.TypeOf(product)

	assert.Nil(t, err)
	assert.Equal(t, rt.Kind(), reflect.Struct)
}
