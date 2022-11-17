# TomTomCheck
The TomTom API is another version of the clasic Google Maps API, it can be used to do a wide ranges of different API calls. In this tool we check most of these API calls if we can perform them, the only downside is for some API calls there is account specific information required like history which is imposible to test via a default API query.

## Usage
```
go run .\main.go <API Key>
```
## Remidation
The remidation is pretty simple, TomTom has a nice setting called "Domain Whitelisting". When enabled before making the API call TomTom will first verify if the source making the API call is an allowed source and will return a 403 forbidden if this is not the case.


