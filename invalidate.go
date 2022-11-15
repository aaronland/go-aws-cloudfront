package cloudfront

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	aws_cloudfront "github.com/aws/aws-sdk-go/service/cloudfront"
)

// InvalidatePaths will issue a "CreateInvalidation" request for 'uris' in 'distribution_id'. It will return the
// invalidation ID and caller reference associated with the request.
func InvalidatePaths(ctx context.Context, svc *aws_cloudfront.CloudFront, distribution_id string, uris ...string) (string, string, error) {

	count := len(uris)

	items := make([]*string, count)

	for idx, u := range uris {
		items[idx] = aws.String(u)
	}

	paths := &aws_cloudfront.Paths{
		Items:    items,
		Quantity: aws.Int64(int64(count)),
	}

	enc_items, err := json.Marshal(items)

	if err != nil {
		return "", "", fmt.Errorf("Failed to encode items used to derive reference, %v", err)
	}

	sum := sha256.Sum256(enc_items)
	ref := fmt.Sprintf("%x", sum)

	batch := &aws_cloudfront.InvalidationBatch{
		CallerReference: aws.String(ref),
		Paths:           paths,
	}

	input := &aws_cloudfront.CreateInvalidationInput{
		DistributionId:    aws.String(distribution_id),
		InvalidationBatch: batch,
	}

	rsp, err := svc.CreateInvalidation(input)

	if err != nil {
		return "", "", fmt.Errorf("Failed to create invalidation, %v", err)
	}

	id := *rsp.Invalidation.Id

	return id, ref, nil
}
