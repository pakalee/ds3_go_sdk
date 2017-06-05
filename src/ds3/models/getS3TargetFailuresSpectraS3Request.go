// Copyright 2014-2017 Spectra Logic Corporation. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License"). You may not use
// this file except in compliance with the License. A copy of the License is located at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file.
// This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// This code is auto-generated, do not modify

package models

import (
    "net/url"
    "net/http"
    "ds3/networking"
    "strconv"
)

type GetS3TargetFailuresSpectraS3Request struct {
    errorMessage *string
    pageLength int
    pageOffset int
    pageStartMarker string
    targetFailureType TargetFailureType
    targetId string
    queryParams *url.Values
}

func NewGetS3TargetFailuresSpectraS3Request() *GetS3TargetFailuresSpectraS3Request {
    queryParams := &url.Values{}

    return &GetS3TargetFailuresSpectraS3Request{
        queryParams: queryParams,
    }
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithPageLength(pageLength int) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.pageLength = pageLength
    getS3TargetFailuresSpectraS3Request.queryParams.Set("page_length", strconv.Itoa(pageLength))
    return getS3TargetFailuresSpectraS3Request
}
func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithPageOffset(pageOffset int) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.pageOffset = pageOffset
    getS3TargetFailuresSpectraS3Request.queryParams.Set("page_offset", strconv.Itoa(pageOffset))
    return getS3TargetFailuresSpectraS3Request
}
func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithPageStartMarker(pageStartMarker string) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.pageStartMarker = pageStartMarker
    getS3TargetFailuresSpectraS3Request.queryParams.Set("page_start_marker", pageStartMarker)
    return getS3TargetFailuresSpectraS3Request
}
func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithTargetId(targetId string) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.targetId = targetId
    getS3TargetFailuresSpectraS3Request.queryParams.Set("target_id", targetId)
    return getS3TargetFailuresSpectraS3Request
}
func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithTargetFailureType(targetFailureType TargetFailureType) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.targetFailureType = targetFailureType
    getS3TargetFailuresSpectraS3Request.queryParams.Set("type", targetFailureType.String())
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithErrorMessage(errorMessage *string) *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.errorMessage = errorMessage
    if errorMessage != nil {
        getS3TargetFailuresSpectraS3Request.queryParams.Set("error_message", *errorMessage)
    } else {
        getS3TargetFailuresSpectraS3Request.queryParams.Set("error_message", "")
    }
    return getS3TargetFailuresSpectraS3Request
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) WithLastPage() *GetS3TargetFailuresSpectraS3Request {
    getS3TargetFailuresSpectraS3Request.queryParams.Set("last_page", "")
    return getS3TargetFailuresSpectraS3Request
}

func (GetS3TargetFailuresSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) Path() string {
    return "/_rest_/s3_target_failure"
}

func (getS3TargetFailuresSpectraS3Request *GetS3TargetFailuresSpectraS3Request) QueryParams() *url.Values {
    return getS3TargetFailuresSpectraS3Request.queryParams
}

func (GetS3TargetFailuresSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetS3TargetFailuresSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetS3TargetFailuresSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}