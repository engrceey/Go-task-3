package exporter

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
)

// SocialMedia interface
type SocialMedia interface {
	Feed() []string
	Fame() int
}

//Txt export file
func Txt(u SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("An error occured opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("an error occured writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}

//JSON export file
func JSON(sm SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		return errors.New("error occured opening this file: " + err.Error())
	}
	for _, fd := range sm.Feed() {
		js, err := json.MarshalIndent(fd, "\n", " ")
		if err != nil {
			return errors.New("An error occured here: " + err.Error())
		}

		byte, err := f.Write(js)

		if err != nil {
			return errors.New("error occured while writing to file: " + err.Error())
		}

		fmt.Printf("wrote %d bytes\n", byte)
	}
	return nil
}

//XML export file
func XML(sm SocialMedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("An error occured opening this file: " + err.Error())
	}
	for _, fd := range sm.Feed() {
		x, err := xml.MarshalIndent(fd, "\n", " ")
		if err != nil {
			return errors.New("There was an error here: " + err.Error())
		}
		byte, err := f.Write(x)
		if err != nil {
			return errors.New("There was an error writing to this file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes\n", byte)
	}
	return nil
}
