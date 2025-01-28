## nodekit catchup

Manage Fast-Catchup for your node

### Synopsis

                                                                                                 
<img alt="Terminal Render" src="/assets/nodekit.png" width="65%">                                          
                                                                                                 
                                                                                                 
Fast-Catchup is a feature that allows your node to catch up to the network faster than normal.   
                                                                                                 
Overview:                                                                                        
The entire process should sync a node in minutes rather than hours or days.                      
Actual sync times may vary depending on the number of accounts, number of blocks and the network.
                                                                                                 
Note: Not all networks support Fast-Catchup.                                                     

```
nodekit catchup [flags]
```

### Options

```
  -d, --datadir string   Data directory for the node
  -h, --help             help for catchup
```

### SEE ALSO

* [nodekit](/README.md)	 - Manage Algorand nodes from the command line
* [nodekit catchup debug](/man/nodekit_catchup_debug.md)	 - Display debug information for Fast-Catchup.
* [nodekit catchup start](/man/nodekit_catchup_start.md)	 - Get the latest catchpoint and start catching up.
* [nodekit catchup stop](/man/nodekit_catchup_stop.md)	 - Stop a fast catchup

