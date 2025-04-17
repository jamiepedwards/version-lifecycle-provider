package versionlifecycle

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var _ resource.Resource = &versionResource{}

// NewVersionResource is a helper function to simplify the provider implementation.
func NewVersionResource() resource.Resource {
	return &versionResource{}
}

// versionResource is the resource implementation.
type versionResource struct{}

// versionResourceModel maps the resource schema data.
type versionResourceModel struct {
	AppVersion        types.String `tfsdk:"app_version"`
	DefaultAppVersion types.String `tfsdk:"default_app_version"`
	CurrentAppVersion types.String `tfsdk:"current_app_version"`
}

// Metadata returns the resource type name.
func (r *versionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_version"
}

// Schema defines the schema for the resource.
func (r *versionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"app_version": schema.StringAttribute{
				Required: true,
			},
			"default_app_version": schema.StringAttribute{
				Required: true,
			},
			"current_app_version": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *versionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Info(ctx, "Creating version resource")
}

// Read refreshes the Terraform state with the latest data.
func (r *versionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Info(ctx, "Reading version resource")
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *versionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Info(ctx, "Updating version resource")
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *versionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Info(ctx, "Deleting version resource")
} 