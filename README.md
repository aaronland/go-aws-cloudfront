# go-aws-cloudfront

Go package providing opinionated tools and methods for working with the `aws-sdk-go/service/cloudfront` package.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/go-aws-cloudfront.svg)](https://pkg.go.dev/github.com/aaronland/go-aws-cloudfront)

## Tools

```
$> make cli
go build -mod vendor -o bin/invalidate cmd/invalidate/main.go
```

### invalidate

Invalidate one or more URIs from a CloudFront distribution.

```
$> ./bin/invalidate -h
Invalidate one or more URIs from a CloudFront distribution.
Usage:
	 ./bin/invalidate uri(N) uri(N)
  -client-uri string
    	A valid client URI in the form of 'aws://?region={AWS_REGION}&credentials={CREDENTIALS}' where '{CREDENTIAL}' is expected to be a valid aaronland/go-aws-session credential string.
  -distribution-id string
    	A valid AWS CloudFront distribution ID.
```

## See also

* https://github.com/aws/aws-sdk-go/service/cloudfront
* https://github.com/aaronland/go-aws-session