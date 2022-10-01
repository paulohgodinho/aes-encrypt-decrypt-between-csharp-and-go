$env:Key = '2s5v8y/B?E(H+MbQeThWmYq3t6w9z$C&'
$env:IV = 'XRFLuBIzjpZ8$5F#'

Write-Output "[Encrypting With Go]"
Push-Location ./encrypt/go
$encryptedWithGoResult = go run main.go $args[0]
Write-Output "Result: $encryptedWithGoResult"
Pop-Location

Write-Output "[Decrypt With C#]"
Push-Location ./decrypt/csharp
$decryptWithCSharpResult = dotnet run $encryptedWithGoResult
Write-Output "Result: $decryptWithCSharpResult"
Pop-Location

Write-Output "------------------------"

Write-Output "[Encrypt With C#]"
Push-Location ./encrypt/csharp
$encryptWithCSharpResult = dotnet run $args[0]
Write-Output "Result: $encryptWithCSharpResult"
Pop-Location

Write-Output "[Decrypt With Go]"
Push-Location ./decrypt/go
$decryptedWithGoResult = go run main.go $encryptWithCSharpResult
Write-Output "Result: $decryptedWithGoResult"
Pop-Location

