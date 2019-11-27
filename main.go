package main

import (
	"fmt"
	"io/ioutil"
	"log"
	pb "github.com/hidaiy/protocol-buffer-sample/dest/tutorial"
	"github.com/golang/protobuf/proto"
)

func main() {
	book := &pb.AddressBook{}
	p := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	book.People = append(book.People, p)

	fname := "./addressbook.txt"

	// Write the new address book back to disk.
	fmt.Println("Marshal ----------------------------------------")
	write(book, fname)
	bytes, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Failed to read a file.")
	}
	fmt.Println(string(bytes))

	fmt.Println("Unmarshal ----------------------------------------")
	read(fname)
}
func write(book *pb.AddressBook, fname string) {
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
}

func read(fname string) {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Printf("book: %#v\n", book)
}
