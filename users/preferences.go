// This file was auto-generated by Fern from our API Definition.

package users

import (
	json "encoding/json"
	fmt "fmt"
	v3 "github.com/trycourier/courier-go/v3"
	core "github.com/trycourier/courier-go/v3/core"
)

type UserPreferencesGetResponse struct {
	Topic *TopicPreference `json:"topic,omitempty"`

	_rawJSON json.RawMessage
}

func (u *UserPreferencesGetResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserPreferencesGetResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserPreferencesGetResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserPreferencesGetResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserPreferencesListResponse struct {
	Paging *v3.Paging `json:"paging,omitempty"`
	// The Preferences associated with the user_id.
	Items []*TopicPreference `json:"items,omitempty"`

	_rawJSON json.RawMessage
}

func (u *UserPreferencesListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserPreferencesListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserPreferencesListResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserPreferencesListResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserPreferencesUpdateResponse struct {
	Message string `json:"message"`

	_rawJSON json.RawMessage
}

func (u *UserPreferencesUpdateResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UserPreferencesUpdateResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserPreferencesUpdateResponse(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserPreferencesUpdateResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserPreferencesUpdateParams struct {
	Status v3.PreferenceStatus `json:"status,omitempty"`
	// The Channels a user has chosen to receive notifications through for this topic
	CustomRouting    []v3.ChannelClassification `json:"custom_routing,omitempty"`
	DefaultStatus    v3.PreferenceStatus        `json:"default_status,omitempty"`
	HasCustomRouting *bool                      `json:"has_custom_routing,omitempty"`
}
