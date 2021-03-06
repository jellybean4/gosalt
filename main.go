// Copyright © 2016 jellybean4
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
  log "github.com/Sirupsen/logrus"
  "github.com/jellybean4/gosalt/cmd"
  "github.com/jellybean4/order_init"
)

func main() {
  if err := order.ExecFunc(); err !=  nil {
    log.Fatalf("execute init func failed, reason %s", err.Error())
  }
  cmd.Execute()
}
