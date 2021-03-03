package msg

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Msg struct {
	Email Email
	ACK int
}

type Email struct {
	To, From, Date, Title, Content string
}

// Compose returns a new Email struct filled with user input data
func ComposeEmail() Email {
	email := new(Email)

	fmt.Println("Compose New Message:")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("To: ")
	scanner.Scan()
	email.To = scanner.Text()

	fmt.Print("From: ")
	scanner.Scan()
	email.From = scanner.Text()

	email.Date = time.Now().Format("01-02-2006 15:04:05")

	fmt.Print("Title: ")
	scanner.Scan()
	email.Title = scanner.Text()

	fmt.Print("Content: ")
	scanner.Scan()
	email.Content = scanner.Text()

	return *email
}

// Print prints msg to the console
func (email Email) Print() {
	fmt.Println("--------- New mail --------- ")
	fmt.Println("To:", email.To)
	fmt.Println("From:", email.From)
	fmt.Println("Date:", email.Date)
	fmt.Println("Title:", email.Title)
	fmt.Println("Content:", email.Content)
	fmt.Println("---------------------------- ")
}
