/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package protocol

/////////////////////////////
// Result Impletment of RPC
/////////////////////////////

// RPCResult ...
type AsyncRPCResult struct {
	Attrs map[string]string
	Err   error
	Rest  interface{}
}

// SetError ...
func (r *AsyncRPCResult) SetError(err error) {
	r.Err = err
}

func (r *AsyncRPCResult) Error() error {
	return r.Err
}

// SetResult ...
func (r *AsyncRPCResult) SetResult(rest interface{}) {
	r.Rest = rest
}

// Result ...
func (r *AsyncRPCResult) Result() interface{} {
	return r.Rest
}

// SetAttachments ...
func (r *AsyncRPCResult) SetAttachments(attr map[string]string) {
	r.Attrs = attr
}

// Attachments ...
func (r *AsyncRPCResult) Attachments() map[string]string {
	return r.Attrs
}

// AddAttachment ...
func (r *AsyncRPCResult) AddAttachment(key, value string) {
	r.Attrs[key] = value
}

// Attachment ...
func (r *AsyncRPCResult) Attachment(key, defaultValue string) string {
	v, ok := r.Attrs[key]
	if !ok {
		v = defaultValue
	}
	return v
}
