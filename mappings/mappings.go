package mappings

type TagMappings struct {
	MappedTags map[string]string
}

func (mappings *TagMappings) Init() {
	mappings.MappedTags = map[string]string{
		"axiom": "testing_default",
		"test":  "custom_value_name",
	}
}
