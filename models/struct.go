package models

import (
	"time"
)

type Tag string

type Time time.Time

type Markdown struct {
	Title       string `json:"title"`
	Date        Time   `json:"date"`
	Description string `json:"description"`
	Tags        []Tag  `json:"tags"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Path        string `json:"path"`
}

type MarkdownDetails struct {
	Markdown
	Body string
}
type MarkdownList []Markdown

type MarkdownPagination struct {
	Markdowns   MarkdownList
	Total       int
	CurrentPage int
	PageNumber  []int
}

type Category struct {
	Title            string
	Number           int
	MarkdownFileList MarkdownList
}

type Categories []Category

func (t *Time) UnmarshalJSON(b []byte) error {
	date, err := time.ParseInLocation("\"2006-01-02 15:04\"", string(b), time.Local)
	if err != nil {
		return nil
	}
	*t = Time(date)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {

	return []byte(t.Format("\"2006-01-02 15:04\"")), nil
}

func (t Time) Format(layout string) string {
	return time.Time(t).Format(layout)
}

func (m MarkdownList) Len() int { return len(m) }

func (m MarkdownList) Less(i, j int) bool { return time.Time(m[i].Date).After(time.Time(m[j].Date)) }

func (m MarkdownList) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
