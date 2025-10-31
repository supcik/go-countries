// SPDX-FileCopyrightText: 2025 Jacques Supcik <jacques.supcik@hefr.ch>
//
// SPDX-License-Identifier: MIT

// Package countries provides data structures for representing country information.
// It also includes functions to retrieve country data from the REST Countries API (https://restcountries.com/).

package countries

// NativeNameT represents the native name structure
type NativeNameT struct {
	Official string `json:"official"` // Official  country name
	Common   string `json:"common"`   // Common country name
}

// Name represents the name structure. It includes the native names, and a map of native names in different languages.
type Name struct {
	NativeNameT
	NativeName map[string]NativeNameT `json:"nativeName"` // Native country name
}

// Currency represents the currency structure with its name and symbol.
type Currency struct {
	Name   string `json:"name"`   // Currency name
	Symbol string `json:"symbol"` // Currency symbol
}

// Idd represents the international direct dialing information, including the root and suffixes.
type Idd struct {
	Root     string   `json:"root"`     // International direct dialing root
	Suffixes []string `json:"suffixes"` // International direct dialing suffixes
}

// CapitalInfo represents the capital information, specifically the latitude and longitude coordinates.
type CapitalInfo struct {
	Latlng []float64 `json:"latlng"` // Capital latitude and longitude
}

// Demonyms represents the name of the inhabitants, with separate fields for male and female forms.
type Demonyms struct {
	F string `json:"f"` // Female demonym
	M string `json:"m"` // Male demonym
}

// Car represents the driving side and distinguished (oval) signs of a country.
type Car struct {
	Signs []string `json:"signs"` // Car distinguished (oval) signs
	Side  string   `json:"side"`  // Car driving side
}

// PostalCode represents the postal code format and regex pattern.
type PostalCode struct {
	Format string `json:"format"` // Postal code format
	Regex  string `json:"regex"`  // Postal code regex pattern
}

// Country represents the main structure for a country, including various details such as name, codes, currencies, languages, and more.
type Country struct {
	Name         Name                `json:"name"`
	Tld          []string            `json:"tld"`          // Internet top level domains
	Cca2         string              `json:"cca2"`         // ISO 3166-1 alpha-2 two-letter country codes
	Ccn3         string              `json:"ccn3"`         // ISO 3166-1 numeric code (UN M49)
	Cca3         string              `json:"cca3"`         // ISO 3166-1 alpha-3 three-letter country codes
	Cioc         string              `json:"cioc"`         // Code of the International Olympic Committee
	Fifa         string              `json:"fifa"`         // FIFA code
	Independent  bool                `json:"independent"`  // ISO 3166-1 independence status (the country is considered a sovereign state)
	Status       string              `json:"status"`       // ISO 3166-1 assignment status
	UnMember     bool                `json:"unMember"`     // UN Member status
	Currencies   map[string]Currency `json:"currencies"`   // List of all currencies
	Idd          Idd                 `json:"idd"`          // International dialing codes
	Capital      []string            `json:"capital"`      // Capital cities
	CapitalInfo  CapitalInfo         `json:"capitalInfo"`  // Capital latitude and longitude
	AltSpellings []string            `json:"altSpellings"` // Alternate spellings of the country name
	Region       string              `json:"region"`       // UN demographic regions
	Subregion    string              `json:"subregion"`    // UN demographic subregions
	Continents   []string            `json:"continents"`   // List of continents the country is on
	Languages    map[string]string   `json:"languages"`    // List of official language
	Translations map[string]Name     `json:"translations"` // List of country name translations
	Latlng       []float64           `json:"latlng"`       // Latitude and longitude
	Landlocked   bool                `json:"landlocked"`   // Landlocked country
	Borders      []string            `json:"borders"`      // Border countries
	Area         float64             `json:"area"`         // Geographical size
	Demonyms     Demonyms            `json:"demonyms"`     // Inhabitants of the country
	Flag         string              `json:"flag"`         // flag emoji
	Flags        map[string]string   `json:"flags"`        // Flagpedia links to svg and png flags
	CoatOfArms   map[string]string   `json:"coatOfArms"`   // MainFacts.com links to svg and png images
	Population   int                 `json:"population"`   // Country population
	Maps         map[string]string   `json:"maps"`         // Link to Google maps and Open Street maps
	Gini         map[string]float64  `json:"gini"`         // Worldbank Gini index
	Car          Car                 `json:"car"`          // Car driving side and distinguished (oval) signs
	PostalCode   PostalCode          `json:"postalCode"`   // Country postal codes
	StartOfWeek  string              `json:"startOfWeek"`  // Day of the start of week (Sunday/Monday/Saturday)
	Timezones    []string            `json:"timezones"`    // Timezones
}
