package mail_test

import (
	"hostListen/base"
	"hostListen/pkg/mail"
	"testing"
)

func TestXxx(t *testing.T) {
	mail.Init(base.MAIL_USER, base.MAIL_PWD, base.MAIL_FROM)
	mail.Mail.Send([]string{base.MAIL_TEST_TO}, "ckie onen mail test", "ckie onen mail test mail test")
}
