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
- [ ] Display system-wide resource usage (CPU, memory, disk)
- [ ] Show quick stats (total containers, images, volumes, networks)
- [ ] Track recent Docker activity
- [ ] Display Docker daemon status

## Containers
- [ ] Search and filter containers
- [ ] Real-time monitoring (CPU%, memory, network IO)
- [ ] Start/stop/restart/remove/pause containers
- [ ] Exec into containers (`docker exec -it`)
- [ ] Stream logs with search and filter
- [ ] Visualize port mappings
- [ ] Group containers by Compose project

## Images
- [ ] Pull/push/build/remove images
- [ ] Image repository browser
- [ ] Vulnerability scanning (Trivy/Grype)
- [ ] Tag management and cleanup
- [ ] Size analysis and optimization

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
- [ ] Volume management with usage stats
- [ ] Volume backup and restore
- [ ] Cleanup unused volumes and networks
- [ ] Visualize network topology
- [ ] Distinguish bind mounts vs. named volumes
- [ ] Visualize network security policies

## Build & Registry
- [ ] Manage build contexts
- [ ] Edit Dockerfiles with syntax support
- [ ] Configure private registries
- [ ] Manage and clean build cache

## Monitoring & Logs
- [ ] Centralized log viewer with search
- [ ] Real-time metrics dashboard
- [ ] Configure alerts and notifications
- [ ] Export metrics and logs


---

## 🗺️ Roadmap

- [ ] MVP: Container and Image management
- [ ] Docker Compose project control
- [ ] Volume/network tooling
- [ ] Compose editor with Git
- [ ] Full metrics/logs dashboard
- [ ] Registry integration & vulnerability scan

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

