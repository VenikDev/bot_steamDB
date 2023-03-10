### ISteamEconomy/GetAssetPrices
Чтобы получить цены на игры из разных регионов через **Steam API**, 
можно использовать метод 
```
"ISteamEconomy/GetAssetPrices"
```
Данный метод позволяет получить цены на игровые предметы в 
различных валютах и регионах. Необходимо передать параметры при 
вызове метода, которые определяют игру и регион, для которого 
нужны цены.

### Пример
Пример вызова метода 
```
"ISteamEconomy/GetAssetPrices"
```
через Steam API для получения цен на игровые предметы в регионе 
Россия (ID региона 5) для игры с идентификатором 570 (Dota 2) 
на языке программирования Golang:

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

const (
    steamAPIEndpoint = "https://api.steampowered.com/ISteamEconomy/GetAssetPrices/v1"
)

type Result struct {
    Assets []struct {
        Name   string `json:"name"`
        Prices []struct {
            Currency int     `json:"currency"`
            Value    float32 `json:"value"`
        } `json:"prices"`
    } `json:"assets"`
}

func GetAssetPrices(gameId, currency int) (*Result, error) {
    params := map[string]interface{}{
        "appid":    gameId,
        "currency": currency,
    }
    queryParams := "?" + encodeQueryParams(params)
    requestUrl := steamAPIEndpoint + queryParams

    req, err := http.NewRequest("GET", requestUrl, nil)
    if err != nil {
        return nil, err
    }

    client := http.DefaultClient
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    var result Result
    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
        return nil, err
    }

    return &result, nil
}

func encodeQueryParams(params map[string]interface{}) string {
    query := ""
    for k, v := range params {
        query += fmt.Sprintf("%s=%v&", k, v)
    }
    query = query[:len(query)-1]
    return query
}

func main() {
    gameId := 570   // Dota 2 app id
    currency := 5  // Region code 5 for Russia

    result, err := GetAssetPrices(gameId, currency)
    if err != nil {
        fmt.Println("Failed to get asset prices: ", err)
        return
    }

    for _, item := range result.Assets {
        fmt.Printf("Item name: %s, Price: %f\n", item.Name, item.Prices[currency-1].Value)
    }
}
```

В данном примере мы используем функцию `GetAssetPrices` 
для получения цен на игровые предметы, передавая ей в качестве 
аргументов идентификатор игры и ID региона. Мы формируем запрос к API, 
передавая полученные параметры и обрабатываем ответ, извлекая цены на 
игровые предметы. Наконец, мы выводим полученные цены на игровые предметы 
в консоли.