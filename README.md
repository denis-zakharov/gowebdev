# http package

- `http.Handler` is an interface with the `ServeHTTP` method.
- `http.HanderFunc` is a function *type* that accepts same args as `ServeHTTP`.
Also implements `http.Handler`.

`http.Handle("/path", http.Hanlder)`

`http.HandleFunc("/path", pathHandler)` is a wrapper to convert the function type to
the `Handler` type.


# encoding json

Serialize data into a JSON string all at once.
```
d, err := json.Marshal(products) // deserialization in one chunk into memory
w.Write(d)
```

Serialize data into a JSON stream directly to a Writer object.
```
e := json.NewEncoder(w)
e.Encode(products)
```

# gorilla mux

A convenient wrapper to the net/http server with routing and middleware support.

https://github.com/gorilla/mux

# struct validator

Validation of struct fields with additional annotations: required, conditions, custom
validating functions.

https://github.com/go-playground/validator

# Swagger/OpenAPI

## documentation
Install `go install github.com/go-swagger/go-swagger/cmd/swagger@latest`

Then you can generate the Swagger documentation:
```
swagger generate spec -o ./swagger.yaml --scan-models
```

The documentation should be available at the `/swagger.yaml` or `/docs` paths.

## client autogeneration
