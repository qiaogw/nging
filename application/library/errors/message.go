/*

   Copyright 2016 Wenhui Shen <www.webx.top>

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

*/
package errors

var DefaultNopMessage Messager = &NopMessage{}

type NopMessage struct {
}

func (n *NopMessage) Error() string {
	return ``
}
func (n *NopMessage) Success() string {
	return ``
}

func (s *NopMessage) String() string {
	return ``
}

type Messager interface {
	Successor
	error
}

func IsMessage(err interface{}) bool {
	_, y := err.(Messager)
	return y
}

func Message(err interface{}) Messager {
	if v, y := err.(Messager); y {
		return v
	}
	return DefaultNopMessage
}
