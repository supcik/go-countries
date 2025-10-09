// SPDX-FileCopyrightText: 2025 Jacques Supcik <jacques.supcik@hefr.ch>
//
// SPDX-License-Identifier: MIT

package countries

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const ApiVersion = "v3.1"
const BaseUrl = "https://restcountries.com/" + ApiVersion + "/"

type ErrorMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func make_url(path string, queryValues url.Values, fields ...string) *url.URL {
	u, _ := url.Parse(BaseUrl)
	u = u.JoinPath(path)

	q := u.Query()
	for key, values := range queryValues {
		for _, value := range values {
			q.Add(key, value)
		}
	}

	q.Add("fields", strings.Join(fields, ","))
	u.RawQuery = q.Encode()
	return u
}

func requestSingle(path string, queryValues url.Values, fields ...string) (Country, error) {
	u := make_url(path, queryValues, fields...)

	resp, err := http.Get(u.String())
	if err != nil {
		return Country{}, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal("Error closing response body:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		var errorMsg ErrorMessage
		err := json.NewDecoder(resp.Body).Decode(&errorMsg)

		if err != nil {
			return Country{}, err
		}
		return Country{}, fmt.Errorf("API request failed with status %s : %v", resp.Status, errorMsg.Message)
	}

	var country Country
	err = json.NewDecoder(resp.Body).Decode(&country)
	if err != nil {
		return Country{}, err
	}
	return country, nil
}

func request(path string, queryValues url.Values, fields ...string) ([]Country, error) {
	u := make_url(path, queryValues, fields...)

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal("Error closing response body:", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		var errorMsg ErrorMessage
		err := json.NewDecoder(resp.Body).Decode(&errorMsg)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("API request failed with status %s : %v", resp.Status, errorMsg.Message)
	}

	var countries []Country
	err = json.NewDecoder(resp.Body).Decode(&countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

// All retrieves data for all countries. You must specify the fields you need (up to 10 fields)
// when calling this method, otherwise you'll get an error.
func All(fields ...string) ([]Country, error) {
	return request("all", nil, fields...)
}

// ByName retrieves data for countries matching the given name. It can be the common or official
// name. If you want to get an exact match, use the ByFullName function.
func ByName(name string, fields ...string) ([]Country, error) {
	return request("name/"+url.PathEscape(name), nil, fields...)
}

// ByFullName retrieves data for countries matching the exact full name. It can be the common or
// official name.
func ByFullName(name string, fields ...string) ([]Country, error) {
	return request("name/"+url.PathEscape(name), url.Values{"fullText": {"true"}}, fields...)
}

// ByCode retrieves data for the country matching the given country code. You can use either
// cca2 (ISO 3166-1 alpha-2) ccn3 (ISO 3166-1 numeric), cca3 (ISO 3166-1 alpha-3) or cioc
// (International Olympic Committee)
func ByCode(code string, fields ...string) (Country, error) {
	return requestSingle("alpha/"+url.PathEscape(code), nil, fields...)
}

// ByCodes retrieves data for multiple countries matching the given country codes. You can use
// the same codes as in ByCode function. The codes should be provided as a slice of strings.
func ByCodes(codes []string, fields ...string) ([]Country, error) {
	return request("alpha", url.Values{"codes": {strings.Join(codes, ",")}}, fields...)
}

// ByCurrency retrieves data for countries using the given currency code or name.
func ByCurrency(currency string, fields ...string) ([]Country, error) {
	return request("currency/"+url.PathEscape(currency), nil, fields...)
}

// ByDemonym retrieves data for countries matching how a citizen is called.
func ByDemonym(demonym string, fields ...string) ([]Country, error) {
	return request("demonym/"+url.PathEscape(demonym), nil, fields...)
}

// ByLanguage retrieves data for countries where the given language code or name is spoken.
func ByLanguage(language string, fields ...string) ([]Country, error) {
	return request("lang/"+url.PathEscape(language), nil, fields...)
}

// ByCapital retrieves data for countries with the given capital city.
func ByCapital(capital string, fields ...string) ([]Country, error) {
	return request("capital/"+url.PathEscape(capital), nil, fields...)
}

// ByRegion retrieves data for countries in the given region (e.g. Europe)
// Currently available regions are: Africa, Americas, Antarctic, Asia, Europe and Oceania.
func ByRegion(region string, fields ...string) ([]Country, error) {
	return request("region/"+url.PathEscape(region), nil, fields...)
}

// BySubregion retrieves data for countries in the given subregion (e.g. Western Europe)
func BySubregion(subregion string, fields ...string) ([]Country, error) {
	return request("subregion/"+url.PathEscape(subregion), nil, fields...)
}

// ByTranslation retrieves data for countries matching the given translation of the country name.
func ByTranslation(translation string, fields ...string) ([]Country, error) {
	return request("translation/"+url.PathEscape(translation), nil, fields...)
}

// ByIndependence retrieves data for countries based on their independence status.
func ByIndependence(independent bool, fields ...string) ([]Country, error) {
	return request("independent", url.Values{"status": {fmt.Sprintf("%t", independent)}}, fields...)
}
