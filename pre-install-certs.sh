#!/bin/bash

# Install Script for Let's Encrypt

#----------------------------------------------------------------------

# Check user is root otherwise exit script

if [ "$EUID" -ne 0 ]
then
  printf "\nPlease run as root\n\n";
  exit;
fi;

cd /root;

#----------------------------------------------------------------------

# Input values for variables

printf "\nPlease Enter the domain name to request a certificate from Let's Encrypt\n";
read -p "Domain Name: " domainName;

if [ -z "${domainName}" ]
then
  printf "\nDomain name cannot be empty\n\n";
  exit;
fi;

# Install Certbot

apt install snapd;
snap install core;
snap refresh core;
snap install --classic certbot;  
ln -s /snap/bin/certbot /usr/bin/certbot;
certbot certonly --manual --key-type=ecdsa --elliptic-curve secp384r1 --preferred-challenges=dns --server https://acme-v02.api.letsencrypt.org/directory -d $domainName;
