# Fosskey CLI

Fosskey is a [**F**]ree, [**O**]pen-source, [**S**]ecure, and [**S**]elf-custodial keychain.

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
