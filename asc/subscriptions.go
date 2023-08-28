package asc

import "context"

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
	ReviewNode                string `json:"reviewNode"`
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
			ReviewNode:                reviewNotes,
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
