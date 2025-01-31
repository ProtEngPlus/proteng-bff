package models

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
