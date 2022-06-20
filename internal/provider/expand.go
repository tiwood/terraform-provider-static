package provider

func Expand(tagsMap map[string]interface{}) map[string]*string {
	output := make(map[string]*string, len(tagsMap))

	for i, v := range tagsMap {
		// Validate should have ignored this error already
		value, _ := v.(string)
		output[i] = &value
	}

	return output
}
