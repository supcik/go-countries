package countries_test

import (
	"fmt"

	"github.com/supcik/go-countries"
)

func ExampleByName() {
	countries, err := countries.ByName("Switzerland", "name", "capital", "region")
	if err != nil {
		panic(err)
	}
	for _, country := range countries {
		fmt.Printf("The capital of %s (%s) is %s.\n", country.Name.Common, country.Region, country.Capital[0])
	}
	// Output: The capital of Switzerland (Europe) is Bern.
}
