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

type GetAzureTargetReadPreferenceSpectraS3Request struct {
    azureTargetReadPreference string
    queryParams *url.Values
}

func NewGetAzureTargetReadPreferenceSpectraS3Request(azureTargetReadPreference string) *GetAzureTargetReadPreferenceSpectraS3Request {
    queryParams := &url.Values{}

    return &GetAzureTargetReadPreferenceSpectraS3Request{
        azureTargetReadPreference: azureTargetReadPreference,
        queryParams: queryParams,
    }
}




func (GetAzureTargetReadPreferenceSpectraS3Request) Verb() networking.HttpVerb {
    return networking.GET
}

func (getAzureTargetReadPreferenceSpectraS3Request *GetAzureTargetReadPreferenceSpectraS3Request) Path() string {
    return "/_rest_/azure_target_read_preference/" + getAzureTargetReadPreferenceSpectraS3Request.azureTargetReadPreference
}

func (getAzureTargetReadPreferenceSpectraS3Request *GetAzureTargetReadPreferenceSpectraS3Request) QueryParams() *url.Values {
    return getAzureTargetReadPreferenceSpectraS3Request.queryParams
}

func (GetAzureTargetReadPreferenceSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (GetAzureTargetReadPreferenceSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (GetAzureTargetReadPreferenceSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}