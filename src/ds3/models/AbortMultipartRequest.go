package models

import (
    "net/url"
    "net/http"
    "ds3/net"
)

type AbortMultipartRequest struct {
    bucketName, objectName string
    uploadId string
}

func NewAbortMultipartRequest(bucketName, objectName, uploadId string) *AbortMultipartRequest {
    return &AbortMultipartRequest{
        bucketName: bucketName,
        objectName: objectName,
        uploadId: uploadId,
    }
}

func (AbortMultipartRequest) Verb() net.HttpVerb {
    return net.DELETE
}

func (self AbortMultipartRequest) Path() string {
    return "/" + self.bucketName + "/" + self.objectName
}

func (self AbortMultipartRequest) QueryParams() *url.Values {
    return &url.Values{"uploadId": []string{self.uploadId}}
}

func (AbortMultipartRequest) Header() *http.Header {
    return &http.Header{}
}

func (AbortMultipartRequest) GetContentStream() net.SizedReadCloser {
    return nil
}

