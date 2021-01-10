/*
   Copyright 2021 John McKenzie

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

package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// init will initialize the speed command.
func init() {
	rootCmd.AddCommand(speedCmd)
}

// speedCmd represents the speed command.
var speedCmd = &cobra.Command{
	Use:   "speed <up|down>",
	Short: "Adjust the speed setting up or down",
	Long: `Adjust the speed setting up or down.

Increase the speed:

hiloctl speed up

Decrease the speed:

hiloctl speed down`,
	Args: speedValid,
	Run:  speedRun,
}

// speedRun handles the speed command.
func speedRun(cmd *cobra.Command, args []string) {
	direction := args[0]
	fmt.Println("speed:", direction)
}

// speedValid validates the arguments to the speed command.
func speedValid(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		direction := strings.ToLower(args[0])
		if direction == "down" || direction == "up" {
			return nil
		}
	}
	return errors.New("invalid direction, must be 'up' or 'down'")
}
