# req

**req** is a Go client library for accessing the [Registre des entreprises du Québec](http://www.registreentreprises.gouv.qc.ca/en/default.aspx).

**Documentation:** [![GoDoc](https://godoc.org/github.com/quebec/req?status.svg)](https://godoc.org/github.com/quebec/req)  
**Build Status:** [![TravisCI Build Status](https://travis-ci.org/quebec/req.svg)](https://travis-ci.org/quebec/req)

**Code Coverage:** [![Coverage Status](https://coveralls.io/repos/quebec/req/badge.svg?branch=master&service=github)](https://coveralls.io/github/quebec/req?branch=master)

## Usage

```go
import "github.com/quebec/req"
```

Construct a new **req** client, then use the various functions on the client to access different parts of the registry. For example, if you want to get company information using a unique *NEQ* identifier:

```go
c := req.NewClient(nil)
company, _ := c.GetNEQ("1143920115")
fmt.Println(company.SectionInformationsGenerales.SousSecIdentification.NomEntreprise)
// Output: BOMBARDIER INC.
```

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
