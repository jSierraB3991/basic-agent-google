# AI Code Agent in Go with Gemini

An interactive terminal-based AI coding assistant written in Go that integrates with Google's **Gemini** models via the official `google.golang.org/genai` SDK. The agent is capable of reasoning, answering questions, and executing tool calls (such as reading, creating, and deleting files/directories, listing file info, and running Go/Python tests) directly on your local filesystem.

---

## Features

- 🤖 **Gemini AI Integration**: Powered by Google GenAI SDK supporting advanced function calling and parallel tool executions.
- 🛠️ **Local Development Tools**: Allows Gemini to autonomously execute local file system and testing operations:
  - **File Operations**: Create, read, delete files, and retrieve directory/file info safely.
  - **Directory Operations**: Create and delete directories (with guardrails protecting project root).
  - **Testing**: Run Go and Python unit tests with captured output returned to the model.
- 🖥️ **Interactive CLI**: Easy-to-use REPL terminal interface supporting commands like `quit`, `exit`, or `salir`, as well as clean EOF and SIGINT (Ctrl+C) handling.
- ⚙️ **Configurable Environment & Flags**: Reads configuration from `.env` and supports runtime CLI flags (`-model`, `-verbose`, `-max-steps`).

---

## Project Structure

```text
├── agent/
│   ├── agent.go              # Core agent REPL loop and message handling
│   ├── configurate.go        # Gemini client configuration & model settings
│   ├── request_input.go      # Communicates with Gemini API with tool declarations
│   └── tools/
│       ├── execute_tool_call.go # Dispatches tool function calls from Gemini
│       ├── get_tools.go      # Defines available GenAI function declarations & config
│       ├── tools_directory.go# Directory creation and deletion tools
│       ├── tools_files.go    # File creation, reading, deletion, and info tools
│       └── tools_run_testing.go # Go and Python test execution tools
├── functions/                # Tool implementations and system operations
├── go.mod                    # Go module dependencies
├── go.sum                    # Go checksum file
└── main.go                   # Application entry point
```

---

## Prerequisites

- **Go** (version 1.22 or higher; project tested on Go 1.26+)
- **Gemini API Key** from [Google AI Studio](https://aistudio.google.com/)

---

## Setup & Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/jsierrab3991/basic-agent-google.git
   cd basic-agent-google
   ```

2. **Configure Environment Variables**:
   Create a `.env` file in the root directory of the project and add your Gemini API key (and optionally configure the model or verbose mode):
   ```env
   GEMINI_API_KEY=your_gemini_api_key_here
   GEMINI_MODEL=gemini-2.5-flash  # Optional (defaults to gemini-2.5-flash)
   VERBOSE=true                  # Optional (defaults to false)
   ```

3. **Install Dependencies**:
   ```bash
   go mod download
   ```

---

## Running the Application

Start the interactive agent CLI:

```bash
go run main.go
```

You can also pass optional CLI arguments:

```bash
go run main.go -model gemini-2.5-flash -verbose -max-steps 10
```

Once running, you can interact with the agent in the terminal prompt:

```text
Simple Code agent With gemini 
>: Hello, can you list the files in the current directory?
```

To exit the agent, type `quit`, `exit`, `salir`, or press `Ctrl+D` / `Ctrl+C`.

---

## Available Agent Tools

The Gemini model can dynamically invoke the following tools during your conversation:
1. **`GetFilesInfo`**: Lists files and subdirectories inside a working directory.
2. **`ReadFileContent`**: Reads and returns the complete content of a target file.
3. **`CreateFile`**: Creates or overwrites a file with specified content at a given path.
4. **`DeleteFile`**: Deletes a specified file from the filesystem.
5. **`CreateDir`**: Creates a directory and any necessary parent directories.
6. **`DeleteDir`**: Deletes a directory and its contents recursively (protecting root).
7. **`RunTestInGo`**: Executes Go test suites (package, test file, or `./...`) and captures stdout/stderr.
8. **`RunPythonTests`**: Executes Python unit tests using `pytest` or `unittest`.

---

## License

This project is licensed under the GNU General Public License v3.0 (GNU GPLv3).
