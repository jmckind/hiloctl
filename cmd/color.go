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
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// colorCmd represents the color command.
var colorCmd = &cobra.Command{
	Use:   "color <RGB>",
	Short: "Set the color",
	Long: `Set the color as a 6 character hex value in the form RRGGBB.

For example, set the color to white:

hiloctl color ffffff`,
	Args: colorValid,
	Run:  colorRun,
}

// colorRun handles the color command.
func colorRun(cmd *cobra.Command, args []string) {
	rgb := args[0]
	fmt.Println("color:", rgb)
}

// colorValid validates the arguments to the color command.
// Only one argument is expected and must be a valid hex RGB value.
func colorValid(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		rgb := args[0]
		if len(strings.TrimSpace(rgb)) == 6 {
			_, err := hex.DecodeString(rgb)
			if err == nil {
				return nil
			}
		}
	}
	return errors.New("invalid RGB value")
}

// init will initialize the color command.
func init() {
	rootCmd.AddCommand(colorCmd)
}
