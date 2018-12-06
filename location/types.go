package location

type Locations []Location

type Location struct {
	Name           string  `json:name`           // "Montr\xe9al"
	Province       string  `json:province`       // "Quebec"
	NorthAmericas  bool    `json:northamericas`  // true
	DisplayCountry bool    `json:displayCountry` // false
	Code           string  `json:code`           // "caqc0363"
	ProvCode       string  `json:provcode`       // "QC"
	CountryCode    string  `json:countrycode`    // "CA"
	Country        string  `json:country`        // "Canada"
	Lat            float64 `json:lat`            // 45.50669827
	Long           float64 `json:long`           // -73.55561373
	Type           string  `json:type`           // "city"
	PostalCode     string  `json:postalcode`     // ""
	DistanceKM     string  `json:distanceKM`     // ""
	County         string  `json:county`         // "Communaut\xe9-Urbaine-de-Montr\xe9al"
	GridIndex      string  `json:gridindex`      // ""
	URL            string  `json:url`            // "/ca/weather/quebec/montreal"
}
