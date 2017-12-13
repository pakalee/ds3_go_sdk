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

import "log"

func (s3ObjectCachedNotificationRegistration *S3ObjectCachedNotificationRegistration) parse(xmlNode *XmlNode, aggErr *AggregateError) {

    // Parse Child Nodes
    for _, child := range xmlNode.Children {
        switch child.XMLName.Local {
        case "CreationDate":
            s3ObjectCachedNotificationRegistration.CreationDate = parseString(child.Content)
        case "Format":
            parseEnum(child.Content, &s3ObjectCachedNotificationRegistration.Format, aggErr)
        case "Id":
            s3ObjectCachedNotificationRegistration.Id = parseString(child.Content)
        case "JobId":
            s3ObjectCachedNotificationRegistration.JobId = parseNullableString(child.Content)
        case "LastFailure":
            s3ObjectCachedNotificationRegistration.LastFailure = parseNullableString(child.Content)
        case "LastHttpResponseCode":
            s3ObjectCachedNotificationRegistration.LastHttpResponseCode = parseNullableInt(child.Content, aggErr)
        case "LastNotification":
            s3ObjectCachedNotificationRegistration.LastNotification = parseNullableString(child.Content)
        case "NamingConvention":
            parseEnum(child.Content, &s3ObjectCachedNotificationRegistration.NamingConvention, aggErr)
        case "NotificationEndPoint":
            s3ObjectCachedNotificationRegistration.NotificationEndPoint = parseNullableString(child.Content)
        case "NotificationHttpMethod":
            parseEnum(child.Content, &s3ObjectCachedNotificationRegistration.NotificationHttpMethod, aggErr)
        case "NumberOfFailuresSinceLastSuccess":
            s3ObjectCachedNotificationRegistration.NumberOfFailuresSinceLastSuccess = parseInt(child.Content, aggErr)
        case "UserId":
            s3ObjectCachedNotificationRegistration.UserId = parseNullableString(child.Content)
        default:
            log.Printf("WARNING: unable to parse unknown xml tag '%s' while parsing S3ObjectCachedNotificationRegistration.", child.XMLName.Local)
        }
    }
}