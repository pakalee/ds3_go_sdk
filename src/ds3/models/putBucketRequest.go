package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
)

type PutBucketRequest struct {
    bucketName string
}

func NewPutBucketRequest(bucketName string) *PutBucketRequest {
    return &PutBucketRequest{
        bucketName: bucketName,
    }
}

func (PutBucketRequest) Verb() networking.HttpVerb {
    return networking.PUT
}

func (putBucketRequest *PutBucketRequest) Path() string {
    return "/" + putBucketRequest.bucketName
}

func (putBucketRequest *PutBucketRequest) QueryParams() *url.Values {
    return new(url.Values)
}

func (PutBucketRequest) Header() *http.Header {
    return &http.Header{}
}

func (PutBucketRequest) GetContentStream() networking.SizedReadCloser {
    return nil
}

func (PutBucketRequest) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
