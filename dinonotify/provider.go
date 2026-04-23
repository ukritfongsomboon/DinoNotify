package dinonotify

// Model ===========================
type Severity string

const (
	SeverityCritical Severity = "CRITICAL"
	SeverityError    Severity = "ERROR"
	SeverityWarning  Severity = "WARNING"
	SeverityInfo     Severity = "INFO"
	SeveritySuccess  Severity = "SUCCESS"
)

func (s Severity) emoji() string {
	switch s {
	case SeverityCritical:
		return "🔴"
	case SeverityError:
		return "🚨"
	case SeverityWarning:
		return "⚠️"
	case SeverityInfo:
		return "ℹ️"
	case SeveritySuccess:
		return "✅"
	default:
		return ""
	}
}

type MessagePayload struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Message  string `json:"message"`
}

type FileType string

const (
	FileTypeImage FileType = "image"
	FileTypeVideo FileType = "video"
	FileTypeFile  FileType = "file"
)

type FilePayload struct {
	Name       string
	URL        string
	PreviewURL string   // optional, for images only
	Type       FileType // optional, default auto-detect from extension
}

// Interface =======================
type ProviderMessage interface {
	Info(message MessagePayload) error
	Error(message MessagePayload) error
	Success(message MessagePayload) error
	Warning(message MessagePayload) error
	SendFile(file FilePayload) error
}
