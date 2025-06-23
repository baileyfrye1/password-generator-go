package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"

	"github.com/sethvargo/go-diceware/diceware"
	"golang.design/x/clipboard"
)

func RegisterFlags(length *int, special *bool, passphrase *bool, words *int, separator *string) {
	flag.IntVar(length, "length", 16, "Length of the password")
	flag.IntVar(length, "l", 16, "Length of the password (shorthand)")

	flag.BoolVar(special, "special", false, "Include special characters in password")
	flag.BoolVar(special, "s", false, "Include special characters in password (shorthand)")

	flag.BoolVar(passphrase, "passphrase", false, "Return a passphrase instead of a password")
	flag.BoolVar(passphrase, "p", false, "Return a passphrase instead of a password (shorthand)")

	flag.IntVar(words, "words", 4, "Number of words in passphrase")
	flag.IntVar(words, "w", 4, "Number of words in passphrase (shorthand)")

	flag.StringVar(separator, "separator", "-", "Separator between words in passphrase")
	flag.StringVar(separator, "sep", "-", "Separator between words in passphrase (shorthand)")

	flag.Usage = func() {
		fmt.Println("Usage: passgen [options]")
		fmt.Println("\nOptions:")
		fmt.Println("  -length, -l int          Length of the password (default 16)")
		fmt.Println("  -special, -s bool        Include special characters in password")
		fmt.Println("  -passphrase, -p bool     Return a passphrase instead of a password")
		fmt.Println("  -words, -w int           Number of words in passphrase (default 4)")
		fmt.Println(
			"  -separator, -sep string  Separator between words in passphrase (default \"-\")",
		)
	}
}

func HandlePassphrase(words int, separator string) error {
	wordList, err := diceware.Generate(words)
	if err != nil {
		return fmt.Errorf("Error generating passphrase: %w\n", err)
	}

	passphrase := strings.Join(wordList, separator)

	clipboard.Write(clipboard.FmtText, []byte(passphrase))
	fmt.Println("Passphrase added to clipboard")

	fmt.Println(passphrase)
	return nil
}

func HandlePassword(special bool, length int) error {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	if special {
		charset += "!@#$%^&*()-_=+[]{}<>?/|"
	}

	password, err := generatePassword(length, charset)
	if err != nil {
		return fmt.Errorf("Error creating password: %w\n", err)
	}

	clipboard.Write(clipboard.FmtText, []byte(password))
	fmt.Println("Password added to clipboard")

	fmt.Println(password)
	return nil
}

func generatePassword(length int, charset string) (string, error) {
	result := make([]byte, length)

	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}

	return string(result), nil
}
