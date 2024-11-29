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
  printf "Please run commands: \"cd /root; git clone https://github.com/Ellwould/subnetcalc\"\n";
  printf "and run install script again\n\n";
  exit;
fi;

#----------------------------------------------------------------------

# Input values for variables

printf "\nPlease enter certificate directory in path for Let's Encrypt certificates.\nFor example if the path were /etc/letsencrypt/live/example.com\nYou would enter example.com\nIf no path exists enter the domain name to request a certificate from Let's Encrypt\n";
read -p "Certificate Directory: " certDirectory;

printf "\nPlease enter the FQDN, this could be the same as the certificate directory just entered.\n";
read -p "FQDN: " FQDN;

printf "\nPlease enter IPv4 address of server, if no public IPv4 address, 127.0.0.1 can be used.\n";
read -p "Public IPv4 Address: " IPv4;

printf "\nPlease enter IPv6 address of server, if no public IPv6 address, ::1 can be used.\n";
read -p "Public IPv6 Address: " IPv6;

# Check Let's Encrypt certificate directory exists and request
# certificates if they do not exist.
# Check FQDN has been input

if [ -z "${certDirectory}" ] || [ -z "${FQDN}" ]
then
  printf "\nCertificate directory and Fully Qualified Domain Name (FQDN) cannot be empty\n";
  exit;
elif [ -z "${IPv4}" ] || [ -z "${IPv6}" ]
then
  printf "\nIPv4 and IPv6 cannot be empty\n";
  exit;
elif [ ! -d "/etc/letsencrypt/live/$certDirectory" ]
then
  printf "\nPlease run \"pre-install-cert.sh\" script first\n\n";
  exit;
fi;

#----------------------------------------------------------------------

# Copy unit files and restart systemd deamon

cp /root/subnetcalc/systemd/subnetcalc.service /usr/lib/systemd/system/;
cp /root/subnetcalc/systemd/subnetresult.service /usr/lib/systemd/system;
systemctl daemon-reload;

#----------------------------------------------------------------------

# Remove any previous version of Go, download and install Go 1.21.4 

wget -P /root https://go.dev/dl/go1.23.3.linux-amd64.tar.gz;
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.23.3.linux-amd64.tar.gz;

# Create HTML/CSS directory and copy HTML/CSS start and end file

mkdir /usr/local/etc/subnetcalc-resource;
cp /root/subnetcalc/html/subnetcalc-start.html /usr/local/etc/subnetcalc-resource/;
cp /root/subnetcalc/html/subnetcalc-end.html /usr/local/etc/subnetcalc-resource/;

# Create and insert FQDN into FQDN.txt

touch /usr/local/etc/subnetcalc-resource/subnetcalc-FQDN.txt;
echo $FQDN > /usr/local/etc/subnetcalc-resource/subnetcalc-FQDN.txt;

# Create Go directories in root home directory

mkdir -p /root/go/{bin,pkg,src/subnethome,src/subnetresult};
mkdir /usr/local/go/src/subnetcalcresource;

# Copy Go source code

cp /root/subnetcalc/go/subnethome.go /root/go/src/subnethome/subnethome.go;
cp /root/subnetcalc/go/subnetresult.go /root/go/src/subnetresult/subnetresult.go;
cp /root/subnetcalc/go/subnetcalcresource.go /usr/local/go/src/subnetcalcresource/subnetcalcresource.go;

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
mv /root/go/src/subnethome/subnethome /usr/local/bin/subnethome;
mv /root/go/src/subnetresult/subnetresult /usr/local/bin/subnetresult;

# Change resource file permissions, owner and group

chown -R root:subnetcalc /usr/local/etc/subnetcalc-resource;
chmod 050 /usr/local/etc/subnetcalc-resource;
chmod 040 /usr/local/etc/subnetcalc-resource/*;

# Srart subnetcalc programs and enable on boot

systemctl start subnetcalc;
systemctl enable subnetcalc;

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

filename="/etc/nginx/nginx.conf";
search="Add_public_IPv4_Address";
replace=$IPv4;
textUpdate;

filename="/etc/nginx/nginx.conf";
search="Add_public_IPv6_Address";
replace=$IPv6;
textUpdate;

# Update Nginx config files with your FQDN and also add's Let's Encrypt cert location

fileFQDNArray=("/etc/nginx/nginx.conf" "/etc/nginx/conf.d/nginx_subnet.conf");

for fileName in ${fileFQDNArray[@]}
do
  filename=$fileName;
  search="Add_FQDN";
  replace=$FQDN;
  if [ -z "${replace}" ]
  then
    echo "FQDN cannot be empty please run install script again";
    exit;
  fi;
  textUpdate;
done;

# Update Nginx TLS file with certificate directroy.

filename="/etc/nginx/conf.d/nginx_tls.conf";
search="Add_Cert_Directory";
replace=$certDirectory;
textUpdate;

systemctl restart nginx;
