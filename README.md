# Go Sony PlayStation store API client

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/whagency/go-playstation-store/master/LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/whagency/go-playstation-store?v=2)](https://goreportcard.com/report/github.com/whagency/go-playstation-store)
[![GoDoc](https://godoc.org/github.com/whagency/go-playstation-store?status.svg)](https://godoc.org/github.com/whagency/go-playstation-store)

## Install

```
go get -u github.com/whagency/go-playstation-store
```

## Getting started

```go
import (
    ps "github.com/whagency/go-playstation-store/v1"
    psData "github.com/whagency/go-playstation-store/v1/constants"
)

client := ps.NewClient(&ps.Config{
    Region:  psData.UnitedStates,
    Timeout: 5 * time.Second,
    Logging: ps.LoggerStdout,
	PageSize: 50
})
```

###### Available loggers

| Logger        | Destination                  |
|---------------|------------------------------|
| LoggerStdout  | Log to Stdout                |
| LoggerFile    | Log to file ./log/wallet.log |
| LoggerOff     | Logger disable               |

## Examples

###### Get a list of available regions

```go
data := client.GetRegions()

for _, item := range data {
    fmt.Println(item)
}
```

###### Get a list of available games from a category

```go
data, err := client.GetCatalogData(psData.PS5Games, 1)
if err != nil {
    panic(err)
}

for _, item := range data.Products {
    fmt.Println(item.Name, item.NpTitleId)
}
```

*Available categories* : `PS4Games, PS5Games, PSPlus, Sales, EAGames, EAPlayEarlyAccess, VR, VR2, FreeGames, Offers`

###### Get detailed information about the game including the concept by its ID

```go
data, err := client.GetProductByID("UP0001-CUSA09311_00-GAME000000000000")
if err != nil {
    panic(err)
}
	
fmt.Println(data.Name, data.Concept, data.NpTitleId)
```

###### Get detailed information about a concept by its ID

```go
data, err := client.GetConceptByID("229601")
if err != nil {
    panic(err)
}

fmt.Println(data.Name, data.DefaultProduct)
```

###### Get all available product editions with prices for a given concept by its ID

```go
data, err := client.GetConceptPricingByID("229601")
if err != nil {
    panic(err)
}

for _, item := range data.SelectableProducts.PurchasableProducts {
    fmt.Println(item.ID, item.Name, item.NpTitleId, item.Price)
}
```

###### Get all available product addons with prices by its "npTitleId"

```go
data, err := client.GetAddOnsProductsByTitleId("CUSA09311_00")
if err != nil {
    panic(err)
}

for _, item := range data.AddOnProducts {
    fmt.Println(item.ID, item.Name, item.NpTitleId, item.Price)
}
```

###### Get all available features such as subscriptions, etc.

```go
data, err := client.GetFeatures()
if err != nil {
    panic(err)
}

for _, item := range data.Offers {
    fmt.Println(item)
}
```

## Support project

You can support Open Source software developers using the link below:

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/whagency)