package scanner

import (
	"bufio"
	"os"
	"strings"
	"fmt"

	"makefile/db/sql"
)

func NewUserInput() (sql.User, error) {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Print("Enter your fullname and phone number (optional) separated by space: ")

	scanner.Scan() 
	if err := scanner.Err(); err != nil {
		return sql.User{}, err
	}

	text := scanner.Text()

	fields := strings.Fields(text)
	
	// FullName, PhoneNumber
	fullName := fields[0]
	var phoneNumber *string
	if len(fields) > 1 {
		phoneNumber = &fields[1]
	}

	user := sql.User{
		FullName: fullName,
		PhoneNumber: phoneNumber,
	}

	return user, nil
}
