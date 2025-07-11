
[go compilation flags cheat sheet](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)

-> faire en sorte que ds le makefile tout le bazar soit installé, genre faire un script pour setup le bazar

### synced folders
> Synced folders enable Vagrant to sync a folder on the host machine to the guest machine, allowing you to continue working on your project's files on your host machine, but use the resources in the guest machine to compile or run your project.

```Ruby
    vmmm2.vm.synced_folder "./src/cli/bin/", "/bin/"
```

[vagrant synced folders](https://developer.hashicorp.com/vagrant/docs/synced-folders)


## On utilise vagrant en fait

- compiler dehors avec arch et mount bin srv and cli
- provision en installant iperf3
- run l'un puis l'autre as daemon
- trouver comment sortir un file avec les logs
et boum! (cv encore etre plein de rebondissement mais bon)
- eventuellement define un network au lieu du default

