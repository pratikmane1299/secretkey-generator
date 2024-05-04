package cmd

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a random secret key",
	Long: `Generate a random secret key with custom options
	For Example:
	passkey generate -l 20 -s
`,
	Run: generatePasskey,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 10, "Length of the passkey to be generated")
	generateCmd.Flags().BoolP("special-chars", "s", false, "Include special characters")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits")
}

func generatePasskey(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	specialChars, _ := cmd.Flags().GetBool("special-chars")
	digits, _ := cmd.Flags().GetBool("digits")

	charset := "abcdefghijklmnopqrstuvwyzABCDEFGHIJKLMNOPQRSTUVWYZ"

	if specialChars {
		charset += "()*&^%$#@?"
	}

	if digits {
		charset += "1234567890"
	}

	passkey := make([]byte, length)

	for i := 0; i < len(passkey); i++ {
		passkey[i] = charset[rand.Intn(len(charset))]
	}

	var sb = strings.Builder{}

	for _, value := range passkey {
		sb.WriteByte(value)
	}

	fmt.Println(sb.String())
}
