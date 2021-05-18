package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/auth0.v5/management"
)

func newPromptConsent() *schema.Resource {

	return &schema.Resource{

		Create: createPromptConsent,
		Read:   readPromptConsent,
		Update: updatePromptConsent,
		Delete: deletePromptConsent,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"consent": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"page_title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"picker_title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"message_multiple_tenants": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"audience_picker_alt_text": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"message_single_tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"accept_button_text": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"decline_button_text": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"invalid_action": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"invalid_audience": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"invalid_scope": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func createPromptConsent(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updatePromptConsent(d, m)
}

func readPromptConsent(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	p, err := api.Prompt.Read()
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.Set("universal_login_experience", p.UniversalLoginExperience)
	d.Set("identifier_first", p.IdentifierFirst)
	return nil
}

func updatePromptConsent(d *schema.ResourceData, m interface{}) error {
	p := buildPromptConsent(d)
	api := m.(*management.Management)
	err := api.Prompt.ReplacePromptConsent(p)
	if err != nil {
		return err
	}
	return readPromptConsent(d, m)
}

func deletePromptConsent(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func buildPromptConsent(d *schema.ResourceData) *management.PromptConsent {
	p := &management.PromptConsent{}

	List(d, "consent").Elem(func(d ResourceData) {
		p.Consent = &management.ScreenConsent{
			PageTitle:              *String(d, "page_title"),
			Title:                  *String(d, "title"),
			PickerTitle:            *String(d, "picker_title"),
			MessageMultipleTenants: *String(d, "message_multiple_tenants"),
			AudiencePickerAltText:  *String(d, "audience_picker_alt_text"),
			MessageSingleTenant:    *String(d, "message_single_tenant"),
			AcceptButtonText:       *String(d, "accept_button_text"),
			DeclineButtonText:      *String(d, "decline_button_text"),
			InvalidAction:          *String(d, "invalid_action"),
			InvalidAudience:        *String(d, "invalid_audience"),
			InvalidScope:           *String(d, "invalid_scope"),
		}
	})

	return p
}
