## Blank Heading
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

 
