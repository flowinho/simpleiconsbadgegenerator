# SimpleIcons Badge Generator

SimpleIcons Badge Generator is a CLI tool written in Go that allows you to generate and download badges for various services using icons from [SimpleIcons](https://simpleicons.org/). The badges can be customized with a label, background color, and font color. The tool also provides an option to generate HTML code to embed the badges.

## Features

- Generate and download badges for multiple services using SimpleIcons.
- Customize badge labels, background colors, and font colors.
- Specify a destination directory for saving badges.
- Optionally generate an HTML file to embed the badges.

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/your-username/simpleicons-badge-generator.git
    cd simpleicons-badge-generator
    ```

2. **Build the application:**

    For macOS:

    ```sh
    GOOS=darwin GOARCH=amd64 go build -o simpleicons-macos main.go
    ```

    For Linux:

    ```sh
    GOOS=linux GOARCH=amd64 go build -o simpleicons-linux main.go
    ```

3. **Make the binary executable:**

    ```sh
    chmod +x simpleicons-macos
    chmod +x simpleicons-linux
    ```

## Usage

```sh
./simpleicons-macos <service1>, <service2>, ...
```

### Example

```sh
./simpleicons-macos Swift, UIKit, Jekyll
```

### Steps

1. **Enter the destination directory:**
    You will be prompted to enter the directory where you want to save the badges. The directory will be created if it doesn't exist.

    ```
    Enter the destination directory: badges
    ```

2. **Enter the text on the badge:**
    For each service, you will be prompted to enter the text that should appear on the badge.

    ```
    Found icon for Swift (computed slug: swift) with color #F05138
    Enter the text on the badge: Swift Language
    ```

3. **Confirm HTML file generation:**
    After downloading all the badges, you will be asked if you want to generate an HTML file to embed the badges.

    ```
    Do you want to create an HTML file to embed the badges? (yes/no): yes
    ```

    If you confirm, an HTML file named `badges.html` will be created in the specified destination directory.

## Output

- The badges will be saved in the specified destination directory with the following naming convention: `<label_text>-badge.svg`.
- If HTML generation is confirmed, an HTML file `badges.html` will be created with `<img>` tags for each badge.

### Example HTML

```html
<img alt="swift" src="swift_language-badge.svg" />
<img alt="uikit" src="uikit_framework-badge.svg" />
<img alt="jekyll" src="jekyll_static_site-badge.svg" />
```

## Error Handling

- If a service is not found, the tool will skip it and continue with the next service.
- If there are any errors during the download or file creation process, an error message will be displayed.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [SimpleIcons](https://simpleicons.org/) for providing a collection of free SVG icons.
- [Shields.io](https://shields.io/) for generating the badges.
