# gtools

`gtools` is a simple yet powerful CLI application that provides practical utilities, currently focused on generating secure tokens. Built to be efficient and easy to use, `gtools` aims to streamline everyday tasks directly from your terminal.

## Features

- **Token Generator**: Generate secure tokens for authentification, testing, or other purposes with customizable length.

## Installation

1. Clone the repository:
```bash
git clone https://github.com/ARTSYS-H/gtools.git
cd gtools
```
2. Build the application:
```bash
go build -o gtools
```
3. Run `gtools` to see available commands:
```bash
./gtools help
```

> **Note**: Make sure you have [Go installed](https://golang.org/doc/install)(version 1.16 or later is recommended)

## Usage

`gtools` currently supports generating secure tokens with a single command. Here's how to use it:

### Generate a Secure Token

To generate a secure token with the default length:

```bash
./gtools token
```

You can also specify the token length with the `--length` flag:

```bash
./gtools token --length 32
```

## Future Development

The `gtools` application is designed to be modular, allowing for easy addition of more useful commands in the future. Stay tuned for updates!

## License

This project is licensed under the MIT License. See the [LICENSE]() file for more details.

## Contact

For any questions, feedback, or suggestions, feel free to reach out.

---

Thank you for using `gtools`!
