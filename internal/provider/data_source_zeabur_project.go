package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/incubator4/terraform-provider-zeabur/internal/api"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &ZeaburProjectDataSource{}

func NewZeaburProjectDataSource() datasource.DataSource {
	return &ZeaburProjectDataSource{}
}

// ZeaburProjectDataSource defines the data source implementation.
type ZeaburProjectDataSource struct {
	client *api.Client
}

// ZeaburProjectDataSourceModel describes the data source data model.
type ZeaburProjectDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Owner  types.String `tfsdk:"owner"`
	Region types.String `tfsdk:"region"`
}

func (p *ZeaburProjectDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (p *ZeaburProjectDataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		MarkdownDescription: "Zeabur project data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,

				MarkdownDescription: "The ID of the project",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The name of the project",
			},
			"owner": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The owner of the project",
			},
			"region": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The region of the project",
			},
		},
	}
}

func (p *ZeaburProjectDataSource) Configure(ctx context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if request.ProviderData == nil {
		return
	}

	client, ok := request.ProviderData.(*api.Client)
	if !ok {
		response.Diagnostics.AddError("invalid provider data", "The provider data is not a valid client")
	}

	p.client = client
}

func (p *ZeaburProjectDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data ZeaburProjectDataSourceModel

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	project, err := p.client.GetProject(ctx, data.Id.ValueString())
	if err != nil {
		response.Diagnostics.AddError("failed to get project", err.Error())
		return
	}

	data.Name = types.StringValue(project.Name)

	tflog.Trace(ctx, "Read project", map[string]interface{}{
		"id":    data.Id,
		"name":  data.Name,
		"owner": data.Owner,
	})

	response.Diagnostics.Append(response.State.Set(ctx, data)...)
}
