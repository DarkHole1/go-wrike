package wrike

import (
	"net/url"
	"strconv"
	"strings"
)

func queryContactsParams2Values(params *QueryContactsParams) url.Values {
	res := url.Values{}

	if params.Me != nil {
		res["me"] = []string{strconv.FormatBool(*params.Me)}
	}

	if params.Deleted != nil {
		res["deleted"] = []string{strconv.FormatBool(*params.Deleted)}
	}

	if params.Metadata != nil {
		if params.Metadata.Value != nil {
			res["metadata"] = []string{`{"key": "` + params.Metadata.Key + `", "value": "` + *params.Metadata.Value + `"}`}
		} else {
			res["metadata"] = []string{`{"key": "` + params.Metadata.Key + `"}`}
		}
	}

	if params.Fields != nil {
		res["fields"] = []string{"[" + strings.Join(params.Fields, ", ") + "]"}
	}

	return res
}

func parseContact(contact map[string]interface{}) Contact {
	var res = Contact{}

	res.ID = contact["id"].(string)
	res.FirstName = contact["firstName"].(string)
	res.LastName = contact["lastName"].(string)
	res.Type = contact["type"].(string)
	res.AvatarURL = contact["avatarUrl"].(string)
	res.Timezone = contact["timezone"].(string)
	res.Locale = contact["locale"].(string)

	if val, ok := contact["me"].(bool); ok {
		res.Me = OptionalBool(val)
	} else {
		res.Me = nil
	}

	if val, ok := contact["memberIds"].([]string); ok {
		res.MemberIDs = val
	} else {
		res.MemberIDs = nil
	}

	if metadata, ok := contact["metadata"].([]interface{}); ok {
		res.Metadata = make([]Metadata, len(metadata))
		for i, d := range metadata {
			data := d.(map[string]interface{})
			res.Metadata[i].Key = data["key"].(string)
			res.Metadata[i].Value = OptionalString(data["value"].(string))
		}
	}

	return res
}
