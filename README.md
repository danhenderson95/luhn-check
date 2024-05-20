# luhn-check

A REST API that takes in a credit card number and checks if the card number is valid according to the Luhn algorithm.

https://en.wikipedia.org/wiki/Luhn_algorithm

## How to use

```
go run main.go
```
```
curl http://localhost:3333/\?value=xxx
```
where "xxx" is the credit card number to test.

## Response

The API returns a JSON object with the shape:
```
{"passes": result}
```
where "result" is a bool

## Example

```
curl http://localhost:3333/\?value=4242424242424242
```
```
{"passes":true}
```
