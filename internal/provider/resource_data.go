package provider

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceData() *schema.Resource {
	return &schema.Resource{
		Description: `
Creates static data assignments.

The **data** map will be copied to the **output** attribute during resource creation. 
**output** values won't change during resource updates, only on re-creation.
`,

		CreateContext: resourceDataCreate,
		ReadContext:   resourceDataRead,
		UpdateContext: resourceDataUpdate,
		DeleteContext: resourceDataDelete,

		Schema: map[string]*schema.Schema{
			"data": {
				Description: "A map of strings which should be statically set as the resources `output`.",
				Type:        schema.TypeMap,
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"triggers": {
				Description: "Can be used to force a re-creation of the resource and therefore will update the `output` attribute.",
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"output": {
				Description: "A copy of the `data` attribute, as it was during resource creation.",
				Type:        schema.TypeMap,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceDataCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	data := d.Get("data")
	d.Set("output", data)
	d.SetId(fmt.Sprintf("%d", rand.Int()))

	return nil
}

func resourceDataRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceDataUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceDataDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
