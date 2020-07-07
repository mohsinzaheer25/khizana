# Khizana

Khizana is an arabic word means locker. This tool is an encrypted personal vault to store secrets on your system.

## Encryption Method

AES 256 Cipher

## Options

Khizana offers different option like `init` `add` `view` `get` `update` `delete` & `destroy`

```cassandraql
Khizana is a personal vault for your sensitive information.

Options:

init				Initialize Your Khizana
view				To View Your Khizana
get KEY				To get Value of The Key
add KEY VALUE			To Add Key Value To Your Khizana
update KEY VALUE		To Updated Value Of The Key
delete KEY			To Remove Key From Khizana
destroy				To Destroy Khizana
help				To Get Help Of Khizana
```

#### Initialization

Initialization is the first step to setup Khizana. Password is mandatory in each action and once you enter password it will get initialized.

```cassandraql
$ khizana init
Enter Password For Your Khizana

Khizana has initialize successfully
```

#### Adding Secret

You can add secret as a `key` `value` to Khizana. Both `key` & `value` is mandatory to add secret.

**Note:** It required `key` to be unique.

```cassandraql
$ khizana add username mohsinzaheer25
Enter Password For Your Khizana

Key Value Added
```

#### View

You can view your secrets using `view` command.

```cassandraql
$ khizana view
Enter Password For Your Khizana

# Khizana
username: mohsinzaheer25
```

#### Getting Secret

Khizana will return the secret based on `key`.

```cassandraql
$ khizana get username
Enter Password For Your Khizana

 mohsinzaheer25
```

#### Updating Secret

You can simply update your secret by using `update` command and passing `key` & `value`.

```cassandraql
$ khizana update username mohsinzaheer
Enter Password For Your Khizana

username value updated

```

You can check if `key` got updated or not using `view` command.

```cassandraql
$ khizana view
Enter Password For Your Khizana

# Khizana
username: mohsinzaheer
```

#### Deleting Secret

You can delete a `key` by using `delete` commanding and passing `key` to it. It will ask if you really want to delete or not and if it is `yes` then it will delete the key.

```cassandraql
$ khizana delete username
Enter Password For Your Khizana
Use the arrow keys to navigate: ↓ ↑ → ←
? Are you sure you want to delete:
  ▸ Yes
    No

✔ Yes

username key deleted
```

You can always check with `view` command for changes.

```cassandraql
$ khizana view
Enter Password For Your Khizana

# Khizana
```

#### Destroy

If you want to destroy the vault then you can use `destroy` command to destroy it.

\*\* Note: Make a backup of your secret, once it is destroyed it can't be recovered.

```cassandraql
Enter Password For Your Khizana
Use the arrow keys to navigate: ↓ ↑ → ←
? Are you sure you want to destroy Khizana:
  ▸ Yes
    No

✔ Yes
Khizana destroyed
```

## Download

This tool is available for Linux and Mac operating systems and can be download using below urls.

[Linux](https://github.com/mohsinzaheer25/khizana/releases/download/1.0/khizana-linux-amd64)

[Mac](https://github.com/mohsinzaheer25/khizana/releases/download/1.0/khizana-mac-amd64)

## Contribute

Contributions are more than welcome, if you are interested please send an email to mohsinzaheer25@hotmail.com until contribution guidelines get ready.
