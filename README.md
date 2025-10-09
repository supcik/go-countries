# Go REST Countries API Wrapper

Go wrapper for the REST Countries API (API v3.1).

This small, dependency-light package provides typed access to https://restcountries.com/ (v3.1). It mirrors the API endpoints and returns a convenient `Country` struct for JSON responses.

## Key features

- Lightweight, idiomatic Go API around REST Countries v3.1
- Typed `Country` model covering commonly used fields (name, codes, capital, currencies, flags, population, etc.)
- Convenience functions for the most common endpoints: `All`, `ByName`, `ByCode`, `ByCapital`, `ByRegion`, and more

## Installation

The module path is `github.com/supcik/countries`. To add it to your project use Go modules:

```bash
go get github.com/supcik/countries@latest
```

or import it directly in your code and run `go mod tidy`:

```go
import "github.com/supcik/countries"
```

## Quick example

Create a small `main.go` to fetch a country by ISO code and print its capital:

```go
package main

import (
	"fmt"
	"log"

	"github.com/supcik/countries"
)

func main() {
	// ByCode returns a single Country for an ISO code like "CH" or "FRA"
	c, err := countries.ByCode("CH", "name", "capital", "cca2", "population")
	if err != nil {
		log.Fatalf("failed to fetch country: %v", err)
	}

	fmt.Printf("Country: %s\n", c.Name.Common)
	if len(c.Capital) > 0 {
		fmt.Printf("Capital: %s\n", c.Capital[0])
	}
	fmt.Printf("Population: %d\n", c.Population)
}
```

## Notes about fields

The REST Countries API supports a `fields` query parameter to limit the
returned fields (and reduce payload). This wrapper mirrors that
behaviour: most functions accept a variadic `fields ...string`
parameter. When calling `All()` you must specify the fields you need
(the underlying API rejects requests that do not include a fields filter
when returning all countries). Example fields: `name`, `cca2`,
`capital`, `currencies`, `flags`, `population`.

Examples

- Get all countries with only name and cca2:

```go
countries.All("name", "cca2")
```

- Search by name (partial match):

```go
countries.ByName("united", "name", "cca3", "capital")
```

- Exact full name match:

```go
countries.ByFullName("United Kingdom of Great Britain and Northern Ireland", "name", "capital")
```

- Multiple codes in one request:

```go
codes := []string{"CH", "DE", "FR"}
countries.ByCodes(codes, "name", "capital")
```

## Contributing

Contributions are welcome. Open an issue for bugs or feature requests and submit PRs for fixes or improvements. Keep changes small and focused and include tests where appropriate.

## License

This project is licensed under the MIT License — see `LICENSE.txt` for details.

## Acknowledgements and references

- REST Countries: https://restcountries.com/
- Related libraries (for reference):
  - github.com/chriscross0/go-restcountries/v2
  - github.com/joaomlopes/gocountries
  - github.com/alediaferia/gocountries

## Contact

Maintainer: Jacques Supcik — jacques.supcik@hefr.ch
