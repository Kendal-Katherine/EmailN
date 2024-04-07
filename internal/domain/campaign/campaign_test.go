package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	/* SUBSTITUIMOS ESSE MONTE DE IF/ELSE por assert
	if campaign.ID != "1" {
		t.Errorf("expected 1")
	} else if campaign.Name != name {
		t.Errorf("expected correct name")
	} else if campaign.Content != content {
		t.Errorf("expected correct content")
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("expected correct contacts")
	}*/
	
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}
