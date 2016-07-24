[![Build Status](https://travis-ci.org/bhavikkumar/level-three-rest.svg?branch=master)](https://travis-ci.org/bhavikkumar/level-three-rest) [![Coverage Status](https://coveralls.io/repos/github/bhavikkumar/level-three-rest/badge.svg?branch=master)](https://coveralls.io/github/bhavikkumar/level-three-rest?branch=master)

## This is still under development

# level-three-rest
This is a example of how to build a REST API which has all functions required to reach the "glory of REST" as described by [Martin Fowler](http://martinfowler.com/articles/richardsonMaturityModel.html).

The goal is to ensure that standards are followed were possible and to have a DRY implementation.

## Dependency Management
We are using [glide](https://github.com/Masterminds/glide) to manage all dependencies.
- Run `go get -u github.com/Masterminds/glide`
- Navigate to project
- Run `glide install`
- Run `go build`

## Dependencies
We are using the [vestigo](https://github.com/husobee/vestigo) router, it is not the fastest but it is RFC 2616 compliant.

## Features
 - [x] CORS
 - [x] OPTIONS
 - [x] Location header on HTTP 201 response
  - Possible using http.ResponseWriter.Header().Add(...)
 - [x] HTTP/2
 - [x] TLS
 - [x] HSTS

### TODO
 - [ ] HATEOAS Links
 - [ ] Rate Limiting
 - [ ] Caching
 - [ ] Versioning
 - [ ] Allow client to limit fields in response
 - [ ] Search Pagination
 - [ ] Authentication/Authorisation
 - [ ] Support multiple Content-Type
 - [ ] Vendor Content-Type
 - Code clean up
