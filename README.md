# solo.to-checker

This checks the availability of usernames using the Solo.to API. It reads a list of usernames from a text file, checks each username, and saves the available ones to a new file.

## Features

- **Username Checking**: Verifies if a username is available, reserved, or blocked.
- **File Input/Output**: Reads usernames from a file (`names.txt`) and writes available usernames to another file (`available_usernames.txt`).
- **Error Handling**: Handles errors gracefully during file operations and API requests.

## Requirements

- Go (version 1.15 or higher)
