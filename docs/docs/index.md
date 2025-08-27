# Introduction
- Linux: 
```
~/.config/cruise/config.toml
```

- Windows:
```
C:\Users\<username>\AppData\Roaming\cruise\config.yml
```

- Mac:
```
~/Library/Application Support/cruise/config.yml
```

# File Format
Cruise uses TOML for its configuration file. Making it readable and easy to edit.

## Global Configuration
It describes about the globl configuration of Cruise.

### Export Directory
Exports directory path where exported files and data is saved. Can be set using `export_dir` key.

### Shell
It is the default shell to use for command execution.

## Keybind Configuration
This document describes all available keybindings in Cruise, organized by functional sections.

## Global Keybinds
Denoted by [keybinds.global].

### Page Finder
Used to set keybind for opening the page finder model. Can be set using the `page_finder` key.

Default Value:
```
page_finder = "tab"
```

### Up
Moves selection up in any list view. Can be set by using the `list_up` key.

Default Value:
```
list_up = "up"
```

### Down
Moves selection down in any list view. Can be set by using the `list_down` key.

Default Value:
```
list_down = "down"
```

### Focus Search
Focuses on the search input field. Can be set using the `focus_search` key.

Default Value:
```
focus_search = "/"
```


### Unfocus Search
Exits search mode and returns focus to main content. Can be set using the `unfocus_search` key.

Default Value:
```
unfocus_search = "esc"
```

## Container Keybinds
Denoted by [keybinds.container]

### Start
Starts a stopped container. Can be set using the `start` key.

Default Value:
```
start = "s"
```

### Stop
Stops a started container. Can be set using the `stop` key .

Default Value:
```
stop = "t"
```
### Remove
Deletes the selected container. Can be set using the `remove` key.

Default Value:
```
remove = "d"
```

### Restart
Restarts the selected container. Can be set using the `restart` key.

Default Value:
```
restart = "r"
```

### Pause
Pauses the running container. Can be set using the `pause` key.

Default Value:
```
pause = "p"
```

### Unpause
Resumes the paused container. Can be set using the `unpause` key.

Default Value:
```
unpause = "u"
```

### Execute
Executes a command inside the selected container.  Can be set using the `exec` key.

Default Value:
```
exec = "e"
```

### Show Details
Displays the detailed information about the container. Can be set using the `show_details` key.

Default Value:
```
show_details = "enter"
```

### Exit Details
Closes the container details view. Can be set using the `exit_details` key.

Default Value:
```
exit_details = "esc"
```

### Port Map
View port mappings for the selected container. Can be set using the `port_map` key.

Default Value:
```
port_map = "m"
```

## Images Keybinds
Denoted by [keybinds.images].

### Remove
Deletes the selected image. Can be set by using the `remove` key.

Default Value:
```
remove = "r"
```

### Prune
Prunes all the unused images. Can be set by using the `prune` key.

Default Value:
```
prune = "d"
```

### Push
Pushes the selected image to a registry. Can be set by using the `push` key.

Default Value:
```
push = "p"
```

## Fuzzy Finder(fzf) Keybinds
Denoted by [keybinds.fzf]

### Up
Moves up the fuzzy finder results. Can be set using `up` key.

Default Value:
```
up = "up"
```

### Down
Moves down the fuzzy finder results. Can be set using the `down` key.

Default Value:
```
down = "down"
```

### Enter
Selects the highlighted item. Can be set by using the `enter` key.

Default Value:
```
enter = "enter"
```

### Exit
Exits the fuzzy finder. Can be set using the `exit` key.

Default Value:
```
exit = "exit"
```

## Monitoring Keybinds
Denoted by [keybinds.monitorkeybinds]

### Search
Searches withing monitoring data. Can be set using the `search` key.

Default Value:
```
search = "/"
```

### Exit Search
Exits search mode in monitoring view. Can be set using the `exit_search` key.

Default Value:
```
exit_search = "esc"
```

## Networking Keybinds
Denoted by [keybinds.network]

### Remove
Removes the selected network. Can be set using the `remove` key.

Default Value:
```
remove = "r"
```

### Prune
Prunes all unused networks. Can be set using the `prune` key.

Default Value:
```
prune = "p"
```

### Show Details
Shows the detailed network information. Can be set using the `show_details` key.

Default Value:
```
show_details = "enter"
```

### Exit Details
Exits the detailed network information. Can be set using the `exit_details` key.

Default Value:
```
exit_details = "esc"
```

## Volume Keybinds
Denoted by [keybinds.volume]

### Remove
Removes the selected volume. Can be set using the `remove` key

Default Value:
```
remove = "r"
```
### Prune
Prunes all unused volumes. Can be set using the `prune` key.

Default Value:
```
prune = "p"
```

### Show Details
Shows the detailed volume information. Can be set using the `show_details` key.

Default Value:
```
show_details = "enter"
```

### Exit Details
Exits the detailed volume information. Can be set using the `exit_details` key.

Default Value:
```
exit_details = "esc"
```

## Vulnerability Keybinds
Denoted by [keyinds.vulnerbility]

### Focus Scanners
Switches focus to the vulnerability scanners panel. Can be set using the `foocus_scanner` key.

Default Value:
```
focus_scanner = "S"
```

### Focus List
Switches focus to the vulnerability results list. Can be set using the `focus_list` key.

Default Value:
```
focus_list = "L"
```

## Styles Keybinds
Denoted by [keybinds.styles].

### Text
It is the primary text color used throughout the interface. Can be set using `text` key.

Default Value:
```
text = "#cdd6f4"
```

### Sub Title Text
It is the color for subtitle text elements. Can be set using the `subtitle_text` key.

Default Value:
```
subtitle_text = "#74c7ec"
```

### Subtitle Background
It is the background color for subtitle sections. Can be set using the `subtitle_bg` key.

Default Value:
```
subtitle_bg = "#313244"
```

### Unfocused Border
It is the Border color for unfocused UI elements. Can be set using the `unfocused_border` key.

Default Value:
```
unfocused_border = "#45475a"
```

### Unfocused Border
It is the Border color for focused UI elements. Can be set using the `focused_border` key.

Default Value:
```
focused_border = "#b4befe"
```

### Help Key Background
It is the Background color for help key indicators. Can be set using the `help_key_bg` key.

Default Value:
```
help_key_bg = "#313244"
```

### Help Key Text
It is the Text color for help key indicators. Can be set using the `help_key_text` key.

Default Value:
```
help_key_bg = "#cdd6f4"
```

### Help Descriptions Text
It is the Text color for help descriptions. Can be set using the `help_desc_text` key.

Default Value:
```
help_desc_bg = "#6c7086"
```

### Menu Selected Background
It is the Background color for selected menu items. Can be set using the `menu_selected_bg` key.

Default Value:
```
menu_selected_bg = "#b4befe"
```

### Menu Selected Text
It is the Text color for selected menu items. Can be set using the `menu_selected_text` key.

Default Value:
```
menu_selected_text = "#1e1e2e"
```

### Error Text
It is the Text color for error messages. Can be set using the `error_text` key.

Default Value:
```
error_text = "#f38ba8"
```

### Error Background
It is the Background color for error messages. Can be set using the `error_bg` key.

Default Value:
```
error_bg = "#11111b"
```

### Popup Border
It is the Border color for popup windows and dialogs. Can be set using the `popup_border` key.

Default Value:
```
popup_border = "#74c7ec"
```

### Placeholder Text
It is the Color for placeholder text in input fields. Can be set using the `placeholder_text` key.

Default Value:
```
placeholder_text = "#585b70"
```

### Message Text
It is the Color for general message and notification text. Can be set using the `msg_text` key.

Default Value:
```
msg_text = "#74c7ec"
```