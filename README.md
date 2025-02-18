# secure-files-go-gui

This is a GUI for [secure-files-go](https://github.com/darshanags/secure-files-go) - a file encryption/decryption program written in [Go](https://go.dev/). The [original version](https://github.com/darshanags/secure-files) of the program was written in Zig.

This implements [RFC8439](https://datatracker.ietf.org/doc/html/rfc8439) which uses ChaCha20 as the cipher and Poly1305 for authentication, and [Argon2](https://datatracker.ietf.org/doc/html/rfc9106) for key derivation (KDF) from a given password.

**This is in no shape or form a well-polished program, so use at your own risk.**

## Program Flow
### Encryption
```mermaid
---
config:
  look: classic
  theme: neutral
---
graph TB
A[Start] --> B[\Password\] --> 
C[Argon2 - Generate 128 bit salt] --> 
D["Argon2 - Generate 256 bit User Encryption Key (UEK)<br> using the password and salt"] -->
E["Generate 256 bit Data Encryption Key (DEK) and 96 bit nonce"] -->
F[Chacha20-Poly1305 - Encrypt DEK<br> using UEK and nonce]
F --> G[Store encrypted DEK,<br> nonce, and salt<br> in output file]
E --> H[Chacha20-Poly1305 - Encrypt input file data<br> using DEK]
H --> I[Store encrypted data in output file]
```

> Written with [StackEdit](https://stackedit.io/).
