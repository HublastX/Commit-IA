# CommitAI

CommitAI is a command-line tool built in Go that leverages Large Language Models (LLMs) to analyze your Git code changes and automatically generate semantic commit messages.

## Features

### 🎯 Core Features
-   🤖 **Automatic commit message generation** - AI-powered commit analysis
-   🌐 **Dual processing modes** - Local and remote LLM processing
-   🔄 **Easy configuration management** - Simple setup and updates
-   🌍 **Multi-language support** - Commit messages in any language
-   🏷️ **Custom commit tags** - Force specific commit types

### ✨ New in Version 2.0
-   📋 **Multiple commit formats** - Support for native Git or custom standards
-   🎨 **Personal prompts** - Define your own commit patterns and styles
-   🔌 **Universal model support** - Any local model or free web version compatibility
-   📦 **Easy installation** - Available via NPM or compiled binaries
-   🔧 **Native Git integration** - Enhanced Git support with advanced customization
-   🖥️ **Improved interface** - Better menu navigation and user experience
-   ⚡ **Performance boost** - Faster execution and optimized CLI performance

## Installation

CommitAI offers multiple installation methods to suit your preferences:

### 🚀 Quick Install (Recommended)

**Option 1: NPM Global Install**
```bash
npm install -g commit-ai-hub
```

**Option 2: From Repository**
```bash
git clone https://github.com/HublastX/Commit-IA
cd Commit-IA
npm install -g
```

### 🔧 Manual Compilation

**Prerequisites:**
-   Go 1.24 or higher
-   Node.js 20+ (for NPM installation)

**Steps:**
1. Clone the repository:
   ```bash
   git clone https://github.com/HublastX/Commit-IA
   cd Commit-IA
   ```

2. Compile for your platform using Go:
   ```bash
   go build -o commitai ./cmd
   ```

3. Move binary to your PATH or use local installer scripts

### 🛠️ Alternative Installation Methods

**Legacy Installer Scripts (Linux/macOS):**
```bash
chmod +x ./install.sh
./install.sh
```

**Legacy Installer Scripts (Windows):**
```cmd
.\install.bat
```

### 🔧 Troubleshooting Installation

If you encounter build errors during manual compilation, install the required development packages:

<details>
<summary><b>Windows</b></summary>

1. Install Scoop package manager (PowerShell):
   ```powershell
   Set-ExecutionPolicy RemoteSigned -scope CurrentUser
   iwr -useb get.scoop.sh | iex
   ```

2. Install dependencies:
   ```powershell
   scoop install mingw
   ```

</details>

<details>
<summary><b>Ubuntu/Debian</b></summary>

```bash
sudo apt update && sudo apt install -y \
    gcc libc6-dev libx11-dev xorg-dev \
    libxtst-dev libpng-dev libxcursor-dev \
    libxrandr-dev libxinerama-dev libdbus-1-dev \
    tesseract-ocr
```

</details>

<details>
<summary><b>Arch Linux/Manjaro</b></summary>

```bash
sudo pacman -Syu && sudo pacman -S --needed \
    gcc glibc libx11 xorg-server-devel \
    libxtst libpng libxcursor libxrandr \
    libxinerama dbus tesseract
```

</details>

## Usage

After using `git add` to stage your changes, you can use commitai to generate commit messages.

### Basic Command

```bash
commitai
```

### Operating Modes

commitai offers two operating modes:

1. **Remote Web Mode** - Access LLMs remotely without additional configuration (might be slower)
2. **Local Mode** - Run the LLM API locally using Docker (requires provider configuration)

To switch between modes or update configuration:

```bash
commitai --update
```

### Additional Options

| Option     | Description                          | Example                                   |
| ---------- | ------------------------------------ | ----------------------------------------- |
| `-d`       | Add additional context               | `commitai -d "Added login functionality"` |
| `-l`       | Specify commit language              | `commitai -l "English"`                   |
| `-t`       | Force specific commit tag            | `commitai -t "feat"`                      |
| `--update` | Update configuration or switch modes | `commitai --update`                       |

### Examples

Generate a commit with additional context:

```bash
commitai -d "Created user login feature with OAuth support"
```

Generate a commit message in English:

```bash
commitai -l "English"
```

Force a specific commit tag:

```bash
commitai -t "feat"
```

Update configuration or switch between local and web modes:

```bash
commitai --update
```

## Local LLM Configuration

When selecting the Local mode, simply configure through the CLI:

1. Run `commitai --update` to access configuration
2. Select **Local Mode** when prompted
3. Choose your LLM `provider` (Google, OpenAI, Anthropic, etc.)
4. Select an available `model` for your chosen provider
5. Enter your `API key` for the provider

That's it! No Docker setup required - the tool will handle everything automatically.

## Configuration Details

-   Web mode uses Google's `gemini-flash-2` model by default
-   All LLM configurations are stored in the `Bot` directory
-   Custom configurations (provider, model, API token) are managed in `Bot/config/config.json`

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

We recommend using Husky for commit validation:

1. Ensure Node.js 22 is installed
2. Run `npm install` to set up Husky, which will validate your commits before submission

## License

This project is licensed under the [Apache-2.0 license](LICENSE).
