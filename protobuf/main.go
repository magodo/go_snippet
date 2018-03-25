package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/magodo/go_snippet/protobuf/tutorial"
	"io/ioutil"
	"os"
	"strings"
)

type actionType int32

const (
	ACTION_NULL actionType = iota
	ACTION_HELP
	ACTION_ADD
	ACTION_DELETE
	ACTION_LIST
	ACTION_SEARCH
	ACTION_QUIT
	ACTION_UNKNOWN
)

var bookFile string

func getAction() string {
	in := bufio.NewReader(os.Stdin)
	act, _ := in.ReadString('\n')
	return strings.ToLower(strings.TrimSpace(act))
}

type quitMainError int32

func (err *quitMainError) Error() string {
	return "quit program"
}

func process(book *pb.AddressBook, act string) (err error) {
	err = nil

	switch act {
	case "help", "h", "?":
		showHelp()
	case "add":
		err = addPerson(book)
	case "delete":
		//err = deletePerson(book)
	case "list":
		err = listBook(book)
	case "search":
		//err = searchPerson(book)
	case "quit", "q":
		err = new(quitMainError)
	case "":
	default:
		fmt.Printf("Unknown action: %s\n", act)
	}
	return
}

// process functions

func showHelp() {
	fmt.Printf("Supported action list:\n* help\n* add\n* delete\n* list\n* search\n* quit\n")
}

func addPerson(book *pb.AddressBook) (err error) {
	in := bufio.NewReader(os.Stdin)
	person := &pb.Person{}
	person.Id = new(int32)

	fmt.Print("ID: ")
	if _, err := fmt.Fscanf(in, "%d\n", person.Id); err != nil {
		return err
	}

	fmt.Print("Name: ")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Printf("empty \"Name\" is not allowed\n")
		return
	}
	person.Name = &name

	for {
		fmt.Print("Phone Number (blank to skip): ")
		phoneNumber, _ := in.ReadString('\n')
		phoneNumber = strings.TrimSpace(phoneNumber)
		if phoneNumber == "" {
			break
		}

		fmt.Print("Phone Type (home, work or mobile): ")
		phoneType, _ := in.ReadString('\n')
		phoneType = strings.ToLower(strings.TrimSpace(phoneType))
		var realPhoneType pb.Person_PhoneType
		switch phoneType {
		case "":
			fmt.Printf("empty \"Phone Type\" is no allowed\n")
			continue
		case "home":
			realPhoneType = pb.Person_HOME
		case "work":
			realPhoneType = pb.Person_WORK
		case "mobile":
			realPhoneType = pb.Person_MOBILE
		default:
			fmt.Printf("Unknown phone type: %s\n", phoneType)
			continue
		}

		person.Phones = append(person.Phones, &pb.Person_PhoneNumber{
			Number: &phoneNumber,
			Type:   &realPhoneType,
		})
	}

	fmt.Print("Email (blank to skip): ")
	email, _ := in.ReadString('\n')
	if email = strings.ToLower(strings.TrimSpace(email)); email != "" {
		person.Email = &email
	}

	book.People = append(book.People, person)
	return
}

func listBook(book *pb.AddressBook) error {
	for _, person := range book.People {
		fmt.Printf("ID: %d\n", *person.Id)
		fmt.Printf("\tName: %s\n", *person.Name)
		for _, phone := range person.Phones {
			var ptype string
			switch *phone.Type {
			case pb.Person_HOME:
				ptype = "home"
			case pb.Person_WORK:
				ptype = "work"
			case pb.Person_MOBILE:
				ptype = "mobile"
			}
			fmt.Printf("\tPhone (%s): %s\n", ptype, *phone.Number)
		}
		if person.Email != nil {
			fmt.Printf("\tEmail: %s\n", *person.Email)
		}
	}
	return nil
}

//

func main() {

	_, err := os.Stat(bookFile)
	if err != nil {
		fmt.Printf("%s not exist, ceating...\n", bookFile)
		_, err = os.Create(bookFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Create file %s failed: %s\n", bookFile, err)
			return
		}
	}

	bin, err := ioutil.ReadFile(bookFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read file failed: %s\n", err)
		return
	}

	book := &pb.AddressBook{}
	if err := proto.Unmarshal(bin, book); err != nil {
		fmt.Fprintf(os.Stderr, "Can't unmarshal %s: %s\n", bookFile, err)
		return
	}
	defer func() {
		bin, _ := proto.Marshal(book)
		ioutil.WriteFile(bookFile, bin, 0644)
	}()

	for {
		fmt.Print("What to do: ")
		action := getAction()
		if err != nil {
			break
		}

		err := process(book, action)
		if err != nil {
			if _, ok := err.(*quitMainError); ok {
				break
			}
			fmt.Fprintf(os.Stderr, "Process failed: %s\n", err)
			return
		}
	}
	return
}

func init() {
	bookFile = *flag.String("book", "addressbook.data", "Filename of address book")
	flag.Parse()
}
