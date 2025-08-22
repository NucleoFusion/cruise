# 🚢 Cruise – A Docker TUI Client

> Terminal UI for managing Docker and Docker Compose projects with style and speed.

**Cruise** is a powerful, intuitive, and fully-featured TUI (Terminal User Interface) for interacting with Docker. Built with Go and [Bubbletea](https://github.com/charmbracelet/bubbletea), it offers a visually rich, mouse-less experience for managing containers, images, volumes, networks, Compose stacks, and more — all from your terminal.

---


## 🚧 Tech Stack

- **Go** – High performance, robust concurrency
- **Bubbletea** – Elegant terminal UI framework
- **Charm ecosystem** – [Lipgloss](https://github.com/charmbracelet/lipgloss), [Bubbles](https://github.com/charmbracelet/bubbles), [Glamour](https://github.com/charmbracelet/glamour)
- **Docker SDK for Go** – Deep Docker integration
- **Trivy / Grype** – Vulnerability scanning
- **YAML parser** – Safe Compose editing
- Optional: **nvim** or fallback editor

---

## ✨ Features

## Dashboard
- [X] Display system-wide resource usage (CPU, memory, disk)
- [X] Show quick stats (total containers, images, volumes, networks)
- [X] Track recent Docker activity
- [X] Display Docker daemon status

## Containers
- [X] Search and filter containers
- [X] Real-time monitoring (CPU%, memory, network IO)
- [X] Start/stop/restart/remove/pause containers
- [X] Exec into containers (`docker exec -it`)
- [X] Stream logs with search and filter
- [X] Visualize port mappings
- [ ] Container Details popup (realtime)
- ~[ ] Group containers by Compose project~

## Images
- [X] Pull/push/build/remove images
- ~[ ] Image repository browser~
- [X] Vulnerability scanning (Trivy/Grype)
- [X] Tag management and cleanup
- [X] Size analysis and optimization

## Docker Compose

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

## Volumes & Networks
- [X] Volume management with usage stats
- [X] Volume backup and restore
- [X] Cleanup unused volumes and networks
- [X] Visualize network topology

## Build & Registry
- [ ] Manage build contexts
- [ ] Edit Dockerfiles with syntax support
- [ ] Configure private registries
- [ ] Manage and clean build cache

## Monitoring & Logs
- [X] Centralized log viewer with search
- [X] Real-time metrics dashboard
- [ ] Configure alerts and notifications
- [ ] Export metrics and logs


---

## V1 Roadmap

### Containers 
- [X] Port mapping visualization
- [ ] Remove Container Dropdown, instead make a VP, that displays all details for that container. (Graphs)

### Vulnerability 
- [X] Vulnerability scanning (Trivy/Grype)
- [X] Individual 
- [X] Toggle b/w trivy and grype
- [ ] Export (When config done, define a export folder) 

### Volumes & Networks
- [X] Volume management with usage stats & backup/restore & prune
- [X] Network topology visualization

### Monitoring & Logs
- [X] Centralized log aggregation with search
- [ ] Alert configuration (notifications)
- [ ] logs export (When config done)

## Misc
- [ ] Config File
- [ ] Docs
- [ ] UI Polish (use lipgloss where possible)
- [ ] Mouse Support

---

## 📦 Installation

Coming soon...

---

## 💬 Contributing

Contributions, feedback, and feature requests are welcome!

1. Clone the repo
2. Run with `go run .` (or build via `go build`)
3. Hack away 🚀

Please check out [CONTRIBUTING.md](CONTRIBUTING.md) for more.

---

## 📄 License

MIT License – see [LICENSE](LICENSE) for details.

---

## 🧑‍💻 Author

Built with ❤️ by [Nucleo](https://github.com/NucleoFusion)

---

