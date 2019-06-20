package wrike

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func jsonRequest(method, url, token string, data url.Values) (map[string]interface{}, error) {
	if method != "GET" {
		data["method"] = []string{method}
	}
	url = "https://www.wrike.com/api/v4/" + url + "?" + data.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "bearer "+token)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var res map[string]interface{}

	err = json.Unmarshal(body, &res)

	return res, err
}

// Refresh - Refreshes oauth token
func (api *API) Refresh() error {
	if len(api.ID) == 0 || len(api.Secret) == 0 {
		panic("Attempting to call refresh at instance without ID or Secret")
	}

	resp, err := http.PostForm("https://www.wrike.com/oauth2/token", url.Values{
		"client_id":     {api.ID},
		"client_secret": {api.Secret},
		"grant_type":    {"refresh_token"},
		"refresh_token": {api.RefreshToken},
	})

	if err != nil {
		return err
	}

	byteBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return err
	}

	var body map[string]interface{}
	err = json.Unmarshal(byteBody, &body)

	if err != nil {
		return err
	}

	access, refresh, err := checkOAuthError(body)

	if err != nil {
		return err
	}

	api.Token = access
	api.RefreshToken = refresh

	return nil
}

// QueryContacts - List contacts of all users and user groups in current account.
func (api *API) QueryContacts(params *QueryContactsParams) ([]Contact, error) {
	url := queryContactsParams2Values(params)
	resp, err := jsonRequest("GET", "contacts", api.Token, url)
	if err != nil {
		return nil, err
	}

	data, err := checkError(resp)
	if err != nil {
		if val, ok := err.(Error); ok {
			if val.ErrorShort == "not_authorized" {
				if len(api.RefreshToken) != 0 {
					err = api.Refresh()

					if err != nil {
						return nil, err
					}

					return api.QueryContacts(params)
				}

				return nil, err
			}
		} else {
			return nil, err
		}
	}

	res := make([]Contact, len(data))

	for i, c := range data {
		res[i] = parseContact(c.(map[string]interface{}))
	}

	return res, nil
}

// QueryTasks - Search among all tasks in current account.
func (api API) QueryTasks(params *QueryTasksParams) ([]Task, error) {
	url := queryTaskParams2Values(params)
	resp, err := jsonRequest("GET", "tasks", api.Token, url)
	if err != nil {
		return nil, err
	}

	data, err := checkError(resp)
	if err != nil {
		if val, ok := err.(Error); ok {
			if val.ErrorShort == "not_authorized" {
				if len(api.RefreshToken) != 0 {
					err = api.Refresh()

					if err != nil {
						return nil, err
					}

					return api.QueryTasks(params)
				}

				return nil, err
			}
		} else {
			return nil, err
		}
	}

	res := make([]Task, len(data))

	for i, c := range data {
		res[i] = parseTask(c.(map[string]interface{}))
	}

	return res, nil
}

// GetTasks - Returns complete information about multiple tasks
func (api API) GetTasks(taskids []string, params *GetTasksParams) ([]Task, error) {
	url := getTaskParams2Values(params)
	tasks := strings.Join(taskids, ",")
	resp, err := jsonRequest("GET", "tasks/"+tasks, api.Token, url)
	if err != nil {
		return nil, err
	}

	data, err := checkError(resp)
	if err != nil {
		if val, ok := err.(Error); ok {
			if val.ErrorShort == "not_authorized" {
				if len(api.RefreshToken) != 0 {
					err = api.Refresh()

					if err != nil {
						return nil, err
					}

					return api.GetTasks(taskids, params)
				}

				return nil, err
			}
		} else {
			return nil, err
		}
	}

	res := make([]Task, len(data))

	for i, c := range data {
		res[i] = parseTask(c.(map[string]interface{}))
	}

	return res, nil
}

// GetTask - Returns complete information about single task
func (api API) GetTask(taskid string, params *GetTasksParams) (*Task, error) {
	tasks, err := api.GetTasks([]string{taskid}, params)
	if err != nil {
		return nil, err
	}
	return &tasks[0], nil
}

// ModifyTask - Updates a task
func (api API) ModifyTask(taskid string, params *ModifyTaskParams) (*Task, error) {
	url := modifyTaskParams2Values(params)
	resp, err := jsonRequest("PUT", "tasks/"+taskid, api.Token, url)
	if err != nil {
		return nil, err
	}

	data, err := checkError(resp)
	if err != nil {
		if val, ok := err.(Error); ok {
			if val.ErrorShort == "not_authorized" {
				if len(api.RefreshToken) != 0 {
					err = api.Refresh()

					if err != nil {
						return nil, err
					}

					return api.ModifyTask(taskid, params)
				}

				return nil, err
			}
		} else {
			return nil, err
		}
	}

	res := make([]Task, len(data))

	for i, c := range data {
		res[i] = parseTask(c.(map[string]interface{}))
	}

	return &res[0], nil
}

// GetWorkflows - Returns list of workflows with custom statuses.
func (api API) GetWorkflows() ([]Workflow, error) {
	resp, err := jsonRequest("GET", "workflows", api.Token, url.Values{})
	if err != nil {
		return nil, err
	}

	data, err := checkError(resp)
	if err != nil {
		if val, ok := err.(Error); ok {
			if val.ErrorShort == "not_authorized" {
				if len(api.RefreshToken) != 0 {
					err = api.Refresh()

					if err != nil {
						return nil, err
					}

					return api.GetWorkflows()
				}

				return nil, err
			}
		} else {
			return nil, err
		}
	}

	res := make([]Workflow, len(data))

	for i, c := range data {
		res[i] = parseWorkflow(c.(map[string]interface{}))
	}

	return res, nil
}

// QueryFolders - Returns a list of folders
func (api *API) QueryFolders(params *QueryFoldersParams) ([]Folder, error) {
	resp, err := jsonRequest("GET", "folders", api.Token, url.Values{})
	if err != nil {
		return nil, err
	}

	data, err := checkError(resp)
	if err != nil {
		if val, ok := err.(Error); ok {
			if val.ErrorShort == "not_authorized" {
				if len(api.RefreshToken) != 0 {
					err = api.Refresh()

					if err != nil {
						return nil, err
					}

					return api.QueryFolders(params)
				}

				return nil, err
			}
		} else {
			return nil, err
		}
	}

	res := make([]Folder, len(data))

	for i, c := range data {
		res[i] = parseFolder(c.(map[string]interface{}))
	}

	return res, nil
}
