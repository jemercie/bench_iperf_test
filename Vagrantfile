Vagrant.configure("2") do |config|
  config.vm.box = "generic/ubuntu2004"
  config.vm.provider "libvirt" do |libvirt|
    libvirt.memory = 1024
    libvirt.cpus = 1
  end
  config.vm.define "vm1kk" do |vmmm1|
    vmmm1.vm.network "private_network", ip: "192.168.33.10"
    vmmm1.vm.synced_folder "./src/srv/bin/", "/srv"
    vmmm1.vm.synced_folder "./src/srv/log/", "/log"
  end
  config.vm.define "vm2kk" do |vmmm2|
    vmmm2.vm.network "private_network", ip: "192.168.33.11"
    vmmm2.vm.synced_folder "./src/cli/bin/", "/cli"
    vmmm2.vm.synced_folder "./src/cli/log/", "/log"
  end
end


