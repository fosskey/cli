# Fosskey CLI

Fosskey is a [**F**]ree, [**O**]pen-source, [**S**]ecure, and [**S**]elf-custodial keychain.

## Why Fosskey?

Would you store your passwords and private keys in the cloud? If you're using your browser's password manager or a third-party tool, you're probably storing them on their server. It could be encrypted, but the level of privacy is still questionable. Fosskey allows you to securely manage your secrets with zero trust in a third party and without compromising your security or privacy. The source code is open for anyone to audit, and it's free forever for everyone. The passwords are only stored locally on your device using [XChaCha20-Poly1305][chacha20-poly1305] encryption along with [Argon2id][argon2] key derivation function.

## Usage

### Insert a new secret

```
⚡ foss insert Gmail

Enter master key: [···]
Enter new secret: [···]

Gmail is now inserted into the vault
```

### List all

```
⚡ foss ls

Enter master key: [···]

Vault
├──Coinbase
├──Gmail
└──Twitter
```

### Fetch

```
⚡ foss fetch Gmail

Enter master key: [···]

MyGma!lP@55
```

### Update

```
⚡ foss update Gmail

Enter master key: [···]
Enter new secret: [···]

Gmail is now updated in the vault
```

### Delete

```
⚡ foss delete Gmail

Enter master key: [···]

Gmail is now deleted from the vault
```

### Change master key

```
⚡ foss rekey

Enter old master key: [···]
Enter new master key: [···]

Masterkey is now changed
```

[chacha20-poly1305]: https://en.wikipedia.org/wiki/ChaCha20-Poly1305
[argon2]: https://en.wikipedia.org/wiki/Argon2
