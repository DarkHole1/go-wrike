package wrike

// API - Main class
type API struct {
	Token, RefreshToken, ID, Secret string
}

// OptionalBool - Creates pointer at bool
func OptionalBool(x bool) *bool {
	b := x
	return &b
}

// OptionalString - Creates pointer at string
func OptionalString(x string) *string {
	s := x
	return &s
}

// Metadata - Represents metadata
type Metadata struct {
	Key   string
	Value *string
}

// DateRange - Represents range for dates
type DateRange struct {
	Start, Equal, End *string
}

// QueryContactsParams - Params for QueryContacts
type QueryContactsParams struct {
	Me       *bool
	Metadata *Metadata
	Deleted  *bool
	Fields   []string
}

// QueryTasksParams - Params for QueryTasks
type QueryTasksParams struct {
	Descendants, SubTasks                                                     *bool
	Title, Importance, Permalink, Type, SortField, SortOrder, NextPageToken   *string
	Status, Authors, Responsibles, CustomStatuses, Fields                     []string
	StartDate, DueDate, SheduledDate, CreatedDate, UpdatedDate, CompletedDate *DateRange
	Limit, PageSize                                                           *int
	Metadata                                                                  *Metadata
}

// GetTasksParams - Params for GetTasks and GetTask
type GetTasksParams struct {
	Fields []string
}

// ModifyTaskParams - Params for ModifyTask
type ModifyTaskParams struct {
	Title, Description, Status, Importance, PriorityBefore, PriorityAfter, CustomStatus                                                              *string
	AddParents, RemoveParents, AddShareds, RemoveShareds, AddResponsibles, RemoveResponsibles, AddFollowers, AddSuperTasks, RemoveSuperTasks, Fields []string
	Follow, Restore                                                                                                                                  *bool
	Metadata                                                                                                                                         []Metadata
}

// QueryFoldersParams - Params for QueryFolders
type QueryFoldersParams struct {
	Permalink                     *string
	Descendants, Project, Deleted *bool
	UpdatedDate                   *DateRange
	Fields                        []string
	Metadata                      *Metadata
}

// Contact - Represents single contact
type Contact struct {
	ID, FirstName, LastName, Type, AvatarURL, Timezone, Locale string
	Profiles                                                   []Profile
	Me                                                         *bool
	MemberIDs                                                  []string
	Metadata                                                   []Metadata
}

// Profile - Represents contact's profile
type Profile struct {
	AccountID, Role        string
	Email                  *string
	External, Admin, Owner bool
}

// Task - Represents task at wrike
type Task struct {
	ID, AccountID, Title, Status, Importance, CreatedDate, UpdatedDate, Scope, CustomStatusID, Permalink, Priority string
	Description, BriefDescription, CompletedDate                                                                   *string
	ParentIDs, SuperParentIDs, SharedIDs, ResponsibleIDs, AuthorIDs                                                []string
	Metadata                                                                                                       []Metadata
}

// Workflow - Represents single workflow
type Workflow struct {
	ID, Name         string
	Standard, Hidden bool
	CustomStatuses   []CustomStatus
}

// CustomStatus - Represents custom status
type CustomStatus struct {
	ID, Name, Group        string
	Color                  *string
	StandardName, Standard bool
}

// Project - Represents project details
type Project struct {
	AuthorID                                                       string
	CustomStatusID, StartDate, EndDate, CreatedDate, CompletedDate *string
	OwnerIDs                                                       []string
}

// Folder - Represents single folder
type Folder struct {
	ID, Title, Scope string
	Color            *string
	ChildIDs         []string
	Project          Project
}
