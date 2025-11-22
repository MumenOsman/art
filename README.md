## Art Decoder/Encoder CLI Tool

This is a command-line utility built in Go that converts compressed art data (using repetition codes like `[5 #]`) into expanded text-based art, and also offers a mode to compress raw text-art back into the encoded format.

-----

## Installation and Setup

### Prerequisites

You must have **Go** (version 1.18 or higher recommended) installed on your system.
### Running the Tool

1.  **Clone the Repository (If applicable):**
    ```bash
    git clone https://github.com/MumenOsman/art.git
    cd art
    ```
2.  **Initialize the Module:**
    ```bash
    go mod tidy
    ```
3.  **Run the Program:**
    You must run the program using `go run .` from the root directory to ensure all internal packages (like `helpers`) are correctly linked.

-----

## Usage

The tool operates on a single input string provided as a command-line argument. The mode is controlled by optional flags.

### Command Structure

```bash
go run . [FLAGS] "<ENCODED_OR_RAW_ART>"
```

### Modes and Flags

| Flag        | Purpose                                                                                                                 | Default Mode (No flags) |
| :---------- | :---------------------------------------------------------------------------------------------------------------------- | :---------------------- |
| --encode    | Converts **raw text-art** into the compressed `[N Pattern]` format.                                                     | **Decode Mode**         |
| --multi     | Enables multi-line processing. Input is split by the newline (`\n`) character and each line is processed independently. | **Single-Line Mode**    |
| --help / -h | Displays the usage message and exits.                                                                                   | N/A                     |

-----

## Core Functionality

### 1\. Decoder Mode (Default)

The decoder expands consecutive character sequences defined by the format `[N Pattern]`.

**Format Rule:** `[Count Pattern]`

| Command     | Input                              | Output      | Description                                      |
| :---------- | :--------------------------------- | :---------- | :----------------------------------------------- |
| Single-Line | `go run . "[3 #]A[2 -_]"`          | `###A-_-_`  | Expands `[3 #]` to `###` and `[2 -_]` to `-_-_`. |
| Multi-Line  | `go run . --multi "[2 #]\nA[4 *]"` | `##\nA****` | Decodes each line separately.                    |

### 2\. Encoder Mode (`--encode`)

The encoder converts raw text into the compressed format by identifying the longest repeating sequences.

| Command | Input | Output | Description |
| :--- | :--- | :--- | :--- |
| Single-Line | `go run . --encode "AB AB AB"` | `[3 AB]` | Compresses a repeating pattern of length 2. |
| Multi-Line | `go run . --encode --multi "11\n@@@"` | `[2 1]\n[3 @]` | Encodes each line independently. |

-----

## Error Handling (Mandatory)

The program must display **`Error`** followed by a newline for any malformed encoded sequence.

| Failure Condition | Example Input |
| :--- | :--- |
| **Missing Count** | `go run . "[ #]"` |
| **Non-Numeric Count** | `go run . "[X #]"` |
| **Missing Separator Space** | `go run . "[5#]"` |
| **Empty Pattern** | `go run . "[5 ]"` |
| **Unbalanced Brackets** | `go run . "[5 #"` or `go run . "]5 #]"` |
| **Pattern Contains Brackets** | `go run . "[3 [A]"` |

If in multi-line mode (`--multi`), the program returns "Error" on the first line that fails decoding.