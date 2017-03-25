package sdk

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetProducts(t *testing.T) {
	products, err := GetProducts(ProductQueryParams{})
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
