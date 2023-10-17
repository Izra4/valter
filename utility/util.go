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

func SendMailsBook() {
	m := gomail.NewMessage()
	m.SetHeader("From", "aryaizra2@gmail.com")
	m.SetHeader("To", "akunuplay7@gmail.com")
	m.SetHeader("Subject", "Booking Receipt")

	body := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Checkout Summary</title>
		</head>
		<body>
			<h1>Your E-commerce Checkout Summary</h1>
			<p>Dear Customer,</p>
			<p>Thank you for shopping with us! Here is a summary of your recent order:</p>

			<table style="border-collapse: collapse; width: 100%;">
				<tr>
					<th style="border: 1px solid #dddddd; text-align: left; padding: 8px;">Product</th>
					<th style="border: 1px solid #dddddd; text-align: left; padding: 8px;">Quantity</th>
					<th style="border: 1px solid #dddddd; text-align: left; padding: 8px;">Price</th>
				</tr>
				<tr>
					<td style="border: 1px solid #dddddd; padding: 8px;">Product Name 1</td>
					<td style="border: 1px solid #dddddd; padding: 8px;">1</td>
					<td style="border: 1px solid #dddddd; padding: 8px;">$19.99</td>
				</tr>
				<tr>
					<td style="border: 1px solid #dddddd; padding: 8px;">Product Name 2</td>
					<td style="border: 1px solid #dddddd; padding: 8px;">2</td>
					<td style="border: 1px solid #dddddd; padding: 8px;">$29.99</td>
				</tr>
				<!-- Add more rows for additional products -->

				<tr>
					<td style="border: 1px solid #dddddd; padding: 8px;"><strong>Total:</strong></td>
					<td style="border: 1px solid #dddddd; padding: 8px;"></td>
					<td style="border: 1px solid #dddddd; padding: 8px;"><strong>$79.97</strong></td>
				</tr>
			</table>

			<p>Thank you for your purchase. Your order will be shipped to the following address:</p>
			<address>
				John Doe<br>
				123 Main Street<br>
				City, State 12345<br>
			</address>

			<p>If you have any questions or need further assistance, please don't hesitate to contact our customer support.</p>

			<p>Thank you for shopping with us!</p>
		</body>
		</html>
	`

	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "aryaizra2@gmail.com", os.Getenv("PASS"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
