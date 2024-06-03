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
var _ datasource.DataSource = &ZeaburUserDataSource{}

func NewZeaburUserDataSource() datasource.DataSource {
	return &ZeaburUserDataSource{}
}

// ZeaburUserDataSource defines the data source implementation.
type ZeaburUserDataSource struct {
	client *api.Client
}

// ZeaburUserDataSourceModel describes the data source data model.
type ZeaburUserDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Email    types.String `tfsdk:"email"`
	Username types.String `tfsdk:"username"`
}

func (u *ZeaburUserDataSource) Metadata(ctx context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_user"
}

func (u *ZeaburUserDataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		MarkdownDescription: "Zeabur user data source",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The ID of the user",
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The name of the user",
			},
			"email": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The email of the user",
			},
			"username": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The username of the user",
			},
		},
	}
}

func (u *ZeaburUserDataSource) Configure(ctx context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if request.ProviderData == nil {
		return
	}

	client, ok := request.ProviderData.(*api.Client)
	if !ok {
		response.Diagnostics.AddError("invalid provider data", "The provider data is not a valid client")
	}

	u.client = client
}

func (u *ZeaburUserDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data ZeaburUserDataSourceModel

	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	user, err := u.client.GetUserInfo(ctx)
	if err != nil {
		response.Diagnostics.AddError("failed to get user", err.Error())
		return
	}

	data.Id = types.StringValue(user.ID)
	data.Name = types.StringValue(user.Name)
	data.Email = types.StringValue(user.Email)
	data.Username = types.StringValue(user.Username)

	tflog.Trace(ctx, "Read project", map[string]interface{}{
		"id":       data.Id,
		"name":     data.Name,
		"email":    data.Email,
		"username": data.Username,
	})

	response.Diagnostics.Append(response.State.Set(ctx, data)...)
}
