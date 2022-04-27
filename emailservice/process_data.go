package emailservice

import (
	"encoding/json"
	"errors"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func processData(data []byte, dryRun bool) error {
	dataObj := PostData{}
	err := json.Unmarshal(data, &dataObj)
	if err != nil {
		return errors.New("can't send mail:" + err.Error())
	}

	recipients := dataObj.Recipients
	subject := dataObj.Subject
	markdownBody := dataObj.MarkdownBody
	markdownBodyBytes := []byte(markdownBody)
	htmlBody := dataObj.HtmlBody
	htmlBodyBytes := []byte(htmlBody)

	if len(htmlBody) < 1 {
		extensions := parser.CommonExtensions | parser.AutoHeadingIDs
		parser := parser.NewWithExtensions(extensions)
		htmlBodyBytes = markdown.ToHTML(markdownBodyBytes, parser, nil)
	}

	if dryRun {
		return nil
	}
	plaintext := markdownBody
	if len(plaintext) == 0 {
		plaintext = htmlBody
	}
	err = sendMail(recipients, subject, string(htmlBodyBytes), plaintext)
	if err != nil {

		return errors.New("can't send mail: " + err.Error())
	}
	return nil
}
