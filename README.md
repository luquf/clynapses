# Clynapses

Command line interface client for https://synapses.polytechnique.fr/

# Installation

Install the golang and the dep package manager  
```console
foo@bar:~$ sudo apt install golang-go go-dep  
```

Configure the $GOPATH and $PATH env variables in your `~/.bashrc` and move to your $GOPATH  
```console
foo@bar:~$ export GOPATH=/gopath/src
foo@bar:~$ export PATH=PATH:$GOPATH/bin
foo@bar:~$ cd $GOPATH/src
```

Clone the git repository
```console
foo@bar:~$ git clone https://github.com/luquf/clynapses  
```

Move inside the directory  
```console
foo@bar:~$ cd clynapses
```

Add your ICS url on the config/config.go file and install dependencies  
```console
foo@bar:~$ dep ensure --vendor-only   
```

Compile and install the binary  
```console
foo@bar:~$ go install  
```
