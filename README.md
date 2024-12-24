# ⌨️ NodeKit

<div align="center">
    <img alt="Terminal Render" src="/assets/Banner.gif" width="65%">
</div>

<div align="center">
    <a target="_blank" href="https://github.com/algorandfoundation/nodekit/actions/workflows/test.yaml">
        <img alt="CI Badge" src="https://github.com/algorandfoundation/nodekit/actions/workflows/test.yaml/badge.svg"/>
    </a>
    <a target="_blank" href="https://github.com/algorandfoundation/nodekit">
        <img alt="CD Badge" src="https://img.shields.io/badge/CD-TODO-red"/>
    </a>
    <a target="_blank" href="https://github.com/algorandfoundation/nodekit/stargazers">
        <img alt="Repository Stars Badge" src="https://img.shields.io/github/stars/algorandfoundation/nodekit?color=7B1E7A&logo=star&style=flat" />
    </a>
    <img alt="Repository Visitors Badge" src="https://api.visitorbadge.io/api/visitors?path=https%3A%2F%2Fgithub.com%2Falgorandfoundation%2Fnodekit&countColor=%237B1E7A&style=flat" />
</div>

---

Terminal UI for managing Algorand nodes.
Built with [bubbles](https://github.com/charmbracelet/bubbles) & [bubbletea](https://github.com/charmbracelet/bubbletea)

> [!CAUTION]
> This project is in alpha state and under heavy development. We do not recommend performing actions (e.g. key management) on participation nodes connected to public networks.

# 🚀 Get Started

Download the latest release by running

```bash
curl -fsSL https://raw.githubusercontent.com/algorandfoundation/nodekit/refs/heads/main/install.sh | bash
```

Follow the instructions 

# ℹ️ Advanced Usage

## 🧑‍💻 Commands

The default command will launch the full TUI application

```bash
./nodekit
```


### Help

Display the usage information for the command

```bash
./nodekit help
```

## ⚙️ Configuration

The TUI first looks for the environment for an `ALGORAND_DATA` variable.

The application supports the `datadir` flag for configuration.

```bash
./nodekit --datadir /var/lib/algorand
```



