package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	//Criamos essa constante para evitar a repetição de código
	name = "Campaign X"
	content = "body"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	
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

	println(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

// CRIANDO TESTE PARA VER SE O ID NÃO É NULO, EXEMPLO PARA PODERMOS PERCEBER QUE PODEMOS TESTAR AS COISAS SEPARADAS
func Test_NewCampaign_IDIsNotNill(t *testing.T) {

	assert := assert.New(t)
	
	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)

}

// Testando o time, tem um bug aula 71.
func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {

	assert := assert.New(t)
	
	//colocando uma variável para combater o bug
	now := time.Now().Add(-time.Minute)//verifica se o tempo de agora  é menor que o da criação do email

	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)

}
