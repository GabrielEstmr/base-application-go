package main_configs_ff_lib_resources

type FeaturesData struct {
	key          string
	group        string
	description  string
	defaultValue bool
}

func NewFeaturesData(
	key string,
	group string,
	description string,
	defaultValue bool) *FeaturesData {
	return &FeaturesData{
		key:          key,
		group:        group,
		description:  description,
		defaultValue: defaultValue}
}

func (this *FeaturesData) IsEmpty() bool {
	return *this == FeaturesData{}
}

func (this *FeaturesData) GetKey() string {
	return this.key
}

func (this *FeaturesData) GetGroup() string {
	return this.group
}

func (this *FeaturesData) GetDescription() string {
	return this.description
}

func (this *FeaturesData) GetDefaultValue() bool {
	return this.defaultValue
}
