Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"
  config.vm.network "private_network", ip: "10.20.0.5"
  
  config.vm.provider "virtualbox" do |vb|
    vb.customize [ "modifyvm", :id, "--uartmode1", "disconnected" ]
  end

  config.vm.provision "shell", inline: <<-SHELL
    # Update virtual machine
    yum update -y
    yum upgrade -y
    
    # Install epel-release
    yum install epel-release -y

    sed -i 's/PasswordAuthentication no/PasswordAuthentication yes/g' /etc/ssh/sshd_config    
    systemctl restart sshd.service
  SHELL
end
