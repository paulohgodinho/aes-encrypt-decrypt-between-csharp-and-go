using System;
using System.IO;
using System.Security.Cryptography;
using System.Text;

byte[] encryptedString = Convert.FromBase64String(args[0]);

string keyString = "";
string ivString = "";
if (Environment.GetEnvironmentVariable("Key") != null && Environment.GetEnvironmentVariable("IV") != null)
{
    keyString = Environment.GetEnvironmentVariable("Key")!;
    ivString = Environment.GetEnvironmentVariable("IV")!;
}

byte[] key = Encoding.UTF8.GetBytes(keyString!);
byte[] iv = Encoding.UTF8.GetBytes(ivString!);

var result = DecryptStringFromBytes_Aes(encryptedString, key, iv);
Console.WriteLine(result);

static string DecryptStringFromBytes_Aes(byte[] cipherText, byte[] Key, byte[] IV)
{
    // Check arguments.
    if (cipherText == null || cipherText.Length <= 0)
        throw new ArgumentNullException("cipherText");
    if (Key == null || Key.Length <= 0)
        throw new ArgumentNullException("Key");
    if (IV == null || IV.Length <= 0)
        throw new ArgumentNullException("IV");

    string plaintext = "";

    using (Aes aesAlg = Aes.Create())
    {
        // Console.WriteLine(aesAlg.Mode);
        // Console.WriteLine(aesAlg.BlockSize);

        aesAlg.Key = Key;
        aesAlg.IV = IV;
        aesAlg.Padding = PaddingMode.PKCS7;

        // Create a decryptor to perform the stream transform.
        ICryptoTransform decryptor = aesAlg.CreateDecryptor(aesAlg.Key, aesAlg.IV);
        // Create the streams used for decryption.
        using (MemoryStream msDecrypt = new MemoryStream(cipherText))
        {
            using (CryptoStream csDecrypt = new CryptoStream(msDecrypt, decryptor, CryptoStreamMode.Read))
            {
                using (StreamReader srDecrypt = new StreamReader(csDecrypt))
                {
                    // Read the decrypted bytes from the decrypting stream
                    // and place them in a string.
                    plaintext = srDecrypt.ReadToEnd();
                }
            }
        }
    }

    return plaintext;
}