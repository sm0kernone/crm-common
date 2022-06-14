package timezone

import "strings"

type CountryTimezone struct {
	Country  string
	Timezone string
}

var CountryTimeZones = []CountryTimezone{
	{
		Country:  "Abkhazia",
		Timezone: "Europe/Moscow",
	},
	{
		Country:  "Aruba",
		Timezone: "America/Aruba",
	},
	{
		Country:  "Afghanistan",
		Timezone: "Asia/Kabul",
	},
	{
		Country:  "Angola",
		Timezone: "Africa/Luanda",
	},
	{
		Country:  "Anguilla",
		Timezone: "America/Anguilla",
	},
	{
		Country:  "Aland Islands",
		Timezone: "Europe/Mariehamn",
	},
	{
		Country:  "Albania",
		Timezone: "Europe/Tirane",
	},
	{
		Country:  "Andorra",
		Timezone: "Europe/Andorra",
	},
	{
		Country:  "United Arab Emirates",
		Timezone: "Asia/Dubai",
	},
	{
		Country:  "Argentina",
		Timezone: "America/Argentina/Buenos_Aires",
	},
	{
		Country:  "Armenia",
		Timezone: "Asia/Yerevan",
	},
	{
		Country:  "American Samoa",
		Timezone: "Pacific/Pago_Pago",
	},
	{
		Country:  "Antarctica",
		Timezone: "Antarctica/McMurdo",
	},
	{
		Country:  "French Southern and Antarctic Lands",
		Timezone: "Indian/Kerguelen",
	},
	{
		Country:  "Antigua and Barbuda",
		Timezone: "America/Antigua",
	},
	{
		Country:  "Australia",
		Timezone: "Australia/Lord_Howe",
	},
	{
		Country:  "Austria",
		Timezone: "Europe/Vienna",
	},
	{
		Country:  "Azerbaijan",
		Timezone: "Asia/Baku",
	},
	{
		Country:  "Antigua and Barbuda",
		Timezone: "America/Antigua",
	},
	{
		Country:  "Burundi",
		Timezone: "Africa/Bujumbura",
	},
	{
		Country:  "Belgium",
		Timezone: "Europe/Brussels",
	},
	{
		Country:  "Benin",
		Timezone: "Africa/Porto-Novo",
	},
	{
		Country:  "Burkina Faso",
		Timezone: "Africa/Ouagadougou",
	},
	{
		Country:  "Bangladesh",
		Timezone: "Asia/Dhaka",
	},
	{
		Country:  "Bulgaria",
		Timezone: "Europe/Sofia",
	},
	{
		Country:  "Bahrain",
		Timezone: "Asia/Bahrain",
	},
	{
		Country:  "Bahamas",
		Timezone: "America/Nassau",
	},
	{
		Country:  "Bosnia And Herzegovina",
		Timezone: "Europe/Sarajevo",
	},
	{
		Country:  "Saint Barthélemy",
		Timezone: "America/St_Barthelemy",
	},
	{
		Country:  "Belarus",
		Timezone: "Europe/Minsk",
	},
	{
		Country:  "Belize",
		Timezone: "America/Belize",
	},
	{
		Country:  "Bermuda",
		Timezone: "Atlantic/Bermuda",
	},
	{
		Country:  "Bolivia",
		Timezone: "America/La_Paz",
	},
	{
		Country:  "Brazil",
		Timezone: "America/Noronha",
	},
	{
		Country:  "Barbados",
		Timezone: "America/Barbados",
	},
	{
		Country:  "Brunei Darussalam",
		Timezone: "Asia/Brunei",
	},
	{
		Country:  "Bhutan",
		Timezone: "Asia/Thimphu",
	},
	{
		Country:  "Bouvet Island",
		Timezone: "Europe/Oslo",
	},
	{
		Country:  "Botswana",
		Timezone: "Africa/Gaborone",
	},
	{
		Country:  "Central African Republic",
		Timezone: "Africa/Bangui",
	},
	{
		Country:  "Canada",
		Timezone: "America/St_Johns",
	},
	{
		Country:  "Cocos (Keeling) Islands",
		Timezone: "Indian/Cocos",
	},
	{
		Country:  "Switzerland",
		Timezone: "Europe/Zurich",
	},
	{
		Country:  "Chile",
		Timezone: "America/Santiago",
	},
	{
		Country:  "China",
		Timezone: "Asia/Shanghai",
	},
	{
		Country:  "Ivory Coast",
		Timezone: "Africa/Abidjan",
	},
	{
		Country:  "Cameroon",
		Timezone: "Africa/Douala",
	},
	{
		Country:  "Congo (DRC)",
		Timezone: "Africa/Kinshasa",
	},
	{
		Country:  "Congo",
		Timezone: "Africa/Brazzaville",
	},
	{
		Country:  "Cook Islands",
		Timezone: "Pacific/Rarotonga",
	},
	{
		Country:  "Colombia",
		Timezone: "America/Bogota",
	},
	{
		Country:  "Comoros",
		Timezone: "Indian/Comoro",
	},
	{
		Country:  "Cape Verde",
		Timezone: "Atlantic/Cape_Verde",
	},
	{
		Country:  "Costa Rica",
		Timezone: "America/Costa_Rica",
	},
	{
		Country:  "Cuba",
		Timezone: "America/Havana",
	},
	{
		Country:  "Curaçao",
		Timezone: "America/Curacao",
	},
	{
		Country:  "Christmas Island",
		Timezone: "Indian/Christmas",
	},
	{
		Country:  "Cayman Islands",
		Timezone: "America/Cayman",
	},
	{
		Country:  "Cyprus",
		Timezone: "Asia/Nicosia",
	},
	{
		Country:  "Czech Republic",
		Timezone: "Europe/Prague",
	},
	{
		Country:  "Germany",
		Timezone: "Europe/Berlin",
	},
	{
		Country:  "Djibouti",
		Timezone: "Africa/Djibouti",
	},
	{
		Country:  "Dominica",
		Timezone: "America/Dominica",
	},
	{
		Country:  "Denmark",
		Timezone: "Europe/Copenhagen",
	},
	{
		Country:  "Dominican Republic",
		Timezone: "America/Santo_Domingo",
	},
	{
		Country:  "Algeria",
		Timezone: "Africa/Algiers",
	},
	{
		Country:  "Ecuador",
		Timezone: "America/Guayaquil",
	},
	{
		Country:  "Egypt",
		Timezone: "Africa/Cairo",
	},
	{
		Country:  "Eritrea",
		Timezone: "Africa/Asmara",
	},
	{
		Country:  "Western Sahara",
		Timezone: "Africa/El_Aaiun",
	},
	{
		Country:  "Spain",
		Timezone: "Europe/Madrid",
	},
	{
		Country:  "Estonia",
		Timezone: "Europe/Tallinn",
	}, {
		Country:  "Ethiopia",
		Timezone: "Africa/Addis_Ababa",
	},
	{
		Country:  "Finland",
		Timezone: "Europe/Helsinki",
	},
	{
		Country:  "Fiji",
		Timezone: "Pacific/Fiji",
	},
	{
		Country:  "Falkland Islands",
		Timezone: "Atlantic/Stanley",
	},
	{
		Country:  "France",
		Timezone: "Europe/Paris",
	},
	{
		Country:  "Faroe Islands",
		Timezone: "Atlantic/Faroe",
	},
	{
		Country:  "Micronesia",
		Timezone: "Pacific/Chuuk",
	},
	{
		Country:  "Gabon",
		Timezone: "Africa/Libreville",
	},
	{
		Country:  "United Kingdom",
		Timezone: "Europe/London",
	},
	{
		Country:  "Georgia",
		Timezone: "Asia/Tbilisi",
	},
	{
		Country:  "Guernsey",
		Timezone: "Europe/Guernsey",
	},
	{
		Country:  "Ghana",
		Timezone: "Africa/Accra",
	},
	{
		Country:  "Gibraltar",
		Timezone: "Europe/Gibraltar",
	},
	{
		Country:  "Guinea",
		Timezone: "Africa/Conakry",
	},
	{
		Country:  "Guadeloupe",
		Timezone: "America/Guadeloupe",
	},
	{
		Country:  "Gambia",
		Timezone: "Africa/Banjul",
	},
	{
		Country:  "Guinea-Bissau",
		Timezone: "Africa/Bissau",
	},
	{
		Country:  "Equatorial Guinea",
		Timezone: "Africa/Malabo",
	},
	{
		Country:  "Greece",
		Timezone: "Europe/Athens",
	},
	{
		Country:  "Grenada",
		Timezone: "America/Grenada",
	},
	{
		Country:  "Greenland",
		Timezone: "America/Godthab",
	},
	{
		Country:  "Guatemala",
		Timezone: "America/Guatemala",
	},
	{
		Country:  "French Guiana",
		Timezone: "America/Cayenne",
	},
	{
		Country:  "Guam",
		Timezone: "Pacific/Guam",
	},
	{
		Country:  "Guyana",
		Timezone: "America/Guyana",
	},
	{
		Country:  "Hong Kong",
		Timezone: "Asia/Hong_Kong",
	},
	{
		Country:  "Honduras",
		Timezone: "America/Tegucigalpa",
	},
	{
		Country:  "Croatia",
		Timezone: "Europe/Zagreb",
	},
	{
		Country:  "Haiti",
		Timezone: "America/Port-au-Prince",
	},
	{
		Country:  "Hungary",
		Timezone: "Europe/Budapest",
	},
	{
		Country:  "Indonesia",
		Timezone: "Asia/Jakarta",
	},
	{
		Country:  "Isle Of Man",
		Timezone: "Europe/Isle_of_Man",
	},
	{
		Country:  "India",
		Timezone: "Asia/Kolkata",
	},
	{
		Country:  "British Indian Ocean Territory",
		Timezone: "Indian/Chagos",
	},
	{
		Country:  "Ireland",
		Timezone: "Europe/Dublin",
	},
	{
		Country:  "Iran",
		Timezone: "Asia/Tehran",
	},
	{
		Country:  "Iraq",
		Timezone: "Asia/Baghdad",
	},
	{
		Country:  "Iceland",
		Timezone: "Atlantic/Reykjavik",
	},
	{
		Country:  "Israel",
		Timezone: "Asia/Jerusalem",
	},
	{
		Country:  "Italy",
		Timezone: "Europe/Rome",
	},
	{
		Country:  "Jamaica",
		Timezone: "America/Jamaica",
	},
	{
		Country:  "Jersey",
		Timezone: "Europe/Jersey",
	},
	{
		Country:  "Jordan",
		Timezone: "Asia/Amman",
	},
	{
		Country:  "Japan",
		Timezone: "Asia/Tokyo",
	},
	{
		Country:  "Kazakhstan",
		Timezone: "Asia/Almaty",
	},
	{
		Country:  "Kenya",
		Timezone: "Africa/Nairobi",
	},
	{
		Country:  "Kyrgyzstan",
		Timezone: "Asia/Bishkek",
	},
	{
		Country:  "Cambodia",
		Timezone: "Asia/Phnom_Penh",
	},
	{
		Country:  "Kiribati",
		Timezone: "Pacific/Tarawa",
	},
	{
		Country:  "Saint Kitts And Nevis",
		Timezone: "America/St_Kitts",
	},
	{
		Country:  "Korea (South)",
		Timezone: "Asia/Seoul",
	},
	{
		Country:  "Kosovo",
		Timezone: "Europe/Belgrade",
	},
	{
		Country:  "Kuwait",
		Timezone: "Asia/Kuwait",
	},
	{
		Country:  "Laos",
		Timezone: "Asia/Vientiane",
	},
	{
		Country:  "Lebanon",
		Timezone: "Asia/Beirut",
	},
	{
		Country:  "Liberia",
		Timezone: "Africa/Monrovia",
	},
	{
		Country:  "Libya",
		Timezone: "Africa/Tripoli",
	},
	{
		Country:  "Saint Lucia",
		Timezone: "America/St_Lucia",
	},
	{
		Country:  "Liechtenstein",
		Timezone: "Europe/Vaduz",
	},
	{
		Country:  "Sri Lanka",
		Timezone: "Asia/Colombo",
	},
	{
		Country:  "Lesotho",
		Timezone: "Africa/Maseru",
	},
	{
		Country:  "Lithuania",
		Timezone: "Europe/Vilnius",
	},
	{
		Country:  "Luxembourg",
		Timezone: "Europe/Luxembourg",
	},
	{
		Country:  "Latvia",
		Timezone: "Europe/Riga",
	},
	{
		Country:  "Macau",
		Timezone: "Asia/Macau",
	},
	{
		Country:  "Saint Martin",
		Timezone: "America/Marigot",
	},
	{
		Country:  "Morocco",
		Timezone: "Africa/Casablanca",
	},
	{
		Country:  "Monaco",
		Timezone: "Europe/Monaco",
	},
	{
		Country:  "Moldova",
		Timezone: "Europe/Chisinau",
	},
	{
		Country:  "Madagascar",
		Timezone: "Indian/Antananarivo",
	},
	{
		Country:  "Maldives",
		Timezone: "Indian/Maldives",
	},
	{
		Country:  "Mexico",
		Timezone: "America/Mexico_City",
	},
	{
		Country:  "Marshall Islands",
		Timezone: "Pacific/Majuro",
	},
	{
		Country:  "Macedonia",
		Timezone: "Europe/Skopje",
	},
	{
		Country:  "Mali",
		Timezone: "Africa/Bamako",
	},
	{
		Country:  "Malta",
		Timezone: "Europe/Malta",
	},
	{
		Country:  "Myanmar",
		Timezone: "Asia/Rangoon",
	},
	{
		Country:  "Montenegro",
		Timezone: "Europe/Podgorica",
	},
	{
		Country:  "Mongolia",
		Timezone: "Asia/Ulaanbaatar",
	},
	{
		Country:  "Northern Mariana Islands",
		Timezone: "Pacific/Saipan",
	},
	{
		Country:  "Mozambique",
		Timezone: "Africa/Maputo",
	},
	{
		Country:  "Mauritania",
		Timezone: "Africa/Nouakchott",
	},
	{
		Country:  "Montserrat",
		Timezone: "America/Montserrat",
	},
	{
		Country:  "Martinique",
		Timezone: "America/Martinique",
	},
	{
		Country:  "Mauritius",
		Timezone: "Indian/Mauritius",
	},
	{
		Country:  "Malawi",
		Timezone: "Africa/Blantyre",
	},
	{
		Country:  "Malaysia",
		Timezone: "Asia/Kuala_Lumpur",
	},
	{
		Country:  "Mayotte",
		Timezone: "Indian/Mayotte",
	},
	{
		Country:  "Namibia",
		Timezone: "Africa/Windhoek",
	},
	{
		Country:  "New Caledonia",
		Timezone: "Pacific/Noumea",
	},
	{
		Country:  "Niger",
		Timezone: "Africa/Niamey",
	},
	{
		Country:  "Norfolk Island",
		Timezone: "Pacific/Norfolk",
	},
	{
		Country:  "Nigeria",
		Timezone: "Africa/Lagos",
	},
	{
		Country:  "Nicaragua",
		Timezone: "America/Managua",
	},
	{
		Country:  "Niue",
		Timezone: "Pacific/Niue",
	},
	{
		Country:  "Netherlands",
		Timezone: "Europe/Amsterdam",
	},
	{
		Country:  "Netherlands Antilles",
		Timezone: "America/Winnipeg",
	},
	{
		Country:  "Norway",
		Timezone: "Europe/Oslo",
	},
	{
		Country:  "Nepal",
		Timezone: "Asia/Kathmandu",
	},
	{
		Country:  "Nauru",
		Timezone: "Pacific/Nauru",
	},
	{
		Country:  "New Zealand",
		Timezone: "Pacific/Auckland",
	},
	{
		Country:  "Oman",
		Timezone: "Asia/Muscat",
	},
	{
		Country:  "Pakistan",
		Timezone: "Asia/Karachi",
	},
	{
		Country:  "Panama",
		Timezone: "America/Panama",
	},
	{
		Country:  "Pitcairn Islands",
		Timezone: "Pacific/Pitcairn",
	},
	{
		Country:  "Peru",
		Timezone: "America/Lima",
	},
	{
		Country:  "Philippines",
		Timezone: "Asia/Manila",
	},
	{
		Country:  "Palau",
		Timezone: "Pacific/Palau",
	},
	{
		Country:  "Papua New Guinea",
		Timezone: "Pacific/Port_Moresby",
	},
	{
		Country:  "Poland",
		Timezone: "Europe/Warsaw",
	},
	{
		Country:  "Puerto Rico",
		Timezone: "America/Puerto_Rico",
	},
	{
		Country:  "Korea (North)",
		Timezone: "Asia/Pyongyang",
	},
	{
		Country:  "Portugal",
		Timezone: "Europe/Lisbon",
	},
	{
		Country:  "Paraguay",
		Timezone: "America/Asuncion",
	},
	{
		Country:  "Palestine",
		Timezone: "Asia/Gaza",
	},
	{
		Country:  "French Polynesia",
		Timezone: "Pacific/Tahiti",
	},
	{
		Country:  "Qatar",
		Timezone: "Asia/Qatar",
	},
	{
		Country:  "Reunion",
		Timezone: "Indian/Reunion",
	},
	{
		Country:  "Romania",
		Timezone: "Europe/Bucharest",
	},
	{
		Country:  "Russian Federation",
		Timezone: "Europe/Kaliningrad",
	},
	{
		Country:  "Rwanda",
		Timezone: "Africa/Kigali",
	},
	{
		Country:  "Saudi Arabia",
		Timezone: "Asia/Riyadh",
	},
	{
		Country:  "Sudan",
		Timezone: "Africa/Khartoum",
	},
	{
		Country:  "Senegal",
		Timezone: "Africa/Dakar",
	},
	{
		Country:  "Singapore",
		Timezone: "Asia/Singapore",
	},
	{
		Country:  "South Georgia",
		Timezone: "Atlantic/South_Georgia",
	},
	{
		Country:  "Svalbard and Jan Mayen",
		Timezone: "Arctic/Longyearbyen",
	},
	{
		Country:  "Solomon Islands",
		Timezone: "Pacific/Guadalcanal",
	},
	{
		Country:  "Sierra Leone",
		Timezone: "Africa/Freetown",
	},
	{
		Country:  "El Salvador",
		Timezone: "America/El_Salvador",
	},
	{
		Country:  "San Marino",
		Timezone: "Europe/San_Marino",
	},
	{
		Country:  "Somalia",
		Timezone: "Africa/Mogadishu",
	},
	{
		Country:  "Saint Pierre And Miquelon",
		Timezone: "America/Miquelon",
	},
	{
		Country:  "Serbia",
		Timezone: "Europe/Belgrade",
	},
	{
		Country:  "Sudan (South)",
		Timezone: "Africa/Juba",
	},
	{
		Country:  "São Tomé and Príncipe",
		Timezone: "Africa/Sao_Tome",
	},
	{
		Country:  "Suriname",
		Timezone: "America/Paramaribo",
	},
	{
		Country:  "Slovakia",
		Timezone: "Europe/Bratislava",
	},
	{
		Country:  "Slovenia",
		Timezone: "Europe/Ljubljana",
	},
	{
		Country:  "Sweden",
		Timezone: "Europe/Stockholm",
	},
	{
		Country:  "Swaziland",
		Timezone: "Africa/Mbabane",
	},
	{
		Country:  "Sint Maarten",
		Timezone: "America/Lower_Princes",
	},
	{
		Country:  "Seychelles",
		Timezone: "Indian/Mahe",
	},
	{
		Country:  "Syria",
		Timezone: "Asia/Damascus",
	},
	{
		Country:  "Turks And Caicos Islands",
		Timezone: "America/Grand_Turk",
	},
	{
		Country:  "Chad",
		Timezone: "Africa/Ndjamena",
	},
	{
		Country:  "Togo",
		Timezone: "Africa/Lome",
	},
	{
		Country:  "Thailand",
		Timezone: "Asia/Bangkok",
	},
	{
		Country:  "Tajikistan",
		Timezone: "Asia/Dushanbe",
	},
	{
		Country:  "Tokelau",
		Timezone: "Pacific/Fakaofo",
	},
	{
		Country:  "Turkmenistan",
		Timezone: "Asia/Ashgabat",
	},
	{
		Country:  "East Timor",
		Timezone: "Asia/Dili",
	},
	{
		Country:  "Tonga",
		Timezone: "Pacific/Tongatapu",
	},
	{
		Country:  "Trinidad And Tobago",
		Timezone: "America/Port_of_Spain",
	},
	{
		Country:  "Tunisia",
		Timezone: "Africa/Tunis",
	},
	{
		Country:  "Turkey",
		Timezone: "Europe/Istanbul",
	}, {

		Country:  "Tuvalu",
		Timezone: "Pacific/Funafuti",
	},
	{
		Country:  "Taiwan",
		Timezone: "Asia/Taipei",
	},
	{
		Country:  "Tanzania",
		Timezone: "Africa/Dar_es_Salaam",
	},
	{
		Country:  "Uganda",
		Timezone: "Africa/Kampala",
	},
	{
		Country:  "Ukraine",
		Timezone: "Europe/Kiev",
	},
	{
		Country:  "United States Minor Outlying Islands",
		Timezone: "Pacific/Johnston",
	},
	{
		Country:  "Uruguay",
		Timezone: "America/Montevideo",
	}, {
		Country:  "United States",
		Timezone: "America/New_York",
	}, {
		Country:  "Uzbekistan",
		Timezone: "Asia/Samarkand",
	},
	{
		Country:  "Vatican City",
		Timezone: "Europe/Vatican",
	},
	{
		Country:  "Saint Vincent And the Grenadines",
		Timezone: "America/St_Vincent",
	},
	{
		Country:  "Venezuela",
		Timezone: "America/Caracas",
	},
	{
		Country:  "Virgin Islands British",
		Timezone: "America/Tortola",
	},
	{
		Country:  "United States Virgin Islands",
		Timezone: "America/St_Thomas",
	},
	{
		Country:  "Vietnam",
		Timezone: "Asia/Ho_Chi_Minh",
	},
	{
		Country:  "Vanuatu",
		Timezone: "Pacific/Efate",
	},
	{
		Country:  "Wallis and Futuna",
		Timezone: "Pacific/Wallis",
	},
	{
		Country:  "Samoa",
		Timezone: "Pacific/Apia",
	},
	{
		Country:  "Yemen",
		Timezone: "Asia/Aden",
	},
	{
		Country:  "South Africa",
		Timezone: "Africa/Johannesburg",
	},
	{
		Country:  "Zambia",
		Timezone: "Africa/Lusaka",
	},
	{
		Country:  "Zimbabwe",
		Timezone: "Africa/Harare",
	},
}

func GetTimezoneByCountry(country string) string {
	for _, countryTz := range CountryTimeZones {
		if strings.Trim(strings.ToLower(countryTz.Country), " ") == strings.Trim(strings.ToLower(country), " ") {
			return countryTz.Timezone
		}
	}

	return "Europe/London"
}
