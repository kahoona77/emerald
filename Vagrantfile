# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

# Our custom installation routine
$script = <<SCRIPT
set -x
echo Get the base system up to date
sudo apt-get update && sudo apt-get -y upgrade && sudo apt-get autoclean -y && sudo apt-get autoremove -y
echo Fix the locale
export LC_ALL="en_US.UTF-8"
export LC_TYPE="UTF-8"
sudo dpkg-reconfigure locales
if [ $(dpkg-query -W -f='${Status}' mongodb-org 2>/dev/null | grep -c "ok installed") -eq 0 ];
then
  echo Install MongoDB
  sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10
  echo "deb http://repo.mongodb.org/apt/ubuntu "$(lsb_release -sc)"/mongodb-org/3.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.0.list
  sudo apt-get update
  sudo apt-get install -y mongodb-org
  sudo apt-get install -y mongodb-org=3.0.5 mongodb-org-server=3.0.5 mongodb-org-shell=3.0.5 mongodb-org-mongos=3.0.5 mongodb-org-tools=3.0.5
  sudo service mongod start
fi
echo All done...
SCRIPT



Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  # Every Vagrant virtual environment requires a box to build off of.
  config.vm.box = "ubuntu/trusty64"

  # Forward the MongoDB default port
  config.vm.network "forwarded_port", guest: 27017, host: 27017

  # Forward additional MongoDB ports for replication and sharding
  config.vm.network "forwarded_port", guest: 27018, host: 27018
  config.vm.network "forwarded_port", guest: 27019, host: 27019
  config.vm.network "forwarded_port", guest: 27020, host: 27020
  config.vm.network "forwarded_port", guest: 27021, host: 27021


  # Install our dependencies
  config.vm.provision "shell", inline: $script

end
