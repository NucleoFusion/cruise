# Contributing

Thank you for considering a contribution! We welcome bug reports, feature requests, code, docs, and suggestions.

## Dependencies

- Go 
- Trivy
- Grype
- Docker

## Getting Started

1. Fork the repository.
2. Clone your fork.
  - ` git clone <your-fork-url>.git`
3. Create a branch for the respective issue.
  - `git checkout -b <branch-name> dev`
  - \<branch-name\> should follow name of feat/newpage, bug/#11 or docs/typo etc.
5. When Creating a PR, make sure to reference the respective issue.

## Issues
- Search open/closed issues before opening a new one.
- For bugs, include OS, Docker version, logs, and steps to reproduce.
- For features, describe the problem and proposed solution.

## Making Changes
- Create a feature/fix branch from `main`.
- Use clear, descriptive commit messages.
- Format your code with `gofmt`, and check with `golangci-lint`.
- Commit Message should follow format \<type\>:\<desc\>, where type is like feat, bug, enhancement, chore etc.

## Pull Requests
- Make PR to dev branch only!!
- Each PR should focus on one change.
- Link relevant issues in the PR.
- Respond to reviewer comments.

## Help & Feedback

Questions? Use GitHub Discussions or open an issue with the “question” label.

## Thanks!
All contributors are credited in the [Credits](README.md#credits) section. We appreciate your help!
