package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccPromptSignup(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccPromptSignupCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt_signup.prompt_signup", "language", "en"),
					resource.TestCheckResourceAttr("auth0_prompt_signup.prompt_signup", "signup.0.description", "description"),
				),
			},
			{
				Config: testAccPromptSignupUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt_signup.prompt_signup", "language", "en"),
					resource.TestCheckResourceAttr("auth0_prompt_signup.prompt_signup", "signup.0.description", "description"),
				),
			},
		},
	})
}

const testAccPromptSignupCreate = `

resource "auth0_prompt_signup" "prompt_signup" {
    language = "en"
	signup {
		description = "description"
		
	}
}
`

const testAccPromptSignupUpdate = `

resource "auth0_prompt_signup" "prompt_signup" {
    language = "en"
    signup {
		description = "updated_page_title"
	}
}
`
