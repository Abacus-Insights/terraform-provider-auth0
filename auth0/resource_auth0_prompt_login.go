package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/auth0.v5/management"
)

func newPromptLogin() *schema.Resource {

	return &schema.Resource{

		Create: createPromptLogin,
		Read:   readPromptLogin,
		Update: updatePromptLogin,
		Delete: deletePromptLogin,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"language": {
				Type:     schema.TypeString,
				Required: true,
			},
			"login": {
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
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func createPromptLogin(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updatePromptLogin(d, m)
}

func flattenScreenLogin(pc *management.ScreenLogin) []interface{} {
	m := make(map[string]interface{})
	if pc != nil {
		m["page_title"] = pc.PageTitle
		m["title"] = pc.Title
		m["description"] = pc.Description
	}
	return []interface{}{m}
}

func readPromptLogin(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	pc, err := api.Prompt.ReadPromptLogin(language(d))
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.Set("login", flattenScreenLogin(pc.Login))
	return nil
}

func updatePromptLogin(d *schema.ResourceData, m interface{}) error {
	pc := buildPromptLogin(d)
	api := m.(*management.Management)
	err := api.Prompt.ReplacePromptLogin(language(d), pc)
	if err != nil {
		return err
	}
	return readPromptLogin(d, m)
}

func deletePromptLogin(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func buildPromptLogin(d *schema.ResourceData) *management.PromptLogin {
	pc := &management.PromptLogin{}

	List(d, "login").Elem(func(d ResourceData) {
		pc.Login = &management.ScreenLogin{
			PageTitle:   d.Get("page_title").(string),
			Title:       d.Get("title").(string),
			Description: d.Get("description").(string),
		}
	})

	return pc
}
