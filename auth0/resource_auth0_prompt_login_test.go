package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccPromptLogin(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccPromptLoginCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt_login.prompt_login", "language", "en"),
					resource.TestCheckResourceAttr("auth0_prompt_login.prompt_login", "login.0.title", "title"),
				),
			},
			{
				Config: testAccPromptLoginUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt_login.prompt_login", "language", "en"),
					resource.TestCheckResourceAttr("auth0_prompt_login.prompt_login", "login.0.title", "title"),
				),
			},
		},
	})
}

const testAccPromptLoginCreate = `

resource "auth0_prompt_login" "prompt_login" {
    language = "en"
	login {
		title = "title"
		
	}
}
`

const testAccPromptLoginUpdate = `

resource "auth0_prompt_login" "prompt_login" {
    language = "en"
    login {
		title = "updated_page_title"
	}
}
`
