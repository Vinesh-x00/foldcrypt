# Foldcrypt

Foldcrypt is a Go-based command-line tool that allows you to securely encrypt and decrypt directories using AES encryption. 
It's designed to safeguard your sensitive data.

Key Features 

- AES Encryption: Utilizes Advanced Encryption Standard for robust data protection.
- Directory Encryption: Encrypts entire directories, preserving their structure.
- Command-Line Interface: Easy-to-use CLI for efficient encryption and decryption.
- Shredding: shread files or directories for secure deletion.

## Installation

You need to have Flutter 22.3+ installed

```bash
  git clone https://github.com/Vinesh-x00/foldcrypt.git
```

Go to the project directory
```bash
  cd foldcrypt
```

Create binary
```bash
go build -o foldcrypt
```

## Usage

```bash
foldcrypt lock [options] [directory]
foldcrypt relock [options] [directory]
foldcrypt unlock [options] [directory] [file]
foldcrypt shread [options] [directory | file]
```

Use `foldcrypt help` for more information

## Tools used
- golang
- cobra CLI tool

