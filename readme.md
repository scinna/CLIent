# Scinna CLIent

The Scinna CLIent is a simple `command line interface` client made to be used either standalone-ly (Simply using it as a CLI tool), or to be integrated with other software.

The main example is to be able to use Scinna as intended until the scinnapse client is ready to be used, or if you don't like it.


## Usage

First, you need to log in. Simply run the client with this command
```
$ scinna-client
```

This will ask you for your credentials and store your token and the configuration in your $HOME/.config/scinna/CLIent.json directory.

Then you'll be able to upload pictures with the command below, just be sure to fill the arguments as you need 

```
$ scinna-client -t "Title of the picture" -d "Description of the picture" -v "UNLISTED" {FILE PATH}
```

The CLIent will also accept piped pictures to be uploaded. This will let you use a third-party software to take the picture. Here's an example with maim
```
$ maim -s | scinna-client -t "$(date +'%Y%m%d-%H%M%S')" -v "UNLISTED" -b
```

The `-b` argument let the CLIent know he can copy the URL into the clipboard after the upload (It requires xsel/xclip on Linux).
