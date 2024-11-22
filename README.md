### README for `lookup-cli`

#### Overview

`lookup-cli` is a command-line tool written in Go that allows you to lookup specific fields in a YAML file based on provided names. The tool is designed to be simple and efficient, providing clear output messages when inputs are missing or invalid.

#### Usage

To use `lookup-cli`, run the following command:

```bash
lookup-cli <name> <output_field>
```

- `<name>`: The name to lookup in the YAML file.
- `<output_field>`: The field to return for the given name (e.g., `age`, `occupation`).

**Example Usage:**

```bash
$ lookup-cli Alice age
18

$ lookup-cli Bob occupation
unemployed

$ lookup-cli Charlie occupation
Field not found

$ lookup-cli Eve age
Name not found

$ lookup-cli
Usage: lookup-cli <name> <output_field>
```

#### Building the Project

##### Regular Build

1. **Prerequisites:**

   - Ensure you have Go installed (version 1.23 or later).
   - Install the required dependencies by running:
     ```bash
     go mod tidy
     ```

2. **Build the Executable:**

   ```bash
   go build -o lookup-cli
   ```

3. **Run the Executable:**
   ```bash
   ./lookup-cli <name> <output_field>
   ```

##### Containerized Build

1. **Prerequisites:**

   - Ensure you have Docker installed.

2. **Create a Dockerfile:**

   ```Dockerfile
   # Use the official Go image as a base
   FROM golang:1.23

   # Set the working directory
   WORKDIR /app

   # Copy the source code into the container
   COPY . .

   # Download dependencies
   RUN go mod tidy

   # Build the executable
   RUN go build -o lookup-cli

   # Set the entry point for the container
   ENTRYPOINT ["./lookup-cli"]
   ```

3. **Build the Docker Image:**

   ```bash
   docker build -t lookup-cli .
   ```

4. **Run the Docker Container:**
   ```bash
   docker run -it lookup-cli <name> <output_field>
   ```

#### Project Structure

- `lookup_test.go`: Contains unit tests for the `lookup-cli` functionality.
- `lookup.go`: Contains the main logic for the CLI, including command initialization and execution.
- `main.go`: The entry point of the application.
- `go.mod`: Defines the Go module and its dependencies.
- `Software Engineer Code Test.txt`: Contains the task description and requirements.

#### Additional Notes

- The YAML file used by the tool is specified via the `--file` or `-f` flag, with a default value of `data.yaml`.
- Ensure that the YAML file is correctly formatted and located in the expected path.

#### Contributing

Feel free to fork the repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

#### License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

This README provides a comprehensive guide on how to use, build, and containerize the `lookup-cli` tool. If you have any questions or need further assistance, please refer to the project's documentation or open an issue on the repository.
