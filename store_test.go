package sdk

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStores(t *testing.T) {
	stores, err := GetStores(StoreQueryParams{Sort: "city:asc"})
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
