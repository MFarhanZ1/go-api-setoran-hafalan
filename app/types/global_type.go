package types

type (
	GlobalServiceIntroduceType struct {
		Response    bool   `json:"response"`
		Message     string `json:"message"`
		Version     string `json:"version"`
		Contributor string `json:"contributor"`
		Timestamp   string `json:"timestamp"`
	}
)