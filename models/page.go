package models

import "os"

type Page struct {
	Title string
	Body  []byte
}

// Save takes a page and saves it as a .txt file on root
// level of the project
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
