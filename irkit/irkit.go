package irkit

import "go.uber.org/zap"

type Device struct {
	ID         int                     `json:"id"`
	Name       string                  `json:"name"`
	Alias      string                  `json:"alias"`
	Status     bool                    `json:"status"`
	DescStatus string                  `json:"desc_status"`
	OnCommand  func(*zap.Logger) error `json:"on_command"`
	OffCommand func(*zap.Logger) error `json:"off_command"`
}
