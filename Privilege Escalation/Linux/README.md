## Briefing
+ Privilege escalation seems to be simple but requires a lot of reconnaissance on the compromised system, once you can able to get a simple shell then escalating into it into a root shell is easy  
+ But you may require some misconfigurations or combination of misconfigurations to spawn a root shell 
+ Mostly every privilege escalation in Linux is due to active access control violations  
+ Access control and user permissions are interlinked when focusing on escalations in linux 
 ### Users and Groups
+ Users are configured in /etc/passwd file where user password hashes are stored in /etc/shadow file 
+ Users are identified with an integer user ID called UID , Root user is a special type of account in Linux it has a UID of 0 where the system grants all permissions on every file 
+ Groups are configured in /etc/group file.
+ Users have can have multiple groups , Users can have a primary group and multiple secondary or supplementary groups by default users primary group will be their username 
+ Groups can have multiple users every file 
+ Every directory defines it permissions in terms of a user 
### Directories
+ All files and Directories can only have a single owner, only owner can change permissions to a file or a directory 
+ when a execute flag is set then only you can enter into the directory or you can cd into the directory when a read is set you can just list the contents in the directory when write flag is set you can create subdirectories in the directory 
+ Read =4, Write = 2, Execute = 1 
### SUID/SGID
+ SUID, when set files will inherit the permissions of the file owner.  
+ SGID set on a file, it allows the file to be executed as the group that owns the file (similar to SUID), If set on a directory, any files created in the directory will have their group ownership set to that of the directory owner
+ The last special permission has been dubbed the "sticky bit." This permission does not affect individual files. However, at the directory level, it restricts file deletion. Only the owner (and root) of a file can remove the file within that directory.

### ID's
+ Three types of ID's, real ID, effective ID saved UID.  
+ The real user ID is The Who they actually are. It's defined in the PASSWD file.  
+ Real ID is actually used less often to the check. It users identity.  
+ A user's effective ID is normally equal to to their real ID.  
+ When executing a process as a another user, the effective ID is set to that user's real ID.  
+ The effective ID is used in most access control decisions to verify user. Such as whoami  command will use effective ID.  
+ Saved UID: It is used when a process is running with elevated privileges (generally root) needs to do some under-privileged work, this can be achieved by temporarily switching to a non-privileged account, While performing under-privileged work, the effective UID is changed to some lower privilege value, and the euid is saved to saved userID(suid), so that it can be used for switching back to a privileged account when the task is completed. 

### Spawning Shells
+ Executing /bin/sh or /bin/bash
+ "rootbash" SUID: Create a copy of the /bin/bash executable file (I usually rename it rootbash), make sure it is owned by the root user, **and has the SUID bit set.**, root shell can be spawned by simply executing the rootbash file with the -p flag
+ msfvenom can also be used to craft a reverse shell
+ A tool to generate various ways to spawn a reverse shell: https://github.com/mthbernardes/rsg
### Tools
+ Tools allow us to automate the reconnaissance that can identify potential privilege escalations.
+ https://github.com/diego-treitos/linux-smart-enumeration 
+     curl "https://github.com/diego-treitos/linux-smart-enumeration/raw/master/lse.sh" -Lo lse.sh;chmod 700 lse.sh
+ https://github.com/rebootuser/LinEnum
+ https://github.com/linted/linuxprivchecker
+ https://github.com/AlessandroZ/BeRoot
+ http://pentestmonkey.net/tools/audit/unix-privesc-check

## Exploits
1. [Kernel Exploits](https://github.com/sumo2001/Prep-Notes/blob/main/Privilege%20Escalation/Linux/README.md#kernel-exploits)
2. [Service Exploits](https://github.com/sumo2001/Prep-Notes/blob/main/Privilege%20Escalation/Linux/README.md#service-exploits)
3. [Weak File Permissions](https://github.com/sumo2001/Prep-Notes/blob/main/Privilege%20Escalation/Linux/README.md#weak-file-permissions)
4. [Sudo](https://github.com/sumo2001/Prep-Notes/blob/main/Privilege%20Escalation/Linux/README.md#sudo)

### Kernel Exploits
+ Kernels are the core of any operating system, a layer between application software and the actual computer hardware.
+ The kernel has complete control over the operating system, a kernel vulnerability can results in root shell
+ Finding and using kernel exploits is usually a simple process:
+      1. Enumerate kernel version (uname -a).
       2. Find matching exploits (Google, ExploitDB, GitHub).
       3. Compile and run
 + **Beware though, as Kernel exploits can often be unstable and may be one-shot or cause a system crash, as it has to be the last option to consider**
 + We can use searchsploit to find certian kernel exploit after successful enumeration
 + But, using [Linux Exploit Suggester 2](https://github.com/jondonas/linux-exploit-suggester-2) will help us to find kernel exploits efficiently.
 + Example:
     + Lets consider enumeration reveals the kernel version as "2.6.32"
     + searchsploit linux kernel 2.6 priv esc debian
     + ./linux-exploit-suggester-2.pl –k 2.6.32
     + This will help us to esclate and get a root shell, ps: reading the exploit code/steps is very important
 ### Service Exploits
+ Services are simply programs that run in the background, accepting input or performing regular tasks, running as root and they are vulnerable?, then voila!
```
ps aux | grep "^root" 
<program> -- version
Identify the exploit
```
+ Consider an example, say services ssh and myql and running as root on a remote machine, enabling port forwarding mysql with random port,and connecting to local machines, spawns a root shell
    + A root process may be bound to an internal port, through which it communicates.
    + If for some reason, an exploit cannot run locally on the target machine, the port can be forwarded using SSH to your local machine
    + "ssh -R local-port:127.0.0.1:target-port username@local-machine"
### Weak File Permissions
  + find /etc -maxdepth 1 **-writable/-readable** -type f
  + Find all directories which can be written to: "find / -executable -writable -type d 2> /dev/null"
 #### The /etc/shadow : contains user password hashes
 + If we are able to read the contents of the /etc/shadow file, we might be able to crack the root user’s password hash.
 + If we can able to write we can replace the hash with a known hash
       + mkpasswd -m sha-512 newpassword
       + replace the above output ;)
 #### The /etc/passwd
+ If we can write to /etc/passwd, we can easily enter a known password hash for the root user, and then use the su command to switch to the root user.
+ ![image](https://user-images.githubusercontent.com/51809378/139635790-c00af860-4abd-493d-ba99-42d8b5d02f97.png)
+ we can create a new user but assign them the root user ID (0). This works because Linux allows multiple entries for the same user ID, as long as the usernames are different.
+ ![image](https://user-images.githubusercontent.com/51809378/139635838-3d71776e-4a40-42a8-8cda-06e69a87ab41.png)
+ In some versions of Linux, it is possible to simply delete the “x”, which Linux interprets as the user having no password
+ ![image](https://user-images.githubusercontent.com/51809378/139635677-11a2a039-945a-4c5d-8f3b-0d4e91a3d466.png)
#### Insecure Backups
+ It is always worth exploring the file system looking for readable backup files. Some common places include user home directories, the / (root) directory, /tmp, and /var/backups. Example : SSH private keys
+ 



 
