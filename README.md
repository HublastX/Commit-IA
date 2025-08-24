# CommitAI

CommitAI is a command-line tool built in Go that leverages Large Language Models (LLMs) to analyze your Git code changes and automatically generate semantic commit messages.

## Features

-   🤖 Automatic commit message generation
-   🌐 Support for both local and remote LLM processing
-   🔄 Easy configuration management
-   🌍 Multi-language commit message support
-   🏷️ Custom commit tag selection

## Installation

### Prerequisites

-   Go 1.24 or higher
-   Docker (optional, for local LLM mode)

### Installation Steps

### Troubleshooting Installation

If you encounter build errors during installation, you may need to install additional development packages:

<details>
<summary><b>Windows</b></summary>

1.  Install Scoop package manager (in PowerShell):

    ```powershell
    Set-ExecutionPolicy RemoteSigned -scope CurrentUser
    iwr -useb get.scoop.sh | iex
    ```

2.  Install required dependencies:

    ```powershell
    scoop install mingw
    ```

    </details>

<details>
<summary><b>Ubuntu/Debian-based Systems</b></summary>

```bash
sudo apt update
sudo apt install -y \
    gcc \
    libc6-dev \
    libx11-dev \
    xorg-dev \
    libxtst-dev \
    libpng-dev \
    libxcursor-dev \
    libxrandr-dev \
    libxinerama-dev \
    libdbus-1-dev \
    tesseract-ocr
```

</details>

<details>
<summary><b>Arch Linux/Manjaro</b></summary>

```bash
sudo pacman -Syu
sudo pacman -S --needed \
     gcc \
     glibc \
     libx11 \
     xorg-server-devel \
     libxtst \
     libpng \
     libxcursor \
     libxrandr \
     libxinerama \
     dbus \
     tesseract
```

</details>

#### Linux and macOS

1. Clone the repository:

    ```bash
    git clone https://github.com/HublastX/Commit-IA
    cd Commit-IA
    ```

2. Make the installer executable and run it:

    ```bash
    chmod +x ./install.sh
    ./install.sh
    ```

#### Windows

1. Clone the repository:

    ```bash
    git clone https://github.com/HublastX/Commit-IA
    cd Commit-IA
    ```

2. Run the installer:

    ```bash
    .\install.bat
    ```

The binary will be compiled and installed, making `commitai` available from anywhere in your system.

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

When selecting the Local mode, follow these steps:

1. Select your LLM `provider` (Google, OpenAI, etc.)
2. Choose an available `model` for your selected provider
3. Enter your `API key` for the chosen provider
4. Start the local API server using Docker:

```bash
docker compose up --build
```

The local API server will run on port 6013 by default.

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
