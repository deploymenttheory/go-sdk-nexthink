package knowledge_bases

import (
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-nexthink/nexthink/internal/validation"
)

func ValidateID(id string) error { return validation.PathSegment(id) }
func ValidateCreate(r *CreateRequest) error {
	if r == nil {
		return fmt.Errorf("knowledge base request is required")
	}
	if err := ValidateID(r.ContentID); err != nil {
		return err
	}
	if strings.TrimSpace(r.FileName) == "" || strings.TrimSpace(r.ITSMHost) == "" {
		return fmt.Errorf("fileName and itsmHost are required")
	}
	if r.TargetGroup != "technical-support" && r.TargetGroup != "employee" {
		return fmt.Errorf("targetGroup must be technical-support or employee")
	}
	return nil
}
func ValidateUpload(r *UploadRequest) error {
	if r == nil || strings.TrimSpace(r.OriginalFileName) == "" || len(r.Data) == 0 {
		return fmt.Errorf("filename and CSV bytes are required")
	}
	return nil
}
func ValidateMultipart(r MultipartContext) error {
	if err := ValidateID(r.ContentID); err != nil {
		return err
	}
	if strings.TrimSpace(r.FileName) == "" || strings.TrimSpace(r.UploadID) == "" {
		return fmt.Errorf("fileName and uploadId are required")
	}
	return nil
}
func ValidatePart(r *UploadPartRequest) error {
	if r == nil {
		return fmt.Errorf("part request is required")
	}
	if err := ValidateMultipart(r.MultipartContext); err != nil {
		return err
	}
	if r.PartNumber < 1 || strings.TrimSpace(r.EncodedChunk) == "" {
		return fmt.Errorf("positive part number and encoded chunk are required")
	}
	return nil
}
func ValidateComplete(r *CompleteMultipartRequest) error {
	if r == nil {
		return fmt.Errorf("completion request is required")
	}
	if err := ValidateMultipart(r.MultipartContext); err != nil {
		return err
	}
	if len(r.Parts) == 0 {
		return fmt.Errorf("at least one uploaded part is required")
	}
	previous := 0
	for _, part := range r.Parts {
		if part.PartNumber <= previous || part.ETag == "" || part.Checksum == "" {
			return fmt.Errorf("parts require increasing positive numbers, etags and checksums")
		}
		previous = part.PartNumber
	}
	return nil
}
