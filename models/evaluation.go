package models

type RerunEvaluationInput struct {
	Plugins    []string `json:"plugins" validate:"required"`
	MutationId string   `json:"mutation_id"`
}
