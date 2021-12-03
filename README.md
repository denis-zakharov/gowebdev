# http package

- `http.Handler` is an interface with the `ServeHTTP` method.
- `http.HanderFunc` is a function *type* that accepts same args as `ServeHTTP`.
Also implements `http.Handler`.

`http.Handle("/path", http.Hanlder)`

`http.HandleFunc("/path", pathHandler)` is a wrapper to convert the function type to
the `Handler` type.