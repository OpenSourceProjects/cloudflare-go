// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// UserAuditLogService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserAuditLogService] method
// instead.
type UserAuditLogService struct {
	Options []option.RequestOption
}

// NewUserAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserAuditLogService(opts ...option.RequestOption) (r *UserAuditLogService) {
	r = &UserAuditLogService{}
	r.Options = opts
	return
}

// Gets a list of audit logs for a user account. Can be filtered by who made the
// change, on which zone, and the timeframe of the change.
func (r *UserAuditLogService) List(ctx context.Context, query UserAuditLogListParams, opts ...option.RequestOption) (res *UserAuditLogListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/audit_logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Union satisfied by [UserAuditLogListResponseObject] or
// [UserAuditLogListResponseQ6tjyqK4APIResponseCommon].
type UserAuditLogListResponse interface {
	implementsUserAuditLogListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*UserAuditLogListResponse)(nil)).Elem(), "")
}

type UserAuditLogListResponseObject struct {
	Errors   interface{}                            `json:"errors,nullable"`
	Messages []interface{}                          `json:"messages"`
	Result   []UserAuditLogListResponseObjectResult `json:"result"`
	Success  bool                                   `json:"success"`
	JSON     userAuditLogListResponseObjectJSON     `json:"-"`
}

// userAuditLogListResponseObjectJSON contains the JSON metadata for the struct
// [UserAuditLogListResponseObject]
type userAuditLogListResponseObjectJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserAuditLogListResponseObject) implementsUserAuditLogListResponse() {}

type UserAuditLogListResponseObjectResult struct {
	// A string that uniquely identifies the audit log.
	ID     string                                     `json:"id"`
	Action UserAuditLogListResponseObjectResultAction `json:"action"`
	Actor  UserAuditLogListResponseObjectResultActor  `json:"actor"`
	// The source of the event.
	Interface string `json:"interface"`
	// An object which can lend more context to the action being logged. This is a
	// flexible value and varies between different actions.
	Metadata interface{} `json:"metadata"`
	// The new value of the resource that was modified.
	NewValue string `json:"newValue"`
	// The value of the resource before it was modified.
	OldValue string                                       `json:"oldValue"`
	Owner    UserAuditLogListResponseObjectResultOwner    `json:"owner"`
	Resource UserAuditLogListResponseObjectResultResource `json:"resource"`
	// A UTC RFC3339 timestamp that specifies when the action being logged occured.
	When time.Time                                `json:"when" format:"date-time"`
	JSON userAuditLogListResponseObjectResultJSON `json:"-"`
}

// userAuditLogListResponseObjectResultJSON contains the JSON metadata for the
// struct [UserAuditLogListResponseObjectResult]
type userAuditLogListResponseObjectResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Actor       apijson.Field
	Interface   apijson.Field
	Metadata    apijson.Field
	NewValue    apijson.Field
	OldValue    apijson.Field
	Owner       apijson.Field
	Resource    apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseObjectResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserAuditLogListResponseObjectResultAction struct {
	// A boolean that indicates if the action attempted was successful.
	Result bool `json:"result"`
	// A short string that describes the action that was performed.
	Type string                                         `json:"type"`
	JSON userAuditLogListResponseObjectResultActionJSON `json:"-"`
}

// userAuditLogListResponseObjectResultActionJSON contains the JSON metadata for
// the struct [UserAuditLogListResponseObjectResultAction]
type userAuditLogListResponseObjectResultActionJSON struct {
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseObjectResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserAuditLogListResponseObjectResultActor struct {
	// The ID of the actor that performed the action. If a user performed the action,
	// this will be their User ID.
	ID string `json:"id"`
	// The email of the user that performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IP string `json:"ip"`
	// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
	Type UserAuditLogListResponseObjectResultActorType `json:"type"`
	JSON userAuditLogListResponseObjectResultActorJSON `json:"-"`
}

// userAuditLogListResponseObjectResultActorJSON contains the JSON metadata for the
// struct [UserAuditLogListResponseObjectResultActor]
type userAuditLogListResponseObjectResultActorJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	IP          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseObjectResultActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
type UserAuditLogListResponseObjectResultActorType string

const (
	UserAuditLogListResponseObjectResultActorTypeUser       UserAuditLogListResponseObjectResultActorType = "user"
	UserAuditLogListResponseObjectResultActorTypeAdmin      UserAuditLogListResponseObjectResultActorType = "admin"
	UserAuditLogListResponseObjectResultActorTypeCloudflare UserAuditLogListResponseObjectResultActorType = "Cloudflare"
)

type UserAuditLogListResponseObjectResultOwner struct {
	// Identifier
	ID   string                                        `json:"id"`
	JSON userAuditLogListResponseObjectResultOwnerJSON `json:"-"`
}

// userAuditLogListResponseObjectResultOwnerJSON contains the JSON metadata for the
// struct [UserAuditLogListResponseObjectResultOwner]
type userAuditLogListResponseObjectResultOwnerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseObjectResultOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserAuditLogListResponseObjectResultResource struct {
	// An identifier for the resource that was affected by the action.
	ID string `json:"id"`
	// A short string that describes the resource that was affected by the action.
	Type string                                           `json:"type"`
	JSON userAuditLogListResponseObjectResultResourceJSON `json:"-"`
}

// userAuditLogListResponseObjectResultResourceJSON contains the JSON metadata for
// the struct [UserAuditLogListResponseObjectResultResource]
type userAuditLogListResponseObjectResultResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseObjectResultResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserAuditLogListResponseQ6tjyqK4APIResponseCommon struct {
	Errors   []UserAuditLogListResponseQ6tjyqK4APIResponseCommonError   `json:"errors,required"`
	Messages []UserAuditLogListResponseQ6tjyqK4APIResponseCommonMessage `json:"messages,required"`
	Result   UserAuditLogListResponseQ6tjyqK4APIResponseCommonResult    `json:"result,required"`
	// Whether the API call was successful
	Success UserAuditLogListResponseQ6tjyqK4APIResponseCommonSuccess `json:"success,required"`
	JSON    userAuditLogListResponseQ6tjyqK4APIResponseCommonJSON    `json:"-"`
}

// userAuditLogListResponseQ6tjyqK4APIResponseCommonJSON contains the JSON metadata
// for the struct [UserAuditLogListResponseQ6tjyqK4APIResponseCommon]
type userAuditLogListResponseQ6tjyqK4APIResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseQ6tjyqK4APIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r UserAuditLogListResponseQ6tjyqK4APIResponseCommon) implementsUserAuditLogListResponse() {}

type UserAuditLogListResponseQ6tjyqK4APIResponseCommonError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    userAuditLogListResponseQ6tjyqK4APIResponseCommonErrorJSON `json:"-"`
}

// userAuditLogListResponseQ6tjyqK4APIResponseCommonErrorJSON contains the JSON
// metadata for the struct [UserAuditLogListResponseQ6tjyqK4APIResponseCommonError]
type userAuditLogListResponseQ6tjyqK4APIResponseCommonErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseQ6tjyqK4APIResponseCommonError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserAuditLogListResponseQ6tjyqK4APIResponseCommonMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    userAuditLogListResponseQ6tjyqK4APIResponseCommonMessageJSON `json:"-"`
}

// userAuditLogListResponseQ6tjyqK4APIResponseCommonMessageJSON contains the JSON
// metadata for the struct
// [UserAuditLogListResponseQ6tjyqK4APIResponseCommonMessage]
type userAuditLogListResponseQ6tjyqK4APIResponseCommonMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseQ6tjyqK4APIResponseCommonMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [UserAuditLogListResponseQ6tjyqK4APIResponseCommonResultUnknown],
// [UserAuditLogListResponseQ6tjyqK4APIResponseCommonResultArray] or
// [shared.UnionString].
type UserAuditLogListResponseQ6tjyqK4APIResponseCommonResult interface {
	ImplementsUserAuditLogListResponseQ6tjyqK4APIResponseCommonResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserAuditLogListResponseQ6tjyqK4APIResponseCommonResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserAuditLogListResponseQ6tjyqK4APIResponseCommonResultArray []interface{}

func (r UserAuditLogListResponseQ6tjyqK4APIResponseCommonResultArray) ImplementsUserAuditLogListResponseQ6tjyqK4APIResponseCommonResult() {
}

// Whether the API call was successful
type UserAuditLogListResponseQ6tjyqK4APIResponseCommonSuccess bool

const (
	UserAuditLogListResponseQ6tjyqK4APIResponseCommonSuccessTrue UserAuditLogListResponseQ6tjyqK4APIResponseCommonSuccess = true
)

type UserAuditLogListParams struct {
	// Finds a specific log by its ID.
	ID     param.Field[string]                       `query:"id"`
	Action param.Field[UserAuditLogListParamsAction] `query:"action"`
	Actor  param.Field[UserAuditLogListParamsActor]  `query:"actor"`
	// Limits the returned results to logs older than the specified date. This can be a
	// date string `2019-04-30` or an absolute timestamp that conforms to RFC3339.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Changes the direction of the chronological sorting.
	Direction param.Field[UserAuditLogListParamsDirection] `query:"direction"`
	// Indicates that this request is an export of logs in CSV format.
	Export param.Field[bool] `query:"export"`
	// Indicates whether or not to hide user level audit logs.
	HideUserLogs param.Field[bool] `query:"hide_user_logs"`
	// Defines which page of results to return.
	Page param.Field[float64] `query:"page"`
	// Sets the number of results to return per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Limits the returned results to logs newer than the specified date. This can be a
	// date string `2019-04-30` or an absolute timestamp that conforms to RFC3339.
	Since param.Field[time.Time]                  `query:"since" format:"date-time"`
	Zone  param.Field[UserAuditLogListParamsZone] `query:"zone"`
}

// URLQuery serializes [UserAuditLogListParams]'s query parameters as `url.Values`.
func (r UserAuditLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserAuditLogListParamsAction struct {
	// Filters by the action type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [UserAuditLogListParamsAction]'s query parameters as
// `url.Values`.
func (r UserAuditLogListParamsAction) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserAuditLogListParamsActor struct {
	// Filters by the email address of the actor that made the change.
	Email param.Field[string] `query:"email" format:"email"`
	// Filters by the IP address of the request that made the change by specific IP
	// address or valid CIDR Range.
	IP param.Field[string] `query:"ip"`
}

// URLQuery serializes [UserAuditLogListParamsActor]'s query parameters as
// `url.Values`.
func (r UserAuditLogListParamsActor) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Changes the direction of the chronological sorting.
type UserAuditLogListParamsDirection string

const (
	UserAuditLogListParamsDirectionDesc UserAuditLogListParamsDirection = "desc"
	UserAuditLogListParamsDirectionAsc  UserAuditLogListParamsDirection = "asc"
)

type UserAuditLogListParamsZone struct {
	// Filters by the name of the zone associated to the change.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [UserAuditLogListParamsZone]'s query parameters as
// `url.Values`.
func (r UserAuditLogListParamsZone) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
