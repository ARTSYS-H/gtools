/*
Copyright © 2024 Lucas Hadey <lhflyextra300@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ARTSYS-H/gtools/internal/tokenGenerator"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// Variables pour les Flags
var length int
var noSymbols bool
var bClipboard bool

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Token Generator",
	Long:  `Generate random string with the chars you want, uppercase or lowercase letters, numbers and/or symbols.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("token called")
		token := tokenGenerator.GetToken(length, noSymbols)
		if bClipboard {
			err := clipboard.WriteAll(token)
			if err != nil {
				fmt.Println("Error with the clipboard module:", err)
				os.Exit(1)
			}
		} else {
			fmt.Println(token)
		}
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tokenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tokenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	tokenCmd.PersistentFlags().IntVarP(&length, "length", "l", 12, "Length of the token")
	tokenCmd.PersistentFlags().BoolVarP(&noSymbols, "no-symbols", "n", false, "If --no-symbols or -n is specified, do not use any non-alphanumeric characters in the generate token")
	tokenCmd.PersistentFlags().BoolVarP(&bClipboard, "clip", "c", false, "If --clip or -c is specified, do not print the password but instead copy it to the clipboard")
}
