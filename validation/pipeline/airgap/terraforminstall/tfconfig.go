package terraforminstall

// The json/yaml config key for the terraform install
const (
	TerraformConfigurationFileKey = "terraform"
)

// TerraformConfig is a struct that has standalone airgap config created in the pipeline
type TerraformConfig struct {
	StandaloneAirgapConfig StandaloneAirgapConfig `json:"standalone" yaml:"standaloneg"`
}

type StandaloneAirgapConfig struct {
	AirgapInternalFQDN string `json:"airgapInternalFQDN" yaml:"airgapInternalFQDN"`
	BootstrapPassword  string `json:"bootstrapPassword" yaml:"bootstrapPassword"`
	PrivateRegistry    string `json:"privateRegistry" yaml:"privateRegistry"`
}
