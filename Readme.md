# go-deployments CLI

A simple Go-based CLI server for executing predefined deployment or utility commands over HTTP with Basic Auth.

## Features

* Define a list of named commands in a YAML configuration.
* HTTP Basic Authentication to protect endpoint access.
* Execute commands on the server via `curl` or any HTTP client.
* Lightweight and easy to configure.

## Installation

Ensure you have Go 1.18+ installed and `$GOPATH/bin` or `$GOBIN` is on your `PATH`.

```bash
# Install the latest release of the CLI
go install github.com/mdtosif/go-deployments/cmd/go-deployments@latest
```

This places the `go-deployments` binary in your Go bin directory.

## Configuration

Create a `config.yaml` file in the same directory where you run the CLI (or specify a custom path via `--config`).

```yaml
# config.yaml

# List of services (commands) to expose
services:
  - name: touch-new-file
    cmd: "touch new_file.txt"

  - name: touch-new-file-2
    cmd: "touch new_file_2.txt"

  - name: touch-new-file-3
    cmd: "touch new_file_3.txt"

# Basic Auth credentials
auth:
  username: alice         # HTTP Basic Auth username
  password: secret123     # HTTP Basic Auth password

slack:
  webhook-url: https://hooks.slack.com/services/...

# Server port
port: 8081
```

### Configuration Fields

* **services**: Array of objects, each with:

  * `name` (string): Identifier used in the HTTP path.
  * `cmd` (string): Shell command to execute on the server.
* **auth**:

  * `username` (string): Basic Auth username.
  * `password` (string): Basic Auth password.
    
* **slack**:
  * `webhook-url` (string): Webhook url to send alert message in slack channel if any error in any command

* **port** (int): TCP port to listen on.

## Usage

Start the server:

```bash
# Default: looks for ./config.yaml
go-deployments

# Or specify a custom config file and port:
go-deployments --config=./path/to/config.yaml --port=9090
```

### Triggering a Command

To run a specific service command, send a GET request to `/deploy/{service-name}`:

```bash
curl http://alice:secret123@localhost:8081/deploy/touch-new-file
```

* **Response**: stdout/stderr of the executed command.

## Flags

```text
Usage of go-deployments:
  -config string
        Path to config file (default "internal/config/config.yaml")
```

You can invoke `-h`, `-help`, or `--help` to see this usage.

## Slack Alert example with webhook
![image](https://github.com/user-attachments/assets/acc2335b-6a41-4df4-a3e5-69f9e7193996)



## Security Note

* **Basic Auth**: Credentials are sent in plain text (Base64). Use HTTPS or a secure network.
* **Command Execution**: Defined commands run with the same privileges as the server process. Restrict commands to safe operations only.

## License

MIT License. See [LICENSE](LICENSE) for details.
