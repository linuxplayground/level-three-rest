[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)]()

## This is still under development

# go-level-three-rest
This is a example of how to build a REST API which has all functions required to reach the "glory of REST" as described by [Martin Fowler](http://martinfowler.com/articles/richardsonMaturityModel.html).

The goal is to ensure that standards are followed were possible and to have a DRY implementation.

## Dependency Management
We are using glide to manage all dependencies. To install all the dependencies run `glide install`

## The Router
We are using the [vestigo](https://github.com/husobee/vestigo) router, it is not the fastest but it is RFC 2616 compliant.
 - Supports CORS
 - Supports OPTIONS
 - Ability to add Location header to 201s using http.ResponseWriter.Header().Add(...)
