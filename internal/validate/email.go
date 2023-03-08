package validate

import "net/mail"

func IsMail(input string) bool {
	_, err := mail.ParseAddress(input)
	return err == nil
}
