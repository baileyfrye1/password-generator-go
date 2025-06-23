# Go Password/Passphrase Generator
Password/passphrase generator script built in Go. Uses the github.com/sethvargo/go-diceware/diceware package to handle passphrase creation and supports the following flags for customization:
- -length, -l int Change password length (default **16**)
- -special, -s bool Add special characters to password (default **false**)
- -passphrase, -p bool Generate a passphrase instead of a password (default **false**)
- -words, -w int Change number of words in passphrase (default **4**)
- -separator, -sep string Change separator between words in passphrase (default **-**)
