package campaign

import "time"

type Contacts struct{
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content string
	Contacts []Contacts
}

func NewCampaign(name string, content string, emails []string) Campaign{
	
	contacts := make([]Contacts, len(emails))
	for index, email := range emails{
		contacts[index].Email = email
	}
	
	return Campaign{
		ID: "1",
		Name: name,
		Content:  content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		}
	}