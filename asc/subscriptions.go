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
	Relationships *Relationship               `json:"relationships"`
	Type          string                      `json:"type"`
}

// SubscriptionGroupAttributes defines https://developer.apple.com/documentation/appstoreconnectapi/subscriptiongroup/attributes
type SubscriptionGroupAttributes struct {
	ReferenceName string `json:"referenceName"`
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

// CreateSubscriptionGroup creates a subscription group for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_subscription_group
func (s *SubscriptionsService) CreateSubscriptionGroup(ctx context.Context, appID, groupName string) (*SubscriptionGroupResponse, *Response, error) {
	res := new(SubscriptionGroupResponse)
	resp, err := s.client.post(ctx, "v1/subscriptionGroups", newRequestBody(SubscriptionGroupData{
		Attributes:    SubscriptionGroupAttributes{ReferenceName: groupName},
		Relationships: &Relationship{Data: &RelationshipData{ID: appID, Type: "apps"}},
		Type:          "subscriptionGroups",
	}), res)
	return res, resp, err
}
