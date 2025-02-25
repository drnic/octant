/*
 * Copyright (c) 2019 VMware, Inc. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

 package errors

 import (
	 "fmt"
	 "time"
 
	 "github.com/google/uuid"
 )
 
 type GenericError struct {
	 id          string
	 timestamp   time.Time
	 err         error
 }
 
 var _ InternalError = (*GenericError)(nil)
 
 func NewGenericError(err error) *GenericError {
	 id, _ := uuid.NewUUID()
 
	 return &GenericError{
		 err:         err,
		 timestamp:   time.Now(),
		 id:          id.String(),
	 }
 }
 
 // ID returns the error unique ID.
 func (o *GenericError) ID() string {
	 return o.id
 }
 
 // Timestamp returns the error timestamp.
 func (o *GenericError) Timestamp() time.Time {
	 return o.timestamp
 }
 
 // Error returns an error string.
 func (o *GenericError) Error() string {
	 return fmt.Sprintf("%s", o.err)
 }