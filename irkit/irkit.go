package irkit

type Device struct {
	ID         int          `json:"id"`
	Name       string       `json:"name"`
	Alias      string       `json:"alias"`
	Status     bool         `json:"status"`
	DescStatus string       `json:"desc_status"`
	OnCommand  func() error `json:"on_command"`
	OffCommand func() error `json:"off_command"`
}
