// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type ExternalProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/external/2.3.1/docs#alias ExternalProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

