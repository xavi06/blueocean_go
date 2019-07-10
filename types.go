package goblueoc

// Pipeline struct
type Pipeline struct {
	Class                     interface{} `json:"_class"`
	Links                     interface{} `json:"_links"`
	Actions                   interface{} `json:"actions"`
	DisplayName               string      `json:"displayName"`
	EstimatedDurationInMillis int         `json:"estimatedDurationInMillis"`
	FullDisplayName           string      `json:"fullDisplayName"`
	FullName                  string      `json:"fullName"`
	LatestRun                 interface{} `json:"latestRun"`
	Name                      string      `json:"name"`
	Oganization               string      `json:"organization"`
	Parameters                []Parameter `json:"parameters"`
	Permissions               interface{} `json:"permissions"`
	WeatherScore              int         `json:"weatherScore"`
}

// Parameter struct
type Parameter struct {
	Class                 interface{} `json:"_class"`
	DefaultParameterValue interface{} `json:"defaultParameterValue"`
	Description           string      `json:"description"`
	Name                  string      `json:"name"`
	Type                  string      `json:"type"`
}

// PipelineRun struct
type PipelineRun struct {
	Class                     interface{} `json:"_class"`
	Links                     interface{} `json:"_links"`
	Actions                   interface{} `json:"actions"`
	ArtifactsZipFile          interface{} `json:"artifactsZipFile"`
	CauseOfBlockage           interface{} `json:"causeOfBlockage"`
	Causes                    []Cause     `json:"causes"`
	ChangeSet                 []ChangeSet `json:"changeSet"`
	Description               string      `json:"description"`
	DurationInMillis          int         `json:"durationInMillis"`
	EnQueueTime               int         `json:"enQueueTime"`
	EndTime                   string      `json:"endTime"`
	EstimatedDurationInMillis int         `json:"estimatedDurationInMillis"`
	ID                        string      `json:"id"`
	Name                      string      `json:"name"`
	Organization              string      `json:"organization"`
	Pipeline                  string      `json:"pipeline"`
	Replayable                bool        `json:"replayable"`
	Result                    string      `json:"result"`
	RunSummary                string      `json:"runSummary"`
	StartTime                 string      `json:"startTime"`
	State                     string      `json:"state"`
	Type                      string      `json:"type"`
	Branch                    string      `json:"branch"`
	CommitID                  string      `json:"commitId"`
	CommitURL                 string      `json:"commitUrl"`
	PullRequest               string      `json:"pullRequest"`
}

// Cause struct
type Cause struct {
	Class            interface{} `json:"_class"`
	ShortDescription string      `json:"shortDescription"`
	UserID           string      `json:"userId"`
	UserName         string      `json:"userName"`
}

// ChangeSet struct
type ChangeSet struct {
	Class         interface{} `json:"_class"`
	Links         interface{} `json:"_links"`
	AffectedPaths []string    `json:"affectedPaths"`
	Author        Author      `json:"author"`
	CommitID      string      `json:"commitId"`
	Issues        interface{} `json:"issues"`
	Msg           string      `json:"msg"`
	Timestamp     string      `json:"timestamp"`
	URL           interface{} `json:"url"`
}

// Author struct
type Author struct {
	Class      interface{} `json:"_class"`
	Links      interface{} `json:"_links"`
	Avatar     interface{} `json:"avator"`
	Email      string      `json:"email"`
	FullName   string      `json:"fullName"`
	ID         string      `json:"id"`
	Permission interface{} `json:"permission"`
}

// Node struct
type Node struct {
	Class              interface{} `json:"_class"`
	Links              interface{} `json:"_links"`
	Actions            interface{} `json:"actions"`
	DisplayDescription string      `json:"displayDescription"`
	DisplayName        string      `json:"displayName"`
	DurationInMillis   int         `json:"durationInMillis"`
	ID                 string      `json:"id"`
	Input              string      `json:"input"`
	Result             string      `json:"result"`
	StartTime          string      `json:"startTime"`
	State              string      `json:"state"`
	Type               string      `json:"type"`
	CauseOfBlockage    string      `json:"causeOfBlockage"`
	Edges              interface{} `json:"edges"`
	FirstParent        interface{} `json:"firstParent"`
	Restartable        bool        `json:"restartable"`
}

// Step struct
type Step struct {
	Class              interface{} `json:"_class"`
	Links              interface{} `json:"_links"`
	Actions            interface{} `json:"actions"`
	DisplayDescription string      `json:"displayDescription"`
	DisplayName        string      `json:"displayName"`
	DurationInMillis   int         `json:"durationInMillis"`
	ID                 string      `json:"id"`
	Input              string      `json:"input"`
	Result             string      `json:"result"`
	StartTime          string      `json:"startTime"`
	State              string      `json:"state"`
	Type               string      `json:"type"`
}

// Favorite struct
type Favorite struct {
	Class interface{} `json:"_class"`
	Links interface{} `json:"_links"`
	Item  Pipeline    `json:"item"`
}
