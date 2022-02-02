package cloudfront

import (
	"context"
	"fmt"
	"github.com/aaronland/go-aws-session"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	aws_cloudfront "github.com/aws/aws-sdk-go/service/cloudfront"
)

// GetServiceWithDSN returns a new *aws_cloudfront.CloudFront with an aws-sdk-go `Session` instance derived from 'dsn'
func GetServiceWithDSN(ctx context.Context, dsn string) (*aws_cloudfront.CloudFront, error) {

	sess, err := session.NewSessionWithDSN(dsn)

	if err != nil {
		return nil, fmt.Errorf("Failed to create session, %w", err)
	}

	return GetServiceWithSession(ctx, sess)
}

// GetServiceWithDSN returns a new *aws_cloudfront.CloudFront with 'sess'
func GetServiceWithSession(ctx context.Context, sess *aws_session.Session) (*aws_cloudfront.CloudFront, error) {
	svc := aws_cloudfront.New(sess)
	return svc, nil
}
