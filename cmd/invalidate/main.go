// Invalidate one or more URIs from a CloudFront distribution.
// For example:
// 	$> ./bin/invalidate -cloudfront-dsn 'region=us-west-2 credentials=session' -distribution-id {DISTRIBUTION_ID} /data/151/194/395/1/1511943951.geojson
package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aaronland/go-aws-cloudfront"
	"github.com/aws/aws-sdk-go/aws"
	aws_cloudfront "github.com/aws/aws-sdk-go/service/cloudfront"
	"log"
	"os"
)

func main() {

	cloudfront_dsn := flag.String("cloudfront-dsn", "", "A valid aaronland/go-aws-session URI. Required elements are: region, credentials.")
	distribution_id := flag.String("distribution-id", "", "A valid AWS CloudFront distribution ID.")
	ref := flag.String("reference", "", "An optional reference label to assign to your invaidation. If empty a reference will be derived from the URIs passed to the tool.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Invalidate one or more URIs from a CloudFront distribution.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s uri(N) uri(N)\n", os.Args[0])
		flag.PrintDefaults()
	}
	
	flag.Parse()

	ctx := context.Background()

	svc, err := cloudfront.GetServiceWithDSN(ctx, *cloudfront_dsn)

	if err != nil {
		log.Fatalf("Failed to create service, %v", err)
	}

	uris := flag.Args()
	count := len(uris)

	items := make([]*string, count)

	for idx, u := range uris {
		items[idx] = aws.String(u)
	}

	paths := &aws_cloudfront.Paths{
		Items:    items,
		Quantity: aws.Int64(int64(count)),
	}

	if *ref == "" {

		enc_items, err := json.Marshal(items)

		if err != nil {
			log.Fatalf("Failed to encode items used to derive reference, %v", err)
		}

		sum := sha256.Sum256(enc_items)
		*ref = fmt.Sprintf("%x", sum)
	}

	batch := &aws_cloudfront.InvalidationBatch{
		CallerReference: aws.String(*ref),
		Paths:           paths,
	}

	input := &aws_cloudfront.CreateInvalidationInput{
		DistributionId:    aws.String(*distribution_id),
		InvalidationBatch: batch,
	}

	rsp, err := svc.CreateInvalidation(input)

	if err != nil {
		log.Fatalf("Failed to create invalidation, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(rsp)

	if err != nil {
		log.Fatalf("Failed to encode response, %v", err)
	}

}
