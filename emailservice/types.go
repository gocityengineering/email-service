package emailservice

type PostData struct {
	Recipients   []*string `json:"recipients"` // array of pointers to strings to match AWS SDK
	Subject      string    `json:"subject"`
	MarkdownBody string    `json:"markdownBody"`
	HtmlBody     string    `json:"htmlBody"`
}
