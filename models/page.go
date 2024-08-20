package models

import "os"

type Page struct {
	Title string
	Body  []byte
}

// Save writes a text file to root level
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return os.WriteFile("./"+filename, p.Body, 0600)
}

// LoadPage tries to read text file with the specified name
// if the file exists, the function returns a page pointer
// with the file contents as page contents
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
