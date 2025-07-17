#!/bin/bash

# Install Script for Subnetcalc

#----------------------------------------------------------------------

# Check user is root otherwise exit script

if [ "$EUID" -ne 0 ]
then
  printf "\nPlease run as root\n\n";
  exit;
fi;

cd /root;

#----------------------------------------------------------------------

# Check subnetcalc has been cloned from GitHub

if [ ! -d "/root/subnetcalc" ]
then
  printf "\nDirectory subnetcalc does not exist in /root.\n";
  printf "Please run commands: \"cd /root; git clone https://github.com/ellwould/subnetcalc\"\n";
  printf "and run install script again\n\n";
  exit;
fi;

#----------------------------------------------------------------------

# Copy unit files and restart systemd deamon

cp /root/subnetcalc/systemd/subnetcalc.service /usr/lib/systemd/system/;
cp /root/subnetcalc/systemd/subnetresult.service /usr/lib/systemd/system/;
systemctl daemon-reload;

#----------------------------------------------------------------------

# Install wget

apt update;
apt install wget;

#----------------------------------------------------------------------

# Remove any previous version of Go, download and install Go 1.24.5 

wget -P /root https://go.dev/dl/go1.24.5.linux-amd64.tar.gz;
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.5.linux-amd64.tar.gz;

# Create subnetcalc directories

mkdir -p /etc/subnetcalc/html-css;

# Copy Subnetcalc configuration file

cp /root/subnetcalc/env/subnetcalc.env /etc/subnetcalc/subnetcalc.env;

# Copy HTML/CSS start and end file

cp /root/subnetcalc/html-css/subnetcalc-start.html /etc/subnetcalc/html-css;
cp /root/subnetcalc/html-css/subnetcalc-end.html /etc/subnetcalc/html-css;

# Create Go directories in root home directory

mkdir -p /root/go/{bin,pkg,src/subnethome,src/subnetresult};

# Copy Go source code

cp /root/subnetcalc/go/subnethome.go /root/go/src/subnethome/subnethome.go;
cp /root/subnetcalc/go/subnetresult.go /root/go/src/subnetresult/subnetresult.go;

# Create Go mod for subnethome

export PATH=$PATH:/usr/local/go/bin;
cd /root/go/src/subnethome;
go mod init root/go/src/subnethome;
go mod tidy;

# Create Go mod for subnetresult

export PATH=$PATH:/usr/local/go/bin;
cd /root/go/src/subnetresult;
go mod init root/go/src/subnetresult;
go mod tidy;

# Compile Go programmes

cd /root/go/src/subnethome;
go build subnethome.go;
cd /root/go/src/subnetresult;
go build subnetresult.go;
cd /root;

# Create system user named subnetcalc with no shell, no home directory and lock account

useradd -r -s /bin/false subnetcalc;
usermod -L subnetcalc;

# Change executables file permissions, owner, group and move executables

chown root:subnetcalc /root/go/src/subnethome/subnethome;
chmod 050 /root/go/src/subnethome/subnethome;
chown root:subnetcalc /root/go/src/subnetresult/subnetresult;
chmod 050 /root/go/src/subnetresult/subnetresult;
mv /root/go/src/subnethome/subnethome /usr/bin/subnethome;
mv /root/go/src/subnetresult/subnetresult /usr/bin/subnetresult;

# Change resource file permissions, owner and group

chown -R root:subnetcalc /etc/subnetcalc;
chmod 050 /etc/subnetcalc;
chmod 040 /etc/subnetcalc/subnetcalc.env;
chmod 050 /etc/subnetcalc/html-css;
chmod 040 /etc/subnetcalc/html-css/*;

# Srart subnetcalc programs and enable on boot

systemctl start subnetcalc;
systemctl enable subnetcalc;
