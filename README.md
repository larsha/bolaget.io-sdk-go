[![Build Status](https://travis-ci.org/larsha/bolaget.io-sdk-go.svg?branch=master)](https://travis-ci.org/larsha/bolaget.io-sdk-go)

# bolaget.io-sdk-go

## Versions
This SDK adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html). And it will follow [bolaget.io](https://github.com/larsha/bolaget.io#versions) versions.

| Version       | Linked to             |
| ------------- |:---------------------:|
| `^1.0.0`      | https://bolaget.io/v1 |

## Import
```go
import (
  b "github.com/larsha/bolaget.io-sdk-go"
)
```

## Get products
Query params available here [bolaget.io](https://github.com/larsha/bolaget.io)
```go
products, err := b.GetProducts(b.ProductQueryParams{Limit: 1})
```

## Get stores
Query params available here [bolaget.io](https://github.com/larsha/bolaget.io)
```go
stores, err := b.GetStores(b.StoreQueryParams{Limit: 1})
```

## Get product
```go
product, err := b.GetProduct("123")
```

## Get store
```go
store, err := b.GetStore("123")
```
