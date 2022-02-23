package auth0

import (
	"net/http"

	"github.com/auth0/go-auth0/management"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func newPromptSignup() *schema.Resource {

	return &schema.Resource{

		Create: createPromptSignup,
		Read:   readPromptSignup,
		Update: updatePromptSignup,
		Delete: deletePromptSignup,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"language": {
				Type:     schema.TypeString,
				Required: true,
			},
			"signup": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func createPromptSignup(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updatePromptSignup(d, m)
}

func flattenScreenSignup(pc *management.ScreenSignup) []interface{} {
	m := make(map[string]interface{})
	if pc != nil {
		m["description"] = pc.Description
	}
	return []interface{}{m}
}

func readPromptSignup(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	pc, err := api.Prompt.ReadPromptSignup(language(d))
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.Set("signup", flattenScreenSignup(pc.Signup))
	return nil
}

func updatePromptSignup(d *schema.ResourceData, m interface{}) error {
	pc := buildPromptSignup(d)
	api := m.(*management.Management)
	err := api.Prompt.ReplacePromptSignup(language(d), pc)
	if err != nil {
		return err
	}
	return readPromptSignup(d, m)
}

func deletePromptSignup(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func buildPromptSignup(d *schema.ResourceData) *management.PromptSignup {
	pc := &management.PromptSignup{}

	List(d, "signup").Elem(func(d ResourceData) {
		pc.Signup = &management.ScreenSignup{
			Description: d.Get("description").(string),
		}
	})

	return pc
}
