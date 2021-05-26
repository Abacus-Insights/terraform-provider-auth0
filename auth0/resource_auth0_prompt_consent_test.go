package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccPromptConsent(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccPromptConsentCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.page_title", "page_title"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.title", "title"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.picker_title", "picker_title"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.message_multiple_tenants", "message_multiple_tenants"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.audience_picker_alt_text", "audience_picker_alt_text"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.message_single_tenant", "message_single_tenant"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.accept_button_text", "accept_button_text"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.decline_button_text", "decline_button_text"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.invalid_action", "invalid_action"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.invalid_audience", "invalid_audience"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.invalid_scope", "invalid_scope"),
				),
			},
			{
				Config: testAccPromptConsentUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.page_title", "updated_page_title"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.title", "updated_title"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.picker_title", "updated_picker_title"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.message_multiple_tenants", "updated_message_multiple_tenants"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.audience_picker_alt_text", "updated_audience_picker_alt_text"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.message_single_tenant", "updated_message_single_tenant"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.accept_button_text", "updated_accept_button_text"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.decline_button_text", "updated_decline_button_text"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.invalid_action", "updated_invalid_action"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.invalid_audience", "updated_invalid_audience"),
					resource.TestCheckResourceAttr("auth0_prompt_consent.prompt_consent", "consent.0.invalid_scope", "updated_invalid_scope"),
				),
			},
		},
	})
}

const testAccPromptConsentCreate = `

resource "auth0_prompt_consent" "prompt_consent" {
	consent {
		language = "en"
		  page_title = "page_title"
		  title = "title"
		  picker_title = "picker_title"
		  message_multiple_tenants = "message_multiple_tenants"
		  audience_picker_alt_text = "audience_picker_alt_text"
		  message_single_tenant = "message_single_tenant"
		  accept_button_text = "accept_button_text"
		  decline_button_text = "decline_button_text"
		  invalid_action = "invalid_action"
		  invalid_audience = "invalid_audience"
		  invalid_scope = "invalid_scope"
		}
	}
`

const testAccPromptConsentUpdate = `

resource "auth0_prompt_consent" "prompt_consent" {
	consent {
		language = "en"
		page_title = "updated_page_title"
		title = "updated_title"
		picker_title = "updated_picker_title"
		message_multiple_tenants = "updated_message_multiple_tenants"
		audience_picker_alt_text = "updated_audience_picker_alt_text"
		message_single_tenant = "updated_message_single_tenant"
		accept_button_text = "updated_accept_button_text"
		decline_button_text = "updated_decline_button_text"
		invalid_action = "updated_invalid_action"
		invalid_audience = "updated_invalid_audience"
		invalid_scope = "updated_invalid_scope"
	}
}
`
