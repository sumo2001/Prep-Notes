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
### SGID/SUID
+ SUID, when set files will get executed with the permissions of the file owner.  
+ SGID when set on a file, the files will get executed with the privileges of the file group. When set on a directory, files created within the directory will inherit that group. Of the directory itself. 

### ID's
+ Three types of ID's, real ID, effective ID saved UID.  
+ The real user ID is The Who they actually are. It's defined in the PASSWD file.  
+ Real ID is actually used less often to the check. It users identity.  
+ A user's effective ID is normally equal to to their real ID.  
+ When executing a process as a another user, the effective ID is set to that user's real ID.  
+ The effective ID is used in most access control decisions to verify user. Such as whoami  command will use effective ID.  
+ Saved UID will be used to ensure that SUID processes can temporarily switch a users effective ID back to their real ID and back again without losing the track of their original effective ID. 

 
