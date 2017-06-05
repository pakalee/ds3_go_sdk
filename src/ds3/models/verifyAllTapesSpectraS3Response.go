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
    "ds3/networking"
    "net/http"
)

type VerifyAllTapesSpectraS3Response struct {
    TapeFailureList *TapeFailureList
    Headers *http.Header
}

func NewVerifyAllTapesSpectraS3Response(webResponse networking.WebResponse) (*VerifyAllTapesSpectraS3Response, error) {
    expectedStatusCodes := []int { 204, 207 }

    switch code := webResponse.StatusCode(); code {
    case 204:
        return &VerifyAllTapesSpectraS3Response{Headers: webResponse.Header()}, nil
    case 207:
        var body VerifyAllTapesSpectraS3Response
        if err := readResponseBody(webResponse, &body.TapeFailureList); err != nil {
            return nil, err
        }
        body.Headers = webResponse.Header()
        return &body, nil
    default:
        return nil, buildBadStatusCodeError(webResponse, expectedStatusCodes)
    }
}