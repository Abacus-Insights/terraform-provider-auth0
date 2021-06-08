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
							Default:  "",
						},
						"title": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"picker_title": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"message_multiple_tenants": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"audience_picker_alt_text": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"message_single_tenant": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"accept_button_text": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"decline_button_text": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"invalid_action": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"invalid_audience": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"invalid_scope": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
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
		if pc.PageTitle != "" {
			m["page_title"] = pc.PageTitle
		}
		if pc.Title != "" {
			m["title"] = pc.Title
		}
		if pc.PickerTitle != "" {
			m["picker_title"] = pc.PickerTitle
		}
		if pc.MessageMultipleTenants != "" {
			m["message_multiple_tenants"] = pc.MessageMultipleTenants
		}
		if pc.AudiencePickerAltText != "" {
			m["audience_picker_alt_text"] = pc.AudiencePickerAltText
		}
		if pc.MessageSingleTenant != "" {
			m["message_single_tenant"] = pc.MessageSingleTenant
		}
		if pc.AcceptButtonText != "" {
			m["accept_button_text"] = pc.AcceptButtonText
		}
		if pc.DeclineButtonText != "" {
			m["decline_button_text"] = pc.DeclineButtonText
		}
		if pc.InvalidAction != "" {
			m["invalid_action"] = pc.InvalidAction
		}
		if pc.InvalidAudience != "" {
			m["invalid_audience"] = pc.InvalidAudience
		}
		if pc.InvalidScope != "" {
			m["invalid_scope"] = pc.InvalidScope
		}
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
	rpc, err := api.Prompt.ReadPromptConsent(language(d))
	if pc.Consent.PageTitle == "" && rpc.Consent.PageTitle != "" {
		pc.Consent.PageTitle = rpc.Consent.PageTitle
	}
	if pc.Consent.Title == "" && rpc.Consent.Title != "" {
		pc.Consent.Title = rpc.Consent.Title
	}
	if pc.Consent.PickerTitle == "" && rpc.Consent.PickerTitle != "" {
		pc.Consent.PickerTitle = rpc.Consent.PickerTitle
	}
	if pc.Consent.MessageMultipleTenants == "" && rpc.Consent.MessageMultipleTenants != "" {
		pc.Consent.MessageMultipleTenants = rpc.Consent.MessageMultipleTenants
	}
	if pc.Consent.AudiencePickerAltText == "" && rpc.Consent.AudiencePickerAltText != "" {
		pc.Consent.AudiencePickerAltText = rpc.Consent.AudiencePickerAltText
	}
	if pc.Consent.MessageSingleTenant == "" && rpc.Consent.MessageSingleTenant != "" {
		pc.Consent.MessageSingleTenant = rpc.Consent.MessageSingleTenant
	}
	if pc.Consent.AcceptButtonText == "" && rpc.Consent.AcceptButtonText != "" {
		pc.Consent.AcceptButtonText = rpc.Consent.AcceptButtonText
	}
	if pc.Consent.DeclineButtonText == "" && rpc.Consent.DeclineButtonText != "" {
		pc.Consent.DeclineButtonText = rpc.Consent.DeclineButtonText
	}
	if pc.Consent.InvalidAction == "" && rpc.Consent.InvalidAction != "" {
		pc.Consent.InvalidAction = rpc.Consent.InvalidAction
	}
	if pc.Consent.InvalidScope == "" && rpc.Consent.InvalidScope != "" {
		pc.Consent.InvalidScope = rpc.Consent.InvalidScope
	}

	if err != nil {
		return err
	}
	err = api.Prompt.ReplacePromptConsent(language(d), pc)
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
