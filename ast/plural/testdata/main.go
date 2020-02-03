package demo

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

var a int

func resourceArmDedicatedHost() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"location": {},

			"resource_group_name": {},

			"host_group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"sku_name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"DSv3-Type1",
					"ESv3-Type1",
					"FSv2-Type2",
				}, false),
			},

			"platform_fault_domain": {
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},

			"auto_replace_on_failure": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"license_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hosts": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"foo": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"property": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"foo": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"tags": {},
		},
	}
}
