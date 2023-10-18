package utility

import (
	"fmt"
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

func GenerateInv() string {
	rand.Seed(time.Now().UnixNano())
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, 6)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return "INV-" + string(b)
}

func SendMailsBook(email, address, tanggal, due, invoiceNumber, description, name string) {
	m := gomail.NewMessage()
	m.SetHeader("From", "aryaizra2@gmail.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Booking Receipt")

	body := `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <!-- <script src="https://cdn.tailwindcss.com"></script> -->
    <title>Document</title>
  </head>
  <body class="p-8 space-y-12" style="padding: 29px;">
    <div class="flexBetween" style="display: flex; justify-content: space-between;">
      <h1 class="font-bold" style="font-weight: bold;">GIME</h1>
      <div class="flexColumn" style="display: flex; flex-direction: column;">
        <span>REG: 123000123000</span>
        <span>Gime@aitech.co | +6213123213123</span>
      </div>
    </div>

    <section class="flexBetween" style="display: flex; justify-content: space-between;">
      <div>
        <h2 class="font-bold" style="font-weight: bold;">` + name + `</h2>
        <div class="flexColumn" style="display: flex; flex-direction: column; width: 30vw;">
          <div class="flexBetween" style="display: flex; justify-content: space-between;">
            <span>INVOICE NUMBER</span>
            <span>` + invoiceNumber + `</span>
          </div>
          <div class="flexBetween" style="display: flex; justify-content: space-between;">
            <span>INVOICE DATE</span>
            <span>` + tanggal + `</span>
          </div>
          <div class="flexBetween" style="display: flex; justify-content: space-between;">
            <span>DUE</span>
            <span>` + due + `</span>
          </div>
        </div>
      </div>

      <div>
        <h2 class="font-bold" style="font-weight: bold;">Address</h2>
        <div class="flexColumn" style="display: flex; flex-direction: column;">
          <span style="width: 20vw;">` + address + `</span>
        </div>
      </div>
    </section>

    <section>
      <table class="w-full" style="width: 100%; margin: 24px 0; border-collapse: collapse;">
        <tr>
          <th style="border: 1px solid;">Description</th>
          <th style="border: 1px solid;">Qty</th>
          <th style="border: 1px solid;">Price</th>
          <th style="border: 1px solid;">GST</th>
          <th style="border: 1px solid;">Amount</th>
        </tr>

        <tr>
          <td style="border: 1px solid;">` + description + `</td>
          <td style="border: 1px solid;">1</td>
          <td style="border: 1px solid;">$5,455.00</td>
          <td style="border: 1px solid;">$545.50</td>
          <td style="border: 1px solid;">$6,000.50</td>
        </tr>
        
      </table>
    </section>

    <section class="w-full" style="width: 100%; margin: 24px 0; display: flex; justify-content: flex-end;">
      <div
        class="bg-neutral-100 p-4" style="background-color: #F9F9FA; border: 2px solid #DFE4EA; padding: 16px; width: 50vw; border-radius: 15px;"
      >
        <div class="flexBetween" style="display: flex; justify-content: space-between;">
          <span>Sub total (excl. GST)</span>
          <span>$5,455.00</span>
        </div>
        <div class="flexBetween" style="display: flex; justify-content: space-between;">
          <span>Total GST:</span>
          <span>$545.50</span>
        </div>
        <div class="flexBetween" style="display: flex; justify-content: space-between;">
          <span>Credit card fee (if using):</span>
          <span>$25.00</span>
        </div>
        <div class="pt-4 font-bold flexBetween" style="display: flex; justify-content: space-between; font-weight: bold;">
          <p>Amount due on ` + due + `:</p>
          <p>6,025.50 NZD</p>
        </div>
      </div>
    </section>

    <div class="flexBetween" style="display: flex; justify-content: space-between;">
      <div class="space-y-6">
        <h2 class="font-bold" style="font-weight: bold;">PAYMENT INSTRUCTION</h2>

        <div class="flexColumn" style="display: flex; flex-direction: column; list-style: none; text-indent: 0;">
          <span>Bank name: ABC Bank</span>
          <span>Account number: 12-1234-123456-12</span>
          <span class="font-bold" style="font-weight: bold;">
            Please use as ` + invoiceNumber + ` as a reference number
          </span>
        </div>

        <p>For any questions please contact us at Gime@aitech.co</p>
      </div>

      <div class="flexColumn" style="display: flex; flex-direction: column;">
        <span>Pay Online</span>
        <a href="https://buy.stripe.com/" class="underline font-bold" style="text-decoration: underline; font-weight: bold;">https://buy.stripe.com/</a>
      </div>
    </div>
  </body>
</html>`

	m.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "aryaizra2@gmail.com", os.Getenv("PASS"))

	// Send the email to Bob, Cora, and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}

func FormatDate(t time.Time) string {
	bulan := map[int]string{
		1:  "Jan",
		2:  "Feb",
		3:  "Mar",
		4:  "Apr",
		5:  "Mei",
		6:  "Jun",
		7:  "Jul",
		8:  "Agu",
		9:  "Sep",
		10: "Okt",
		11: "Nov",
		12: "Des",
	}

	day := t.Day()
	month := t.Month()
	year := t.Year()

	abbreviatedMonth, ok := bulan[int(month)]
	if !ok {
		abbreviatedMonth = month.String()
	}

	formattedDate := fmt.Sprintf("%d %s %d", day, abbreviatedMonth, year)
	return formattedDate
}

func FormatDueDate(t time.Time) string {
	bulan := map[int]string{
		1:  "Jan",
		2:  "Feb",
		3:  "Mar",
		4:  "Apr",
		5:  "Mei",
		6:  "Jun",
		7:  "Jul",
		8:  "Agu",
		9:  "Sep",
		10: "Okt",
		11: "Nov",
		12: "Des",
	}

	dueDate := t.AddDate(0, 1, 0)

	day := dueDate.Day()
	month := dueDate.Month()
	year := dueDate.Year()

	abbreviatedMonth, ok := bulan[int(month)]
	if !ok {
		abbreviatedMonth = month.String()
	}

	formattedDueDate := fmt.Sprintf("%d %s %d", day, abbreviatedMonth, year)
	return formattedDueDate
}
