package ipinfo

// GetCountryName gets the full name of a country from its code, e.g.
// "PK" -> "Pakistan".
func GetCountryName(country string) string {
	return countriesMap[country]
}

// GetCountryFlagEmoji gets the emoji flag of a country from its code, e.g.
// "PK" -> "🇵🇰".
func GetCountryFlagEmoji(country string) string {
	return countriesFlags[country].Emoji
}

// GetCountryFlagUnicode gets the unicode of an emoji from country code, e.g.
// "PK" -> "U+1F1F5 U+1F1F0".
func GetCountryFlagUnicode(country string) string {
	return countriesFlags[country].Unicode
}

// GetCountryCurrencyCode gets the currency code of a country from its code, e.g.
// "PK" -> "PKR".
func GetCountryCurrencyCode(country string) string {
	return countriesCurrencies[country].Code
}

// GetCountryCurrencySymbol gets the symbol of currency from country code, e.g.
// "PK" -> "₨".
func GetCountryCurrencySymbol(country string) string {
	return countriesCurrencies[country].Symbol
}

// IsEU takes the country code and returns `true`
// if the country is a member of the EU, e.g. "SE" -> true
func IsEU(country string) bool {
	for _, val := range euCountries {
		if val == country {
			return true
		}
	}
	return false
}

type CountryFlag struct {
	Emoji   string `json:"emoji,omitempty" csv:"emoji"`
	Unicode string `json:"unicode,omitempty" csv:"unicode"`
}

type CountryCurrency struct {
	Code   string `json:"code,omitempty" csv:"code"`
	Symbol string `json:"symbol,omitempty" csv:"symbol"`
}

var countriesMap = map[string]string{
	"BD": "Bangladesh",
	"BE": "Belgium",
	"BF": "Burkina Faso",
	"BG": "Bulgaria",
	"BA": "Bosnia and Herzegovina",
	"BB": "Barbados",
	"WF": "Wallis and Futuna",
	"BL": "Saint Barthelemy",
	"BM": "Bermuda",
	"BN": "Brunei",
	"BO": "Bolivia",
	"BH": "Bahrain",
	"BI": "Burundi",
	"BJ": "Benin",
	"BT": "Bhutan",
	"JM": "Jamaica",
	"BV": "Bouvet Island",
	"BW": "Botswana",
	"WS": "Samoa",
	"BQ": "Bonaire, Saint Eustatius and Saba ",
	"BR": "Brazil",
	"BS": "Bahamas",
	"JE": "Jersey",
	"BY": "Belarus",
	"BZ": "Belize",
	"RU": "Russia",
	"RW": "Rwanda",
	"RS": "Serbia",
	"TL": "East Timor",
	"RE": "Reunion",
	"TM": "Turkmenistan",
	"TJ": "Tajikistan",
	"RO": "Romania",
	"TK": "Tokelau",
	"GW": "Guinea-Bissau",
	"GU": "Guam",
	"GT": "Guatemala",
	"GS": "South Georgia and the South Sandwich Islands",
	"GR": "Greece",
	"GQ": "Equatorial Guinea",
	"GP": "Guadeloupe",
	"JP": "Japan",
	"GY": "Guyana",
	"GG": "Guernsey",
	"GF": "French Guiana",
	"GE": "Georgia",
	"GD": "Grenada",
	"GB": "United Kingdom",
	"GA": "Gabon",
	"SV": "El Salvador",
	"GN": "Guinea",
	"GM": "Gambia",
	"GL": "Greenland",
	"GI": "Gibraltar",
	"GH": "Ghana",
	"OM": "Oman",
	"TN": "Tunisia",
	"JO": "Jordan",
	"HR": "Croatia",
	"HT": "Haiti",
	"HU": "Hungary",
	"HK": "Hong Kong",
	"HN": "Honduras",
	"HM": "Heard Island and McDonald Islands",
	"VE": "Venezuela",
	"PR": "Puerto Rico",
	"PS": "Palestinian Territory",
	"PW": "Palau",
	"PT": "Portugal",
	"SJ": "Svalbard and Jan Mayen",
	"PY": "Paraguay",
	"IQ": "Iraq",
	"PA": "Panama",
	"PF": "French Polynesia",
	"PG": "Papua New Guinea",
	"PE": "Peru",
	"PK": "Pakistan",
	"PH": "Philippines",
	"PN": "Pitcairn",
	"PL": "Poland",
	"PM": "Saint Pierre and Miquelon",
	"ZM": "Zambia",
	"EH": "Western Sahara",
	"EE": "Estonia",
	"EG": "Egypt",
	"ZA": "South Africa",
	"EC": "Ecuador",
	"IT": "Italy",
	"VN": "Vietnam",
	"SB": "Solomon Islands",
	"ET": "Ethiopia",
	"SO": "Somalia",
	"ZW": "Zimbabwe",
	"SA": "Saudi Arabia",
	"ES": "Spain",
	"ER": "Eritrea",
	"ME": "Montenegro",
	"MD": "Moldova",
	"MG": "Madagascar",
	"MF": "Saint Martin",
	"MA": "Morocco",
	"MC": "Monaco",
	"UZ": "Uzbekistan",
	"MM": "Myanmar",
	"ML": "Mali",
	"MO": "Macao",
	"MN": "Mongolia",
	"MH": "Marshall Islands",
	"MK": "Macedonia",
	"MU": "Mauritius",
	"MT": "Malta",
	"MW": "Malawi",
	"MV": "Maldives",
	"MQ": "Martinique",
	"MP": "Northern Mariana Islands",
	"MS": "Montserrat",
	"MR": "Mauritania",
	"IM": "Isle of Man",
	"UG": "Uganda",
	"TZ": "Tanzania",
	"MY": "Malaysia",
	"MX": "Mexico",
	"IL": "Israel",
	"FR": "France",
	"IO": "British Indian Ocean Territory",
	"SH": "Saint Helena",
	"FI": "Finland",
	"FJ": "Fiji",
	"FK": "Falkland Islands",
	"FM": "Micronesia",
	"FO": "Faroe Islands",
	"NI": "Nicaragua",
	"NL": "Netherlands",
	"NO": "Norway",
	"NA": "Namibia",
	"VU": "Vanuatu",
	"NC": "New Caledonia",
	"NE": "Niger",
	"NF": "Norfolk Island",
	"NG": "Nigeria",
	"NZ": "New Zealand",
	"NP": "Nepal",
	"NR": "Nauru",
	"NU": "Niue",
	"CK": "Cook Islands",
	"XK": "Kosovo",
	"CI": "Ivory Coast",
	"CH": "Switzerland",
	"CO": "Colombia",
	"CN": "China",
	"CM": "Cameroon",
	"CL": "Chile",
	"CC": "Cocos Islands",
	"CA": "Canada",
	"CG": "Republic of the Congo",
	"CF": "Central African Republic",
	"CD": "Democratic Republic of the Congo",
	"CZ": "Czech Republic",
	"CY": "Cyprus",
	"CX": "Christmas Island",
	"CR": "Costa Rica",
	"CW": "Curacao",
	"CV": "Cape Verde",
	"CU": "Cuba",
	"SZ": "Swaziland",
	"SY": "Syria",
	"SX": "Sint Maarten",
	"KG": "Kyrgyzstan",
	"KE": "Kenya",
	"SS": "South Sudan",
	"SR": "Suriname",
	"KI": "Kiribati",
	"KH": "Cambodia",
	"KN": "Saint Kitts and Nevis",
	"KM": "Comoros",
	"ST": "Sao Tome and Principe",
	"SK": "Slovakia",
	"KR": "South Korea",
	"SI": "Slovenia",
	"KP": "North Korea",
	"KW": "Kuwait",
	"SN": "Senegal",
	"SM": "San Marino",
	"SL": "Sierra Leone",
	"SC": "Seychelles",
	"KZ": "Kazakhstan",
	"KY": "Cayman Islands",
	"SG": "Singapore",
	"SE": "Sweden",
	"SD": "Sudan",
	"DO": "Dominican Republic",
	"DM": "Dominica",
	"DJ": "Djibouti",
	"DK": "Denmark",
	"VG": "British Virgin Islands",
	"DE": "Germany",
	"YE": "Yemen",
	"DZ": "Algeria",
	"US": "United States",
	"UY": "Uruguay",
	"YT": "Mayotte",
	"UM": "United States Minor Outlying Islands",
	"LB": "Lebanon",
	"LC": "Saint Lucia",
	"LA": "Laos",
	"TV": "Tuvalu",
	"TW": "Taiwan",
	"TT": "Trinidad and Tobago",
	"TR": "Turkey",
	"LK": "Sri Lanka",
	"LI": "Liechtenstein",
	"LV": "Latvia",
	"TO": "Tonga",
	"LT": "Lithuania",
	"LU": "Luxembourg",
	"LR": "Liberia",
	"LS": "Lesotho",
	"TH": "Thailand",
	"TF": "French Southern Territories",
	"TG": "Togo",
	"TD": "Chad",
	"TC": "Turks and Caicos Islands",
	"LY": "Libya",
	"VA": "Vatican",
	"VC": "Saint Vincent and the Grenadines",
	"AE": "United Arab Emirates",
	"AD": "Andorra",
	"AG": "Antigua and Barbuda",
	"AF": "Afghanistan",
	"AI": "Anguilla",
	"VI": "U.S. Virgin Islands",
	"IS": "Iceland",
	"IR": "Iran",
	"AM": "Armenia",
	"AL": "Albania",
	"AO": "Angola",
	"AQ": "Antarctica",
	"AS": "American Samoa",
	"AR": "Argentina",
	"AU": "Australia",
	"AT": "Austria",
	"AW": "Aruba",
	"IN": "India",
	"AX": "Aland Islands",
	"AZ": "Azerbaijan",
	"IE": "Ireland",
	"ID": "Indonesia",
	"UA": "Ukraine",
	"QA": "Qatar",
	"MZ": "Mozambique",
}

var euCountries = []string{
	"IE", "AT", "LT", "LU", "LV", "DE", "DK", "SE", "SI", "SK", "CZ", "CY",
	"NL", "FI", "FR", "MT", "ES", "IT", "EE", "PL", "PT", "HU", "HR", "GR",
	"RO", "BG", "BE",
}

var countriesFlags = map[string]CountryFlag{
	"AD": {"🇦🇩", "U+1F1E6 U+1F1E9"},
	"AE": {"🇦🇪", "U+1F1E6 U+1F1EA"},
	"AF": {"🇦🇫", "U+1F1E6 U+1F1EB"},
	"AG": {"🇦🇬", "U+1F1E6 U+1F1EC"},
	"AI": {"🇦🇮", "U+1F1E6 U+1F1EE"},
	"AL": {"🇦🇱", "U+1F1E6 U+1F1F1"},
	"AM": {"🇦🇲", "U+1F1E6 U+1F1F2"},
	"AO": {"🇦🇴", "U+1F1E6 U+1F1F4"},
	"AQ": {"🇦🇶", "U+1F1E6 U+1F1F6"},
	"AR": {"🇦🇷", "U+1F1E6 U+1F1F7"},
	"AS": {"🇦🇸", "U+1F1E6 U+1F1F8"},
	"AT": {"🇦🇹", "U+1F1E6 U+1F1F9"},
	"AU": {"🇦🇺", "U+1F1E6 U+1F1FA"},
	"AW": {"🇦🇼", "U+1F1E6 U+1F1FC"},
	"AX": {"🇦🇽", "U+1F1E6 U+1F1FD"},
	"AZ": {"🇦🇿", "U+1F1E6 U+1F1FF"},
	"BA": {"🇧🇦", "U+1F1E7 U+1F1E6"},
	"BB": {"🇧🇧", "U+1F1E7 U+1F1E7"},
	"BD": {"🇧🇩", "U+1F1E7 U+1F1E9"},
	"BE": {"🇧🇪", "U+1F1E7 U+1F1EA"},
	"BF": {"🇧🇫", "U+1F1E7 U+1F1EB"},
	"BG": {"🇧🇬", "U+1F1E7 U+1F1EC"},
	"BH": {"🇧🇭", "U+1F1E7 U+1F1ED"},
	"BI": {"🇧🇮", "U+1F1E7 U+1F1EE"},
	"BJ": {"🇧🇯", "U+1F1E7 U+1F1EF"},
	"BL": {"🇧🇱", "U+1F1E7 U+1F1F1"},
	"BM": {"🇧🇲", "U+1F1E7 U+1F1F2"},
	"BN": {"🇧🇳", "U+1F1E7 U+1F1F3"},
	"BO": {"🇧🇴", "U+1F1E7 U+1F1F4"},
	"BQ": {"🇧🇶", "U+1F1E7 U+1F1F6"},
	"BR": {"🇧🇷", "U+1F1E7 U+1F1F7"},
	"BS": {"🇧🇸", "U+1F1E7 U+1F1F8"},
	"BT": {"🇧🇹", "U+1F1E7 U+1F1F9"},
	"BV": {"🇧🇻", "U+1F1E7 U+1F1FB"},
	"BW": {"🇧🇼", "U+1F1E7 U+1F1FC"},
	"BY": {"🇧🇾", "U+1F1E7 U+1F1FE"},
	"BZ": {"🇧🇿", "U+1F1E7 U+1F1FF"},
	"CA": {"🇨🇦", "U+1F1E8 U+1F1E6"},
	"CC": {"🇨🇨", "U+1F1E8 U+1F1E8"},
	"CD": {"🇨🇩", "U+1F1E8 U+1F1E9"},
	"CF": {"🇨🇫", "U+1F1E8 U+1F1EB"},
	"CG": {"🇨🇬", "U+1F1E8 U+1F1EC"},
	"CH": {"🇨🇭", "U+1F1E8 U+1F1ED"},
	"CI": {"🇨🇮", "U+1F1E8 U+1F1EE"},
	"CK": {"🇨🇰", "U+1F1E8 U+1F1F0"},
	"CL": {"🇨🇱", "U+1F1E8 U+1F1F1"},
	"CM": {"🇨🇲", "U+1F1E8 U+1F1F2"},
	"CN": {"🇨🇳", "U+1F1E8 U+1F1F3"},
	"CO": {"🇨🇴", "U+1F1E8 U+1F1F4"},
	"CR": {"🇨🇷", "U+1F1E8 U+1F1F7"},
	"CU": {"🇨🇺", "U+1F1E8 U+1F1FA"},
	"CV": {"🇨🇻", "U+1F1E8 U+1F1FB"},
	"CW": {"🇨🇼", "U+1F1E8 U+1F1FC"},
	"CX": {"🇨🇽", "U+1F1E8 U+1F1FD"},
	"CY": {"🇨🇾", "U+1F1E8 U+1F1FE"},
	"CZ": {"🇨🇿", "U+1F1E8 U+1F1FF"},
	"DE": {"🇩🇪", "U+1F1E9 U+1F1EA"},
	"DJ": {"🇩🇯", "U+1F1E9 U+1F1EF"},
	"DK": {"🇩🇰", "U+1F1E9 U+1F1F0"},
	"DM": {"🇩🇲", "U+1F1E9 U+1F1F2"},
	"DO": {"🇩🇴", "U+1F1E9 U+1F1F4"},
	"DZ": {"🇩🇿", "U+1F1E9 U+1F1FF"},
	"EC": {"🇪🇨", "U+1F1EA U+1F1E8"},
	"EE": {"🇪🇪", "U+1F1EA U+1F1EA"},
	"EG": {"🇪🇬", "U+1F1EA U+1F1EC"},
	"EH": {"🇪🇭", "U+1F1EA U+1F1ED"},
	"ER": {"🇪🇷", "U+1F1EA U+1F1F7"},
	"ES": {"🇪🇸", "U+1F1EA U+1F1F8"},
	"ET": {"🇪🇹", "U+1F1EA U+1F1F9"},
	"FI": {"🇫🇮", "U+1F1EB U+1F1EE"},
	"FJ": {"🇫🇯", "U+1F1EB U+1F1EF"},
	"FK": {"🇫🇰", "U+1F1EB U+1F1F0"},
	"FM": {"🇫🇲", "U+1F1EB U+1F1F2"},
	"FO": {"🇫🇴", "U+1F1EB U+1F1F4"},
	"FR": {"🇫🇷", "U+1F1EB U+1F1F7"},
	"GA": {"🇬🇦", "U+1F1EC U+1F1E6"},
	"GB": {"🇬🇧", "U+1F1EC U+1F1E7"},
	"GD": {"🇬🇩", "U+1F1EC U+1F1E9"},
	"GE": {"🇬🇪", "U+1F1EC U+1F1EA"},
	"GF": {"🇬🇫", "U+1F1EC U+1F1EB"},
	"GG": {"🇬🇬", "U+1F1EC U+1F1EC"},
	"GH": {"🇬🇭", "U+1F1EC U+1F1ED"},
	"GI": {"🇬🇮", "U+1F1EC U+1F1EE"},
	"GL": {"🇬🇱", "U+1F1EC U+1F1F1"},
	"GM": {"🇬🇲", "U+1F1EC U+1F1F2"},
	"GN": {"🇬🇳", "U+1F1EC U+1F1F3"},
	"GP": {"🇬🇵", "U+1F1EC U+1F1F5"},
	"GQ": {"🇬🇶", "U+1F1EC U+1F1F6"},
	"GR": {"🇬🇷", "U+1F1EC U+1F1F7"},
	"GS": {"🇬🇸", "U+1F1EC U+1F1F8"},
	"GT": {"🇬🇹", "U+1F1EC U+1F1F9"},
	"GU": {"🇬🇺", "U+1F1EC U+1F1FA"},
	"GW": {"🇬🇼", "U+1F1EC U+1F1FC"},
	"GY": {"🇬🇾", "U+1F1EC U+1F1FE"},
	"HK": {"🇭🇰", "U+1F1ED U+1F1F0"},
	"HM": {"🇭🇲", "U+1F1ED U+1F1F2"},
	"HN": {"🇭🇳", "U+1F1ED U+1F1F3"},
	"HR": {"🇭🇷", "U+1F1ED U+1F1F7"},
	"HT": {"🇭🇹", "U+1F1ED U+1F1F9"},
	"HU": {"🇭🇺", "U+1F1ED U+1F1FA"},
	"ID": {"🇮🇩", "U+1F1EE U+1F1E9"},
	"IE": {"🇮🇪", "U+1F1EE U+1F1EA"},
	"IL": {"🇮🇱", "U+1F1EE U+1F1F1"},
	"IM": {"🇮🇲", "U+1F1EE U+1F1F2"},
	"IN": {"🇮🇳", "U+1F1EE U+1F1F3"},
	"IO": {"🇮🇴", "U+1F1EE U+1F1F4"},
	"IQ": {"🇮🇶", "U+1F1EE U+1F1F6"},
	"IR": {"🇮🇷", "U+1F1EE U+1F1F7"},
	"IS": {"🇮🇸", "U+1F1EE U+1F1F8"},
	"IT": {"🇮🇹", "U+1F1EE U+1F1F9"},
	"JE": {"🇯🇪", "U+1F1EF U+1F1EA"},
	"JM": {"🇯🇲", "U+1F1EF U+1F1F2"},
	"JO": {"🇯🇴", "U+1F1EF U+1F1F4"},
	"JP": {"🇯🇵", "U+1F1EF U+1F1F5"},
	"KE": {"🇰🇪", "U+1F1F0 U+1F1EA"},
	"KG": {"🇰🇬", "U+1F1F0 U+1F1EC"},
	"KH": {"🇰🇭", "U+1F1F0 U+1F1ED"},
	"KI": {"🇰🇮", "U+1F1F0 U+1F1EE"},
	"KM": {"🇰🇲", "U+1F1F0 U+1F1F2"},
	"KN": {"🇰🇳", "U+1F1F0 U+1F1F3"},
	"KP": {"🇰🇵", "U+1F1F0 U+1F1F5"},
	"KR": {"🇰🇷", "U+1F1F0 U+1F1F7"},
	"KW": {"🇰🇼", "U+1F1F0 U+1F1FC"},
	"KY": {"🇰🇾", "U+1F1F0 U+1F1FE"},
	"KZ": {"🇰🇿", "U+1F1F0 U+1F1FF"},
	"LA": {"🇱🇦", "U+1F1F1 U+1F1E6"},
	"LB": {"🇱🇧", "U+1F1F1 U+1F1E7"},
	"LC": {"🇱🇨", "U+1F1F1 U+1F1E8"},
	"LI": {"🇱🇮", "U+1F1F1 U+1F1EE"},
	"LK": {"🇱🇰", "U+1F1F1 U+1F1F0"},
	"LR": {"🇱🇷", "U+1F1F1 U+1F1F7"},
	"LS": {"🇱🇸", "U+1F1F1 U+1F1F8"},
	"LT": {"🇱🇹", "U+1F1F1 U+1F1F9"},
	"LU": {"🇱🇺", "U+1F1F1 U+1F1FA"},
	"LV": {"🇱🇻", "U+1F1F1 U+1F1FB"},
	"LY": {"🇱🇾", "U+1F1F1 U+1F1FE"},
	"MA": {"🇲🇦", "U+1F1F2 U+1F1E6"},
	"MC": {"🇲🇨", "U+1F1F2 U+1F1E8"},
	"MD": {"🇲🇩", "U+1F1F2 U+1F1E9"},
	"ME": {"🇲🇪", "U+1F1F2 U+1F1EA"},
	"MF": {"🇲🇫", "U+1F1F2 U+1F1EB"},
	"MG": {"🇲🇬", "U+1F1F2 U+1F1EC"},
	"MH": {"🇲🇭", "U+1F1F2 U+1F1ED"},
	"MK": {"🇲🇰", "U+1F1F2 U+1F1F0"},
	"ML": {"🇲🇱", "U+1F1F2 U+1F1F1"},
	"MM": {"🇲🇲", "U+1F1F2 U+1F1F2"},
	"MN": {"🇲🇳", "U+1F1F2 U+1F1F3"},
	"MO": {"🇲🇴", "U+1F1F2 U+1F1F4"},
	"MP": {"🇲🇵", "U+1F1F2 U+1F1F5"},
	"MQ": {"🇲🇶", "U+1F1F2 U+1F1F6"},
	"MR": {"🇲🇷", "U+1F1F2 U+1F1F7"},
	"MS": {"🇲🇸", "U+1F1F2 U+1F1F8"},
	"MT": {"🇲🇹", "U+1F1F2 U+1F1F9"},
	"MU": {"🇲🇺", "U+1F1F2 U+1F1FA"},
	"MV": {"🇲🇻", "U+1F1F2 U+1F1FB"},
	"MW": {"🇲🇼", "U+1F1F2 U+1F1FC"},
	"MX": {"🇲🇽", "U+1F1F2 U+1F1FD"},
	"MY": {"🇲🇾", "U+1F1F2 U+1F1FE"},
	"MZ": {"🇲🇿", "U+1F1F2 U+1F1FF"},
	"NA": {"🇳🇦", "U+1F1F3 U+1F1E6"},
	"NC": {"🇳🇨", "U+1F1F3 U+1F1E8"},
	"NE": {"🇳🇪", "U+1F1F3 U+1F1EA"},
	"NF": {"🇳🇫", "U+1F1F3 U+1F1EB"},
	"NG": {"🇳🇬", "U+1F1F3 U+1F1EC"},
	"NI": {"🇳🇮", "U+1F1F3 U+1F1EE"},
	"NL": {"🇳🇱", "U+1F1F3 U+1F1F1"},
	"NO": {"🇳🇴", "U+1F1F3 U+1F1F4"},
	"NP": {"🇳🇵", "U+1F1F3 U+1F1F5"},
	"NR": {"🇳🇷", "U+1F1F3 U+1F1F7"},
	"NU": {"🇳🇺", "U+1F1F3 U+1F1FA"},
	"NZ": {"🇳🇿", "U+1F1F3 U+1F1FF"},
	"OM": {"🇴🇲", "U+1F1F4 U+1F1F2"},
	"PA": {"🇵🇦", "U+1F1F5 U+1F1E6"},
	"PE": {"🇵🇪", "U+1F1F5 U+1F1EA"},
	"PF": {"🇵🇫", "U+1F1F5 U+1F1EB"},
	"PG": {"🇵🇬", "U+1F1F5 U+1F1EC"},
	"PH": {"🇵🇭", "U+1F1F5 U+1F1ED"},
	"PK": {"🇵🇰", "U+1F1F5 U+1F1F0"},
	"PL": {"🇵🇱", "U+1F1F5 U+1F1F1"},
	"PM": {"🇵🇲", "U+1F1F5 U+1F1F2"},
	"PN": {"🇵🇳", "U+1F1F5 U+1F1F3"},
	"PR": {"🇵🇷", "U+1F1F5 U+1F1F7"},
	"PS": {"🇵🇸", "U+1F1F5 U+1F1F8"},
	"PT": {"🇵🇹", "U+1F1F5 U+1F1F9"},
	"PW": {"🇵🇼", "U+1F1F5 U+1F1FC"},
	"PY": {"🇵🇾", "U+1F1F5 U+1F1FE"},
	"QA": {"🇶🇦", "U+1F1F6 U+1F1E6"},
	"RE": {"🇷🇪", "U+1F1F7 U+1F1EA"},
	"RO": {"🇷🇴", "U+1F1F7 U+1F1F4"},
	"RS": {"🇷🇸", "U+1F1F7 U+1F1F8"},
	"RU": {"🇷🇺", "U+1F1F7 U+1F1FA"},
	"RW": {"🇷🇼", "U+1F1F7 U+1F1FC"},
	"SA": {"🇸🇦", "U+1F1F8 U+1F1E6"},
	"SB": {"🇸🇧", "U+1F1F8 U+1F1E7"},
	"SC": {"🇸🇨", "U+1F1F8 U+1F1E8"},
	"SD": {"🇸🇩", "U+1F1F8 U+1F1E9"},
	"SE": {"🇸🇪", "U+1F1F8 U+1F1EA"},
	"SG": {"🇸🇬", "U+1F1F8 U+1F1EC"},
	"SH": {"🇸🇭", "U+1F1F8 U+1F1ED"},
	"SI": {"🇸🇮", "U+1F1F8 U+1F1EE"},
	"SJ": {"🇸🇯", "U+1F1F8 U+1F1EF"},
	"SK": {"🇸🇰", "U+1F1F8 U+1F1F0"},
	"SL": {"🇸🇱", "U+1F1F8 U+1F1F1"},
	"SM": {"🇸🇲", "U+1F1F8 U+1F1F2"},
	"SN": {"🇸🇳", "U+1F1F8 U+1F1F3"},
	"SO": {"🇸🇴", "U+1F1F8 U+1F1F4"},
	"SR": {"🇸🇷", "U+1F1F8 U+1F1F7"},
	"SS": {"🇸🇸", "U+1F1F8 U+1F1F8"},
	"ST": {"🇸🇹", "U+1F1F8 U+1F1F9"},
	"SV": {"🇸🇻", "U+1F1F8 U+1F1FB"},
	"SX": {"🇸🇽", "U+1F1F8 U+1F1FD"},
	"SY": {"🇸🇾", "U+1F1F8 U+1F1FE"},
	"SZ": {"🇸🇿", "U+1F1F8 U+1F1FF"},
	"TC": {"🇹🇨", "U+1F1F9 U+1F1E8"},
	"TD": {"🇹🇩", "U+1F1F9 U+1F1E9"},
	"TF": {"🇹🇫", "U+1F1F9 U+1F1EB"},
	"TG": {"🇹🇬", "U+1F1F9 U+1F1EC"},
	"TH": {"🇹🇭", "U+1F1F9 U+1F1ED"},
	"TJ": {"🇹🇯", "U+1F1F9 U+1F1EF"},
	"TK": {"🇹🇰", "U+1F1F9 U+1F1F0"},
	"TL": {"🇹🇱", "U+1F1F9 U+1F1F1"},
	"TM": {"🇹🇲", "U+1F1F9 U+1F1F2"},
	"TN": {"🇹🇳", "U+1F1F9 U+1F1F3"},
	"TO": {"🇹🇴", "U+1F1F9 U+1F1F4"},
	"TR": {"🇹🇷", "U+1F1F9 U+1F1F7"},
	"TT": {"🇹🇹", "U+1F1F9 U+1F1F9"},
	"TV": {"🇹🇻", "U+1F1F9 U+1F1FB"},
	"TW": {"🇹🇼", "U+1F1F9 U+1F1FC"},
	"TZ": {"🇹🇿", "U+1F1F9 U+1F1FF"},
	"UA": {"🇺🇦", "U+1F1FA U+1F1E6"},
	"UG": {"🇺🇬", "U+1F1FA U+1F1EC"},
	"UM": {"🇺🇲", "U+1F1FA U+1F1F2"},
	"US": {"🇺🇸", "U+1F1FA U+1F1F8"},
	"UY": {"🇺🇾", "U+1F1FA U+1F1FE"},
	"UZ": {"🇺🇿", "U+1F1FA U+1F1FF"},
	"VA": {"🇻🇦", "U+1F1FB U+1F1E6"},
	"VC": {"🇻🇨", "U+1F1FB U+1F1E8"},
	"VE": {"🇻🇪", "U+1F1FB U+1F1EA"},
	"VG": {"🇻🇬", "U+1F1FB U+1F1EC"},
	"VI": {"🇻🇮", "U+1F1FB U+1F1EE"},
	"VN": {"🇻🇳", "U+1F1FB U+1F1F3"},
	"VU": {"🇻🇺", "U+1F1FB U+1F1FA"},
	"WF": {"🇼🇫", "U+1F1FC U+1F1EB"},
	"WS": {"🇼🇸", "U+1F1FC U+1F1F8"},
	"XK": {"🇽🇰", "U+1F1FD U+1F1F0"},
	"YE": {"🇾🇪", "U+1F1FE U+1F1EA"},
	"YT": {"🇾🇹", "U+1F1FE U+1F1F9"},
	"ZA": {"🇿🇦", "U+1F1FF U+1F1E6"},
	"ZM": {"🇿🇲", "U+1F1FF U+1F1F2"},
	"ZW": {"🇿🇼", "U+1F1FF U+1F1FC"},
}

var countriesCurrencies = map[string]CountryCurrency{
	"AD": {"EUR", "€"},
	"AE": {"AED", "د.إ"},
	"AF": {"AFN", "؋"},
	"AG": {"XCD", "$"},
	"AI": {"XCD", "$"},
	"AL": {"ALL", "L"},
	"AM": {"AMD", "֏"},
	"AO": {"AOA", "Kz"},
	"AQ": {"", "$"},
	"AR": {"ARS", "$"},
	"AS": {"USD", "$"},
	"AT": {"EUR", "€"},
	"AU": {"AUD", "$"},
	"AW": {"AWG", "ƒ"},
	"AX": {"EUR", "€"},
	"AZ": {"AZN", "₼"},
	"BA": {"BAM", "KM"},
	"BB": {"BBD", "$"},
	"BD": {"BDT", "৳"},
	"BE": {"EUR", "€"},
	"BF": {"XOF", "CFA"},
	"BG": {"BGN", "лв"},
	"BH": {"BHD", ".د.ب"},
	"BI": {"BIF", "FBu"},
	"BJ": {"XOF", "CFA"},
	"BL": {"EUR", "€"},
	"BM": {"BMD", "$"},
	"BN": {"BND", "$"},
	"BO": {"BOB", "$b"},
	"BQ": {"USD", "$"},
	"BR": {"BRL", "R$"},
	"BS": {"BSD", "$"},
	"BT": {"BTN", "Nu."},
	"BV": {"NOK", "kr"},
	"BW": {"BWP", "P"},
	"BY": {"BYR", "Br"},
	"BZ": {"BZD", "BZ$"},
	"CA": {"CAD", "$"},
	"CC": {"AUD", "$"},
	"CD": {"CDF", "FC"},
	"CF": {"XAF", "FCFA"},
	"CG": {"XAF", "FCFA"},
	"CH": {"CHF", "CHF"},
	"CI": {"XOF", "CFA"},
	"CK": {"NZD", "$"},
	"CL": {"CLP", "$"},
	"CM": {"XAF", "FCFA"},
	"CN": {"CNY", "¥"},
	"CO": {"COP", "$"},
	"CR": {"CRC", "₡"},
	"CU": {"CUP", "₱"},
	"CV": {"CVE", "$"},
	"CW": {"ANG", "ƒ"},
	"CX": {"AUD", "$"},
	"CY": {"EUR", "€"},
	"CZ": {"CZK", "Kč"},
	"DE": {"EUR", "€"},
	"DJ": {"DJF", "Fdj"},
	"DK": {"DKK", "kr"},
	"DM": {"XCD", "$"},
	"DO": {"DOP", "RD$"},
	"DZ": {"DZD", "دج"},
	"EC": {"USD", "$"},
	"EE": {"EUR", "€"},
	"EG": {"EGP", "£"},
	"EH": {"MAD", "MAD"},
	"ER": {"ERN", "Nfk"},
	"ES": {"EUR", "€"},
	"ET": {"ETB", "Br"},
	"FI": {"EUR", "€"},
	"FJ": {"FJD", "$"},
	"FK": {"FKP", "£"},
	"FM": {"USD", "$"},
	"FO": {"DKK", "kr"},
	"FR": {"EUR", "€"},
	"GA": {"XAF", "FCFA"},
	"GB": {"GBP", "£"},
	"GD": {"XCD", "$"},
	"GE": {"GEL", "ლ"},
	"GF": {"EUR", "€"},
	"GG": {"GBP", "£"},
	"GH": {"GHS", "GH₵"},
	"GI": {"GIP", "£"},
	"GL": {"DKK", "kr"},
	"GM": {"GMD", "D"},
	"GN": {"GNF", "FG"},
	"GP": {"EUR", "€"},
	"GQ": {"XAF", "FCFA"},
	"GR": {"EUR", "€"},
	"GS": {"GBP", "£"},
	"GT": {"GTQ", "Q"},
	"GU": {"USD", "$"},
	"GW": {"XOF", "CFA"},
	"GY": {"GYD", "$"},
	"HK": {"HKD", "$"},
	"HM": {"AUD", "$"},
	"HN": {"HNL", "L"},
	"HR": {"HRK", "kn"},
	"HT": {"HTG", "G"},
	"HU": {"HUF", "Ft"},
	"ID": {"IDR", "Rp"},
	"IE": {"EUR", "€"},
	"IL": {"ILS", "₪"},
	"IM": {"GBP", "£"},
	"IN": {"INR", "₹"},
	"IO": {"USD", "$"},
	"IQ": {"IQD", "ع.د"},
	"IR": {"IRR", "﷼"},
	"IS": {"ISK", "kr"},
	"IT": {"EUR", "€"},
	"JE": {"GBP", "£"},
	"JM": {"JMD", "J$"},
	"JO": {"JOD", "JD"},
	"JP": {"JPY", "¥"},
	"KE": {"KES", "KSh"},
	"KG": {"KGS", "лв"},
	"KH": {"KHR", "៛"},
	"KI": {"AUD", "$"},
	"KM": {"KMF", "CF"},
	"KN": {"XCD", "$"},
	"KP": {"KPW", "₩"},
	"KR": {"KRW", "₩"},
	"KW": {"KWD", "KD"},
	"KY": {"KYD", "$"},
	"KZ": {"KZT", "₸"},
	"LA": {"LAK", "₭"},
	"LB": {"LBP", "£"},
	"LC": {"XCD", "$"},
	"LI": {"CHF", "CHF"},
	"LK": {"LKR", "₨"},
	"LR": {"LRD", "$"},
	"LS": {"LSL", "M"},
	"LT": {"LTL", "Lt"},
	"LU": {"EUR", "€"},
	"LV": {"EUR", "€"},
	"LY": {"LYD", "LD"},
	"MA": {"MAD", "MAD"},
	"MC": {"EUR", "€"},
	"MD": {"MDL", "lei"},
	"ME": {"EUR", "€"},
	"MF": {"EUR", "€"},
	"MG": {"MGA", "Ar"},
	"MH": {"USD", "$"},
	"MK": {"MKD", "ден"},
	"ML": {"XOF", "CFA"},
	"MM": {"MMK", "K"},
	"MN": {"MNT", "₮"},
	"MO": {"MOP", "MOP$"},
	"MP": {"USD", "$"},
	"MQ": {"EUR", "€"},
	"MR": {"MRO", "UM"},
	"MS": {"XCD", "$"},
	"MT": {"EUR", "€"},
	"MU": {"MUR", "₨"},
	"MV": {"MVR", "Rf"},
	"MW": {"MWK", "MK"},
	"MX": {"MXN", "$"},
	"MY": {"MYR", "RM"},
	"MZ": {"MZN", "MT"},
	"NA": {"NAD", "$"},
	"NC": {"XPF", "₣"},
	"NE": {"XOF", "CFA"},
	"NF": {"AUD", "$"},
	"NG": {"NGN", "₦"},
	"NI": {"NIO", "C$"},
	"NL": {"EUR", "€"},
	"NO": {"NOK", "kr"},
	"NP": {"NPR", "₨"},
	"NR": {"AUD", "$"},
	"NU": {"NZD", "$"},
	"NZ": {"NZD", "$"},
	"OM": {"OMR", "﷼"},
	"PA": {"PAB", "B/."},
	"PE": {"PEN", "S/."},
	"PF": {"XPF", "₣"},
	"PG": {"PGK", "K"},
	"PH": {"PHP", "₱"},
	"PK": {"PKR", "₨"},
	"PL": {"PLN", "zł"},
	"PM": {"EUR", "€"},
	"PN": {"NZD", "$"},
	"PR": {"USD", "$"},
	"PS": {"ILS", "₪"},
	"PT": {"EUR", "€"},
	"PW": {"USD", "$"},
	"PY": {"PYG", "Gs"},
	"QA": {"QAR", "﷼"},
	"RE": {"EUR", "€"},
	"RO": {"RON", "lei"},
	"RS": {"RSD", "Дин."},
	"RU": {"RUB", "₽"},
	"RW": {"RWF", "R₣"},
	"SA": {"SAR", "﷼"},
	"SB": {"SBD", "$"},
	"SC": {"SCR", "₨"},
	"SD": {"SDG", "ج.س."},
	"SE": {"SEK", "kr"},
	"SG": {"SGD", "S$"},
	"SH": {"SHP", "£"},
	"SI": {"EUR", "€"},
	"SJ": {"NOK", "kr"},
	"SK": {"EUR", "€"},
	"SL": {"SLL", "Le"},
	"SM": {"EUR", "€"},
	"SN": {"XOF", "CFA"},
	"SO": {"SOS", "S"},
	"SR": {"SRD", "$"},
	"SS": {"SSP", "£"},
	"ST": {"STD", "Db"},
	"SV": {"USD", "$"},
	"SX": {"ANG", "ƒ"},
	"SY": {"SYP", "£"},
	"SZ": {"SZL", "E"},
	"TC": {"USD", "$"},
	"TD": {"XAF", "FCFA"},
	"TF": {"EUR", "€"},
	"TG": {"XOF", "CFA"},
	"TH": {"THB", "฿"},
	"TJ": {"TJS", "SM"},
	"TK": {"NZD", "$"},
	"TL": {"USD", "$"},
	"TM": {"TMT", "T"},
	"TN": {"TND", "د.ت"},
	"TO": {"TOP", "T$"},
	"TR": {"TRY", "₺"},
	"TT": {"TTD", "TT$"},
	"TV": {"AUD", "$"},
	"TW": {"TWD", "NT$"},
	"TZ": {"TZS", "TSh"},
	"UA": {"UAH", "₴"},
	"UG": {"UGX", "USh"},
	"UM": {"USD", "$"},
	"US": {"USD", "$"},
	"UY": {"UYU", "$U"},
	"UZ": {"UZS", "лв"},
	"VA": {"EUR", "€"},
	"VC": {"XCD", "$"},
	"VE": {"VEF", "Bs"},
	"VG": {"USD", "$"},
	"VI": {"USD", "$"},
	"VN": {"VND", "₫"},
	"VU": {"VUV", "VT"},
	"WF": {"XPF", "₣"},
	"WS": {"WST", "WS$"},
	"XK": {"EUR", "€"},
	"YE": {"YER", "﷼"},
	"YT": {"EUR", "€"},
	"ZA": {"ZAR", "R"},
	"ZM": {"ZMK", "ZK"},
	"ZW": {"ZWL", "$"},
}
