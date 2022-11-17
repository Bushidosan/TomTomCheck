// Security audit for TomTom API keys, performs a variaty of different checks.
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func main() {
	figure.NewFigure("TomTom Check", "", true).Print()

	if len(os.Args) != 2 {
		fmt.Println("APIKey is required: TomTomCheck <API Key>")
		os.Exit(0)
	}

	APIKey := os.Args[1]
	color.Yellow("\n\nAPI Key used is: %s \n\n", APIKey)

	APICalls := map[string]string{
		"Map Display API":                       "https://a.api.tomtom.com/map/1/tile/basic/main/1/0/0.png?key=" + APIKey,
		"Routing API":                           "https://api.tomtom.com/routing/1/calculateRoute/52.50931,13.42936:52.50274,13.43872/json?instructionsType=text&language=en-US&vehicleHeading=90&sectionType=traffic&report=effectiveSettings&routeType=eco&traffic=true&avoid=unpavedRoads&travelMode=car&vehicleMaxSpeed=120&vehicleCommercial=false&vehicleEngineType=combustion&key=" + APIKey,
		"Matrix Routing v2 API":                 "https://api.tomtom.com/routing/matrix/2/async/00-00000000-0000-0000-0000-000000000000-0000?key=" + APIKey,
		"Search API":                            "https://api.tomtom.com/search/2/search/36.98844,-121.97483.json?key=" + APIKey,
		"Geocodoing API":                        "https://api.tomtom.com/search/2/geocode/De Ruijterkade 154, 1011 AC, Amsterdam.json?key=" + APIKey,
		"Reverse Geocoding API":                 "https://api.tomtom.com/search/2/reverseGeocode/52.157831,5.223776.json?radius=100&key=" + APIKey,
		"Batch Search API":                      "https://api.tomtom.com/search/2/batch/45e0909c-625a-4822-a060-8f7f88498c0e?key=" + APIKey,
		"EV Charging Stations Availability API": "https://api.tomtom.com/search/2/chargingAvailability.json?chargingAvailability=00112233-4455-6677-8899-aabbccddeeff&connectorSet=IEC62196Type2CableAttached&minPowerKW=22.2&maxPowerKW=43.2&key=" + APIKey,
		"Traffic API":                           "https://api.tomtom.com/traffic/services/4/incidentViewport/-939584.4813015489,-23954526.723651607,14675583.153020501,25043442.895825107/2/-939584.4813015489,-23954526.723651607,14675583.153020501,25043442.895825107/2/true/xml?key=" + APIKey,
		"Geofencing API":                        "https://api.tomtom.com/geofencing/1/projects/44de824d-c368-46cf-a234-a6792682dfd6/fences?key=" + APIKey,
	}

	// Domain Whitelisting check, if enabled exit.
	color.Blue("Testing for Domain Whitelisting\n\n")
	resp, _ := http.Get("https://api.tomtom.com/map/1/copyrights/caption.json?key=" + APIKey)
	if resp.StatusCode == http.StatusForbidden {
		color.Green("[-] API Key not vulnerable, Domain Whitelisting is enabled")
		os.Exit(1)
	} else {
		color.Red("[!] Domain Whitelisting is not enabled, testing API endpoints.\n")
	}

	// Testing different categories of API endpoints, documentation mentioned that these are the different option someone can enable/disable.
	for name, request := range APICalls {
		fmt.Printf("\n\n[-] Testing: %s\n", name)
		resp, _ := http.Get(request)
		if resp.StatusCode == http.StatusForbidden {
			color.Green("\n[-] API Key is not vulnerable for: " + name)
		} else {
			color.Red("\n[!] API Key is vulnerable for: " + name)
		}
	}
}
