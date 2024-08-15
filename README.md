# File Encrypter

File Encrypter can encrypt and decrypt individual files or all files in a directory.

> [!WARNING]
> This utilizes the XOR cipher, considered a rather weak encryption algorithm.

## Usage

1. Come up with a **random key** that will be used to encrypt the files. You must remember this string as all data would be **lost permanently** without it.

2. Execute the following command in the terminal:
```bash
ec.exe <path> <key>
```

`path` - The name of the file or directory

`key` - Your secret string that is used to encrypt and decrypt the file(s)

Example:
```bash
ec.exe secret.txt SecretKey123
```

Your data is now encrypted and can only be restored with your chosen key. To decrypt the file, execute the same command.

## Building

First of all, make sure to have the Golang CLI installed and it is added to PATH.
To then compile the program and convert it to an executable, run the following command:

```bash
go build -o ec.exe main.go
```

Alternatively, there are precompiled binaries that can be found under [releases](https://github.com/amueller0/file-encrypter/releases).
This does not require the installation of Go.
