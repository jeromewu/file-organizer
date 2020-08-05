# file-organizer

Organize files within specific folder to yr-mo folders

> Only works in Linux

## Usage

```
$ ./file-organizer -i <INPUT_FOLDER> [OPTION]
```

Options:
- `-i <INPUT_FOLDER>` (required): input folder, subfolder will be scanned as well
- `-o <OUTPUT_FOLDER>`: output folder, default is **out**
- `-m`: without this flag, the files will be copied, and moved if this flag is on, default is off

## Examples

**Organize files inside Pictures and copy files to out/**

```
$ ./file-organizer -i Pictures
```

**Organize files inside Pictures and copy files to Picture-out/**

```
$ ./file-organizer -i Pictures -o Picture-out
```

**Organize files inside Pictures and move files to Picture-out/**

```
$ ./file-organizer -i Pictures -o Picture-out -m
```
