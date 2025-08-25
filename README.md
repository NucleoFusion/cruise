# Cruise – Docker TUI Client

> Terminal UI for managing Docker with style and speed.

**Cruise** is a powerful, intuitive, and fully-featured TUI (Terminal User Interface) for interacting with Docker. Built with Go and [Bubbletea](https://github.com/charmbracelet/bubbletea), it offers a visually rich, keyboard-first experience for managing containers, images, volumes, networks, logs and more — all from your terminal.


## Description

Ever felt that docker CLI is too lengthy or limited? Find yourself executing commands again and again for stats? Or wrote a full multiline command just for a typo to ruin it? Well... Fret no more. Cruise - Is a TUI Docker Client, fitting easily in your terminal-first dev workflow, while making repetitive Docker work easy and fun.

> How is _cruise_ different from existing solutions?

Existing applications are limited in what they do, they serve as mostly a monitoring service, _not_ a management service let alone a Client.

With Cruise you can:
- Manage Lifecycles of Containers, Images, Volumes, Networks.
- Have a centralized Monitoring service
- Scan images for vulnerabilities
- Get Detailed view on Docker Artifacts
- and more to come!

### 🚧 Tech Stack

- **Go** – High performance, robust concurrency
- **Bubbletea** – Elegant terminal UI framework
- **Charm ecosystem** – [Lipgloss](https://github.com/charmbracelet/lipgloss), [Bubbles](https://github.com/charmbracelet/bubbles), [Glamour](https://github.com/charmbracelet/glamour)
- **Docker SDK for Go** – Deep Docker integration
- **Trivy / Grype** – Vulnerability scanning
- **Viper** – Configuration management


## Usage

<details>
  <summary>Screenshots</summary>
</details>

Once [installed](#installation). You can run the app normally.

```
cruise
```

## Installation

Coming soon...


## Contributing

Please check out [CONTRIBUTING.md](CONTRIBUTING.md) for more.


## License

MIT License – see [LICENSE](LICENSE) for details.

## Credits

Built by [Nucleo](https://github.com/NucleoFusion).

Inspiration and Advice from [SourcewareLab](https://github.com/SourcewareLab).

Special Thanks to [Hegedus Mark](https://github.com/hegedus-mark) and [Mongy](https://github.com/A-Cer23).

## TODO Features

<details>
  
## Images
- ~[ ] Image repository browser~ (v2)

## Docker Compose (v2)

### Dashboard
- [ ] List all Compose projects
- [ ] Up/down/restart/recreate operations
- [ ] Visualize Compose service dependencies
- [ ] Manage environment variables

### Service Dashboard
- [ ] Start/stop/restart/scale services
- [ ] Real-time service monitoring
- [ ] Network visualization
- [ ] Aggregated service logs with filters

### Compose Editor
- [ ] Built-in editor with nvim or fallback
- [ ] Syntax highlighting and error detection
- [ ] Git integration for version control

## Build & Registry
- [ ] Manage build contexts
- [ ] Edit Dockerfiles with syntax support
- [ ] Configure private registries
- [ ] Manage and clean build cache

## Monitoring & Logs
- [ ] Configure alerts and notifications
- [ ] Export metrics and logs

</details>


## V1 Roadmap

<details>
  
### Vulnerability 
- [ ] Export (When config done, define a export folder) 

### Monitoring & Logs
- [ ] logs export (When config done)

## Misc
- [ ] Docs
- [ ] Mouse Support

</details>
