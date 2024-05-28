package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/incubator4/terraform-provider-zeabur/internal/api"
	"time"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ZeaburProjectResource{}
var _ resource.ResourceWithImportState = &ZeaburProjectResource{}

func NewZeaburProjectResource() resource.Resource {
	return &ZeaburProjectResource{}
}

// ZeaburProjectResource defines the resource implementation.
type ZeaburProjectResource struct {
	client *api.Client
}

// ZeaburProjectResourceModel defines the model for the resource.
type ZeaburProjectResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Region      types.String `tfsdk:"region"`
	Description types.String `tfsdk:"description"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

func (p *ZeaburProjectResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	//TODO implement me
	panic("implement me")
}

func (p *ZeaburProjectResource) Metadata(ctx context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_project"
}

func (p *ZeaburProjectResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"region": schema.StringAttribute{
				Required: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *ZeaburProjectResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.Client)

	if !ok {
		resp.Diagnostics.AddError("Invalid provider data", "Expected *api.Client, got something else")
		return
	}

	p.client = client
}

func (p *ZeaburProjectResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan ZeaburProjectResourceModel
	diags := request.Plan.Get(ctx, &plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	// Create zeabur project
	project, err := p.client.CreateProject(ctx, plan.Region.ValueString(), plan.Name.ValueStringPointer())
	if err != nil {
		response.Diagnostics.AddError("Error creating project", err.Error())
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(project.ID)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = response.State.Set(ctx, plan)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (p *ZeaburProjectResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state ZeaburProjectResourceModel
	diags := request.State.Get(ctx, &state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return

	}

	// Get refreshed project details
	project, err := p.client.GetProject(ctx, state.ID.ValueString())
	if err != nil {
		response.Diagnostics.AddError("Error getting project", err.Error())
		return

	}
	state.Name = types.StringValue(project.Name)

	diags = response.State.Set(ctx, state)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}
}

func (p *ZeaburProjectResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	//TODO implement me
	panic("implement me")
}

func (p *ZeaburProjectResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var data ZeaburProjectResourceModel

	response.Diagnostics.Append(request.State.Get(ctx, &data)...)

	if response.Diagnostics.HasError() {
		return
	}

	err := p.client.DeleteProject(ctx, data.ID.ValueString())
	if err != nil {
		response.Diagnostics.AddError("Error deleting project", err.Error())
		return
	}
}
