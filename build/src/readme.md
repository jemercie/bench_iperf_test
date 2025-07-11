## On utilise vagrant en fait

- compiler dehors avec arch et mount bin srv and cli
- provision en installant iperf3
- run l'un puis l'autre as daemon
- trouver comment sortir un file avec les logs
et boum! (cv encore etre plein de rebondissement mais bon)






### TODO:
2. setup VMs 4
3. setup Network 2
4. faire marcher cli et serv sur leurs vms respectives 2
5. faire un output clean 1
6. intégrer relay/gateway 5
-> ca a l'air complique et pas reproducible

1. nix shell
2. define les vms et faire des calls exterieurs avec dagger
3. define des vms pour mac et linux et get l'env dans 


### dans dagger fonctions pour:
1. generate xml and define vms with `virt-install` & `--print-xml`-> retreive l'os de l'env ou quoi pour que ce soit adapté
2. define un network et le lancer
3. lancer les deux vms

[man virt-install](https://linux.die.net/man/1/virt-install)

virt-install
--name=srv

- 4G par vm
-> lancer le bordel et faire des htop pour voir

### define un network
1. avec un xml ou une cmd comme au dessus
2. libvirt network.Create()?? [libvirt doc explaining how to precise the network xml definition]

[documentation](https://wiki.libvirt.org/VirtualNetworking.html)


- `virt-builder` [doc](https://libguestfs.org/virt-builder.1.html) to build a custom image

https://docs.gitlab.com/runner/executors/custom_examples/libvirt/

## cmds
- `nix develop`

[article: executer un truc au demarrage de 4 façons](https://www.malekal.com/linux-executer-script-commande-demarrage/)

## [issues with vagrant and apple](https://github.com/vagrant-libvirt/vagrant-libvirt/issues/1205)

