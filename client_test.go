// SPDX-FileCopyrightText: 2025 Jacques Supcik <jacques.supcik@hefr.ch>
//
// SPDX-License-Identifier: MIT

package countries

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func asCca3Map(countries []Country) map[string]Country {
	countryMap := make(map[string]Country)
	for _, country := range countries {
		countryMap[country.Cca3] = country
	}
	return countryMap
}

func TestAll(t *testing.T) {
	countries, err := All("cca3", "name")
	assert.NoError(t, err)
	assert.Greater(t, len(countries), 200)
	c := asCca3Map(countries)
	assert.Contains(t, c, "CHE")
	assert.Contains(t, c, "DEU")
	assert.Contains(t, c, "FRA")
}

func TestByName(t *testing.T) {
	countries, err := ByName("Switzerland", "name", "capital", "region")
	assert.NoError(t, err)
	assert.Len(t, countries, 1)
	assert.Equal(t, "Switzerland", countries[0].Name.Common)
	assert.Equal(t, "Bern", countries[0].Capital[0])
	assert.Equal(t, "Europe", countries[0].Region)
}

func TestByInvalidName(t *testing.T) {
	countries, err := ByName("InvalidCountryName", "name")
	assert.Error(t, err)
	assert.Nil(t, countries)
}

func TestByFullName(t *testing.T) {
	countries, err := ByFullName("Switzerland")
	assert.NoError(t, err)
	assert.Len(t, countries, 1)
	assert.Equal(t, "Switzerland", countries[0].Name.Common)
}

func TestByCode(t *testing.T) {
	country, err := ByCode("CH", "name")
	assert.NoError(t, err)
	assert.Equal(t, "Switzerland", country.Name.Common)
}

func TestByInvalidCode(t *testing.T) {
	country, err := ByCode("XX", "name")
	assert.Error(t, err)
	assert.Equal(t, Country{}, country)
}

func TestByCodes(t *testing.T) {
	countries, err := ByCodes([]string{"CH", "DE"}, "cca3", "name")
	assert.NoError(t, err)
	assert.Len(t, countries, 2)
	c := asCca3Map(countries)
	assert.Contains(t, c, "CHE")
	assert.Contains(t, c, "DEU")
}

func TestByCurrency(t *testing.T) {
	countries, err := ByCurrency("CHF", "cca3", "name")
	assert.NoError(t, err)
	assert.Len(t, countries, 2)
	c := asCca3Map(countries)
	assert.Contains(t, c, "CHE")
	assert.Contains(t, c, "LIE")
}

func TestByDemonym(t *testing.T) {
	countries, err := ByDemonym("Swiss", "name")
	assert.NoError(t, err)
	assert.Len(t, countries, 1)
	assert.Equal(t, "Switzerland", countries[0].Name.Common)
}

func TestByLanguage(t *testing.T) {
	countries, err := ByLanguage("french", "cca3", "name")
	assert.NoError(t, err)
	c := asCca3Map(countries)
	assert.Contains(t, c, "FRA")
	assert.Contains(t, c, "CAN")
	assert.Contains(t, c, "BEL")
	assert.Contains(t, c, "CHE")
}

func TestByCapital(t *testing.T) {
	countries, err := ByCapital("Bern", "name")
	assert.NoError(t, err)
	assert.Len(t, countries, 1)
	assert.Equal(t, "Switzerland", countries[0].Name.Common)
}

func TestByRegion(t *testing.T) {
	countries, err := ByRegion("Europe", "cca3", "name")
	assert.NoError(t, err)
	assert.Greater(t, len(countries), 40)
	c := asCca3Map(countries)
	assert.Contains(t, c, "CHE")
	assert.Contains(t, c, "DEU")
	assert.Contains(t, c, "FRA")
}

func TestBySubregion(t *testing.T) {
	countries, err := BySubregion("Western Europe", "cca3", "name")
	assert.NoError(t, err)
	assert.Greater(t, len(countries), 5)
	c := asCca3Map(countries)
	assert.Contains(t, c, "CHE")
	assert.Contains(t, c, "DEU")
	assert.Contains(t, c, "FRA")
}

func TestByTranslation(t *testing.T) {
	countries, err := ByTranslation("スイス", "name")
	assert.NoError(t, err)
	assert.Len(t, countries, 1)
	assert.Equal(t, "Switzerland", countries[0].Name.Common)
}

func TestByIndependance(t *testing.T) {
	countries, err := ByIndependence(true, "cca3", "name")
	assert.NoError(t, err)
	assert.Greater(t, len(countries), 150)
	c := asCca3Map(countries)
	assert.Contains(t, c, "CHE")
	assert.Contains(t, c, "DEU")
	assert.Contains(t, c, "FRA")

	countries, err = ByIndependence(false, "cca3", "name")
	assert.NoError(t, err)
	assert.Greater(t, len(countries), 10)
	c = asCca3Map(countries)
	assert.Contains(t, c, "ATA") // Antarctica
	assert.Contains(t, c, "BVT") // Bouvet Island
}
