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
)

type GetTapeDensityDirectiveSpectraS3Request struct {
    tapeDensityDirective string
    queryParams *url.Values
}

func NewGetTapeDensityDirectiveSpectraS3Request(tapeDensityDirective string) *GetTapeDensityDirectiveSpectraS3Request {
    queryParams := &url.Values{}

    return &GetTapeDensityDirectiveSpectraS3Request{
        tapeDensityDirective: tapeDensityDirective,
        queryParams: queryParams,
    }
}




func (GetTapeDensityDirectiveSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getTapeDensityDirectiveSpectraS3Request *GetTapeDensityDirectiveSpectraS3Request) Path() string {
    return "/_rest_/tape_density_directive/" + getTapeDensityDirectiveSpectraS3Request.tapeDensityDirective
}

func (getTapeDensityDirectiveSpectraS3Request *GetTapeDensityDirectiveSpectraS3Request) QueryParams() *url.Values {
    return getTapeDensityDirectiveSpectraS3Request.queryParams
}

func (GetTapeDensityDirectiveSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetTapeDensityDirectiveSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetTapeDensityDirectiveSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}