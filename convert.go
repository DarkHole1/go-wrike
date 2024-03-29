package wrike

import (
	"net/url"
	"strconv"
	"strings"
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
		res["updatedDate"] = []string{dateRange2String(params.UpdatedDate)}
	}

	if params.CompletedDate != nil {
		res["completedDate"] = []string{dateRange2String(params.CreatedDate)}
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

func modifyTaskParams2Values(params *ModifyTaskParams) url.Values {
	res := url.Values{}

	if params.Title != nil {
		res["title"] = []string{*params.Title}
	}

	if params.Description != nil {
		res["descrption"] = []string{*params.Description}
	}

	if params.Status != nil {
		res["status"] = []string{*params.Status}
	}

	if params.Importance != nil {
		res["importance"] = []string{*params.Importance}
	}

	if params.PriorityBefore != nil {
		res["priorityBefore"] = []string{*params.PriorityBefore}
	}

	if params.PriorityAfter != nil {
		res["priorityAfter"] = []string{*params.PriorityAfter}
	}

	if params.CustomStatus != nil {
		res["customStatus"] = []string{*params.CustomStatus}
	}

	if params.AddParents != nil {
		res["addParents"] = []string{stringArray2String(params.AddParents)}
	}

	if params.RemoveParents != nil {
		res["removeParents"] = []string{stringArray2String(params.RemoveParents)}
	}

	if params.AddShareds != nil {
		res["addShareds"] = []string{stringArray2String(params.AddShareds)}
	}

	if params.RemoveShareds != nil {
		res["removeShareds"] = []string{stringArray2String(params.RemoveShareds)}
	}

	if params.AddResponsibles != nil {
		res["addResponsibles"] = []string{stringArray2String(params.AddResponsibles)}
	}

	if params.RemoveResponsibles != nil {
		res["removeResponsibles"] = []string{stringArray2String(params.RemoveResponsibles)}
	}

	if params.AddFollowers != nil {
		res["addFollowers"] = []string{stringArray2String(params.AddFollowers)}
	}

	if params.AddSuperTasks != nil {
		res["addSuperTasks"] = []string{stringArray2String(params.AddFollowers)}
	}

	if params.RemoveSuperTasks != nil {
		res["removeSuperTasks"] = []string{stringArray2String(params.RemoveSuperTasks)}
	}

	if params.Fields != nil {
		res["fields"] = []string{stringArray2String(params.Fields)}
	}

	if params.Follow != nil {
		res["follow"] = []string{strconv.FormatBool(*params.Follow)}
	}

	if params.Restore != nil {
		res["restore"] = []string{strconv.FormatBool(*params.Restore)}
	}

	if params.Metadata != nil {
		subres := make([]string, len(params.Metadata))
		for i, meta := range params.Metadata {
			subres[i] = metadata2String(&meta)
		}
		res["metadata"] = []string{"[" + strings.Join(subres, ",") + "]"}
	}

	return res
}

func queryFoldersParams2Values(params *QueryFoldersParams) url.Values {
	res := url.Values{}

	if params.Permalink != nil {
		res["permalink"] = []string{*params.Permalink}
	}

	if params.Descendants != nil {
		res["descendants"] = []string{strconv.FormatBool(*params.Descendants)}
	}

	if params.Project != nil {
		res["project"] = []string{strconv.FormatBool(*params.Project)}
	}

	if params.Deleted != nil {
		res["deleted"] = []string{strconv.FormatBool(*params.Deleted)}
	}

	if params.UpdatedDate != nil {
		res["updatedDate"] = []string{dateRange2String(params.UpdatedDate)}
	}

	if params.Fields != nil {
		res["fields"] = []string{stringArray2String(params.Fields)}
	}

	if params.Metadata != nil {
		res["metadata"] = []string{metadata2String(params.Metadata)}
	}

	return res
}

func createCommentParams2Values(params *CreateCommentParams) url.Values {
	res := url.Values{}

	res["text"] = []string{params.Text}

	if params.PlainText != nil {
		res["plainText"] = []string{strconv.FormatBool(*params.PlainText)}
	}

	return res
}

func parseMetadata(meta map[string]interface{}) Metadata {
	var res Metadata
	res.Key = meta["key"].(string)
	res.Value = OptionalString(meta["value"].(string))
	return res
}

func parseProfile(profile map[string]interface{}) Profile {
	var res Profile

	res.AccountID = profile["accountId"].(string)
	res.Role = profile["role"].(string)

	if val, ok := profile["email"].(string); ok {
		res.Email = OptionalString(val)
	}

	res.External = profile["external"].(bool)
	res.Admin = profile["admin"].(bool)
	res.Owner = profile["owner"].(bool)

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

	if val, ok := contact["profiles"].([]interface{}); ok {
		res.Profiles = make([]Profile, len(val))
		for i, profile := range val {
			res.Profiles[i] = parseProfile(profile.(map[string]interface{}))
		}
	}

	if val, ok := contact["me"].(bool); ok {
		res.Me = OptionalBool(val)
	}

	if val, ok := contact["memberIds"].([]interface{}); ok {
		res.MemberIDs = make([]string, len(val))
		for i, s := range val {
			res.MemberIDs[i] = s.(string)
		}
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

	if val, ok := task["parentIds"].([]interface{}); ok {
		res.ParentIDs = make([]string, len(val))
		for i, s := range val {
			res.ParentIDs[i] = s.(string)
		}
	}

	if val, ok := task["superParentIds"].([]interface{}); ok {
		res.SuperParentIDs = make([]string, len(val))
		for i, s := range val {
			res.SuperParentIDs[i] = s.(string)
		}
	}

	if val, ok := task["sharedIds"].([]interface{}); ok {
		res.SharedIDs = make([]string, len(val))
		for i, s := range val {
			res.SharedIDs[i] = s.(string)
		}
	}

	if val, ok := task["responsibleIds"].([]interface{}); ok {
		res.ResponsibleIDs = make([]string, len(val))
		for i, s := range val {
			res.ResponsibleIDs[i] = s.(string)
		}
	}

	if val, ok := task["authorIds"].([]interface{}); ok {
		res.AuthorIDs = make([]string, len(val))
		for i, s := range val {
			res.AuthorIDs[i] = s.(string)
		}
	}

	if metadata, ok := task["metadata"].([]map[string]interface{}); ok {
		res.Metadata = make([]Metadata, len(metadata))
		for i, meta := range metadata {
			res.Metadata[i] = parseMetadata(meta)
		}
	}

	return res
}

func parseWorkflow(workflow map[string]interface{}) Workflow {
	var res Workflow

	res.ID = workflow["id"].(string)
	res.Name = workflow["name"].(string)
	res.Standard = workflow["standard"].(bool)
	res.Hidden = workflow["hidden"].(bool)

	cs := workflow["customStatuses"].([]interface{})
	res.CustomStatuses = make([]CustomStatus, len(cs))

	for i, status := range cs {
		res.CustomStatuses[i] = parseCustomStatus(status.(map[string]interface{}))
	}

	return res
}

func parseCustomStatus(status map[string]interface{}) CustomStatus {
	var res CustomStatus

	res.ID = status["id"].(string)
	res.Name = status["name"].(string)
	res.Group = status["group"].(string)

	if val, ok := status["color"].(string); ok {
		res.Color = OptionalString(val)
	}

	res.StandardName = status["standardName"].(bool)
	res.Standard = status["standard"].(bool)

	return res
}

func parseProject(project map[string]interface{}) Project {
	var res Project

	res.AuthorID = project["authorId"].(string)

	if val, ok := project["customStatusId"].(string); ok {
		res.CustomStatusID = OptionalString(val)
	}

	if val, ok := project["startDate"].(string); ok {
		res.StartDate = OptionalString(val)
	}

	if val, ok := project["endDate"].(string); ok {
		res.EndDate = OptionalString(val)
	}

	if val, ok := project["createdDate"].(string); ok {
		res.CreatedDate = OptionalString(val)
	}

	if val, ok := project["completedDate"].(string); ok {
		res.CompletedDate = OptionalString(val)
	}

	if val, ok := project["ownerIds"].([]interface{}); ok {
		res.OwnerIDs = make([]string, len(val))
		for i, s := range val {
			res.OwnerIDs[i] = s.(string)
		}
	}

	return res
}

func parseFolder(folder map[string]interface{}) Folder {
	var res Folder

	res.ID = folder["id"].(string)
	res.Title = folder["title"].(string)
	res.Scope = folder["scope"].(string)

	if val, ok := folder["Color"].(string); ok {
		res.Color = OptionalString(val)
	}

	if val, ok := folder["childIds"].([]interface{}); ok {
		res.ChildIDs = make([]string, len(val))
		for i, s := range val {
			res.ChildIDs[i] = s.(string)
		}
	}

	if val, ok := folder["project"].(map[string]interface{}); ok {
		res.Project = parseProject(val)
	}

	return res
}

func parseComment(comment map[string]interface{}) Comment {
	var res Comment

	res.ID = comment["id"].(string)
	res.AuthorID = comment["authorId"].(string)
	res.Text = comment["text"].(string)
	res.CreatedDate = comment["createdDate"].(string)

	if val, ok := comment["taskId"].(string); ok {
		res.TaskID = OptionalString(val)
	}

	if val, ok := comment["folderId"].(string); ok {
		res.FolderID = OptionalString(val)
	}

	return res
}
