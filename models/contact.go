package models

import (
	"fmt"

	u "../utils"
	"github.com/jinzhu/gorm"
)

// Contact struct for represent contacts
type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"` //The user that this contact belongs to
}

/*
 This struct function validate the required parameters sent through the http request body
returns message and true if the requirement is met
*/

// Validate validate Contact struct
func (contact *Contact) Validate() (map[string]interface{}, bool) {

	if contact.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if contact.UserID < 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameter are present
	return u.Message(true, "success"), true
}

// Create create contact
func (contact *Contact) Create() map[string]interface{} {

	if resp, ok := contact.Validate(); !ok {
		return resp
	}

	getDB().Create(contact)

	resp := u.Message(true, "Success")
	resp["contact"] = contact
	return resp
}

func getContact(id uint) *Contact {

	contact := &Contact{}
	err := getDB().Table("contacts").Where("id = ?", id).First(contact).Error

	if err != nil {
		return nil
	}

	return contact
}

func getContacts(user uint) []*Contact {

	contacts := make([]*Contact, 0)
	err := getDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return contacts
}
