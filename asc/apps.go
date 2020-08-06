package asc

import (
	"fmt"
)

// AppsService handles communication with build-related methods of the App Store Connect API
//
// https://developer.apple.com/documentation/appstoreconnectapi/apps
// https://developer.apple.com/documentation/appstoreconnectapi/app_metadata
type AppsService service

// Platform defines model for Platform.
type Platform string

// List of Platform
const (
	PlatformIOS   Platform = "IOS"
	PlatformMACOS Platform = "MAC_OS"
	PlatformTVOS  Platform = "TV_OS"
)

// App defines model for App.
type App struct {
	Attributes *struct {
		AvailableInNewTerritories *bool   `json:"availableInNewTerritories,omitempty"`
		BundleID                  *string `json:"bundleId,omitempty"`
		ContentRightsDeclaration  *string `json:"contentRightsDeclaration,omitempty"`
		IsOrEverWasMadeForKids    *bool   `json:"isOrEverWasMadeForKids,omitempty"`
		Name                      *string `json:"name,omitempty"`
		PrimaryLocale             *string `json:"primaryLocale,omitempty"`
		Sku                       *string `json:"sku,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		AppInfos *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"appInfos,omitempty"`
		AppStoreVersions *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"appStoreVersions,omitempty"`
		AvailableTerritories *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"availableTerritories,omitempty"`
		BetaAppLocalizations *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"betaAppLocalizations,omitempty"`
		BetaAppReviewDetail *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"betaAppReviewDetail,omitempty"`
		BetaGroups *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"betaGroups,omitempty"`
		BetaLicenseAgreement *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"betaLicenseAgreement,omitempty"`
		Builds *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"builds,omitempty"`
		EndUserLicenseAgreement *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"endUserLicenseAgreement,omitempty"`
		GameCenterEnabledVersions *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"gameCenterEnabledVersions,omitempty"`
		InAppPurchases *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"inAppPurchases,omitempty"`
		PreOrder *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"preOrder,omitempty"`
		PreReleaseVersions *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"preReleaseVersions,omitempty"`
		Prices *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"prices,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppUpdateRequest defines model for AppUpdateRequest.
type AppUpdateRequest struct {
	Attributes    *AppUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                         `json:"id"`
	Relationships *AppUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                         `json:"type"`
}

// AppUpdateRequestAttributes are attributes for AppUpdateRequest
type AppUpdateRequestAttributes struct {
	AvailableInNewTerritories *bool   `json:"availableInNewTerritories,omitempty"`
	BundleID                  *string `json:"bundleId,omitempty"`
	ContentRightsDeclaration  *string `json:"contentRightsDeclaration,omitempty"`
	PrimaryLocale             *string `json:"primaryLocale,omitempty"`
}

// AppUpdateRequestRelationships are relationships for AppUpdateRequest
type AppUpdateRequestRelationships struct {
	AvailableTerritories *struct {
		Data *[]RelationshipsData `json:"data,omitempty"`
	} `json:"availableTerritories,omitempty"`
	Prices *struct {
		Data *[]RelationshipsData `json:"data,omitempty"`
	} `json:"prices,omitempty"`
}

// AppResponse defines model for AppResponse.
type AppResponse struct {
	Data     App            `json:"data"`
	Included *[]interface{} `json:"included,omitempty"`
	Links    DocumentLinks  `json:"links"`
}

// AppsResponse defines model for AppsResponse.
type AppsResponse struct {
	Data     []App              `json:"data"`
	Included *[]interface{}     `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// InAppPurchase defines model for InAppPurchase.
type InAppPurchase struct {
	Attributes *struct {
		InAppPurchaseType *string `json:"inAppPurchaseType,omitempty"`
		ProductID         *string `json:"productId,omitempty"`
		ReferenceName     *string `json:"referenceName,omitempty"`
		State             *string `json:"state,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		Apps *struct {
			Data  *[]RelationshipsData `json:"data,omitempty"`
			Links *RelationshipsLinks  `json:"links,omitempty"`
			Meta  *PagingInformation   `json:"meta,omitempty"`
		} `json:"apps,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// InAppPurchaseResponse defines model for InAppPurchaseResponse.
type InAppPurchaseResponse struct {
	Data  InAppPurchase `json:"data"`
	Links DocumentLinks `json:"links"`
}

// InAppPurchasesResponse defines model for InAppPurchasesResponse.
type InAppPurchasesResponse struct {
	Data  []InAppPurchase    `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// ListAppsQuery are query options for ListApps
type ListAppsQuery struct {
	FieldsApps                          []string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements         []string `url:"fields[betaLicenseAgreements],omitempty"`
	FieldsPreReleaseVersions            []string `url:"fields[preReleaseVersions],omitempty"`
	FieldsBetaAppReviewDetails          []string `url:"fields[betaAppReviewDetails],omitempty"`
	FieldsBetaAppLocalizations          []string `url:"fields[betaAppLocalizations],omitempty"`
	FieldsBuilds                        []string `url:"fields[builds],omitempty"`
	FieldsBetaGroups                    []string `url:"fields[betaGroups],omitempty"`
	FieldsEndUserLicenseAgreements      []string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsAppStoreVersions              []string `url:"fields[appStoreVersions],omitempty"`
	FieldsTerritories                   []string `url:"fields[territories],omitempty"`
	FieldsAppPrices                     []string `url:"fields[appPrices],omitempty"`
	FieldsAppPreOrders                  []string `url:"fields[appPreOrders],omitempty"`
	FieldsAppInfos                      []string `url:"fields[appInfos],omitempty"`
	FieldsPerfPowerMetrics              []string `url:"fields[perfPowerMetrics],omitempty"`
	FieldsInAppPurchases                []string `url:"fields[inAppPurchases],omitempty"`
	FilterBundleID                      []string `url:"filter[bundleId],omitempty"`
	FilterID                            []string `url:"filter[id],omitempty"`
	FilterName                          []string `url:"filter[name],omitempty"`
	FilterSKU                           []string `url:"filter[sku],omitempty"`
	FilterAppStoreVersions              []string `url:"filter[appStoreVersions],omitempty"`
	FilterAppStoreVersionsPlatform      []string `url:"filter[appStoreVersionsPlatform],omitempty"`
	FilterAppStoreVersionsAppStoreState []string `url:"filter[appStoreVersionsAppStoreState],omitempty"`
	FilterGameCenterEnabledVersions     []string `url:"filter[gameCenterEnabledVersions],omitempty"`
	Include                             []string `url:"include,omitempty"`
	Limit                               int      `url:"limit,omitempty"`
	LimitPreReleaseVersions             int      `url:"limit[preReleaseVersions],omitempty"`
	LimitBuilds                         int      `url:"limit[builds],omitempty"`
	LimitBetaGroups                     int      `url:"limit[betaGroups],omitempty"`
	LimitBetaAppLocalizations           int      `url:"limit[betaAppLocalizations],omitempty"`
	LimitPrices                         int      `url:"limit[prices],omitempty"`
	LimitAvailableTerritories           int      `url:"limit[availableTerritories],omitempty"`
	LimitAppStoreVersions               int      `url:"limit[appStoreVersions],omitempty"`
	LimitAppInfos                       int      `url:"limit[appInfos],omitempty"`
	LimitGameCenterEnabledVersions      int      `url:"limit[gameCenterEnabledVersions],omitempty"`
	LimitInAppPurchases                 int      `url:"limit[inAppPurchases],omitempty"`
	Sort                                []string `url:"sort,omitempty"`
	ExistsGameCenterEnabledVersions     []string `url:"exists[gameCenterEnabledVersions],omitempty"`
	Cursor                              string   `url:"cursor,omitempty"`
}

// GetAppQuery are query options for GetApp
type GetAppQuery struct {
	FieldsApps                      []string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements     []string `url:"fields[betaLicenseAgreements],omitempty"`
	FieldsPreReleaseVersions        []string `url:"fields[preReleaseVersions],omitempty"`
	FieldsBetaAppReviewDetails      []string `url:"fields[betaAppReviewDetails],omitempty"`
	FieldsBetaAppLocalizations      []string `url:"fields[betaAppLocalizations],omitempty"`
	FieldsBuilds                    []string `url:"fields[builds],omitempty"`
	FieldsBetaGroups                []string `url:"fields[betaGroups],omitempty"`
	FieldsEndUserLicenseAgreements  []string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsAppStoreVersions          []string `url:"fields[appStoreVersions],omitempty"`
	FieldsTerritories               []string `url:"fields[territories],omitempty"`
	FieldsAppPrices                 []string `url:"fields[appPrices],omitempty"`
	FieldsAppPreOrders              []string `url:"fields[appPreOrders],omitempty"`
	FieldsAppInfos                  []string `url:"fields[appInfos],omitempty"`
	FieldsPerfPowerMetrics          []string `url:"fields[perfPowerMetrics],omitempty"`
	FieldsGameCenterEnabledVersions []string `url:"fields[gameCenterEnabledVersions],omitempty"`
	FieldsInAppPurchases            []string `url:"fields[inAppPurchases],omitempty"`
	Include                         []string `url:"include,omitempty"`
	LimitPreReleaseVersions         int      `url:"limit[preReleaseVersions],omitempty"`
	LimitBuilds                     int      `url:"limit[builds],omitempty"`
	LimitBetaGroups                 int      `url:"limit[betaGroups],omitempty"`
	LimitBetaAppLocalizations       int      `url:"limit[betaAppLocalizations],omitempty"`
	LimitPrices                     int      `url:"limit[prices],omitempty"`
	LimitAvailableTerritories       int      `url:"limit[availableTerritories],omitempty"`
	LimitAppStoreVersions           int      `url:"limit[appStoreVersions],omitempty"`
	LimitAppInfos                   int      `url:"limit[appInfos],omitempty"`
	LimitGameCenterEnabledVersions  int      `url:"limit[gameCenterEnabledVersions],omitempty"`
	LimitInAppPurchases             int      `url:"limit[inAppPurchases],omitempty"`
}

// ListInAppPurchasesQuery are query options for ListInAppPurchases
type ListInAppPurchasesQuery struct {
	FieldsApps              []string `url:"fields[apps],omitempty"`
	FieldsInAppPurchases    []string `url:"fields[inAppPurchases],omitempty"`
	FilterCanBeSubmitted    []string `url:"filter[canBeSubmitted],omitempty"`
	FilterInAppPurchaseType []string `url:"filter[inAppPurchaseType],omitempty"`
	Limit                   int      `url:"limit,omitempty"`
	Include                 []string `url:"include,omitempty"`
	Sort                    []string `url:"sort,omitempty"`
	Cursor                  string   `url:"cursor,omitempty"`
}

// GetInAppPurchaseQuery are query options for GetInAppPurchase
type GetInAppPurchaseQuery struct {
	FieldsInAppPurchases []string `url:"fields[inAppPurchases],omitempty"`
	Include              []string `url:"include,omitempty"`
	LimitApps            int      `url:"limit[apps],omitempty"`
}

// ListApps finds and lists apps added in App Store Connect.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_apps
func (s *AppsService) ListApps(params *ListAppsQuery) (*AppsResponse, *Response, error) {
	res := new(AppsResponse)
	resp, err := s.client.get("apps", params, res)
	return res, resp, err
}

// GetApp gets information about a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_information
func (s *AppsService) GetApp(id string, params *GetAppQuery) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s", id)
	res := new(AppResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// UpdateApp updates app information including bundle ID, primary locale, price schedule, and global availability.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app
func (s *AppsService) UpdateApp(id string, body *AppUpdateRequest) (*AppResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s", id)
	res := new(AppResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}

// RemoveBetaTestersFromApp removes one or more beta testers' access to test any builds of a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/remove_beta_testers_from_all_groups_and_builds_of_an_app
func (s *AppsService) RemoveBetaTestersFromApp(id string, linkages *[]RelationshipsData) (*Response, error) {
	url := fmt.Sprintf("apps/%s/relationships/betaTesters", id)
	return s.client.delete(url, linkages)
}

// ListInAppPurchasesForApp lists the in-app purchases that are available for your app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_in-app_purchases_for_an_app
func (s *AppsService) ListInAppPurchasesForApp(id string, params *ListInAppPurchasesQuery) (*InAppPurchasesResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/inAppPurchases", id)
	res := new(InAppPurchasesResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetInAppPurchase gets information about an in-app purchase.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_in-app_purchase_information
func (s *AppsService) GetInAppPurchase(id string, params *GetInAppPurchaseQuery) (*InAppPurchaseResponse, *Response, error) {
	url := fmt.Sprintf("inAppPurchases/%s", id)
	res := new(InAppPurchaseResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}
