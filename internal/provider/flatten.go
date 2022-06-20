package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Flatten(tagMap map[string]*string) map[string]interface{} {
	// If tagsMap is nil, len(tagsMap) will be 0.
	output := make(map[string]interface{}, len(tagMap))

	for i, v := range tagMap {
		if v == nil {
			continue
		}

		output[i] = *v
	}

	return output
}

func FlattenAndSet(d *schema.ResourceData, tagMap map[string]*string, attrib string) error {
	flattened := Flatten(tagMap)
	if err := d.Set(attrib, flattened); err != nil {
		return fmt.Errorf("setting `%s`: %s", attrib, err)
	}

	return nil
}
