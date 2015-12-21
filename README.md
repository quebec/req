# req

**req** is a Go client library for accessing the [Registre des entreprises du Québec](http://www.registreentreprises.gouv.qc.ca/en/default.aspx).

**Documentation:** [![GoDoc](https://godoc.org/github.com/quebec/req?status.svg)](https://godoc.org/github.com/quebec/req)  
**Build Status:** [![TravisCI Build Status](https://travis-ci.org/quebec/req.svg)](https://travis-ci.org/quebec/req)

## Usage

```go
import "github.com/quebec/req"
```

Construct a new **req** client, then use the various functions on the client to access different parts of the registry.

## Roadmap

This library is being initially developed for one of my internal project,
so API methods will likely be implemented in the order that they are
needed by my project. Eventually, I would like to cover the entire API,
so contributions are of course [always welcome][contributing]. The
calling pattern is pretty well established, so adding new methods is relatively
straightforward.

[contributing]: CONTRIBUTING.md


## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.