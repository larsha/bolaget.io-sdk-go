# bolaget.io-sdk-go


## Usage
```go
import (
  b "bolaget.io-sdk-go"
)
```

## Get products, filter by query params available here https://github.com/larsha/bolaget.io
```go
products, err := b.GetProducts(b.ProductQueryParams{Limit: 1})
```

## Get stores, filter by query params available here https://github.com/larsha/bolaget.io
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
