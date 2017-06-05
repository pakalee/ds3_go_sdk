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

type GetCacheFilesystemSpectraS3Request struct {
    cacheFilesystem string
    queryParams *url.Values
}

func NewGetCacheFilesystemSpectraS3Request(cacheFilesystem string) *GetCacheFilesystemSpectraS3Request {
    queryParams := &url.Values{}

    return &GetCacheFilesystemSpectraS3Request{
        cacheFilesystem: cacheFilesystem,
        queryParams: queryParams,
    }
}




func (GetCacheFilesystemSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getCacheFilesystemSpectraS3Request *GetCacheFilesystemSpectraS3Request) Path() string {
    return "/_rest_/cache_filesystem/" + getCacheFilesystemSpectraS3Request.cacheFilesystem
}

func (getCacheFilesystemSpectraS3Request *GetCacheFilesystemSpectraS3Request) QueryParams() *url.Values {
    return getCacheFilesystemSpectraS3Request.queryParams
}

func (GetCacheFilesystemSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetCacheFilesystemSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetCacheFilesystemSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}