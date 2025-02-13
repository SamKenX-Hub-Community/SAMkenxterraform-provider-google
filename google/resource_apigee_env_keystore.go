// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceApigeeEnvKeystore() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeEnvKeystoreCreate,
		Read:   resourceApigeeEnvKeystoreRead,
		Delete: resourceApigeeEnvKeystoreDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeEnvKeystoreImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"env_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The Apigee environment group associated with the Apigee environment,
in the format 'organizations/{{org_name}}/environments/{{env_name}}'.`,
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The name of the newly created keystore.`,
			},
			"aliases": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `Aliases in this keystore.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeEnvKeystoreCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandApigeeEnvKeystoreName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}{{env_id}}/keystores")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EnvKeystore: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating EnvKeystore: %s", err)
	}

	// Store the ID now
	id, err := ReplaceVars(d, config, "{{env_id}}/keystores/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating EnvKeystore %q: %#v", d.Id(), res)

	return resourceApigeeEnvKeystoreRead(d, meta)
}

func resourceApigeeEnvKeystoreRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}{{env_id}}/keystores/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ApigeeEnvKeystore %q", d.Id()))
	}

	if err := d.Set("aliases", flattenApigeeEnvKeystoreAliases(res["aliases"], d, config)); err != nil {
		return fmt.Errorf("Error reading EnvKeystore: %s", err)
	}
	if err := d.Set("name", flattenApigeeEnvKeystoreName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading EnvKeystore: %s", err)
	}

	return nil
}

func resourceApigeeEnvKeystoreDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := ReplaceVars(d, config, "{{ApigeeBasePath}}{{env_id}}/keystores/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting EnvKeystore %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "EnvKeystore")
	}

	log.Printf("[DEBUG] Finished deleting EnvKeystore %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeEnvKeystoreImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := ParseImportId([]string{
		"(?P<env_id>.+)/keystores/(?P<name>.+)",
		"(?P<env_id>.+)/(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := ReplaceVars(d, config, "{{env_id}}/keystores/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeEnvKeystoreAliases(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenApigeeEnvKeystoreName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandApigeeEnvKeystoreName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
