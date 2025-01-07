## algorun

Manage Algorand nodes from the command line

### Synopsis

                                                                                                    
<img alt="Terminal Render" src="/assets/nodekit.png" width="65%">                                             
                                                                                                    
                                                                                                    
Manage Algorand nodes from the command line                                                         
                                                                                                    
Overview:                                                                                           
Welcome to Algorun, a TUI for managing Algorand nodes.                                              
A one stop shop for managing Algorand nodes, including node creation, configuration, and management.
                                                                                                    
Note: This is still a work in progress. Expect bugs and rough edges.                                

```
algorun [flags]
```

### Options

```
  -d, --datadir string   Data directory for the node
  -h, --help             help for algorun
```

### SEE ALSO

* [algorun bootstrap](/man/algorun_bootstrap.md)	 - Initialize a fresh node
* [algorun catchup](/man/algorun_catchup.md)	 - Manage Fast-Catchup for your node
* [algorun configure](/man/algorun_configure.md)	 - Change settings on the system (WIP)
* [algorun debug](/man/algorun_debug.md)	 - Display debugging information
* [algorun install](/man/algorun_install.md)	 - Install the node daemon
* [algorun start](/man/algorun_start.md)	 - Start the node daemon
* [algorun stop](/man/algorun_stop.md)	 - Stop the node daemon
* [algorun uninstall](/man/algorun_uninstall.md)	 - Uninstall the node daemon
* [algorun upgrade](/man/algorun_upgrade.md)	 - Upgrade the node daemon

###### Auto generated by spf13/cobra on 6-Jan-2025

### Installing

Connect to your server and run the installation script which will bootstrap your node.

```bash
curl -fsSL https://nodekit.run/install.sh | bash
```
