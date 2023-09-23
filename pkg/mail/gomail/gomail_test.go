package gomail

import (
	"context"
	"os"
	"strconv"
	"testing"
)

func TestGoMail_SendResetPasswordEmail(t *testing.T) {
	host, port, from, password := getEmailEnvironment(t)
	gm := New(host, port, from, password)

	err := gm.SendResetPasswordEmail(context.Background(), "test@mail.com", "12341")
	if err != nil {
		t.Error(err)
	}
}

func getEmailEnvironment(t *testing.T) (host string, port int, from string, password string) {
	host, ok := os.LookupEnv("SMTP_HOST")
	if !ok {
		t.Fatal("smtp host env not found")
	}

	portStr, ok := os.LookupEnv("SMTP_PORT")
	if !ok {
		t.Fatal("smtp port env not found")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		t.Fatal("invalid port env")
	}

	from, ok = os.LookupEnv("SMTP_FROM")
	if !ok {
		t.Fatal("smtp from env not found")
	}

	password, ok = os.LookupEnv("SMTP_PASSWORD")
	if !ok {
		t.Fatal("smtp password env not found")
	}

	return
}
