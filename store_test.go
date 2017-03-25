package sdk

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetStores(t *testing.T) {
	stores, err := GetStores(StoreQueryParams{})
	rt := reflect.TypeOf(stores)

	assert.Nil(t, err)
	assert.Equal(t, rt.Kind(), reflect.Slice)
}

func TestGetStore(t *testing.T) {
	product, err := GetStore("0146")
	rt := reflect.TypeOf(product)

	assert.Nil(t, err)
	assert.Equal(t, rt.Kind(), reflect.Struct)
}
