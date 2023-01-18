package client

import (
	"fmt"
)

type FormatType string

const (
	FormatTypeCSV  = "csv"
	FormatTypeJSON = "json"
)

type Spec struct {
	Bucket         string     `json:"bucket,omitempty"`
	Path           string     `json:"path,omitempty"`
	Format         FormatType `json:"format,omitempty"`
	IncludeHeaders bool       `json:"include_headers,omitempty"`
	Delimiter      rune       `json:"delimiter,omitempty"`
	NoRotate       bool       `json:"no_rotate,omitempty"`
}

func (s *Spec) SetDefaults() {
	if s.Delimiter == 0 {
		s.Delimiter = ','
	}
}

func (s *Spec) Validate() error {
	if s.Bucket == "" {
		return fmt.Errorf("bucket is required")
	}
	if s.Path == "" {
		return fmt.Errorf("path is required")
	}
	if s.Format == "" {
		return fmt.Errorf("format is required")
	}

	return nil
}
