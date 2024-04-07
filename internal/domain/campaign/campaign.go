package campaign

import (
	"time"

	"github.com/rs/xid"
)

type Contacts struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contacts
}

func NewCampaign(name string, content string, emails []string) *Campaign {

	contacts := make([]Contacts, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:      xid.New().String(),
		Name:    name,
		Content: content,
		//Para testar o createdOn é necessário comentar ele
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}
}
