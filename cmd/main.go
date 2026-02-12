package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"gopkg.in/yaml.v2"
)

type Page struct {
	Title    string                 `yaml:"title"`
	Template string                 `yaml:"template"`
	Path     string                 `yaml:"path"`
	Meta     map[string]any 		`yaml:"meta"`
	Content  template.HTML
}

type HomePage struct {
	Page
	Articles []Article
}

type Article struct {
	Path  string
	Title string
	Meta  map[string]any
}

const (
	staticDir    = "web/static"
	templatesDir = "web/templates"
	contentDir   = "web/content"

	outputDir    = "_dist"
)

func main() {
	log.Println("Building static site...")
	if err := buildSite(); err != nil {
		log.Fatalf("Failed to build site: %v", err)
	}
	log.Println("Static site generation complete.")
}

func buildSite() error {
	if err := os.RemoveAll(outputDir); err != nil {
		return fmt.Errorf("failed to clean output directory: %w", err)
	}
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	if err := copyDir(staticDir, outputDir); err != nil {
		return fmt.Errorf("failed to copy static directory: %w", err)
	}
	log.Printf("Static assets copied: %s -> %s\n", staticDir, outputDir)

	contentFiles, err := os.ReadDir(contentDir)
	if err != nil {
		return fmt.Errorf("failed to read content directory: %w", err)
	}

	var articles []Article
	var pages []Page

	for _, fileInfo := range contentFiles {
		if !fileInfo.IsDir() && filepath.Ext(fileInfo.Name()) == ".md" {
			filePath := filepath.Join(contentDir, fileInfo.Name())
			page, err := parseMarkdownFile(filePath)
			if err != nil {
				log.Printf("Error parsing Markdown file %s: %v", filePath, err)
				continue
			}

			if page.Meta == nil {
				page.Meta = make(map[string]any)
			}

			page.Meta["Year"] = time.Now().Year()

			if page.Template == "article" {
				articles = append(articles, Article{
					Path:  page.Path,
					Title: page.Title,
					Meta:  page.Meta,
				})
			}
			pages = append(pages, page)
		}
	}

	// Sort projects or handle them as needed
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Title < articles[j].Title
	})

	for _, page := range pages {
		outputPath := filepath.Join(page.Path, "index.html")
		templateName := page.Template + ".html"

		var data any
		if page.Template == "home" {
			data = HomePage{
				Page:     page,
				Articles: articles,
			}
		} else {
			data = page
		}

		if err := renderPage(templateName, outputPath, data); err != nil {
			return fmt.Errorf("failed to render page %s: %w", outputPath, err)
		}
	}

	return nil
}

func renderPage(templateName, outputPath string, data any) error {
	fullOutputPath := filepath.Join(outputDir, outputPath)
	if err := os.MkdirAll(filepath.Dir(fullOutputPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory for %s: %w", fullOutputPath, err)
	}

	tmpl, err := template.ParseFiles(filepath.Join(templatesDir, "main.html"), filepath.Join(templatesDir, templateName))
	if err != nil {
		return fmt.Errorf("failed to parse templates for %s: %w", templateName, err)
	}

	file, err := os.Create(fullOutputPath)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", fullOutputPath, err)
	}
	defer file.Close()

	if err := tmpl.ExecuteTemplate(file, "main.html", data); err != nil {
		return fmt.Errorf("failed to execute template for %s: %w", outputPath, err)
	}

	log.Printf("Generated: %s", fullOutputPath)
	return nil
}

func parseMarkdownFile(filePath string) (Page, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Page{}, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return Page{}, fmt.Errorf("invalid front matter format in %s", filePath)
	}

	var page Page
	if err := yaml.Unmarshal([]byte(parts[1]), &page); err != nil {
		return Page{}, fmt.Errorf("failed to unmarshal front matter in %s: %w", filePath, err)
	}

	htmlContent := markdown.ToHTML([]byte(parts[2]), nil, nil)
	page.Content = template.HTML(htmlContent)

	return page, nil
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)
		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}
		return copyFile(path, dstPath)
	})
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = dstFile.ReadFrom(srcFile)
	return err
}
