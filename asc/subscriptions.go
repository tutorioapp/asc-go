package asc

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
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
	FamilySharable     bool   `json:"familySharable"`
	Name               string `json:"name"`
	ProductID          string `json:"productId"`
	ReviewNotes        string `json:"reviewNote"`
	SubscriptionPeriod string `json:"subscriptionPeriod"`
	GroupLevel         int    `json:"groupLevel"`
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

type SubscriptionsResponse struct {
	Data  []Subscription `json:"data"`
	Links DocumentLinks  `json:"links"`
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

type SubscriptionPriceResponse struct {
	Data  []SubscriptionPrice `json:"data"`
	Links DocumentLinks       `json:"links"`
}

type SubscriptionPrice struct {
	Attributes    SubscriptionPriceAttributes `json:"attributes"`
	ID            string                      `json:"id"`
	Links         ResourceLinks               `json:"links"`
	Relationships any                         `json:"relationships"`
	Type          string                      `json:"type"`
}

type SubscriptionPriceAttributes struct {
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
	PreserveCurrentPrice bool `json:"preserveCurrentPrice"`
	Preserved            bool `json:"preserved,omitempty"`
	StartDate            any  `json:"startDate"`
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

type SubscriptionUpdateRequestData struct {
	Attributes    SubscriptionUpdateAttributes    `json:"attributes"`
	ID            string                          `json:"id"`
	Relationships SubscriptionUpdateRelationships `json:"relationships"`
	Type          string                          `json:"type"`
}

type SubscriptionUpdateAttributes struct {
	FamilySharable     bool   `json:"familySharable"`
	Name               string `json:"name"`
	ReviewNote         string `json:"reviewNote"`
	SubscriptionPeriod string `json:"subscriptionPeriod"`
	GroupLevel         int    `json:"groupLevel"`
}

type SubscriptionUpdateRelationships struct {
	IntroductoryOffers struct {
		Data []*RelationshipData `json:"data"`
	} `json:"introductoryOffers"`
	Prices struct {
		Data []*RelationshipData `json:"data"`
	} `json:"prices"`
	PromotionalOffers struct {
		Data []*RelationshipData `json:"data"`
	} `json:"promotionalOffers"`
}

type SubscriptionPriceUpdateRelationships struct {
	Prices struct {
		Data []*RelationshipData `json:"data"`
	} `json:"prices"`
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
	ID            string                            `json:"id"`
	Relationships struct {
		Subscription struct {
			Data *RelationshipData `json:"data"`
		} `json:"subscription"`
		SubscriptionPricePoint struct {
			Data *RelationshipData `json:"data"`
		} `json:"subscriptionPricePoint"`
		Territory struct {
			Data *RelationshipData `json:"data"`
		} `json:"territory"`
	} `json:"relationships"`
	Type string `json:"type"`
}

type SubscriptionGroupsResponse struct {
	Data  []SubscriptionGroup `json:"data"`
	Links DocumentLinks       `json:"links"`
}

// SubscriptionAvailabilityData defines model for SubscriptionAvailabilityCreateRequest.Data
//
// https://developer.apple.com/documentation/appstoreconnectapi/subscriptionavailabilitycreaterequest/data
type SubscriptionAvailabilityData struct {
	Attributes    SubscriptionAvailabilityAttributes `json:"attributes"`
	Relationships struct {
		AvailableTerritories struct {
			Data []*RelationshipData `json:"data"`
		} `json:"availableTerritories"`
		Subscription struct {
			Data *RelationshipData `json:"data"`
		} `json:"subscription"`
	} `json:"relationships"`
	Type string `json:"type"`
}

// SubscriptionAvailabilityAttributes defines model for SubscriptionAvailabilityCreateRequest.Data.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/subscriptionavailabilitycreaterequest/data/attributes
type SubscriptionAvailabilityAttributes struct {
	AvailableInNewTerritories bool `json:"availableInNewTerritories"`
}

// GetSubscriptionGroups gets a subscription group for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_subscription_group
func (s *SubscriptionsService) GetSubscriptionGroups(ctx context.Context, appID, groupName string) (*SubscriptionGroupsResponse, *Response, error) {
	res := new(SubscriptionGroupsResponse)
	resp, err := s.client.get(ctx, "v1/apps/"+appID+"/subscriptionGroups?filter[referenceName]="+url.QueryEscape(groupName), nil, res)
	return res, resp, err
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

// DeleteSubscriptionGroup deletes a subscription group for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_subscription_group
func (s *SubscriptionsService) DeleteSubscriptionGroup(ctx context.Context, groupID string) (*Response, error) {
	return s.client.delete(ctx, "v1/subscriptionGroups/"+groupID, nil)
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

// DeleteSubscriptionGroupLocalization deletes a subscription group localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_subscription_group_localization
func (s *SubscriptionsService) DeleteSubscriptionGroupLocalization(ctx context.Context, localizationID string) (*Response, error) {
	res := new(Response)
	_, err := s.client.delete(ctx, fmt.Sprintf("v1/subscriptionGroupLocalizations/%s", localizationID), nil)
	return res, err
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
			FamilySharable:     false,
			Name:               name,
			ProductID:          productID,
			ReviewNotes:        reviewNotes,
			SubscriptionPeriod: string(period),
			GroupLevel:         1,
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

// DeleteSubscription deletes a subscription for an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_auto-renewable_subscription
func (s *SubscriptionsService) DeleteSubscription(ctx context.Context, subscriptionID string) (*Response, error) {
	return s.client.delete(ctx, "v1/subscriptions/"+subscriptionID, nil)
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

// DeleteSubscriptionLocalization deletes a subscription localization.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_a_subscription_localization
func (s *SubscriptionsService) DeleteSubscriptionLocalization(ctx context.Context, subscriptionLocalizationID string) (*Response, error) {
	url := fmt.Sprintf("v1/subscriptionLocalizations/%s", subscriptionLocalizationID)
	return s.client.delete(ctx, url, nil)
}

// CreateSubscriptionPriceChange schedules a subscription price change for a specific territory.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_subscription_price_change
// Deprecated: use CreateSubscriptionPriceChange instead
func (s *SubscriptionsService) CreateSubscriptionPriceChange(ctx context.Context, subscriptionID, priceID, regionID string, preserveCurrentPrice bool) (*SubscriptionPriceCreateResponse, *Response, error) {
	res := new(SubscriptionPriceCreateResponse)
	resp, err := s.client.post(ctx, "v1/subscriptionPrices", newRequestBody(SubscriptionPriceCreateData{
		Attributes: SubscriptionPriceCreateAttributes{
			PreserveCurrentPrice: preserveCurrentPrice,
			StartDate:            time.Now().Format("2006-01-02"), // ISO 8601
		},
		Relationships: struct {
			Subscription struct {
				Data *RelationshipData `json:"data"`
			} `json:"subscription"`
			SubscriptionPricePoint struct {
				Data *RelationshipData `json:"data"`
			} `json:"subscriptionPricePoint"`
			Territory struct {
				Data *RelationshipData `json:"data"`
			} `json:"territory"`
		}{
			Subscription: struct {
				Data *RelationshipData `json:"data"`
			}{Data: &RelationshipData{ID: subscriptionID, Type: "subscriptions"}},
			SubscriptionPricePoint: struct {
				Data *RelationshipData `json:"data"`
			}{Data: &RelationshipData{ID: priceID, Type: "subscriptionPricePoints"}},
			Territory: struct {
				Data *RelationshipData `json:"data"`
			}{Data: &RelationshipData{ID: regionID, Type: "territories"}},
		},
		Type: "subscriptionPrices",
	}), res)
	return res, resp, err
}

// SetSubscriptionAvailability sets the availability of a subscription for specified territories.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_the_territory_availability_of_a_subscription
func (s *SubscriptionsService) SetSubscriptionAvailability(ctx context.Context, subscriptionID string, regions []*RelationshipData) (*Response, error) {
	resp, err := s.client.post(ctx, "v1/subscriptionAvailabilities", newRequestBody(SubscriptionAvailabilityData{
		Attributes: SubscriptionAvailabilityAttributes{
			AvailableInNewTerritories: true,
		},
		Relationships: struct {
			AvailableTerritories struct {
				Data []*RelationshipData `json:"data"`
			} `json:"availableTerritories"`
			Subscription struct {
				Data *RelationshipData `json:"data"`
			} `json:"subscription"`
		}{
			AvailableTerritories: struct {
				Data []*RelationshipData `json:"data"`
			}{Data: regions},
			Subscription: struct {
				Data *RelationshipData `json:"data"`
			}{Data: &RelationshipData{ID: subscriptionID, Type: "subscriptions"}},
		},
		Type: "subscriptionAvailabilities",
	}), nil)
	return resp, err
}

// ListSubscriptionsByGroup returns a list of subscriptions for a given subscription group ID.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_subscriptions_for_a_subscription_group
func (s *SubscriptionsService) ListSubscriptionsByGroup(ctx context.Context, id string) (*SubscriptionsResponse, *Response, error) {
	res := new(SubscriptionsResponse)
	resp, err := s.client.get(ctx, "v1/subscriptionGroups/"+id+"/subscriptions", nil, res)
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

type SubscriptionPriceCreateRequest struct {
	Relationships SubscriptionPriceUpdateRelationships `json:"relationships"`
	Type          string                               `json:"type"`
	ID            string                               `json:"id"`
}

type SubscriptionPriceScheduleCreateRequestRelationships struct {
	Prices struct {
		Data []*RelationshipData `json:"data"`
	} `json:"prices"`
}

type SubscriptionPriceInlineCreate struct {
	Attributes    SubscriptionPriceInlineCreateAttributes    `json:"attributes"`
	ID            string                                     `json:"id"`
	Relationships SubscriptionPriceInlineCreateRelationships `json:"relationships"`
	Type          string                                     `json:"type"`
}

type SubscriptionPriceInlineCreateAttributes struct {
}

type SubscriptionPriceInlineCreateRelationships struct {
	InAppPurchasePricePoint struct {
		Data *RelationshipData `json:"data"`
	} `json:"subscriptionPricePoint"`
}

func (s *SubscriptionsService) SetSubscriptionPrices(ctx context.Context, subID string, startTime time.Time, regionPrice map[string]string) (*SubscriptionResponse, *Response, error) {
	var priceData []*RelationshipData
	var include []SubscriptionPriceInlineCreate

	i := 0
	for _, priceID := range regionPrice {
		priceJson, err := base64.RawStdEncoding.DecodeString(priceID)
		if err != nil {
			return nil, nil, err
		}

		var priceInfo map[string]string
		if err := json.Unmarshal(priceJson, &priceInfo); err != nil {
			return nil, nil, err
		}

		priceData = append(priceData, &RelationshipData{ID: "${newprice-" + fmt.Sprintf("%d", i) + "}", Type: "subscriptionPrices"})
		include = append(include, SubscriptionPriceInlineCreate{
			Type: "subscriptionPrices",
			ID:   "${newprice-" + fmt.Sprintf("%d", i) + "}",
			Relationships: SubscriptionPriceInlineCreateRelationships{
				InAppPurchasePricePoint: struct {
					Data *RelationshipData `json:"data"`
				}{Data: &RelationshipData{ID: priceID, Type: "subscriptionPricePoints"}},
			},
		})

		i++
	}

	res := new(SubscriptionResponse)
	resp, err := s.client.patch(ctx, "v1/subscriptions/"+subID, newRequestBodyWithIncluded(SubscriptionPriceCreateRequest{
		Type: "subscriptions",
		ID:   subID,
		Relationships: SubscriptionPriceUpdateRelationships{
			Prices: struct {
				Data []*RelationshipData `json:"data"`
			}{Data: priceData},
		},
	}, include), &res)
	return res, resp, err
}

//// SetSubscriptionPrices bulks sets the prices for a subscription.
////
//// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_auto-renewable_subscription#path-parameters
//func (s *SubscriptionsService) SetSubscriptionPrices(ctx context.Context, name, reviewNotes, subscriptionID string, startTime time.Time, regionPrice map[string]string) (*SubscriptionResponse, *Response, error) {
//	var prices []SubscriptionPriceCreateData
//	//var priceRels []*RelationshipData
//
//	for region, price := range regionPrice {
//
//		//priceJson, err := base64.RawStdEncoding.DecodeString(price)
//		//if err != nil {
//		//	return nil, nil, err
//		//}
//		//
//		//var priceInfo map[string]string
//		//if err := json.Unmarshal(priceJson, &priceInfo); err != nil {
//		//	return nil, nil, err
//		//}
//
//		prices = append(prices, SubscriptionPriceCreateData{
//			Attributes: SubscriptionPriceCreateAttributes{
//				PreserveCurrentPrice: true,
//				StartDate:            startTime.Format("2006-01-02"), // ISO 8601
//			},
//			ID: subscriptionID,
//			Relationships: struct {
//				Subscription struct {
//					Data *RelationshipData `json:"data"`
//				} `json:"subscription"`
//				SubscriptionPricePoint struct {
//					Data *RelationshipData `json:"data"`
//				} `json:"subscriptionPricePoint"`
//				Territory struct {
//					Data *RelationshipData `json:"data"`
//				} `json:"territory"`
//			}{
//				Subscription: struct {
//					Data *RelationshipData `json:"data"`
//				}{Data: &RelationshipData{ID: subscriptionID, Type: "subscriptions"}},
//				SubscriptionPricePoint: struct {
//					Data *RelationshipData `json:"data"`
//				}{Data: &RelationshipData{ID: price, Type: "subscriptionPricePoints"}},
//				Territory: struct {
//					Data *RelationshipData `json:"data"`
//				}{Data: &RelationshipData{ID: region, Type: "territories"}},
//			},
//			Type: "subscriptionPrices",
//		})
//
//		//priceJson, err := base64.RawStdEncoding.DecodeString(price)
//		//if err != nil {
//		//	return nil, nil, err
//		//}
//		//
//		//var priceInfo map[string]string
//		//if err := json.Unmarshal(priceJson, &priceInfo); err != nil {
//		//	return nil, nil, err
//		//}
//		//
//		//priceRels = append(priceRels, &RelationshipData{ID: priceInfo["p"], Type: "subscriptionPrices"})
//	}
//
//	fmt.Println(prices)
//
//	res := new(SubscriptionResponse)
//	resp, err := s.client.patch(ctx, "v1/subscriptions/"+subscriptionID, newRequestBodyWithIncluded(SubscriptionUpdateRequestData{
//		Attributes: SubscriptionUpdateAttributes{
//			FamilySharable:     false,
//			Name:               name,
//			ReviewNote:         reviewNotes,
//			SubscriptionPeriod: string(SubscriptionPeriodP1M),
//			GroupLevel:         1,
//		},
//		ID:   subscriptionID,
//		Type: "subscriptions",
//		Relationships: SubscriptionUpdateRelationships{
//			IntroductoryOffers: struct {
//				Data []*RelationshipData `json:"data"`
//			}{Data: []*RelationshipData{}},
//			Prices: struct {
//				Data []*RelationshipData `json:"data"`
//			}{Data: []*RelationshipData{}},
//			PromotionalOffers: struct {
//				Data []*RelationshipData `json:"data"`
//			}{Data: []*RelationshipData{}},
//		},
//	}, prices), res)
//
//	resJson, _ := json.Marshal(res)
//	fmt.Println(string(resJson))
//	fmt.Println(resp.Status)
//
//	return res, resp, err
//}

// GetSubscriptionPricePoints returns a list of approved prices apple will allow you to set for a subscription.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_subscription_price_point_information
func (s *SubscriptionsService) GetSubscriptionPricePoints(ctx context.Context, id, territory string) (*SubscriptionPricePointsResponse, *Response, error) {
	res := new(SubscriptionPricePointsResponse)
	resp, err := s.client.get(ctx, "v1/subscriptions/"+id+"/pricePoints?include=territory&filter[territory]="+url.QueryEscape(territory)+"&limit=8000", nil, res)
	return res, resp, err
}

// GetSubscriptionPrice returns a list of prices for an auto-renewable subscription, by territory.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_prices_for_a_subscription
func (s *SubscriptionsService) GetSubscriptionPrice(ctx context.Context, id, territory string) (*SubscriptionPriceResponse, *Response, error) {
	res := new(SubscriptionPriceResponse)
	resp, err := s.client.get(ctx, "v1/subscriptions/"+id+"/prices?include=subscriptionPricePoint&filter[territory]="+url.QueryEscape(territory)+"&limit=200", nil, res)
	return res, resp, err
}
