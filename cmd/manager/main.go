/*
 * Copyright 2021 The KunStack Authors.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"kunstack.com/pharos/pkg/log"
	"math/rand"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"time"

	"kunstack.com/pharos/cmd/manager/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	cmd := app.NewControllerCommand(signals.SetupSignalHandler())
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
