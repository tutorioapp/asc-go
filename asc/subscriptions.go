package asc

import (
	"context"
	"net/url"
	"time"
)

// SubscriptionsService handles communication with methods related to Auto-Renewable Subscriptions
//
// https://developer.apple.com/documentation/appstoreconnectapi/app_store/auto-renewable_subscriptions
type SubscriptionsService service

// SubscriptionGroupData defines model for SubscriptionGroupCreateRequest.Data.
//
// https://developer.apple.com/documentation/appstoreconnectapi/subscriptiongroupcreaterequest/data
type SubscriptionGroupData struct {
	Attributes    SubscriptionGroupAttributes `json:"attributes"`
	Relationships struct {
		App struct {
			Data *RelationshipData `json:"data"`
		} `json:"app"`
	} `json:"relationships"`
	Type string `json:"type"`
}

// SubscriptionData defines model for SubscriptionGroupCreateRequest.Data.
//
// https://developer.apple.com/documentation/appstoreconnectapi/subscriptioncreaterequest/data
type SubscriptionData struct {
	Attributes    SubscriptionAttributes `json:"attributes"`
	Relationships struct {
		App struct {
			Data *RelationshipData `json:"data"`
		} `json:"group"`
	} `json:"relationships"`
	Type string `json:"type"`
}

// SubscriptionGroupAttributes defines https://developer.apple.com/documentation/appstoreconnectapi/subscriptiongroup/attributes
type SubscriptionGroupAttributes struct {
	ReferenceName string `json:"referenceName"`
}

// SubscriptionAttributes defines https://developer.apple.com/documentation/appstoreconnectapi/subscriptioncreaterequest/data/attributes
type SubscriptionAttributes struct {
	AvailableInAllTerritories bool   `json:"availableInAllTerritories"`
	FamilySharable            bool   `json:"familySharable"`
	Name                      string `json:"name"`
	ProductID                 string `json:"productId"`
	ReviewNotes               string `json:"reviewNote"`
	SubscriptionPeriod        string `json:"subscriptionPeriod"`
	GroupLevel                int    `json:"groupLevel"`
}

// SubscriptionGroup defines https://developer.apple.com/documentation/appstoreconnectapi/subscriptiongroup
type SubscriptionGroup struct {
	Attributes    SubscriptionGroupAttributes `json:"attributes"`
	ID            string                      `json:"id"`
	Links         ResourceLinks               `json:"links"`
	Relationships any                         `json:"relationships"`
	Type          string                      `json:"type"`
}

type SubscriptionGroupResponse struct {
	Data  SubscriptionGroup `json:"data"`
	Links DocumentLinks     `json:"links"`
}

type Subscription struct {
	Attributes    SubscriptionAttributes `json:"attributes"`
	ID            string                 `json:"id"`
	Links         ResourceLinks          `json:"links"`
	Relationships any                    `json:"relationships"`
	Type          string                 `json:"type"`
}

type SubscriptionResponse struct {
	Data  Subscription  `json:"data"`
	Links DocumentLinks `json:"links"`
}

type SubscriptionPricePointAttributes struct {
	CustomerPrice string `json:"customerPrice"`
	Proceeds      string `json:"proceeds"`
	ProceedsYear2 string `json:"proceedsYear2"`
}

type SubscriptionPricePoint struct {
	Attributes    SubscriptionPricePointAttributes `json:"attributes"`
	ID            string                           `json:"id"`
	Links         ResourceLinks                    `json:"links"`
	Relationships any                              `json:"relationships"`
	Type          string                           `json:"type"`
}

type SubscriptionPricePointsResponse struct {
	Data  []SubscriptionPricePoint `json:"data"`
	Links DocumentLinks            `json:"links"`
}

type SubscriptionLocalization struct {
	Attributes    SubscriptionLocalizationAttributes `json:"attributes"`
	ID            string                             `json:"id"`
	Links         ResourceLinks                      `json:"links"`
	Relationships any                                `json:"relationships"`
	Type          string                             `json:"type"`
}

type SubscriptionLocalizationResponse struct {
	Data  SubscriptionLocalization `json:"data"`
	Links DocumentLinks            `json:"links"`
}

type SubscriptionPriceCreateAttributes struct {
	PreserveCurrentPrice bool      `json:"preserveCurrentPrice"`
	Preserved            bool      `json:"preserved,omitempty"`
	StartDate            time.Time `json:"startDate"`
}

type SubscriptionPriceCreate struct {
	Attributes    SubscriptionPriceCreateAttributes `json:"attributes"`
	ID            string                            `json:"id"`
	Links         ResourceLinks                     `json:"links"`
	Relationships any                               `json:"relationships"`
	Type          string                            `json:"type"`
}

type SubscriptionPriceCreateResponse struct {
	Data  SubscriptionPriceCreate `json:"data"`
	Links DocumentLinks           `json:"links"`
}

type SubscriptionGroupLocalization struct {
	Attributes    SubscriptionGroupLocalizationAttributes `json:"attributes"`
	ID            string                                  `json:"id"`
	Links         ResourceLinks                           `json:"links"`
	Relationships any                                     `json:"relationships"`
	Type          string                                  `json:"type"`
}

type SubscriptionGroupLocalizationResponse struct {
	Data  SubscriptionGroupLocalization `json:"data"`
	Links DocumentLinks                 `json:"links"`
}

type SubscriptionLocalizationAttributes struct {
	Locale      string `json:"locale"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SubscriptionGroupLocalizationAttributes struct {
	Locale        string `json:"locale"`
	CustomAppName string `json:"customAppName"`
	Name          string `json:"name"`
}

// SubscriptionLocalizationData defines model for SubscriptionLocalizationCreateRequest.Data
//
// https://developer.apple.com/documentation/appstoreconnectapi/subscriptionlocalizationcreaterequest/data
type SubscriptionLocalizationData struct {
	Attributes    SubscriptionLocalizationAttributes `json:"attributes"`
	Relationships struct {
		App struct {
			Data *RelationshipData `json:"data"`
		} `json:"subscription"`
	} `json:"relationships"`
	Type string `json:"type"`
}

// SubscriptionGroupLocalizationData defines model for SubscriptionGroupLocalizationCreateRequest.Data
//
// https://developer.apple.com/documentation/appstoreconnectapi/subscriptiongrouplocalizationcreaterequest/data
type SubscriptionGroupLocalizationData struct {
	Attributes    SubscriptionGroupLocalizationAttributes `json:"attributes"`
	Relationships struct {
		App struct {
			Data *RelationshipData `json:"data"`
		} `json:"subscriptionGroup"`
	} `json:"relationships"`
	Type string `json:"type"`
}

// SubscriptionPriceCreateData defines model for SubscriptionPriceCreateRequest.Data
//
// https://developer.apple.com/documentation/appstoreconnectapi/subscriptionpricecreaterequest/data
type SubscriptionPriceCreateData struct {
	Attributes    SubscriptionPriceCreateAttributes `json:"attributes"`
	Relationships struct {
		Subscription struct {
			Data *RelationshipData `json:"data"`
		} `json:"subscription"`
		SubscriptionPricePoint struct {
			Data *RelationshipData `json:"data"`
		} `json:"subscriptionPricePoint"`
	} `json:"relationships"`
	Type string `json:"type"`
}

// CreateSubscriptionGroup creates a subscription group for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_subscription_group
func (s *SubscriptionsService) CreateSubscriptionGroup(ctx context.Context, appID, groupName string) (*SubscriptionGroupResponse, *Response, error) {
	res := new(SubscriptionGroupResponse)
	resp, err := s.client.post(ctx, "v1/subscriptionGroups", newRequestBody(SubscriptionGroupData{
		Attributes: SubscriptionGroupAttributes{ReferenceName: groupName},
		Relationships: struct {
			App struct {
				Data *RelationshipData `json:"data"`
			} `json:"app"`
		}{
			App: struct {
				Data *RelationshipData `json:"data"`
			}{Data: &RelationshipData{ID: appID, Type: "apps"}},
		},
		Type: "subscriptionGroups",
	}), res)
	return res, resp, err
}

// CreateSubscriptionGroupLocalization creates a subscription group localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_subscription_group_localization
func (s *SubscriptionsService) CreateSubscriptionGroupLocalization(ctx context.Context, groupID, locale, appName, groupName string) (*SubscriptionGroupLocalizationResponse, *Response, error) {
	res := new(SubscriptionGroupLocalizationResponse)
	resp, err := s.client.post(ctx, "v1/subscriptionGroupLocalizations", newRequestBody(SubscriptionGroupLocalizationData{
		Attributes: SubscriptionGroupLocalizationAttributes{
			Locale:        locale,
			Name:          groupName,
			CustomAppName: appName,
		},
		Relationships: struct {
			App struct {
				Data *RelationshipData `json:"data"`
			} `json:"subscriptionGroup"`
		}{
			App: struct {
				Data *RelationshipData `json:"data"`
			}{Data: &RelationshipData{ID: groupID, Type: "subscriptionGroups"}},
		},
		Type: "subscriptionGroupLocalizations",
	}), res)
	return res, resp, err
}

type SubscriptionPeriod string

const (
	SubscriptionPeriodP1W SubscriptionPeriod = "ONE_WEEK"
	SubscriptionPeriodP1M SubscriptionPeriod = "ONE_MONTH"
	SubscriptionPeriodP2M SubscriptionPeriod = "TWO_MONTHS"
	SubscriptionPeriodP3M SubscriptionPeriod = "THREE_MONTHS"
	SubscriptionPeriodP6M SubscriptionPeriod = "SIX_MONTHS"
	SubscriptionPeriodP1Y SubscriptionPeriod = "ONE_YEAR"
)

// CreateSubscription creates a subscription for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_auto-renewable_subscription
func (s *SubscriptionsService) CreateSubscription(ctx context.Context, name, productID, groupID, reviewNotes string, period SubscriptionPeriod) (*SubscriptionResponse, *Response, error) {
	res := new(SubscriptionResponse)
	resp, err := s.client.post(ctx, "v1/subscriptions", newRequestBody(SubscriptionData{
		Attributes: SubscriptionAttributes{
			AvailableInAllTerritories: true,
			FamilySharable:            false,
			Name:                      name,
			ProductID:                 productID,
			ReviewNotes:               reviewNotes,
			SubscriptionPeriod:        string(period),
			GroupLevel:                1,
		},
		Relationships: struct {
			App struct {
				Data *RelationshipData `json:"data"`
			} `json:"group"`
		}{
			App: struct {
				Data *RelationshipData `json:"data"`
			}{Data: &RelationshipData{ID: groupID, Type: "subscriptionGroups"}},
		},
		Type: "subscriptions",
	}), res)
	return res, resp, err
}

// CreateSubscriptionLocalization creates a subscription localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_subscription_localization
func (s *SubscriptionsService) CreateSubscriptionLocalization(ctx context.Context, subscriptionID, locale, name, description string) (*SubscriptionLocalizationResponse, *Response, error) {
	res := new(SubscriptionLocalizationResponse)
	resp, err := s.client.post(ctx, "v1/subscriptionLocalizations", newRequestBody(SubscriptionLocalizationData{
		Attributes: SubscriptionLocalizationAttributes{
			Locale:      locale,
			Name:        name,
			Description: description,
		},
		Relationships: struct {
			App struct {
				Data *RelationshipData `json:"data"`
			} `json:"subscription"`
		}{
			App: struct {
				Data *RelationshipData `json:"data"`
			}{Data: &RelationshipData{ID: subscriptionID, Type: "subscriptions"}},
		},
		Type: "subscriptionLocalizations",
	}), res)
	return res, resp, err
}

// CreateSubscriptionPriceChange schedules a subscription price change for a specific territory.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_subscription_price_change
func (s *SubscriptionsService) CreateSubscriptionPriceChange(ctx context.Context, subscriptionID, priceID string, preserveCurrentPrice bool) (*SubscriptionPriceCreateResponse, *Response, error) {
	res := new(SubscriptionPriceCreateResponse)
	resp, err := s.client.post(ctx, "v1/subscriptionPrices", newRequestBody(SubscriptionPriceCreateData{
		Attributes: SubscriptionPriceCreateAttributes{
			PreserveCurrentPrice: preserveCurrentPrice,
			StartDate:            time.Now(),
		},
		Relationships: struct {
			Subscription struct {
				Data *RelationshipData `json:"data"`
			} `json:"subscription"`
			SubscriptionPricePoint struct {
				Data *RelationshipData `json:"data"`
			} `json:"subscriptionPricePoint"`
		}{
			Subscription: struct {
				Data *RelationshipData `json:"data"`
			}{Data: &RelationshipData{ID: subscriptionID, Type: "subscriptions"}},
			SubscriptionPricePoint: struct {
				Data *RelationshipData `json:"data"`
			}{Data: &RelationshipData{ID: priceID, Type: "subscriptionPricePoints"}},
		},
		Type: "subscriptionPrices",
	}), res)
	return res, resp, err
}

// GetSubscription returns the subscription for a given subscription ID.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_subscription_information
func (s *SubscriptionsService) GetSubscription(ctx context.Context, id string) (*SubscriptionResponse, *Response, error) {
	res := new(SubscriptionResponse)
	resp, err := s.client.get(ctx, "v1/subscriptions/"+id, nil, res)
	return res, resp, err
}

// GetSubscriptionPricePoints returns a list of approved prices apple will allow you to set for a subscription.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_subscription_price_point_information
func (s *SubscriptionsService) GetSubscriptionPricePoints(ctx context.Context, id, territory string) (*SubscriptionPricePointsResponse, *Response, error) {
	res := new(SubscriptionPricePointsResponse)
	resp, err := s.client.get(ctx, "v1/subscriptions/"+id+"/pricePoints?include=territory&filter[territory]="+url.QueryEscape(territory)+"&limit=8000", nil, res)
	return res, resp, err
}
