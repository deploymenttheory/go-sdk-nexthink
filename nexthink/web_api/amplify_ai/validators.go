package amplify_ai

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

func required(values ...string) error {
	for _, value := range values {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("required request field is empty")
		}
	}
	return nil
}
func validateGenerateOrFetchAnalysis(r *AnalysisRequest) error {
	if err := required(r.DeviceID); err != nil {
		return err
	}
	if strings.TrimSpace(r.TicketID) == "" && (r.OnTheFly == nil || strings.TrimSpace(r.OnTheFly.Title) == "" && strings.TrimSpace(r.OnTheFly.TicketNumber) == "") {
		return fmt.Errorf("ticketId or onTheFly ticket details are required")
	}
	return nil
}
func validateSubmitFeedback(r *FeedbackRequest) error {
	if err := required(r.ResolutionPlanID); err != nil {
		return err
	}
	if r.Feedback.Rating < 1 || r.Feedback.Rating > 5 {
		return fmt.Errorf("feedback rating must be between 1 and 5")
	}
	return nil
}
func validatePostMetric(r *MetricRequest) error { return required(r.Metric.Name) }
func validateReportTicketRetrievalDuration(r *TicketRetrievalDurationRequest) error {
	if r.DurationSeconds < 0 || math.IsNaN(r.DurationSeconds) || math.IsInf(r.DurationSeconds, 0) {
		return fmt.Errorf("durationSeconds must be a finite non-negative number")
	}
	return required(r.Status)
}
func validateExecuteAction(r *ExecuteActionRequest) error {
	for _, key := range []string{"resolutionPlanId", "resolutionStepId"} {
		var value string
		if err := json.Unmarshal((*r)[key], &value); err != nil {
			return fmt.Errorf("%s must be a string", key)
		}
		if err := required(value); err != nil {
			return err
		}
	}
	for key, value := range *r {
		if !json.Valid(value) {
			return fmt.Errorf("%s must contain valid JSON", key)
		}
	}
	return nil
}
func validateExecuteUserAction(r *ExecuteUserActionRequest) error {
	return required(r.ResolutionPlanID, r.ResolutionStepID)
}
func validateRefreshResolutionStep(r *RefreshResolutionStepRequest) error {
	return required(r.DeviceID, r.ID)
}
func validateUpdateResolutionStep(r *UpdateResolutionStepRequest) error {
	return required(r.Status, r.ID)
}
func validateResolveTicket(r *ResolveTicketRequest) error { return required(r.SysID) }
