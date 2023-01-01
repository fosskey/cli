<p align="center">
  <a href="https://fosskey.com">
    <img alt="Fosskey" src="https://user-images.githubusercontent.com/508043/210156279-fce2059d-4715-46a7-94e9-35f1d67ea431.png" width="143" />
  </a>
</p>

# Fosskey CLI

Fosskey is a [**F**]ree, [**O**]pen-source, [**S**]ecure, and [**S**]elf-custodial keychain.

## How do "they" store our passwords?

Unfortunately, there are still so many websites that store our passwords as plain text or use weak encryption. See for yourself at [Password Storage Disclosures](https://pulse.michalspacek.cz/passwords/storages). If we use the same password for multiple websites, our privacy and identity are most likely vulnerable to data breaches and hacks. It's time to use a unique password for every website. But how would we remember them all? Password manager? 🤔

## Why Fosskey?

We are super skeptical when it comes to our security and privacy. We don't trust closed-source software to keep our private affairs. And neither we trust a third party to safe keep our secrets in their custody.

Browser's built-in password managers and third-party apps store the passwords on their servers. The storage could be encrypted, but the level of privacy is still questionable. The goal of the Fosskey project is to give the power back to the hand of the users. The source code is open for anyone to audit, and it's free forever for everyone. Your secrets and passwords are only stored locally on your device using [XChaCha20-Poly1305][chacha20-poly1305] encryption and with [Argon2id][argon2] key-derivation function.

## How do I Install?

There is no executable file to download. As we mentioned before: don't trust the binary, only trust the source. So you have to compile the code directly from the source code.

**Step 1:**

[Install Go](https://go.dev/) and [install Git](https://git-scm.com/) if you don't have them installed already.

**Step 2:**

Open up your terminal (if you're on macOS or Linux) or command prompt (if you're on Windows). Then run the following command to clone the repo:

```
git clone https://github.com/fosskey/cli.git
cd cli
```

**Step 3:**

If you're on macOS or Linux, run:

```
go build -o bin/foss
mkdir -p $GOPATH/bin
cp bin/foss $GOPATH/bin
```

If you're on Windows, run:

```
go build -o bin\foss.exe
mkdir %GOPATH%\bin
copy bin\foss.exe %GOPATH%\bin
```

**Step 4:**

Now run:

```
foss
```

If the install is successful you should see the usage information.

## How do I use it?

Insert a new secret:

```
⚡ foss insert Gmail

Enter master key: [···]
Enter new secret: [···]

Gmail is now inserted into the vault
```

List all:

```
⚡ foss ls

Enter master key: [···]

Vault
├──Coinbase
├──Gmail
└──Twitter
```

Fetch:

```
⚡ foss fetch Gmail

Enter master key: [···]

MyGma!lP@55
```

Update:

```
⚡ foss update Gmail

Enter master key: [···]
Enter new secret: [···]

Gmail is now updated in the vault
```

Delete:

```
⚡ foss delete Gmail

Enter master key: [···]

Gmail is now deleted from the vault
```

Change the master key:

```
⚡ foss rekey

Enter old master key: [···]
Enter new master key: [···]

Masterkey is now changed
```

## How does it work?

Fosskey does not store the master key. Instead, it uses the Argon2id key-derivation function to generate a 256-bit key from the master key and a 128-bit random salt. And then, the derived key and another 192-bit random nonce are used to encrypt the secret payload (plain text) using XChaCha20-Poly1305 AEAD (Authenticated Encryption with Associated Data). Finally, the nonce, the cipher text, and the salt are glued together and stored in the `.foss/vault` file under your user's home directory (e.g. in `~/.foss/vault` on macOS and Linux, or `C:\Users\yourname\.foss\vault` on Windows).

<p align="left">
  <a href="https://fosskey.com">
    <img alt="Fosskey" src="https://user-images.githubusercontent.com/508043/210265188-671411b3-433c-4713-8734-1b1b8ee07d76.png" width="830" />
  </a>
</p>

## How secure is it?

While using the recommended parameters specified in [RFC 9106][rfc9106-params], the encryption/decryption method took about 0.8 seconds to process on a quad-core Intel processor with 16 GiB of memory. If a master key is composed of 8 characters of upper-case (A-Z), lower-case (a-z) letters and numbers (0-9), and symbols (32), there will be a total of 94 possible characters. Therefore, at least a total of B=nP(r-1) brute-force attacks is required to guess the correct master key. Here "B" is the permutation of (n, r-1). Thus, with the target hardware configuration (quad-core, 16 GiB memory), it will take about 1.3 million computation years to brute-force the 8-character long master key.

[chacha20-poly1305]: https://en.wikipedia.org/wiki/ChaCha20-Poly1305
[argon2]: https://en.wikipedia.org/wiki/Argon2
[rfc9106-params]: https://www.rfc-editor.org/rfc/rfc9106.html#name-parameter-choice
