# TomTomCheck
The TomTom API is another version of the clasic Google Maps API, it can be used to do a wide ranges of different API calls. In this tool we check most of these API calls if we can perform them, the only downside is for some API calls there is account specific information required like history which is imposible to test via a default API query.

## Usage
```
go run .\main.go <API Key>
```
![TomTomCheck](https://user-images.githubusercontent.com/67435928/202419434-a10b2d57-d361-434f-883e-052fc10608e4.png)

## Remidation
The remidation is pretty simple, TomTom has a nice setting called "Domain Whitelisting". When enabled before making the API call TomTom will first verify if the source making the API call is an allowed source and will return a 403 forbidden if this is not the case.

**Key take away?** : Make sure domain whitelisting is enabled otherwise bad people can waste your money.

https://developer.tomtom.com/blog/decoded/how-protect-your-api-key-using-domain-whitelisting
