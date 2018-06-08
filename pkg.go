package service

// Pkg
type Pkg struct {
	Version   int                    `json:"pkg_version"`
	Uuid      string                 `json:"pkg_uuid"`
	PDateTime string                 `json:"pgk_date_time"`
	Client    map[string]interface{} `json:"client"`
}
