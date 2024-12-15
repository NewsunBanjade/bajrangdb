Here is a clear and easy-to-understand `README.md` for your `bajrang` CLI tool:

---

# Bajrang CLI Tool

## Overview

The **Bajrang CLI Tool** is a command-line interface (CLI) for managing a simple local database. The tool allows you to perform basic operations such as **create**, **update**, **delete**, and **get** data in the database file. The database is stored locally in a JSON file and can be accessed through various commands.

### Features:
- **Create**: Add a new entry to the database.
- **Update**: Modify an existing entry.
- **Delete**: Remove an entry from the database.
- **Get**: Retrieve data for a specific entry.

## Installation

### 1. Clone or Download the Repository

First, clone or download the repository to your local machine.

```bash
git clone https://github.com/your-username/bajrang.git
cd bajrang
```

### 2. Install Dependencies

The tool uses the `cobra` library to build the CLI commands. You can install the required dependencies by running:

```bash
go get -u github.com/spf13/cobra@latest
```

### 3. Build the Tool

Once the dependencies are installed, you can build the tool using Go:

```bash
go build -o bajrang main.go
```

This will generate an executable file called `bajrang` in your project directory.

## Usage

You can interact with the database using the following commands:

### Command Structure

```bash
./bajrang [command] [arguments]
```

### Available Commands

1. **Create a New Entry**

   To create a new entry in the database, use the `create` command. It requires an `id` and `data` as arguments.

   Example:

   ```bash
   ./bajrang create <id> <data>
   ```

   - **id**: The unique identifier for the entry.
   - **data**: The content or data for the entry.

   Example:

   ```bash
   ./bajrang create user1 "Sample Data"
   ```

2. **Update an Existing Entry**

   To update an existing entry, use the `update` command. It also requires an `id` and `data`.

   Example:

   ```bash
   ./bajrang update <id> <data>
   ```

   - **id**: The unique identifier for the entry you want to update.
   - **data**: The updated content or data.

   Example:

   ```bash
   ./bajrang update user1 "Updated Data"
   ```

3. **Delete an Entry**

   To delete an entry from the database, use the `delete` command. It only requires an `id`.

   Example:

   ```bash
   ./bajrang delete <id>
   ```

   - **id**: The unique identifier for the entry you want to delete.

   Example:

   ```bash
   ./bajrang delete user1
   ```

4. **Get an Entry**

   To retrieve the data of an existing entry, use the `get` command. It requires an `id`.

   Example:

   ```bash
   ./bajrang get <id>
   ```

   - **id**: The unique identifier for the entry you want to get.

   Example:

   ```bash
   ./bajrang get user1
   ```

## Example Usage

Here’s a quick example of how to use the tool:

1. **Create an entry**:
   ```bash
   ./bajrang create user1 "Sample Data"
   ```

2. **Get the created entry**:
   ```bash
   ./bajrang get user1
   ```

3. **Update the entry**:
   ```bash
   ./bajrang update user1 "Updated Data"
   ```

4. **Delete the entry**:
   ```bash
   ./bajrang delete user1
   ```

## File Structure

```plaintext
bajrang/
├── main.go       # Main Go program with Cobra CLI
└── db/
    └── data.json # Database file where data is stored
```

## How It Works

- The `bajrang` tool interacts with a local JSON file (`data.json`) to store and retrieve data.
- The tool uses a simple in-memory store to hold data, and any changes are saved to the JSON file after each operation.
- Each entry is identified by a unique `id` and associated with `data` that can be created, updated, or deleted.

## Contributing

If you'd like to contribute to this project, feel free to fork it, make your changes, and open a pull request.

## License

This project is open source and available under the [MIT License](LICENSE).
