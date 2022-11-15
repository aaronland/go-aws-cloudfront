package cloudfront

import (
	"context"
	"fmt"
	aa_session "github.com/aaronland/go-aws-session"
	aws_session "github.com/aws/aws-sdk-go/aws/session"
	aws_cloudfront "github.com/aws/aws-sdk-go/service/cloudfront"
	"net/url"
)

// NewSessionWithURI will return a new `aws-sdk-go/aws/session.Session` instance derived
// from 'uri' which is expected to take the form of:
//
//	aws://?region={AWS_REGION}&credentials={CREDENTIALS}
//
// Where '{CREDENTIALS}' is expected to be a valid `aaronland/go-aws-session` credentials
// string.
func NewSessionWithURI(ctx context.Context, uri string) (*aws_session.Session, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	q := u.Query()
	region := q.Get("region")
	credentials := q.Get("credentials")

	dsn := fmt.Sprintf("credentials=%s region=%s", credentials, region)

	sess, err := aa_session.NewSessionWithDSN(dsn)

	return sess, nil
}

// NewClientWithURI will return a new `aws-sdk-go/service/cloudfront.CloudFront` instance
// derived from 'uri' which is expected to take the form of:
//
//	aws://?region={AWS_REGION}&credentials={CREDENTIALS}
//
// Where '{CREDENTIALS}' is expected to be a valid `aaronland/go-aws-session` credentials
// string.
func NewClientWithURI(ctx context.Context, uri string) (*aws_cloudfront.CloudFront, error) {

	sess, err := NewSessionWithURI(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create session, %w", err)
	}

	client := aws_cloudfront.New(sess)
	return client, nil
}
