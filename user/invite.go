// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// InviteService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInviteService] method instead.
type InviteService struct {
	Options []option.RequestOption
}

// NewInviteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInviteService(opts ...option.RequestOption) (r *InviteService) {
	r = &InviteService{}
	r.Options = opts
	return
}

// Lists all invitations associated with my user.
func (r *InviteService) List(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[Invite], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/invites"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all invitations associated with my user.
func (r *InviteService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Invite] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, opts...))
}

// Responds to an invitation.
func (r *InviteService) Edit(ctx context.Context, inviteID string, body InviteEditParams, opts ...option.RequestOption) (res *InviteEditResponse, err error) {
	var env InviteEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if inviteID == "" {
		err = errors.New("missing required invite_id parameter")
		return
	}
	path := fmt.Sprintf("user/invites/%s", inviteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the details of an invitation.
func (r *InviteService) Get(ctx context.Context, inviteID string, opts ...option.RequestOption) (res *InviteGetResponse, err error) {
	var env InviteGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if inviteID == "" {
		err = errors.New("missing required invite_id parameter")
		return
	}
	path := fmt.Sprintf("user/invites/%s", inviteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Invite struct {
	// ID of the user to add to the organization.
	InvitedMemberID string `json:"invited_member_id,required,nullable"`
	// ID of the organization the user will be added to.
	OrganizationID string `json:"organization_id,required"`
	// Invite identifier tag.
	ID string `json:"id"`
	// When the invite is no longer active.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The email address of the user who created the invite.
	InvitedBy string `json:"invited_by"`
	// Email address of the user to add to the organization.
	InvitedMemberEmail string `json:"invited_member_email"`
	// When the invite was sent.
	InvitedOn time.Time `json:"invited_on" format:"date-time"`
	// Organization name.
	OrganizationName string `json:"organization_name"`
	// Roles to be assigned to this user.
	Roles []shared.Role `json:"roles"`
	// Current status of the invitation.
	Status InviteStatus `json:"status"`
	JSON   inviteJSON   `json:"-"`
}

// inviteJSON contains the JSON metadata for the struct [Invite]
type inviteJSON struct {
	InvitedMemberID    apijson.Field
	OrganizationID     apijson.Field
	ID                 apijson.Field
	ExpiresOn          apijson.Field
	InvitedBy          apijson.Field
	InvitedMemberEmail apijson.Field
	InvitedOn          apijson.Field
	OrganizationName   apijson.Field
	Roles              apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Invite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteJSON) RawJSON() string {
	return r.raw
}

// Current status of the invitation.
type InviteStatus string

const (
	InviteStatusPending  InviteStatus = "pending"
	InviteStatusAccepted InviteStatus = "accepted"
	InviteStatusRejected InviteStatus = "rejected"
	InviteStatusExpired  InviteStatus = "expired"
)

func (r InviteStatus) IsKnown() bool {
	switch r {
	case InviteStatusPending, InviteStatusAccepted, InviteStatusRejected, InviteStatusExpired:
		return true
	}
	return false
}

type InviteEditResponse = interface{}

type InviteGetResponse = interface{}

type InviteEditParams struct {
	// Status of your response to the invitation (rejected or accepted).
	Status param.Field[InviteEditParamsStatus] `json:"status,required"`
}

func (r InviteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of your response to the invitation (rejected or accepted).
type InviteEditParamsStatus string

const (
	InviteEditParamsStatusAccepted InviteEditParamsStatus = "accepted"
	InviteEditParamsStatusRejected InviteEditParamsStatus = "rejected"
)

func (r InviteEditParamsStatus) IsKnown() bool {
	switch r {
	case InviteEditParamsStatusAccepted, InviteEditParamsStatusRejected:
		return true
	}
	return false
}

type InviteEditResponseEnvelope struct {
	Result InviteEditResponse             `json:"result"`
	JSON   inviteEditResponseEnvelopeJSON `json:"-"`
}

// inviteEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [InviteEditResponseEnvelope]
type inviteEditResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InviteGetResponseEnvelope struct {
	Result InviteGetResponse             `json:"result"`
	JSON   inviteGetResponseEnvelopeJSON `json:"-"`
}

// inviteGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InviteGetResponseEnvelope]
type inviteGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
