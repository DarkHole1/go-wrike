package wrike

import (
	"net/url"
	"strconv"
)

func metadata2String(metadata *Metadata) string {
	if metadata.Value != nil {
		return `{"key": "` + metadata.Key + `", "value": "` + *metadata.Value + `"}`
	}
	return `{"key": "` + metadata.Key + `"}`
}

func stringArray2String(arr []string) string {
	res := "["
	for _, str := range arr {
		res += `"` + str + `", `
	}
	if len(res) > 2 {
		res = res[:len(res)-2]
	}
	res += "]"
	return res
}

func dateRange2String(dr *DateRange) string {
	res := "{"

	if dr.Start != nil {
		res += `"start": "` + *dr.Start + `", `
	}

	if dr.Equal != nil {
		res += `"equal": "` + *dr.Equal + `", `
	}

	if dr.End != nil {
		res += `"end": "` + *dr.End + `", `
	}

	if len(res) > 2 {
		res = res[:len(res)-2]
	}
	res += "}"
	return res
}

func queryContactsParams2Values(params *QueryContactsParams) url.Values {
	res := url.Values{}

	if params.Me != nil {
		res["me"] = []string{strconv.FormatBool(*params.Me)}
	}

	if params.Deleted != nil {
		res["deleted"] = []string{strconv.FormatBool(*params.Deleted)}
	}

	if params.Metadata != nil {
		res["metadata"] = []string{metadata2String(params.Metadata)}
	}

	if params.Fields != nil {
		res["fields"] = []string{stringArray2String(params.Fields)}
	}

	return res
}

func queryTaskParams2Values(params *QueryTasksParams) url.Values {
	res := url.Values{}

	if params.Descendants != nil {
		res["descendants"] = []string{strconv.FormatBool(*params.Descendants)}
	}

	if params.SubTasks != nil {
		res["subTasks"] = []string{strconv.FormatBool(*params.SubTasks)}
	}

	if params.Title != nil {
		res["title"] = []string{*params.Title}
	}

	if params.Importance != nil {
		res["importance"] = []string{*params.Importance}
	}

	if params.Permalink != nil {
		res["permalink"] = []string{*params.Permalink}
	}

	if params.Type != nil {
		res["type"] = []string{*params.Type}
	}

	if params.SortField != nil {
		res["sortField"] = []string{*params.SortField}
	}

	if params.SortOrder != nil {
		res["sortOrder"] = []string{*params.SortOrder}
	}

	if params.NextPageToken != nil {
		res["nextPageToken"] = []string{*params.NextPageToken}
	}

	if params.Status != nil {
		res["status"] = []string{stringArray2String(params.Status)}
	}

	if params.Authors != nil {
		res["authors"] = []string{stringArray2String(params.Authors)}
	}

	if params.Responsibles != nil {
		res["responsibles"] = []string{stringArray2String(params.Responsibles)}
	}

	if params.CustomStatuses != nil {
		res["customStatuses"] = []string{stringArray2String(params.CustomStatuses)}
	}

	if params.Fields != nil {
		res["fields"] = []string{stringArray2String(params.Fields)}
	}

	if params.StartDate != nil {
		res["startDate"] = []string{dateRange2String(params.StartDate)}
	}

	if params.DueDate != nil {
		res["dueDate"] = []string{dateRange2String(params.DueDate)}
	}

	if params.SheduledDate != nil {
		res["sheduledDate"] = []string{dateRange2String(params.SheduledDate)}
	}

	if params.CreatedDate != nil {
		res["createdDate"] = []string{dateRange2String(params.CreatedDate)}
	}

	if params.UpdatedDate != nil {
		res["updatedDate"] = []string{dateRange2String(params.CreatedDate)}
	}

	if params.CompletedDate != nil {
		res["completedDate"] = []string{dateRange2String(params.UpdatedDate)}
	}

	if params.Limit != nil {
		res["limit"] = []string{strconv.Itoa(*params.Limit)}
	}

	if params.PageSize != nil {
		res["pageSize"] = []string{strconv.Itoa(*params.PageSize)}
	}

	if params.Metadata != nil {
		res["metadata"] = []string{metadata2String(params.Metadata)}
	}

	return res
}

func getTaskParams2Values(params *GetTasksParams) url.Values {
	res := url.Values{}

	if params.Fields != nil {
		res["fields"] = []string{stringArray2String(params.Fields)}
	}

	return res
}

func parseMetadata(meta map[string]interface{}) Metadata {
	var res Metadata
	res.Key = meta["key"].(string)
	res.Value = OptionalString(meta["value"].(string))
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

	if metadata, ok := contact["metadata"].([]map[string]interface{}); ok {
		res.Metadata = make([]Metadata, len(metadata))
		for i, meta := range metadata {
			res.Metadata[i] = parseMetadata(meta)

		}
	}

	return res
}

func parseTask(task map[string]interface{}) Task {
	var res Task

	res.ID = task["id"].(string)
	res.AccountID = task["accountId"].(string)
	res.Title = task["title"].(string)
	res.Status = task["status"].(string)
	res.Importance = task["importance"].(string)
	res.CreatedDate = task["createdDate"].(string)
	res.UpdatedDate = task["updatedDate"].(string)
	res.Scope = task["scope"].(string)
	res.CustomStatusID = task["customStatusId"].(string)
	res.Permalink = task["permalink"].(string)
	res.Priority = task["priority"].(string)

	if val, ok := task["description"].(string); ok {
		res.Description = OptionalString(val)
	}

	if val, ok := task["briefDescription"].(string); ok {
		res.BriefDescription = OptionalString(val)
	}

	if val, ok := task["completedDate"].(string); ok {
		res.CompletedDate = OptionalString(val)
	}

	if val, ok := task["parentIds"].([]string); ok {
		res.ParentIDs = val
	}

	if val, ok := task["superParentIds"].([]string); ok {
		res.SuperParentIDs = val
	}

	if val, ok := task["sharedIds"].([]string); ok {
		res.SharedIDs = val
	}

	if val, ok := task["responsibleIds"].([]string); ok {
		res.ResponsibleIDs = val
	}

	if val, ok := task["authorIds"].([]string); ok {
		res.AuthorIDs = val
	}

	if metadata, ok := task["metadata"].([]map[string]interface{}); ok {
		res.Metadata = make([]Metadata, len(metadata))
		for i, meta := range metadata {
			res.Metadata[i] = parseMetadata(meta)
		}
	}

	return res
}
