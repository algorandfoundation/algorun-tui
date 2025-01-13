---
title: Getting Started with NodeKit
description: Installing NodeKit
sidebar:
  order: 10
---

Welcome to NodeKit, your command-line one-stop-shop for Algorand node running.

NodeKit can help you with:

- Installing and configuring an Algorand node
- Syncing your node with the latest state of the network
- Managing consensus participation keys
- Monitoring your node and the network

To get started with NodeKit, copy-paste this command in your terminal:

```bash
curl -fsSL https://nodekit.run/install.sh | bash
```

<details><summary>Troubleshooting: Command 'curl' not found</summary>
If you get an error about the `curl` command not being found, you need to install the `curl` package.

On Ubuntu systems, you do this with:

```bash
sudo apt install -y curl
```
</details>
<details><summary>Troubleshooting: Command 'bash' not found</summary>
Some versions of Mac OS may not include the required `bash` executable that runs the installer.

If you get an error about `bash` not being available, please install bash on your system manually.

For Mac OS, a popular way to do this is to install [Homebrew](https://brew.sh/) and then install bash using:

```bash
brew install bash
```
</details>

This will detect your operating system and download the appropriate NodeKit executable to your local directory.

It will then immediately start the bootstrap process to get your Algorand node up and running:

![Screenshot of a typical NodeKit installation process](/assets/nodekit-install.png)

Otherwise, read on for guidance on the bootstrapping process.
