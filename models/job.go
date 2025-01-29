package models

type JobInput struct {
	Name         string                 `bson:"name" json:"name" validate:"required"`
	LabResult    LabResult              `bson:"lab_result" json:"lab_result"`
	Options      map[string]interface{} `bson:"options" json:"options" validate:"required"`
	Meta         []string               `bson:"meta" json:"meta"`
	InputProtein string                 `bson:"input_protein" json:"input_protein" validate:"required"`
	RunType      string                 `bson:"run_type" json:"run_type"`
	Description  string                 `bson:"description" json:"description"`
}

type LabResult struct {
	Total     int       `bson:"total" json:"total" validate:"gte=0"`
	Sequences []string  `bson:"sequences" json:"sequences" validate:"required"`
	Scores    []float32 `bson:"scores" json:"scores" validate:"required"`
}

type ConfigurationInput struct {
	JobId            string                 `bson:"job_id" json:"job_id" validate:"required"`
	State            string                 `bson:"state" json:"state"`
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
