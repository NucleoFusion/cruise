# Configuration

Cruise supports a flexible configuration file for customizing themes, keybinds etc. We use a TOML format and the default config file is at: 


- Linux/MacOS:
```
~/.config/cruise/config.toml
```

- Windows:
```
C:\Users\<username>\.cruise\config.toml
```

## File Format

Cruise uses TOML for its configuration file. Making it readable and easy to edit.

An example minimal file with vim motions (hjkl motions)
```toml
[keybinds.global]
up = "k"
down = "j"
```

## Key Sections

the config file has 3 main sections:

- **General**: for general stuff such as shell, export_dir etc
- **Styles**: for customizing styles and colors
- **Keybinds**: for customizing keybinds 
