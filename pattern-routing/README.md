# Pattern Routing

In Go is handled vy the DefaultServeMux method which is an instance of ServeMux. When nil is passed to the
handler parameter for the ListenAndServe function then the DefaultServeMux just indirectly calling
http.DefaultServeMux.HandleFunc(...).

ServeMux allows variable parts in url patterns and stub out the routes we'd like the application to have.

## Choosing a Muxer

Many web apps handle requests to URLs with variable parts such as
`/settings/godric` or `/notes/42`. ServeMux handles HTTP requests by multiplexing over handlers registered to
fixed patterns only, but remember that ServeMux is just one implementation of the `Handler` interface.
net/http wisely defines the foundation Server and Requst, but leaves handling to anything that implements
`ServeHTTP(ResponseWriter, *Request)`.

- Patterns with parameters and easy retrieval of their values
- HTTP method matching rules
- Route reversal to build redirection URLs from a handler

### Example

| HTTP Verb | Path        | Action | Description         |
| --------- | ----------- | ------ | ------------------- |
| GET       | /notes      | list   | list multiple notes |
| POST      | /notes      | create | create a note       |
| GET       | /notes/{id} | read   | get a note          |
| PUT       | /notes/{id} | update | update a note       |
| DELETE    | /notes/{id} | delete | delete a note       |
