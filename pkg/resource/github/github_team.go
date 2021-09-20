// GENERATED, DO NOT EDIT THIS FILE
package github

import "github.com/cloudskiff/driftctl/pkg/resource"

const GithubTeamResourceType = "github_team"

func initGithubTeamMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(GithubTeamResourceType, func(res *resource.Resource) {
		val := res.Attrs
		if defaultMaintainer, exist := val.Get("create_default_maintainer"); !exist || defaultMaintainer == nil {
			(*val)["create_default_maintainer"] = false
		}
		val.SafeDelete([]string{"etag"})
	})
	resourceSchemaRepository.SetHumanReadableAttributesFunc(GithubTeamResourceType, func(res *resource.Resource) map[string]string {
		val := res.Attrs
		attrs := make(map[string]string)
		attrs["Id"] = res.ResourceId()
		if name := val.GetString("name"); name != nil && *name != "" {
			attrs["Name"] = *name
		}
		return attrs
	})
	resourceSchemaRepository.SetFlags(GithubTeamResourceType, resource.FlagDeepMode)
}
