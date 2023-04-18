package dataexternal

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataExternalConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A list of strings, whose first element is the program to run and whose subsequent elements are optional command line arguments to the program.
	//
	// Terraform does not execute the program through a shell, so it is not necessary to escape shell metacharacters nor add quotes around arguments containing spaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/external/2.3.1/docs/data-sources/external#program DataExternal#program}
	Program *[]*string `field:"required" json:"program" yaml:"program"`
	// A map of string values to pass to the external program as the query arguments.
	//
	// If not supplied, the program will receive an empty object as its input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/external/2.3.1/docs/data-sources/external#query DataExternal#query}
	Query *map[string]*string `field:"optional" json:"query" yaml:"query"`
	// Working directory of the program. If not supplied, the program will run in the current directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/external/2.3.1/docs/data-sources/external#working_dir DataExternal#working_dir}
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

