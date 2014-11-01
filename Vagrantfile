# -*- mode: ruby -*-
# vi: set ft=ruby :

# bootstrap script partially copied from https://github.com/Lukx/vagrant-lamp
$script = <<SCRIPT
    apt-get -y install git mc golang
    apt-get -y install mercurial meld

    export GOPATH=~/
    echo export GOPATH=$GOPATH >> ~/.bash_profile
    export PATH="$PATH:$GOPATH/bin"
    echo 'export PATH="$PATH:$GOPATH/bin"' >> ~/.bash_profile

    go get github.com/revel/cmd/revel
SCRIPT

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

    config.vm.box = "ubuntu/trusty64"
    config.vm.network :forwarded_port, host: 6666, guest: 80

    config.vm.synced_folder "./", "/home/vagrant"

    #config.vm.provider "virtualbox" do |v|
    #  v.name = "hackaton_vm" # PROJECT NAME PUT HERE
    #end

    # provisioner config
    config.vm.provision "shell", inline: $script
end