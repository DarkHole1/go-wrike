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
type DateRange struct{}

// QueryContactsParams - Params for QueryContacts
type QueryContactsParams struct {
	Me       *bool
	Metadata *Metadata
	Deleted  *bool
	Fields   []string
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

// QueryTasksParams - Params for QueryTasks
type QueryTasksParams struct {
	Descendats, SubTasks                                                      *bool
	Title, Importance, Permalink, Type, SortField, SortOrder, NextPageToken   *string
	Status, Authors, Responsibles, CustomStatuses, Fields                     []string
	StartDate, DueDate, SheduledDate, CreatedDate, UpdatedDate, CompletedDate *DateRange
	Limit, PageSize                                                           *int
	Metadata                                                                  *Metadata
}
