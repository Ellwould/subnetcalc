#!/bin/bash


# [NOT READY FOR PRODUCTION]

# Install Script for Subnetcalc

#----------------------------------------------------------------------

# Check user is root otherwise exit script

if [ "$EUID" -ne 0 ]
then
  printf "\nPlease run as root\n\n";
  exit;
fi;

#----------------------------------------------------------------------

# Check subnetcalc has been cloned from GitHub

if [ ! -d "/root/subnetcalc" ]
then
  printf "\nDirectory subnetcalc does not exist in /root.\n";
  printf "Please run commands: \"cd /root; git clone https://github.com/Ellwould/subnetcalc\"\n";
  printf "and run install script again\n\n";
  exit;
fi;

#----------------------------------------------------------------------

# Copy unit files and restart systemd deamon

cp /root/subnetcalc/systemd/subnetcalc.service /usr/lib/systemd/system/;
cp /root/subnetcalc/systemd/subnetresult.service /usr/lib/systemd/system;
systemctl daemon-reload;

#----------------------------------------------------------------------

# Remove any previous version of Go, download and install Go 1.21.4 

wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz;
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz;

# Create Go directories in root home directory

mkdir -p /root/go/{bin,pkg,src/subnethome,src/subnetresult};

# Create Go mod for subnetresult

go mod init /root/go/src/subnetresult;
go mod tidy;

#----------------------------------------------------------------------

# Install the latest version of Nginx

apt install curl gnupg2 ca-certificates lsb-release debian-archive-keyring;
curl https://nginx.org/keys/nginx_signing.key | gpg --dearmor | tee /usr/share/keyrings/nginx-archive-keyring.gpg > /dev/null;
gpg --dry-run --quiet --no-keyring --import --import-options import-show /usr/share/keyrings/nginx-archive-keyring.gpg;

printf '________________________________________________________________________________________________________';
printf '\n\nThe Nginx fingerprint above should be\n      573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62\n';
printf 'Does the finger print match? (yes/no):';
read ans
if [ "$ans" = "yes" ] || [ "$ans" = "Yes" ] || [ "$ans" = "y" ] || [ "$ans" = "Y" ]
then
  printf 'Fingerprint matched\n';
else
  printf '________________________________________________________________________________________________________';
  rm /usr/share/keyrings/nginx-archive-keyring.gpg;
  printf '\n\nFor security the Nginx keyring has been removed and the script has stopped\n\n';
  printf '________________________________________________________________________________________________________\n';
  exit;
fi;
echo "deb [signed-by=/usr/share/keyrings/nginx-archive-keyring.gpg] https://nginx.org/packages/ubuntu `lsb_release -cs` nginx" tee /etc/apt/sources.list.d/nginx.list;

apt update;
apt install nginx;

#----------------------------------------------------------------------

# Copy Nginx config files

cp /root/subnetcalc/nginx/nginx.conf /etc/nginx/nginx.conf;
mkdir /etc/nginx/conf.d;
cp /root/subnetcalc/nginx/nginx_* /etc/nginx/conf.d/;

#----------------------------------------------------------------------

# Edit Nginx config files

# Function to search and replace text in files

function textUpdate {
  sed -i "s/$search/$replace/" $filename;
};

# Update Nginx config files with your server IPv4 and IPv6 addresses

searchIpArray=("Add_public_IPv4_Address" "Add_public_IPv6_Address");

for ip in ${searchIpArray[@]}
do
  filename="/etc/nginx/nginx.conf";
  search=$ip;
  echo -e "$ip"', \nif no public IP put 127.0.0.1 for IPv4 and ::1 for IPv6, \nto find the server IP addresse(s) use command "ip addr":' | tr _ " ";
  read -p "" replace;
  if [ -z "${replace}" ];
  then
    echo "IP address cannot be empty please run install script again";
    exit;
  fi;
  textUpdate;
done;

# Update Nginx config files with your FQDN and also add's Let's Encrypt cert location

fileFQDNArray=("/etc/nginx/nginx.conf" "/etc/nginx/conf.d/nginx_subnet.conf" "/etc/nginx/conf.d/nginx_tls.conf");

echo "Please enter FQDN:";
read -p "" replace;

for fileName in ${fileFQDNArray[@]}
do
  filename=$fileName;
  search="Add_FQDN";
  if [ -z "${replace}" ]
  then
    echo "FQDN cannot be empty please run install script again";
    exit;
  fi;
  textUpdate;
done;
