package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	//Criamos essa constante para evitar a repetição de código
	name     = "Campaign X"
	content  = "body hi!"
	contacts = []string{"email1@e.com", "email2@e.com"}
	fake     = faker.New() //criamos uma variável para usar o faker

)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

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

	println(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

// CRIANDO TESTE PARA VER SE O ID NÃO É NULO, EXEMPLO PARA PODERMOS PERCEBER QUE PODEMOS TESTAR AS COISAS SEPARADAS
func Test_NewCampaign_IDIsNotNill(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

// Testando o time, tem um bug aula 71.
func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	//colocando uma variável para combater o bug
	now := time.Now().Add(-time.Minute) //verifica se o tempo de agora  é menor que o da criação do email

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

// CRIAÇÃO DO TDD (testando antes da produção)
func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required with min 5", err.Error())
}
func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name is required with max 24", err.Error())
}
func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required with min 5", err.Error())
}
func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, fake.Lorem().Text(1000), contacts)

	assert.Equal("content is required with max 1024", err.Error())
}
func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, nil)

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("Email is invalid", err.Error())
}
