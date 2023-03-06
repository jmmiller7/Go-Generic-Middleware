## Where does the middleware go?
Router -> Middleware Handlers -> Application/API Handlers

## What is the goal?
Implement middlewares that implement the `http.Handler` interface.

``` 
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

### Avoide unnecessary, verbose code
Defining a custom type to conform to the `http.Handler` interface can sometimes be tedious and unnecessary. Instead, we can create a function that conforms to the `http.HanderFunc` type and use the `http.HandlerFunc()` adapter to plug it in!

```
type HandlerFunc func(ResponseWriter, *Request) 
```
And why can we do this? Because the `HandlerFunc` type just so happens to have a `ServeHttp(ResponseWriter, *Request)` func attached to it!

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
```

## Middleware Handler vs Application Handler
Middleware Handlers are quite similar to Application Handlers. They are both `http.HandlerFunc` implementations. However, they have one key difference: Middleware Handlers are created are closures that use a `next` parameter. This allows you to chain functions together and execute logic before or after calling the next registered middleware.
```
func NewExampleMiddleware(next http.Handler) http.Handler{
    func(w http.Writer, r *http.Request){
        // do some pre-exec logic here, if needed...
        
        // call the next middleware/handler in line
        next(w, r)
        
        // do some post-exec logic here, if needed... 
    }
}
```
Using the pattern above, we can apply as many middleware handlers as we desire!

### What are the commonalities across Go Routers?


### Sources
https://www.alexedwards.net/blog/making-and-using-middleware