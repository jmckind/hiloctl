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

// init will initialize the mode command.
func init() {
	rootCmd.AddCommand(modeCmd)
}

// modeCmd represents the mode command.
var modeCmd = &cobra.Command{
	Use:   "mode <next|prev>",
	Short: "Navigate the mode setting next or previous",
	Long: `Adjust the mode setting next or previous.

Navigate to the next mode:

hiloctl mode next

Navigate to the previous mode:

hiloctl mode prev`,
	Args: modeValid,
	Run:  modeRun,
}

// modeRun handles the mode command.
func modeRun(cmd *cobra.Command, args []string) {
	direction := args[0]
	fmt.Println("mode:", direction)
}

// modeValid validates the arguments to the mode command.
func modeValid(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		direction := strings.ToLower(args[0])
		if direction == "next" || direction == "prev" {
			return nil
		}
	}
	return errors.New("invalid direction, must be 'next' or 'prev'")
}
