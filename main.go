package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// IconData represents the structure of an icon in the SimpleIcons JSON file
type IconData struct {
	Title string `json:"title"`
	Hex   string `json:"hex"`
}

// IconsData represents the structure of the JSON object from SimpleIcons
type IconsData struct {
	Icons []IconData `json:"icons"`
}

// FetchIcons fetches the icon data from the SimpleIcons GitHub repository
func FetchIcons(url string) ([]IconData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch icon data")
	}

	var iconsData IconsData
	err = json.NewDecoder(resp.Body).Decode(&iconsData)
	if err != nil {
		return nil, err
	}

	return iconsData.Icons, nil
}

// FindIconByName searches for an icon by name and returns the icon data if found
func FindIconByName(icons []IconData, name string) *IconData {
	name = strings.ToLower(name)
	for _, icon := range icons {
		if strings.ToLower(icon.Title) == name {
			return &icon
		}
	}
	return nil
}

// Function to strip the leading '#' from color inputs if present
func stripHash(color string) string {
	if strings.HasPrefix(color, "#") {
		return color[1:]
	}
	return color
}

// Function to compute slug from service name
func computeSlug(serviceName string) string {
	slug := strings.ToLower(serviceName)
	slug = strings.ReplaceAll(slug, " ", "")
	slug = strings.ReplaceAll(slug, ".", "dot")
	return slug
}

// Function to download Badge from Shields.io
func downloadBadge(labelText string, iconName string, iconBackgroundColor string, fontColor string, destDir string) (string, error) {
	labelTextSlug := computeSlug(labelText)
	iconNameLower := computeSlug(iconName)
	iconBackgroundColor = stripHash(iconBackgroundColor)
	fontColor = stripHash(fontColor)
	url := fmt.Sprintf("https://img.shields.io/badge/-%s-%s?style=flat-square&logo=%s&logoColor=%s", labelText, iconBackgroundColor, iconNameLower, fontColor)
	
	// Debug output of the constructed URL
	fmt.Printf("Constructed URL: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to download the badge")
	}

	// Create the destination directory if it does not exist
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.MkdirAll(destDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	fileName := fmt.Sprintf("%s-badge.svg", labelTextSlug)
	filePath := filepath.Join(destDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Printf("Badge downloaded successfully as %s\n", filePath)
	return fileName, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./simpleicons-macos <\"service 1\", \"service 2\", ...>")
		return
	}

	servicesInput := os.Args[1:]
	services := strings.Join(servicesInput, " ")

	iconsURL := "https://raw.githubusercontent.com/simple-icons/simple-icons/develop/_data/simple-icons.json"
	icons, err := FetchIcons(iconsURL)
	if err != nil {
		fmt.Printf("Error fetching icons: %v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the destination directory: ")
	destDir, _ := reader.ReadString('\n')
	destDir = strings.TrimSpace(destDir)

	var htmlContent []string

	serviceList := strings.Split(services, ",")
	for _, serviceName := range serviceList {
		serviceName = strings.TrimSpace(serviceName)
		icon := FindIconByName(icons, serviceName)
		if icon != nil {
			slug := computeSlug(icon.Title)
			fmt.Printf("Found icon for %s (computed slug: %s) with color #%s\n", icon.Title, slug, icon.Hex)
			fmt.Print("Enter the text on the badge: ")
			labelText, _ := reader.ReadString('\n')
			labelText = strings.TrimSpace(labelText)

			fontColor := "white" // Assuming white logo color for simplicity
			fileName, err := downloadBadge(labelText, slug, icon.Hex, fontColor, destDir)
			if err != nil {
				fmt.Printf("Error downloading badge for %s: %v\n", serviceName, err)
			} else {
				htmlTag := fmt.Sprintf(`<img alt="%s" src="%s" />`, slug, fileName)
				htmlContent = append(htmlContent, htmlTag)
			}
		} else {
			fmt.Printf("Service '%s' not found. Skipping...\n", serviceName)
		}
	}

	// Ask if HTML code should be created
	fmt.Print("Do you want to create an HTML file to embed the badges? (yes/no): ")
	confirmation, _ := reader.ReadString('\n')
	confirmation = strings.TrimSpace(strings.ToLower(confirmation))

	if confirmation == "yes" {
		htmlFileName := filepath.Join(destDir, "badges.html")
		htmlFile, err := os.Create(htmlFileName)
		if err != nil {
			fmt.Printf("Error creating HTML file: %v\n", err)
			return
		}
		defer htmlFile.Close()

		for _, htmlTag := range htmlContent {
			_, err = htmlFile.WriteString(htmlTag + "\n")
			if err != nil {
				fmt.Printf("Error writing to HTML file: %v\n", err)
				return
			}
		}
		fmt.Printf("HTML file created successfully as %s\n", htmlFileName)
	}
}
