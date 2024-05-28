// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"github.com/incubator4/terraform-provider-zeabur/internal/api"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure ZeaburProvider satisfies various provider interfaces.
var _ provider.Provider = &ZeaburProvider{}
var _ provider.ProviderWithFunctions = &ZeaburProvider{}

// ZeaburProvider defines the provider implementation.
type ZeaburProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ZeaburProviderModel describes the provider data model.
type ZeaburProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
	APIToken types.String `tfsdk:"api_token"`
}

func (p *ZeaburProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "zeabur"
	resp.Version = p.version
}

func (p *ZeaburProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Example provider attribute",
				Optional:            true,
			},
			"api_token": schema.StringAttribute{
				MarkdownDescription: "API token for the provider",
				Optional:            true,
			},
		},
	}
}

func (p *ZeaburProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	apiToken := os.Getenv("ZEABUR_API_TOKEN")

	var data ZeaburProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if data.APIToken.ValueString() != "" {
		apiToken = data.APIToken.ValueString()
	}

	if apiToken == "" {
		resp.Diagnostics.AddError(
			"Missing API Token Configuration",
			"While the API token can be set in the provider configuration, "+
				"it is recommended to set the ZEABUR_API_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }

	// Example client configuration for data sources and resources
	client := api.NewClient(apiToken)
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *ZeaburProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewZeaburProjectResource,
	}
}

func (p *ZeaburProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewZeaburProjectDataSource,
	}
}

func (p *ZeaburProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ZeaburProvider{
			version: version,
		}
	}
}
