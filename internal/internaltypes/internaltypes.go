package internaltypes

import (
	"time"

	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

type WorkflowQuery struct {
	Name string
	Repo string
}

type RunLog struct {
	RunID       string
	Log         string
	Stage       string
	Workflow    string
	CreatedDate time.Time
	UpdatedDate time.Time
}

type Result struct {
	_id         string
	RunID       string
	Workflow    string
	Stage       Stage
	Result      bool
	Skipped     bool
	Outputs     []Env
	Logs        []string
	CreatedDate time.Time
	UpdatedDate time.Time
}

type RunResult struct {
	SkippedStages    uint32
	FailedStages     uint32
	SuccessfulStages uint32
	Workflow         string
	RunID            string
	Outputs          []Env
	Logs             []string
	Result           bool
	RunTime          time.Duration
	StartTime        time.Time
	EndTime          time.Time
}

type Input struct {
	Name     string `mapstructure:"name" validate:"nonzero,nowhitespace"`
	Default  string `mapstructure:"default"`
	Optional bool   `mapstructure:"optional,omitempty"`
}

type Stage struct {
	Stage     string   `mapstructure:"stage" validate:"nonzero"`
	ID        string   `mapstructure:"id,omitempty" validate:"nonzero,nowhitespace"`
	Script    []string `mapstructure:"script" validate:"nonzero"`
	If        string   `mapstructure:"if,omitempty"`
	Clean     bool     `mapstructure:"clean,omitempty"`
	Env       []Env    `mapstructure:"env,omitempty"`
	Artifacts []string `mapstructure:"artifacts,omitempty"`
	Image     string   `mapstructure:"image,omitempty"`
	Needs     string   `mapstructure:"needs,omitempty" validate:"nowhitespace"`
	Import    []Import `mapstructure:"import,omitempty"`
}

type Import struct {
	From      string   `mapstructure:"from" validate:"nonzero,nowhitespace"`
	Artifacts []string `mapstructure:"artifacts" validate:"nonzero"`
}

type Env struct {
	Name  string `mapstructure:"name" validate:"nonzero,nowhitespace"`
	Value string `mapstructure:"value" validate:"nonzero"`
}

type Workflow struct {
	_id         string
	ID          string  `mapstructure:"id" validate:"nonzero,nowhitespace"`
	Image       string  `mapstructure:"image" validate:"nonzero,nowhitespace"`
	Description string  `mapstructure:"description"`
	Env         []Env   `mapstructure:"env"`
	Input       []Input `mapstructure:"input"`
	// Mount       bool    `mapstructure:"mount"`
	Stages []Stage `mapstructure:"stages" validate:"nonzero"`
	Repo   string  `mapstructure:"repository,omitempty"` // To be filled automatically. Not part of YAML.
}

type WorkflowArgument struct {
	Repo     string            `json:"repo" xml:"repo" form:"repo" query:"repo" mapstructure:"repo" validate:"nonzero"`
	Workflow string            `json:"workflow" xml:"workflow" form:"workflow" query:"workflow" mapstructure:"workflow" validate:"nonzero"`
	Args     map[string]string `json:"args" xml:"args" form:"args" query:"args" mapstructure:"args" validate:"nonzero"`
}

type SlackMesseger struct {
	Callback *slack.InteractionCallback
	Slacker  *slacker.Slacker
}
