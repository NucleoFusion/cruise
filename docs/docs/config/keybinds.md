# Keybinds

This section covers all keybinds for _Cruise_. 

## Global Keybinds
Denoted by `[keybinds.global]` toml tag.

### Page Finder
Sets the keybind for opening popup to traverse pages. Can be set using the `page_finder` key. It is global so a general keyboard english character is not recommended.

Default Value:
```toml
page_finder = "tab"
```

### Up
Sets the keybind for moving selection up in all list view. Can be set by using the `list_up` key.

Default Value:
```toml
list_up = "up"
```

### Down
Sets the keybind for moving selection down in all list view. Can be set by using the `list_down` key.

Default Value:
```toml
list_down = "down"
```

### Focus Search
Sets the keybinds for focusing on the search input field. Can be set using the `focus_search` key.

Default Value:
```toml
focus_search = "/"
```


### Unfocus Search
Sets the keybinds for exiting search mode, ie, unfocus on the search input field. Can be set using the `unfocus_search` key.

Default Value:
```toml
unfocus_search = "esc"
```


### Quick Quit 
Sets the keybind for quitting application. Can be set using the `quick_quit` key. 

Default Value:
```toml
quick_quit = "q"
```

## Container 
Denoted by `[keybinds.container]` toml tag.

### Start
Sets the keybind for starting a stopped container. Can be set using the `start` key.

Default Value:
```toml
start = "s"
```

### Stop
Sets the keybind for stopping a container. Can be set using the `stop` key.

Default Value:
``` toml
stop = "t"
```

### Remove
Sets the keybind for removing a container. Can be set using the `remove` key.

Default Value:
``` toml
remove = "d"
```

### Restart
Sets the keybind for restarting a container. Can be set using the `restart` key.

Default Value:
``` toml
restart = "r"
```

### Pause
Sets the keybind for pausing a container. Can be set using the `pause` key.

Default Value:
``` toml
pause = "p"
```

### Unpause
Sets the keybind for unpausing a container. Can be set using the `unpause` key.

Default Value:
``` toml
unpause = "u"
```

### Execute
Sets the keybind for obtaining a shell instance inside a container. Can be set using the `exec` key.

Default Value:
``` toml
exec = "e"
```

### Show Details
Sets the keybind for showing detailed information about a container. Can be set using the `show_details` key.

Default Value:
``` toml
show_details = "enter"
```

### Exit Details
Sets the keybind for exiting the container details view. Can be set using the `exit_details` key.

Default Value:
``` toml
exit_details = "esc"
```

### Port Map
Sets the keybind for viewing port mappings of a container. Can be set using the `port_map` key.

Default Value:
``` toml
port_map = "m"
```

## Images 
Denoted by `[keybinds.images]` toml tag.

### Remove
Sets the keybind for removing an image. Can be set using the `remove` key.

Default Value:
``` toml
remove = "r"
```

### Prune
Sets the keybind for pruning unused images. Can be set using the `prune` key.

Default Value:
``` toml
prune = "d"
```

### Push
Sets the keybind for pushing an image to a registry. Can be set using the `push` key.

Default Value:
``` toml
push = "p"
```


### Sync 

Sets the keybind for refresh/sync images list. Can be set using the `sync` key.

Default Value:
``` toml
sync = "s"
```

## Fuzzy Finder(fzf) 
Denoted by `[keybinds.fzf]` toml tag.

### Up
Sets the keybind for moving up in the fuzzy finder results. Can be set using the `up` key.

Default Value:
``` toml
up = "up"
```

### Down
Sets the keybind for moving down in the fuzzy finder results. Can be set using the `down` key.

Default Value:
``` toml
down = "down"
```

### Enter
Sets the keybind for selecting the highlighted fuzzy finder item. Can be set using the `enter` key.

Default Value:
``` toml
enter = "enter"
```

### Exit
Sets the keybind for exiting the fuzzy finder. Can be set using the `exit` key.

Default Value:
``` toml
exit = "esc"
```

## Monitoring 
Denoted by `[keybinds.monitoring]` toml tag.

### Search
Sets the keybind for searching within monitoring data. Can be set using the `search` key.

Default Value:
``` toml
search = "/"
```

### Exit Search
Sets the keybind for exiting search mode. Can be set using the `exit_search` key.

Default Value:
``` toml
exit_search = "esc"
```


### Export
Sets the keybind for exporting to the export dir. Can be set using the `export` key.

Default Value:
``` toml
export= "e"
```

## Networks 
Denoted by `[keybinds.network]` toml tag.

### Remove
Sets the keybind for removing a network. Can be set using the `remove` key.

Default Value:
``` toml
remove = "r"
```

### Prune
Sets the keybind for pruning unused networks. Can be set using the `prune` key.

Default Value:
``` toml
prune = "p"
```

### Show Details
Sets the keybind for showing detailed information about a network. Can be set using the `show_details` key.

Default Value:
``` toml
show_details = "enter"
```

### Exit Details
Sets the keybind for exiting the detailed network view. Can be set using the `exit_details` key.

Default Value:
``` toml
exit_details = "esc"
```

## Volume
Denoted by `[keybinds.volume]` toml tag.

### Remove
Sets the keybind for removing a volume. Can be set using the `remove` key.

Default Value:
``` toml
remove = "r"
```

### Prune
Sets the keybind for pruning unused volumes. Can be set using the `prune` key.

Default Value:
``` toml
prune = "p"
```

### Show Details
Sets the keybind for showing detailed information about a volume. Can be set using the `show_details` key.

Default Value:
``` toml
show_details = "enter"
```

### Exit Details
Sets the keybind for exiting the detailed volume view. Can be set using the `exit_details` key.

Default Value:
``` toml
exit_details = "esc"
```

## Vulnerability
Denoted by `[keybinds.vulnerability]` toml tag.

### Focus Scanners
Sets the keybind for switching focus to the vulnerability scanners panel. Can be set using the `focus_scanner` key.

Default Value:
``` toml
focus_scanner = "S"
```

### Focus List
Sets the keybind for switching focus to the vulnerability results list. Can be set using the `focus_list` key.

Default Value:
``` toml
focus_list = "L"
```

### Export
Sets the keybind for exporting to the export dir. Can be set using the `export` key.

Default Value:
``` toml
export= "e"
```

