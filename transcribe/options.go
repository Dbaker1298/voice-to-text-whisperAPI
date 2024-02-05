package transcribe

type TranscribeConfig struct {
	Model    string `json:"model"`
	Language string `json:"language"`
	File     string `json:"file"`
}

type TranscribeOption func(*TranscribeConfig)

func WithModel(model string) TranscribeOption {
	return func(tc *TranscribeConfig) {
		tc.Model = model
	}
}

func WithLanguage(lang string) TranscribeOption {
	return func(tc *TranscribeConfig) {
		tc.Language = lang
	}
}

func WithFile(file string) TranscribeOption {
	return func(tc *TranscribeConfig) {
		tc.File = file
	}
}
