// Invalidate one or more URIs from a CloudFront distribution.
// For example:
//
//	$> ./bin/invalidate -cloudfront-dsn 'region=us-west-2 credentials=session' -distribution-id {DISTRIBUTION_ID} /data/151/194/395/1/1511943951.geojson
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/go-aws-cloudfront"
	"log"
	"os"
)

func main() {

	client_uri := flag.String("client-uri", "", "")
	distribution_id := flag.String("distribution-id", "", "A valid AWS CloudFront distribution ID.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Invalidate one or more URIs from a CloudFront distribution.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s uri(N) uri(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	ctx := context.Background()

	cl, err := cloudfront.NewClientWithURI(ctx, *client_uri)

	if err != nil {
		log.Fatalf("Failed to create client, %v", err)
	}

	uris := flag.Args()

	ref, err := cloudfront.InvalidatePaths(ctx, cl, *distribution_id, uris...)

	if err != nil {
		log.Fatalf("Failed to invalidate paths, %v", err)
	}

	fmt.Println(ref)
}
