package utility

import (
	storage_go "github.com/supabase-community/storage-go"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"os"
	"time"
)

func Hashing(pass string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func ComparePass(hashedPass, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	if err != nil {
		return err
	}
	return nil
}

func SendMails(userEmail string, code string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "aryaizra2@gmail.com")
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Code for reset")
	m.SetBody("text/html", "Hello, this is your code for reset the password, <b>"+code+"</b>")

	d := gomail.NewDialer("smtp.gmail.com", 587, "aryaizra2@gmail.com", os.Getenv("PASS"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

const charset = "abcdefghijklmnopqrstuvwxyz"

func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func SupabaseClient() *storage_go.Client {
	return storage_go.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_API"), nil)
}
