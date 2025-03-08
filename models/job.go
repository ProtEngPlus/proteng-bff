package models

import(
	"time"
)

type CreateJobInput struct {
	Job				JobInput          		`bson:"job" json:"job" validate:"required"`
	QueryResult     CopyQueryResultInput    `bson:"query_result" json:"query_result"`
}
type JobInput struct {
	Name             string                 `bson:"name" json:"name" validate:"required"`
	StageId          int                    `bson:"stage_id" json:"stage_id" validate:"gte=0,lte=3"`
	RefJobId         string                 `bson:"ref_job_id" json:"ref_job_id"`
	LabResult        LabResult              `bson:"lab_result" json:"lab_result"`
	Options          map[string]interface{} `bson:"options" json:"options" validate:"required"`
	Artifacts        map[string]interface{} `bson:"artifact" json:"artifact"`
	Meta             []string               `bson:"meta" json:"meta"`
	InputProtein     string                 `bson:"input_protein" json:"input_protein" validate:"required"`
	RunType          string                 `bson:"run_type" json:"run_type" validate:"required"`
	Description      string                 `bson:"description" json:"description"`
	IsNotificationOn bool                   `bson:"is_notification_on" json:"is_notification_on"`
}

type CopyQueryResultInput struct {
	Id           string     `bson:"_id" json:"id"`
    JobId        string     `bson:"job_id" json:"job_id"`
	RunId        int                    `bson:"run_id" json:"run_id"`
	InputProtein string                 `bson:"input_protein" json:"input_protein"`
	State        string  `bson:"state" json:"state"`
	Result       []CopyResultInput       	`bson:"result" json:"result"`
	CreatedAt    time.Time              `bson:"created_at" json:"created_at"`
	CompleteAt   time.Time              `bson:"complete_at" json:"complete_at"`
}

type CopyResultInput struct {
	Id           string 				`bson:"id" json:"id"`
	IsSelected   bool					`bson:"is_selected" json:"is_selected"`
	Sequences    string 				`bson:"sequences" json:"sequences"`
	Score 		 float64      			`bson:"score" json:"score"`
	MaxScore 	 float64				`bson:"max_score" json:"max_score"`
	HspQueryFrom float64 				`bson:"hsp_query_from" json:"hsp_query_from"`
	HspQueryTo   float64                `bson:"hsp_query_to" json:"hsp_query_to"`
	QueryCover   float64 				`bson:"query_cover" json:"query_cover"`
	EValues      float64 				`bson:"e_values" json:"e_values"`
	Accession    string                 `bson:"accession" json:"accession"`
	PercentIdentity float64				`bson:"percent_identity" json:"percent_identity"`
	AccLen       int					`bson:"acc_len" json:"acc_len"`
	Description  string					`bson:"description" json:"description"`
	Organisms 	 string					`bson:"organisms" json:"organisms"`
}
type LabResult struct {
	Total     int       `bson:"total" json:"total" validate:"gte=0"`
	Sequences []string  `bson:"sequences" json:"sequences" validate:"required"`
	Scores    []float32 `bson:"scores" json:"scores" validate:"required"`
}

type ConfigurationInput struct {
	RefJobId         string                 `bson:"ref_job_id" json:"ref_job_id" validate:"required"`
	State            string                 `bson:"state" json:"state"`
	StageId          int                    `bson:"stage_id" json:"stage_id" validate:"gte=0,lte=3"`
	Name             string                 `bson:"name" json:"name" validate:"required"`
	UserId           string                 `bson:"user_id" json:"user_id" validate:"required"`
	LabResult        LabResult              `bson:"lab_result" json:"lab_result"`
	Options          map[string]interface{} `bson:"options" json:"options" validate:"required"`
	Artifacts        map[string]interface{} `bson:"artifact" json:"artifact"`
	Meta             []string               `bson:"meta" json:"meta"`
	InputProtein     string                 `bson:"input_protein" json:"input_protein" validate:"required"`
	RunType          string                 `bson:"run_type" json:"run_type" validate:"required"`
	Description      string                 `bson:"description" json:"description"`
	IsNotificationOn bool                   `bson:"is_notification_on" json:"is_notification_on"`
}
