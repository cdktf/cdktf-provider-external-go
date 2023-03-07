package provider


type ExternalProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/external#alias ExternalProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

