package auth0

import (
	"net/http"

	"github.com/auth0/go-auth0/management"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			"language": {
				Type:     schema.TypeString,
				Required: true,
			},
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

func language(d *schema.ResourceData) string {
	return *String(d, "language")
}

func createPromptConsent(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updatePromptConsent(d, m)
}

func flattenScreenConsent(pc *management.ScreenConsent) []interface{} {
	m := make(map[string]interface{})
	if pc != nil {
		m["page_title"] = pc.PageTitle
		m["title"] = pc.Title
		m["picker_title"] = pc.PickerTitle
		m["message_multiple_tenants"] = pc.MessageMultipleTenants
		m["audience_picker_alt_text"] = pc.AudiencePickerAltText
		m["message_single_tenant"] = pc.MessageSingleTenant
		m["accept_button_text"] = pc.AcceptButtonText
		m["decline_button_text"] = pc.DeclineButtonText
		m["invalid_action"] = pc.InvalidAction
		m["invalid_audience"] = pc.InvalidAudience
		m["invalid_scope"] = pc.InvalidScope
	}
	return []interface{}{m}
}

func readPromptConsent(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	pc, err := api.Prompt.ReadPromptConsent(language(d))
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.Set("consent", flattenScreenConsent(pc.Consent))
	return nil
}

func updatePromptConsent(d *schema.ResourceData, m interface{}) error {
	pc := buildPromptConsent(d)
	api := m.(*management.Management)
	err := api.Prompt.ReplacePromptConsent(language(d), pc)
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
	pc := &management.PromptConsent{}

	List(d, "consent").Elem(func(d ResourceData) {
		pc.Consent = &management.ScreenConsent{
			PageTitle:              d.Get("page_title").(string),
			Title:                  d.Get("title").(string),
			PickerTitle:            d.Get("picker_title").(string),
			MessageMultipleTenants: d.Get("message_multiple_tenants").(string),
			AudiencePickerAltText:  d.Get("audience_picker_alt_text").(string),
			MessageSingleTenant:    d.Get("message_single_tenant").(string),
			AcceptButtonText:       d.Get("accept_button_text").(string),
			DeclineButtonText:      d.Get("decline_button_text").(string),
			InvalidAction:          d.Get("invalid_action").(string),
			InvalidAudience:        d.Get("invalid_audience").(string),
			InvalidScope:           d.Get("invalid_scope").(string),
		}
	})

	return pc
}
