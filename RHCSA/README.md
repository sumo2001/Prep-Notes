Wish me luck
<p align="center"> <img src="https://github.com/sumo2001/Prep-Notes/assets/51809378/a516eb2a-e488-4e38-8027-6d29ca3e4f67"> <p/>

- The /boot cannot reside within the LVM
- There are 6 console screens available during the installation process of RHEL, the installer is known as anaconda which saves its log at /tmp and later moves them to /var/log/
- Wayland is an advanced display protocol used by graphical applications of RHEL, Wayland provides superior graphics capabilities, features, and performance than X.
- There are two components that are critical to the functionality of a graphical environment: a display manager (a.k.a. login manager) and a desktop environment. Both are launched following the completion of the groundwork established by Wayland.
- The login/display manager is the one that presents us with a login page for us to log in, the so-called GDM (Gnome Display Manager)
- File System Categories
  - / - The Root File System
  - /etc - The Etcetera - or extended text configuration
  - /boot - the boot file system
  - /opt - the optional directory - hold additional software that may need to be installed on the system
  - /usr - the UNIX System resources - commands required at system boot
  - /var - The Variable Directory - data that frequently changes 
  - /dev - The Devices File system  - used to store device nodes for physical hardware and virtual devices, managed by udevd
  - /proc - The Procfs File System - cureent state of CPU mem Disks, fs, netwokring snd so on, destroyed at shutdown
  - /run - the repo of data for processes running on the system
  - /sys - info is used to load necessary support for the devices 
- RHEL supports seven types of files: regular, directory, block special device, character special device, symbolic link, named pipe, and socket
- gzip-compressed (-z) And bzip2-compressed (-j) List (-t) the content
- umask - Its purpose is to set default permissions on new files and directories without touching the permissions on existing files and directories.  The default umask value is set to 0022 for all userss including the root user.
- The setuid and setgid bits may be defined on binary executable files to provide non-owners and non-group members the ability to run them with the privileges of the owner or the owning group, respectively



