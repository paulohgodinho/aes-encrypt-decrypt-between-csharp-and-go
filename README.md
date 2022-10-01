# aes-encrypt-decrypt-between-csharp-and-go

This simple project aims to show interoperability between **encrypting and decrypting** libraries of the C# and Go languages.
Cryptography libraries can be **very** opinionated, this project shows how to make these two talk while using **AES in MODE_CBC**

## Motivation

I have a backend in Go that used AES CFB as the mode for its encryption, and I was simply not capable of decrypting its payloads using C# built-in libraries. The desision was made to move to CBC and overcome its input padding issues on both sides, this repository was created to validate that AWS in MODE CBC would work.

## Sample 

```
./RunEncryptionTest.ps1 "ThisIsASecret_NeverTellAnyone"
```

### Output

```
[Encrypting With Go]
Result: Wpt7exoRz9zxrxBedrOgVMKf2zoJ7s8Gaig/UbVF/Ys=
[Decrypt With C#]
Result: ThisIsASecret_NeverTellAnyone
------------------------
[Encrypt With C#]
Result: Wpt7exoRz9zxrxBedrOgVMKf2zoJ7s8Gaig/UbVF/Ys=
[Decrypt With Go]
Result: ThisIsASecret_NeverTellAnyone
```

## Project Structure
You can find the encrypting and decrypting versions of the algorithms in its respective folders
```
aes-encrypt-decrypt-between-csharp-and-go
├─ encrypt
│  ├─ csharp
│  ├─ go
├─ decrypt
│  ├─ csharp
│  ├─ go
```

## Calling directly
The `RunEncryptionTest.ps1` file shows how to call the projects individually, as an example, here is a call to encrypt a string with Go and decrypt it with C#
```
$env:Key = '2s5v8y/B?E(H+MbQeThWmYq3t6w9z$C&'
$env:IV = 'XRFLuBIzjpZ8$5F#'

# Go Encription
Push-Location ./encrypt/go
$encryptedWithGoResult = go run main.go $args[0]
Write-Output "Result: $encryptedWithGoResult"
Pop-Location

# CSharp decryption
Push-Location ./decrypt/csharp
$decryptWithCSharpResult = dotnet run $encryptedWithGoResult
Write-Output "Result: $decryptWithCSharpResult"
Pop-Location
```
