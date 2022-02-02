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
  -cloudfront-dsn string
    	A valid aaronland/go-aws-session URI. Required elements are: region, credentials.
  -distribution-id string
    	A valid AWS CloudFront distribution ID.
  -reference string
    	An optional reference label to assign to your invaidation. If empty a reference will be derived from the URIs passed to the tool.
```

## See also

* https://github.com/aws/aws-sdk-go/service/cloudfront
* https://github.com/aaronland/go-aws-session