# go-shopify

Another Shopify Api Library in Go.

both does not include a license which makes it 

**Warning**: This library is not ready for primetime :-)

[![Build Status](https://travis-ci.org/Receiptful/go-shopify.svg?branch=master)](https://travis-ci.org/Receiptful/go-shopify)

## Install

```console
$ go get github.com/receiptful/go-shopify
```

## Use

```go
import "github.com/receiptful/go-shopify"
```

This gives you access to the `goshopify` package.

### Oauth

If you don't have an access token yet, you can obtain one with the oauth flow.
Something like this will work:

```go
// Create an app somewhere.
app := goshopify.App{
    ApiKey: "abcd",
    ApiSecret: "efgh",
    RedirectUrl: "https://example.com/shopify/callback",
    Scope: "read_products",
}

// Create an oauth-authorize url for and redirect to it.
// In some request handler, you probably want something like this:
func MyHandler(w http.ResponseWriter, r *http.Request) {
    shopName := r.URL.Query().Get("shop")
    authUrl := app.AuthorizeURL(shopName)
    http.Redirect(w, r, authUrl, http.StatusFound)
}

// Fetch a permanent access token in the callback
func MyHandler(w http.ResponseWriter, r *http.Request) {
    if !app.ValidateHmac(r.URL) {
        http.Error(w, "Invalid Signature", http.StatusUnauthorized)
        return
    }

    query := r.URL.Query()
    shopName := query.Get("shop")
    code := query.Get("code")
    token, err := app.GetAccessToken(shopName, code)

    // Do something with the token, like store it in a DB.
}
```

### Api calls

With a permanent access token, you can make API calls like this:

```go
// Create an app somewhere.
app := goshopify.App{
    ApiKey: "abcd",
    ApiSecret: "efgh",
    RedirectUrl: "https://example.com/shopify/callback",
    Scope: "read_products",
}

// Create a new API client
client := goshopify.NewClient(app, "shopname", "token")

// Fetch some data.
products := client.Product.List()
```

Not all endpoints are implemented right now. In those case, feel free to
implement them and make a PR, or you can create your own struct for the data
and use `NewRequest` with the API client. This is how the existing endpoints
are implemented.
