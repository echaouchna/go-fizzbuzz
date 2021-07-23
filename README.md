# go-fizzbuzz

A more general fizzbuzz webserver

## About

### Introduction to fizzbuzz

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of `3` by `fizz`, all multiples of `5` by `buzz`, and all multiples of `15` by `fizzbuzz`.

The output would look like this: `1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...`.

### Idea behind go-fizzbuzz

The goal here is to implement a web server that will expose a REST API endpoint that:

- Accepts five parameters: three integers __int1__, __int2__ and __limit__, and two strings __str1__ and __str2__.
- Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

An additional statistics endpoint was added to allow users to know what's the most frequent request has been. This endpoint:

- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request

## Build

to compile the code:

```sh
make build
```

or to build docker image :

```sh
make build-docker
```

## Usage

### Run the server

```sh
bin/go-fizzbuzz
```

or

```sh
make run
```

or

```sh
docker run -p 10000:10000 echaouchna/go-fizzbuzz
```

### Queries

#### Print a fizzbuzz calculation result

Use curl to test (or any other client) the api

```sh
curl 'localhost:10000/calculate/100?int1=3&int2=5&str1=fizz&str2=buzz'
```

in this example the limit is a path param and it's set to 100, the limit is required

the query params int1, int2, str1, str2 are optional

the table below explains the different parameters of the endpoint

| param | param type | value type | required | default value |
|---|---|---|---|---|
| limit | path param | integer | `true` |  |
| int1 | query param | non zero positive integer | `false` | `3` |
| int2 | query param | non zero positive integer | `false` | `5` |
| str1 | query param | string | `false` | `fizz` |
| str2 | query param | string | `false` | `buzz` |

#### Stats

To print the most used request, you can use the following query:

```sh
curl localhost:10000/most_used
```

## Contribute

To run tests:

```sh
make test
```

## Limitations

- The stats works only for one instance for now
