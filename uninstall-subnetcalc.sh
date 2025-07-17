#!/bin/bash

# Uninstall script for Subnetcalc

#----------------------------------------------------------------------

# Check user is root otherwise exit script

if [ "$EUID" -ne 0 ]
then
  printf "\nPlease run as root\n\n";
  exit;
fi;

cd /root;

#----------------------------------------------------------------------

# Stop Subnetcalc automatically starting on boot

systemctl stop subnetcalc.service;
systemctl disable subnetcalc.service;

# Remove Subnetcalc unit files and reload systemd deamon

rm /usr/lib/systemd/system/subnetcalc.service;
rm /usr/lib/systemd/system/subnetresult.service;
systemctl daemon-reload;

#----------------------------------------------------------------------

# Remove subnethome and subnetresult binaries

rm /usr/bin/subnethome;
rm /usr/bin/subnetresult;

# Remove all other directores and files used by Subnetcalc

rm -r /etc/subnetcalc;

# Remove subnethome and subnetresult source code in root home directory

rm -r /root/go/src/subnethome;
rm -r /root/go/src/subnetresult;

# Remove the user and group subnetcalc from the system

userdel subnetcalc;
