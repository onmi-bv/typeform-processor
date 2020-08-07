package typeform

// WebhookDataResponse ...
type WebhookDataResponse struct {
	ID        string `json:"id,required"`
	FormID    string `json:"form_id,required"`
	Tag       string `json:"tag,required"`
	URL       string `json:"url,required"`
	Enabled   bool   `json:"enabled,required"`
	CreatedAt string `json:"created_at,required"`
	UpdatedAt string `json:"updated_at,required"`
}

// WebhookCreateRequest ...
type WebhookCreateRequest struct {
	URL     string `json:"url,required"`
	Enabled bool   `json:"enabled,required"`
}

// WebhookCreateResponse ...
type WebhookCreateResponse struct {
	ID        string `json:"id,required"`
	FormID    string `json:"form_id,required"`
	Tag       string `json:"tag,required"`
	URL       string `json:"url,required"`
	Enabled   bool   `json:"enabled,required"`
	CreatedAt string `json:"created_at,required"`
	UpdatedAt string `json:"updated_at,required"`
}

// WebhookTypeform ...
type WebhookTypeform struct {
	EventID      string       `json:"event_id,required"`
	EventType    string       `json:"event_type,required"`
	FormResponse FormResponse `json:"form_response,required"`
}

// FormResponse ...
type FormResponse struct {
	FormID      string     `json:"form_id,required"`
	Token       string     `json:"token,required"`
	SubmittedAt string     `json:"submitted_at,required"`
	LandedAt    string     `json:"landed_at,required"`
	Calculated  Calculated `json:"calculated,required"`
	Definition  Definition `json:"definition,required"`
	Answers     []Answer   `json:"answers,required"`
	Hidden      Hidden     `json:"hidden,required"`
}

// Answer ...
type Answer struct {
	Type    string      `json:"type,required"`
	Text    string      `json:"text,omitempty"`
	Field   AnswerField `json:"field"`
	Email   string      `json:"email,omitempty"`
	Date    string      `json:"date,omitempty"`
	Choices Choices     `json:"choices"`
	Number  int64       `json:"number,omitempty"`
	Boolean bool        `json:"boolean,omitempty"`
	Choice  Choice      `json:"choice,omitempty"`
}

// Choice ...
type Choice struct {
	Label string `json:"label"`
}

// Choices ...
type Choices struct {
	Labels []string `json:"labels"`
}

// AnswerField ...
type AnswerField struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Calculated ...
type Calculated struct {
	Score int64 `json:"score"`
}

// Definition ...
type Definition struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Fields []FieldElement `json:"fields"`
}

// FieldElement ...
type FieldElement struct {
	ID                      string `json:"id"`
	Title                   string `json:"title"`
	Type                    string `json:"type"`
	Ref                     string `json:"ref"`
	AllowMultipleSelections bool   `json:"allow_multiple_selections"`
	AllowOtherChoice        bool   `json:"allow_other_choice"`
}

// Hidden ...
type Hidden struct {
	User string `json:"user"`
	T    string `json:"t"`
}
