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

type DeleteTapeFailureNotificationRegistrationSpectraS3Request struct {
    notificationId string
    queryParams *url.Values
}

func NewDeleteTapeFailureNotificationRegistrationSpectraS3Request(notificationId string) *DeleteTapeFailureNotificationRegistrationSpectraS3Request {
    queryParams := &url.Values{}

    return &DeleteTapeFailureNotificationRegistrationSpectraS3Request{
        notificationId: notificationId,
        queryParams: queryParams,
    }
}




func (DeleteTapeFailureNotificationRegistrationSpectraS3Request) Verb() networking.HttpVerb {
    return networking.DELETE
}

func (deleteTapeFailureNotificationRegistrationSpectraS3Request *DeleteTapeFailureNotificationRegistrationSpectraS3Request) Path() string {
    return "/_rest_/tape_failure_notification_registration/" + deleteTapeFailureNotificationRegistrationSpectraS3Request.notificationId
}

func (deleteTapeFailureNotificationRegistrationSpectraS3Request *DeleteTapeFailureNotificationRegistrationSpectraS3Request) QueryParams() *url.Values {
    return deleteTapeFailureNotificationRegistrationSpectraS3Request.queryParams
}

func (DeleteTapeFailureNotificationRegistrationSpectraS3Request) GetChecksum() networking.Checksum {
    return networking.NewNoneChecksum()
}
func (DeleteTapeFailureNotificationRegistrationSpectraS3Request) Header() *http.Header {
    return &http.Header{}
}

func (DeleteTapeFailureNotificationRegistrationSpectraS3Request) GetContentStream() networking.ReaderWithSizeDecorator {
    return nil
}