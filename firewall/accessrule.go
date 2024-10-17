// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// AccessRuleService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessRuleService] method instead.
type AccessRuleService struct {
	Options []option.RequestOption
}

// NewAccessRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessRuleService(opts ...option.RequestOption) (r *AccessRuleService) {
	r = &AccessRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for an account or zone. The rule will apply to all
// zones in the account or zone.
//
// Note: To create an IP Access rule that applies to a single zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *AccessRuleService) New(ctx context.Context, params AccessRuleNewParams, opts ...option.RequestOption) (res *WAFRule, err error) {
	var env AccessRuleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches IP Access rules of an account or zone. These rules apply to all the
// zones in the account or zone. You can filter the results using several optional
// parameters.
func (r *AccessRuleService) List(ctx context.Context, params AccessRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WAFRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Fetches IP Access rules of an account or zone. These rules apply to all the
// zones in the account or zone. You can filter the results using several optional
// parameters.
func (r *AccessRuleService) ListAutoPaging(ctx context.Context, params AccessRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WAFRule] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type AccessRuleCIDRConfigurationParam struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[AccessRuleCIDRConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleCIDRConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleCIDRConfigurationParam) implementsFirewallAccessRuleNewParamsConfigurationUnion() {}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type AccessRuleCIDRConfigurationTarget string

const (
	AccessRuleCIDRConfigurationTargetIPRange AccessRuleCIDRConfigurationTarget = "ip_range"
)

func (r AccessRuleCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type AccessRuleIPConfigurationParam struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccessRuleIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleIPConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleIPConfigurationParam) implementsFirewallAccessRuleNewParamsConfigurationUnion() {}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleIPConfigurationTarget string

const (
	AccessRuleIPConfigurationTargetIP AccessRuleIPConfigurationTarget = "ip"
)

func (r AccessRuleIPConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleIPConfigurationTargetIP:
		return true
	}
	return false
}

type ASNConfigurationParam struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[ASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r ASNConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ASNConfigurationParam) implementsFirewallAccessRuleNewParamsConfigurationUnion() {}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type ASNConfigurationTarget string

const (
	ASNConfigurationTargetASN ASNConfigurationTarget = "asn"
)

func (r ASNConfigurationTarget) IsKnown() bool {
	switch r {
	case ASNConfigurationTargetASN:
		return true
	}
	return false
}

type CountryConfigurationParam struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[CountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r CountryConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CountryConfigurationParam) implementsFirewallAccessRuleNewParamsConfigurationUnion() {}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type CountryConfigurationTarget string

const (
	CountryConfigurationTargetCountry CountryConfigurationTarget = "country"
)

func (r CountryConfigurationTarget) IsKnown() bool {
	switch r {
	case CountryConfigurationTargetCountry:
		return true
	}
	return false
}

type IPV6ConfigurationParam struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[IPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r IPV6ConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPV6ConfigurationParam) implementsFirewallAccessRuleNewParamsConfigurationUnion() {}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type IPV6ConfigurationTarget string

const (
	IPV6ConfigurationTargetIp6 IPV6ConfigurationTarget = "ip6"
)

func (r IPV6ConfigurationTarget) IsKnown() bool {
	switch r {
	case IPV6ConfigurationTargetIp6:
		return true
	}
	return false
}

type AccessRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[AccessRuleNewParamsConfigurationUnion] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[AccessRuleNewParamsMode] `json:"mode,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r AccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
type AccessRuleNewParamsConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccessRuleNewParamsConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleNewParamsConfiguration) implementsFirewallAccessRuleNewParamsConfigurationUnion() {}

// The rule configuration.
//
// Satisfied by [firewall.AccessRuleIPConfigurationParam],
// [firewall.IPV6ConfigurationParam], [firewall.AccessRuleCIDRConfigurationParam],
// [firewall.ASNConfigurationParam], [firewall.CountryConfigurationParam],
// [AccessRuleNewParamsConfiguration].
type AccessRuleNewParamsConfigurationUnion interface {
	implementsFirewallAccessRuleNewParamsConfigurationUnion()
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleNewParamsConfigurationTarget string

const (
	AccessRuleNewParamsConfigurationTargetIP      AccessRuleNewParamsConfigurationTarget = "ip"
	AccessRuleNewParamsConfigurationTargetIp6     AccessRuleNewParamsConfigurationTarget = "ip6"
	AccessRuleNewParamsConfigurationTargetIPRange AccessRuleNewParamsConfigurationTarget = "ip_range"
	AccessRuleNewParamsConfigurationTargetASN     AccessRuleNewParamsConfigurationTarget = "asn"
	AccessRuleNewParamsConfigurationTargetCountry AccessRuleNewParamsConfigurationTarget = "country"
)

func (r AccessRuleNewParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsConfigurationTargetIP, AccessRuleNewParamsConfigurationTargetIp6, AccessRuleNewParamsConfigurationTargetIPRange, AccessRuleNewParamsConfigurationTargetASN, AccessRuleNewParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleNewParamsMode string

const (
	AccessRuleNewParamsModeBlock            AccessRuleNewParamsMode = "block"
	AccessRuleNewParamsModeChallenge        AccessRuleNewParamsMode = "challenge"
	AccessRuleNewParamsModeWhitelist        AccessRuleNewParamsMode = "whitelist"
	AccessRuleNewParamsModeJSChallenge      AccessRuleNewParamsMode = "js_challenge"
	AccessRuleNewParamsModeManagedChallenge AccessRuleNewParamsMode = "managed_challenge"
)

func (r AccessRuleNewParamsMode) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsModeBlock, AccessRuleNewParamsModeChallenge, AccessRuleNewParamsModeWhitelist, AccessRuleNewParamsModeJSChallenge, AccessRuleNewParamsModeManagedChallenge:
		return true
	}
	return false
}

type AccessRuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Result WAFRule `json:"result,required"`
	// Whether the API call was successful
	Success AccessRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessRuleNewResponseEnvelopeJSON    `json:"-"`
}

// accessRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessRuleNewResponseEnvelope]
type accessRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessRuleNewResponseEnvelopeSuccess bool

const (
	AccessRuleNewResponseEnvelopeSuccessTrue AccessRuleNewResponseEnvelopeSuccess = true
)

func (r AccessRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessRuleListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID        param.Field[string]                            `path:"zone_id"`
	Configuration param.Field[AccessRuleListParamsConfiguration] `query:"configuration"`
	// The direction used to sort returned rules.
	Direction param.Field[AccessRuleListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[AccessRuleListParamsMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[AccessRuleListParamsMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
	// The field used to sort returned rules.
	Order param.Field[AccessRuleListParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccessRuleListParams]'s query parameters as `url.Values`.
func (r AccessRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessRuleListParamsConfiguration struct {
	// The target to search in existing rules.
	Target param.Field[AccessRuleListParamsConfigurationTarget] `query:"target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [AccessRuleListParamsConfiguration]'s query parameters as
// `url.Values`.
func (r AccessRuleListParamsConfiguration) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The target to search in existing rules.
type AccessRuleListParamsConfigurationTarget string

const (
	AccessRuleListParamsConfigurationTargetIP      AccessRuleListParamsConfigurationTarget = "ip"
	AccessRuleListParamsConfigurationTargetIPRange AccessRuleListParamsConfigurationTarget = "ip_range"
	AccessRuleListParamsConfigurationTargetASN     AccessRuleListParamsConfigurationTarget = "asn"
	AccessRuleListParamsConfigurationTargetCountry AccessRuleListParamsConfigurationTarget = "country"
)

func (r AccessRuleListParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleListParamsConfigurationTargetIP, AccessRuleListParamsConfigurationTargetIPRange, AccessRuleListParamsConfigurationTargetASN, AccessRuleListParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// The direction used to sort returned rules.
type AccessRuleListParamsDirection string

const (
	AccessRuleListParamsDirectionAsc  AccessRuleListParamsDirection = "asc"
	AccessRuleListParamsDirectionDesc AccessRuleListParamsDirection = "desc"
)

func (r AccessRuleListParamsDirection) IsKnown() bool {
	switch r {
	case AccessRuleListParamsDirectionAsc, AccessRuleListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type AccessRuleListParamsMatch string

const (
	AccessRuleListParamsMatchAny AccessRuleListParamsMatch = "any"
	AccessRuleListParamsMatchAll AccessRuleListParamsMatch = "all"
)

func (r AccessRuleListParamsMatch) IsKnown() bool {
	switch r {
	case AccessRuleListParamsMatchAny, AccessRuleListParamsMatchAll:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleListParamsMode string

const (
	AccessRuleListParamsModeBlock            AccessRuleListParamsMode = "block"
	AccessRuleListParamsModeChallenge        AccessRuleListParamsMode = "challenge"
	AccessRuleListParamsModeWhitelist        AccessRuleListParamsMode = "whitelist"
	AccessRuleListParamsModeJSChallenge      AccessRuleListParamsMode = "js_challenge"
	AccessRuleListParamsModeManagedChallenge AccessRuleListParamsMode = "managed_challenge"
)

func (r AccessRuleListParamsMode) IsKnown() bool {
	switch r {
	case AccessRuleListParamsModeBlock, AccessRuleListParamsModeChallenge, AccessRuleListParamsModeWhitelist, AccessRuleListParamsModeJSChallenge, AccessRuleListParamsModeManagedChallenge:
		return true
	}
	return false
}

// The field used to sort returned rules.
type AccessRuleListParamsOrder string

const (
	AccessRuleListParamsOrderConfigurationTarget AccessRuleListParamsOrder = "configuration.target"
	AccessRuleListParamsOrderConfigurationValue  AccessRuleListParamsOrder = "configuration.value"
	AccessRuleListParamsOrderMode                AccessRuleListParamsOrder = "mode"
)

func (r AccessRuleListParamsOrder) IsKnown() bool {
	switch r {
	case AccessRuleListParamsOrderConfigurationTarget, AccessRuleListParamsOrderConfigurationValue, AccessRuleListParamsOrderMode:
		return true
	}
	return false
}
