/**
 * Copyright 2019 Comcast Cable Communications Management, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"xhttp/xhttpclient"

	"go.uber.org/fx"
)

type ClientIn struct {
	fx.In
	ClientInstrumentsIn
	xhttpclient.UnmarshalIn
}

func provideClient(configKey string) func(ClientIn) (xhttpclient.Interface, error) {
	return func(in ClientIn) (xhttpclient.Interface, error) {
		return xhttpclient.Unmarshal(
			configKey,
			in.RequestCount.Then,
			in.RequestDuration.Then,
			in.RequestsInFlight.Then,
		)(in.UnmarshalIn)
	}
}