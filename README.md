<p align="center">
  <a href="https://fosskey.com">
    <img alt="Fosskey" src="https://user-images.githubusercontent.com/508043/210156279-fce2059d-4715-46a7-94e9-35f1d67ea431.png" height="64" />
  </a>
</p>

# Fosskey CLI

Fosskey is a [**F**]ree, [**O**]pen-source, [**S**]ecure, and [**S**]elf-custodial keychain.

## How do "they" store our passwords?

Unfortunately, there are still so many websites that store our passwords as plain text or use weak encryption. See for yourself at [Password Storage Disclosures](https://pulse.michalspacek.cz/passwords/storages). If you use the same password for multiple websites, your privacy and identity are most likely vulnerable to data breaches and hacks. It's time to use a unique password for every website. But how would you remember them all? Password manager? 🤔

## Why Fosskey?

We are super skeptical when it comes to our security and privacy. We don't trust closed-source software to keep our private affairs. And neither we trust a third party to safe keep our secrets in their custody.

Browser's built-in password managers and third-party apps store the passwords on their servers. The storage could be encrypted, but the level of privacy is still questionable. The goal of the Fosskey project is to give the power back to the hand of the users. The source code is open for anyone to audit, and it's free forever for everyone. The passwords are only stored locally on your device using [XChaCha20-Poly1305][chacha20-poly1305] encryption and with [Argon2id][argon2] key-derivation function.

## Install

There is no executable file to download. As we mentioned before: don't trust the binary, only trust the source. So you have to compile the code directly from the source code.

**Step 1:**

[Install Go](https://go.dev/) and [install Git](https://git-scm.com/) if you don't have them installed already.

**Step 2:**

Clone the repo:

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

## Usage

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

[chacha20-poly1305]: https://en.wikipedia.org/wiki/ChaCha20-Poly1305
[argon2]: https://en.wikipedia.org/wiki/Argon2
