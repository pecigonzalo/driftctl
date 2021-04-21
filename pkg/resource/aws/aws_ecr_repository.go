// GENERATED, DO NOT EDIT THIS FILE
package aws

import (
	"github.com/zclconf/go-cty/cty"

	"github.com/cloudskiff/driftctl/pkg/dctlcty"
)

const AwsEcrRepositoryResourceType = "aws_ecr_repository"

type AwsEcrRepository struct {
	Arn                     *string           `cty:"arn" computed:"true"`
	Id                      string            `cty:"id" computed:"true"`
	ImageTagMutability      *string           `cty:"image_tag_mutability"`
	Name                    *string           `cty:"name"`
	RegistryId              *string           `cty:"registry_id" computed:"true"`
	RepositoryUrl           *string           `cty:"repository_url" computed:"true"`
	Tags                    map[string]string `cty:"tags"`
	EncryptionConfiguration *[]struct {
		EncryptionType *string `cty:"encryption_type"`
		KmsKey         *string `cty:"kms_key" computed:"true"`
	} `cty:"encryption_configuration"`
	ImageScanningConfiguration *[]struct {
		ScanOnPush *bool `cty:"scan_on_push"`
	} `cty:"image_scanning_configuration"`
	Timeouts *struct {
		Delete *string `cty:"delete"`
	} `cty:"timeouts" diff:"-"`
	CtyVal *cty.Value `diff:"-"`
}

func (r *AwsEcrRepository) TerraformId() string {
	return r.Id
}

func (r *AwsEcrRepository) TerraformType() string {
	return AwsEcrRepositoryResourceType
}

func (r *AwsEcrRepository) CtyValue() *cty.Value {
	return r.CtyVal
}

func initAwsEcrRepositoryMetaData() {
	dctlcty.SetMetadata(AwsEcrRepositoryResourceType, AwsEcrRepositoryTags, AwsEcrRepositoryNormalizer)
}

var AwsEcrRepositoryTags = map[string]string{
	"arn":                              `computed:"true"`,
	"id":                               `computed:"true"`,
	"registry_id":                      `computed:"true"`,
	"repository_url":                   `computed:"true"`,
	"encryption_configuration.kms_key": `computed:"true"`,
}

func AwsEcrRepositoryNormalizer(val *rescty.CtyAttributes) {
	val.SafeDelete([]string{"timeouts"})
}
