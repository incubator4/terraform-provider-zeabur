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
var _ datasource.DataSource = &ZeaburServiceDataSource{}

func NewZeaburServiceDataSource() datasource.DataSource {
	return &ZeaburServiceDataSource{}
}

// ZeaburServiceDataSource defines the data source implementation.
type ZeaburServiceDataSource struct {
	client *api.Client
}

type ZeaburServiceDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Owner   types.String `tfsdk:"owner"`
	Project types.String `tfsdk:"project"`
}

func (s *ZeaburServiceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

func (s *ZeaburServiceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("invalid provider data", "The provider data is not a valid client")
	}

	s.client = client
}

func (s *ZeaburServiceDataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		MarkdownDescription: "Zeabur service data source",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the service",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the service",
			},
			"owner": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The owner of the service",
			},
			"project": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The project of the service",
			},
		},
	}
}

func (s *ZeaburServiceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ZeaburServiceDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	service, err := s.client.GetService(ctx,
		"",
		data.Owner.ValueString(),
		data.Project.ValueString(),
		data.Name.ValueString(),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to get service", err.Error())
		return
	}
	data.Id = types.StringValue(service.ID)

	tflog.Trace(ctx, "Read service", map[string]interface{}{
		"id":      data.Id,
		"name":    data.Name,
		"owner":   data.Owner,
		"project": data.Project,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}
