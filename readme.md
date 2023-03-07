## What is a Handler?
A Handler is an interface in Go that is responsible for processing an http request via ServeHttp
```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

```
Handlers can be both Middleware and Application Handlers, depending on their implementation.

## What is Middleware?
Router -> Middleware Handlers -> Application/API Handlers

Middleware is logic/behavior that should be performed before and/or after the intended http request. This includes things such as request logging, auth, checking for headers and error handling.

## Middleware Handler vs Application Handler
Middleware Handlers are quite similar to Application Handlers. They are both `http.Handler` implementations. However, they have one key difference: Middleware Handlers are created as types that contain a `next http.Handler` that shall be called during its execution. This allows you to chain functions together and execute logic before and/or after calling the next registered middleware.

```
type SomeMiddleware{
    next Http.Handler
}
func(s SomeMiddleware) ServeHttp(w http.ResponseWriter, r *http.Request){
    // do some stuff before

    next.ServeHttp(w, r)

    // do some stuff after
}
```
## What is the goal?
Middleware that is meant for reusability should conform to the `http.Handler` interface. Many frameworks provide interfaces and adapters to convert `http.Handler` and `http.HandlerFunc` into their expected types.

### Avoide unnecessary, verbose code
Defining a custom type for each middleware implementation can sometimes be tedious and unnecessary. Instead, we can create a function that conforms to the `http.HanderFunc` type using the `http.HandlerFunc()` adapter to plug it in as a `http.Handler`!

```
type HandlerFunc func(ResponseWriter, *Request) 

// This function means a HandlerFunc is also a Handler! 
func(h HandlerFunc) ServeHttp(ResponseWriter, *Request){
    // serves http! :) 
}

---

func NewSomeMiddleware(next http.Handler) http.Handler{
    return http.HandlerFunc((w http.ResponseWriter, r *http.Request){
        // do some stuff before

        next.ServeHttp(w, r)

        // do some stuff after
    })
}
```
Why can we do this? Because the `HandlerFunc` type just so happens to have a `ServeHttp(ResponseWriter, *Request)` func attached to it, which makes it a `Handler`!

## How can you pass custom data into handlers?
```
func HelloHandler(w ResponseWriter, *Request) {
    w.Write([]byte("Hello, Jake!"))
}
```
In some cases, you may need access to some variables in order for the `HandlerFunc` to complete its job. You can do this using closures!
```
func HelloHandler(name string) http.HandlerFunc{
    func HelloHandler(w ResponseWriter, *Request) {
        w.Write([]byte(fmt.Sprintf("Hello, %s!, name)))
    } 
}

You can apply this same concept when creating middleware.
```
### Sources
https://www.alexedwards.net/blog/making-and-using-middleware