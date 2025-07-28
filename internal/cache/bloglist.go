package cache

import (
	"bufio"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"circuitcamel.com/internal/models"
	"circuitcamel.com/internal/utils"
)

func getBlogPosts() ([]models.BlogPost, error) {
	files, err := filepath.Glob("./content/blog/*.md")
	if err != nil {
		return nil, err
	}
	result := make([]models.BlogPost, len(files))
	for i, v := range files {
		result[i], _ = loadBlogMarkdown(v)
		slug := strings.Split(filepath.Base(v), ".")[0]
		result[i] = models.BlogPost{Body: result[i].Body, Slug: slug,
			Date: result[i].Date, Title: result[i].Title}
	}

	sort.Slice(result, func(i, j int) bool {
		ti, _ := time.Parse("02-01-2006", result[i].Date)
		tj, _ := time.Parse("02-01-2006", result[j].Date)
		return tj.Before(ti)
	})

	return result, nil
}

func loadBlogMarkdown(path string) (models.BlogPost, error) {
	file, err := os.Open(path)
	if err != nil {
		return models.BlogPost{}, err
	}
	defer file.Close()

	var b models.BlogPost
	var mdLines []string
	scanner := bufio.NewScanner(file)
	inMeta := false
	metaStarted := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "---" {
			if !metaStarted {
				inMeta = true
				metaStarted = true
				continue
			} else if inMeta {
				inMeta = false
				continue
			}
		}

		if inMeta {
			if parts := strings.SplitN(line, ":", 2); len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				val := strings.TrimSpace(parts[1])
				switch key {
				case "date":
					t, _ := time.Parse("02-01-2006", val)
					b.Date = t.Format("02-01-2006")
				case "title":
					b.Title = val
				}
			}
		} else if metaStarted {
			mdLines = append(mdLines, line)
		}
	}
	md := strings.Join(mdLines, "\n")
	b.Body = utils.MdToHTML([]byte(md))
	return b, nil
}
