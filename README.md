# Goofy-Goblin
A Malware Framework written in Golang

## Purposes
- Import only the suitable and required modules
- Make malware generation for multiple targets and operating systems easier with reusable modules
- Allow each binary to have a different signicature so that detection and identification will be much harder with the help of obfuscation modules

## Modules
### Chrome Stored Password Extraction on macOS
This module will attempt to trick the user into entering the password for accessing the Chrome keychain so that it will be able to use the key to decrypt all the password stored in the database. It uses the macOS builtin function to ask for the password so that it will look legit.
  
Two strategies are currently supported:  
1. It will attempt to ask for the password every time Chrome is opened.
2. It will spam the user to enter the password every few seconds until the password is entered.

### Full PTY shell for linux and macOS
This module can create a full PTY shell so that you can do tab-autocomplete, ctrl-c without exiting the shell, use vim, sudo and everything.  
Better shell, better hacker!
