package models

type JobInput struct {
	Name         string                 `bson:"name" json:"name" validate:"required"`
	LabResult    LabResult              `bson:"lab_result" json:"lab_result"`
	Options      map[string]interface{} `bson:"options" json:"options" validate:"required"`
	Meta         []string               `bson:"meta" json:"meta"`
	InputProtein string                 `bson:"input_protein" json:"input_protein" validate:"required"`
}

type DuplicateJobInput struct {
	Name         string                 `bson:"name" json:"name" validate:"required"`
	LabResult    LabResult              `bson:"lab_result" json:"lab_result"`
	Options      map[string]interface{} `bson:"options" json:"options" validate:"required"`
	Meta         []string               `bson:"meta" json:"meta"`
	InputProtein string                 `bson:"input_protein" json:"input_protein" validate:"required"`
	RefJobId     string                 `bson:"ref_job_id" json:"ref_job_id"`
}

type LabResult struct {
	Total     int       `bson:"total" json:"total" validate:"gte=0"`
	Sequences []string  `bson:"sequences" json:"sequences" validate:"required"`
	Scores    []float32 `bson:"scores" json:"scores" validate:"required"`
}
