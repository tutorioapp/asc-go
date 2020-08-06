package asc

import (
	"fmt"
)

// AppInfoLocalization defines model for AppInfoLocalization.
type AppInfoLocalization struct {
	Attributes *struct {
		Locale            *string `json:"locale,omitempty"`
		Name              *string `json:"name,omitempty"`
		PrivacyPolicyText *string `json:"privacyPolicyText,omitempty"`
		PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
		Subtitle          *string `json:"subtitle,omitempty"`
	} `json:"attributes,omitempty"`
	ID            string        `json:"id"`
	Links         ResourceLinks `json:"links"`
	Relationships *struct {
		AppInfo *struct {
			Data  *RelationshipsData  `json:"data,omitempty"`
			Links *RelationshipsLinks `json:"links,omitempty"`
		} `json:"appInfo,omitempty"`
	} `json:"relationships,omitempty"`
	Type string `json:"type"`
}

// AppInfoLocalizationCreateRequest defines model for AppInfoLocalizationCreateRequest.
type AppInfoLocalizationCreateRequest struct {
	Attributes    AppInfoLocalizationCreateRequestAttributes    `json:"attributes"`
	Relationships AppInfoLocalizationCreateRequestRelationships `json:"relationships"`
	Type          string                                        `json:"type"`
}

// AppInfoLocalizationCreateRequestAttributes are attributes for AppInfoLocalizationCreateRequest
type AppInfoLocalizationCreateRequestAttributes struct {
	Locale            string  `json:"locale"`
	Name              *string `json:"name,omitempty"`
	PrivacyPolicyText *string `json:"privacyPolicyText,omitempty"`
	PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
	Subtitle          *string `json:"subtitle,omitempty"`
}

// AppInfoLocalizationCreateRequestRelationships are relationships for AppInfoLocalizationCreateRequest
type AppInfoLocalizationCreateRequestRelationships struct {
	AppInfo struct {
		Data RelationshipsData `json:"data"`
	} `json:"appInfo"`
}

// AppInfoLocalizationResponse defines model for AppInfoLocalizationResponse.
type AppInfoLocalizationResponse struct {
	Data  AppInfoLocalization `json:"data"`
	Links DocumentLinks       `json:"links"`
}

// AppInfoLocalizationUpdateRequest defines model for AppInfoLocalizationUpdateRequest.
type AppInfoLocalizationUpdateRequest struct {
	Attributes *AppInfoLocalizationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                      `json:"id"`
	Type       string                                      `json:"type"`
}

// AppInfoLocalizationUpdateRequestAttributes are attributes for AppInfoLocalizationUpdateRequest
type AppInfoLocalizationUpdateRequestAttributes struct {
	Name              *string `json:"name,omitempty"`
	PrivacyPolicyText *string `json:"privacyPolicyText,omitempty"`
	PrivacyPolicyURL  *string `json:"privacyPolicyUrl,omitempty"`
	Subtitle          *string `json:"subtitle,omitempty"`
}

// AppInfoLocalizationsResponse defines model for AppInfoLocalizationsResponse.
type AppInfoLocalizationsResponse struct {
	Data  []AppInfoLocalization `json:"data"`
	Links PagedDocumentLinks    `json:"links"`
	Meta  *PagingInformation    `json:"meta,omitempty"`
}

// ListAppInfoLocalizationsForAppInfoQuery are query options for ListAppInfoLocalizationsForAppInfo
type ListAppInfoLocalizationsForAppInfoQuery struct {
	FieldsAppInfos             []string `url:"fields[appInfos],omitempty"`
	FieldsAppInfoLocalizations []string `url:"fields[appInfoLocalizations],omitempty"`
	Limit                      int      `url:"limit,omitempty"`
	Include                    []string `url:"include,omitempty"`
	FilterLocale               []string `url:"filter[locale],omitempty"`
	Cursor                     string   `url:"cursor,omitempty"`
}

// GetAppInfoLocalizationQuery are query options for GetAppInfoLocalization
type GetAppInfoLocalizationQuery struct {
	FieldsAppInfoLocalizations []string `url:"fields[appInfoLocalizations],omitempty"`
	Include                    []string `url:"include,omitempty"`
}

// ListAppInfoLocalizationsForAppInfo gets a list of localized, app-level information for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_info_localizations_for_an_app_info
func (s *AppsService) ListAppInfoLocalizationsForAppInfo(id string, params *ListAppInfoLocalizationsForAppInfoQuery) (*AppInfoLocalizationsResponse, *Response, error) {
	url := fmt.Sprintf("appInfos/%s/appInfoLocalizations", id)
	res := new(AppInfoLocalizationsResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// GetAppInfoLocalization reads localized app-level information.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_info_localization_information
func (s *AppsService) GetAppInfoLocalization(id string, params *GetAppInfoLocalizationQuery) (*AppInfoLocalizationResponse, *Response, error) {
	url := fmt.Sprintf("appInfoLocalizations/%s", id)
	res := new(AppInfoLocalizationResponse)
	resp, err := s.client.get(url, params, res)
	return res, resp, err
}

// CreateAppInfoLocalization adds app-level localized information for a new locale.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_info_localization
func (s *AppsService) CreateAppInfoLocalization(body *AppInfoLocalizationCreateRequest) (*AppInfoLocalizationResponse, *Response, error) {
	res := new(AppInfoLocalizationResponse)
	resp, err := s.client.post("appInfoLocalizations", body, res)
	return res, resp, err
}

// UpdateAppInfoLocalization modifies localized app-level information for a particular language.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_info_localization
func (s *AppsService) UpdateAppInfoLocalization(id string, body *AppInfoLocalizationUpdateRequest) (*AppInfoLocalizationResponse, *Response, error) {
	url := fmt.Sprintf("appInfoLocalizations/%s", id)
	res := new(AppInfoLocalizationResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}

// DeleteAppInfoLocalization deletes an app information localization that is associated with an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_info_localization
func (s *AppsService) DeleteAppInfoLocalization(id string) (*Response, error) {
	url := fmt.Sprintf("appInfoLocalizations/%s", id)
	return s.client.delete(url, nil)
}
