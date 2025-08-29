# Installation

## Archives

Latest archives (zip/tar.gz files) are available on the github [releases](https://github.com/NucleoFusion/cruise/releases) page.
Just download the file for your system and run the following commands to extract it.

### Linux / MacOS 

```bash
tar -xvf ./path/to/file.tar.gz -C /extract/in/directory/
```

### Windows

```bash
# PowerShell (built-in)
Expand-Archive -Path path\to\file.zip  -DestinationPath "C:\extract\to\directory"

# Command Prompt (if using tar/bsdtar)
tar -xf path\to\file.zip  -C "C:\extract\to\directory"

```

## Debian / RPM

### Debian / Ubuntu (.deb) 

Install the .deb file [_here_](https://github.com/NucleoFusion/cruise/releases).

```bash
sudo apt install ./path/to/file.deb
```
> _maintained by [NucleoFusion](https://github.com/NucleoFusion)_

### Fedora / RHEL (.dnf) 

Install the .rpm file [_here_](https://github.com/NucleoFusion/cruise/releases).

```bash
sudo dnf install ./path/to/file.dnf
```
> _maintained by [NucleoFusion](https://github.com/NucleoFusion)_

## Homebrew

_Cruise_ can be installed via Homebrew using the,

```bash
brew tap NucleoFusion/homebrew-tap
brew install --cask cruise
```


## AUR

Arch users can install _cruise_ through the AUR. using any AUR Helper.

```bash
yay -S cruise
```

## Chocolatey

_Cruise_ is also available via Chocolatey.

```bash
choco install cruise
```
