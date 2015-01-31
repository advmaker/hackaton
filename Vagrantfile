# -*- mode: ruby -*-
# vi: set ft=ruby :

# bootstrap script partially copied from https://github.com/Lukx/vagrant-lamp
$script = <<SCRIPT
    apt-get update
    apt-get -y install git mc golang
    apt-get -y install postgresql postgresql-contrib

    apt-get -y install nodejs npm
    ln -s /usr/bin/nodejs /usr/bin/node
    npm install bower -g
    npm install gulp -g

    mkdir -p /home/vagrant/src/github.com/advmaker/hackaton

    export GOPATH=/home/vagrant
    export PATH="$PATH:/home/vagrant/bin"
    export PORT=8000
    export PGHOST=/var/run/postgresql
    echo 'export GOPATH=/home/vagrant' >> /home/vagrant/.bash_profile
    echo 'export PATH="$PATH:/home/vagrant/bin"' >> /home/vagrant/.bash_profile

    cd /opt/
    git clone https://github.com/pote/gpm.git && cd /opt/gpm
    ./configure
    make install

    cd /opt/
    git clone https://github.com/pote/gvp.git && cd /opt/gvp
    ./configure
    make install

    chown -R vagrant:vagrant /home/vagrant

SCRIPT

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

    config.vm.box = "ubuntu/trusty64"
    config.vm.network "private_network", ip: "192.168.100.100"

    config.vm.synced_folder "./", "/home/vagrant/src/github.com/advmaker/hackaton"

    config.vm.provider "virtualbox" do |v|
      v.name = "hackaton_vm" # PROJECT NAME PUT HERE
    end

    # provisioner config
    config.vm.provision "shell", inline: $script
end