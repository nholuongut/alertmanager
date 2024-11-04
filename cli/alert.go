// Copyright 2018 Nho luong DevOps
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"github.com/alecthomas/kingpin/v2"
)

func configureAlertCmd(app *kingpin.Application) {
	alertCmd := app.Command("alert", "Add or query alerts.").PreAction(requireAlertManagerURL)
	configureQueryAlertsCmd(alertCmd)
	configureAddAlertCmd(alertCmd)
}
