# Roadmap

## Project Vision

Cruise is a terminal-first, extensible container management tool designed to empower developers and operators with unified, real-time control and insight over container workflows. 

We aim to bridge the gaps left by standard CLIs and GUIs, Cruise democratizes cloud-native and local container operations for everyone, anywhere. By breaking the “one CLI at a time” barrier by enabling powerful batch operations, visual dashboards, actionable monitoring, and seamless remote management creating a perfect management tool.

## Guiding Principles

- **Terminal-First Accessibility:**  
  Deliver powerful management and monitoring from within any terminal, empowering a wide range of users.

- **Unified Control Surface:**  
  Aggregate lifecycle, monitoring, and troubleshooting of containers, images, networks, and volumes in a single TUI.

- **Real-Time Visibility:**  
  Provide live dashboards, resource usage, and log insights to reduce manual polling and speed troubleshooting.

- **Workflow Acceleration:**  
  Enable batch operations, macros, and dashboard-driven actions to reduce repetitive CLI work.

- **Extensibility and Pluggability:**  
  Support plugins and multiple container runtimes—adapting to future community needs and environments.

- **Observability and Diagnostics:**  
  Aggregate logs and metrics with rich filtering and export to speed up problem solving.

- **Low-Footprint, Easy Adoption:**  
  Run with minimal dependencies and configuration for instant use anywhere.

- **Open Collaboration:**  
  Community-driven roadmap, transparent governance, and feedback loops as core values.

- **Simplicity for Empowerment:**  
  Lower onboarding friction for beginners while offering depth and efficiency for experts.


## Features

### Ongoing
  
#### Images
- [ ] Image repository browser

#### Docker Compose (v2)

##### Dashboard
- [ ] List all Compose projects
- [ ] Up/down/restart/recreate operations
- [ ] Visualize Compose service dependencies
- [ ] Manage environment variables

##### Service Dashboard
- [ ] Start/stop/restart/scale services
- [ ] Real-time service monitoring
- [ ] Network visualization
- [ ] Aggregated service logs with filters

#####  Compose Editor
- [ ] Built-in editor with nvim or fallback
- [ ] Syntax highlighting and error detection
- [ ] Git integration for version control

####  Build & Registry
- [ ] Manage build contexts
- [ ] Edit Dockerfiles with syntax support
- [ ] Configure private registries
- [ ] Manage and clean build cache

#### Monitoring & Logs
- [ ] Configure alerts and notifications
- [ ] Export metrics and logs
  
### Future Plans

- [ ] Remote/multi-host management (TUI across sockets/SSH)
- [ ] Snapshotting and audit history
- [ ] Workflow macros/automation
- [ ] Additional container engines support (Podman, containerd, ...)

## Get Involved!

Cruise is community-driven! Suggest features, propose improvements, or join our discussions to shape the future of container management from the terminal.
