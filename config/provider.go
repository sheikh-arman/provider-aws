/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/sheikh-arman/provider-aws/config/dynamodb"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "aws"
	modulePath     = "github.com/sheikh-arman/provider-aws"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("kubedb.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			RegionAddition(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		dynamodb.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
