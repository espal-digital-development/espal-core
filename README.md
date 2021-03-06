```diff
- ALPHA NOTICE : This project is in it's early alpha phase
- Please do not attempt to use in production.
- Contribution and feedback would be much appreciated!
```

# Espal: The world's fastest all-round web-based CMS, CRM, B2C, B2B, PIM, PLM, ERP in one

Espal aims to deliver a fresh next-generation all-in-one experience on CMS, CRM, B2C, B2B, PIM, PLM and ERP. The main focus of the project is performance while maintaining usability and style. Espal can easily outperform any modern-day script-based system by a hundred-fold. This means you can run an enterprise-grade system on a very small and light server. We have plans to even integrate load-balancing for big hosting-solutions like DigitalOcean in the near future so you can easily just install one Espal instance and all the balancing and resource-distribution will be taken care of.

## Summary

The aim of Espal is to offer a solution like you've never experience before mainly focussing on the speed part. Performance is the factor in everything that's being added in the system. Code is carefully crafted to ensure that performance stays at it's peek, no matter how many new functionalities get introduced.

Espal also doesn't suffer from the bloated effect that rendered coding language have and allows you to remove source code from live servers and leave only the application binary itself running.

We will also provide you with detailed deployment guides and tips to ensure both the Espal application and your server's safety will be at an enterprise grade at all times. Our philosophy is that the best should be available for everyone; not some highest bidder.

**Espal will be free, forever**. Commercial acquisition offers aren't appreciated. Professional paid support can be requested (deployment, hosting and maintenance) for those who wish so, but new features will always need to be generalized and approved by the Advocate System (see below) and will never be made private for any reason. We believe that collectivity will work best for everyone. For more information please contact us at [espaldd@icloud.com](mailto:espaldd@icloud.com). Of course donations are always welcome too.

## Details

TODO :: Actually go a bit in depth here about the features of the business logic of the CMS, CRM, B2C etc.

## Requirements

- Go 1.16+
- CockroachDB 21.1.5+

## Installation

```
go get -u github.com/espal-digital-development/espal-run
```
```
espal-run -create-project <your-project-name>
```

## Development

The given installation instructions below will specifcally indicate which OS version or variation have been used. Older and/or newer versions or variations might work, but is not guaranteed.

GoLand isn't supported as it's a paid app and we don't want to enforce cost upon the developers. We recommend Visual Studio Code.

TODO :: Explain about using mod replace to run app/module/core together in development mode

### Mac

<details>
<summary>macOS Big Sur 11.0 and macOS Catalina 10.15.6</summary>

- Install Homebrew from https://brew.sh
- Open Terminal and run these commands:
    ```
    brew install go cockroach pngquant jpegoptim gifsicle svgo
    mkdir -p ~/gopath/bin

    # This might depend on whether you use Bash, Zsh or something similar
    echo "export PATH=~/go/bin:\$PATH" >> ~/.bashrc
    echo "export PATH=~/gopath/bin:\$PATH" >> ~/.bashrc
    echo "export GOPATH=~/gopath" >> ~/.bashrc
    source ~/.bashrc
    
    brew cask install visual-studio-code

    mkir ~/Documents/GitHub
    cd ~/Documents/GitHub
    git clone https://github.com/espal-digital-development/espal-core # or your fork
    code espal-core/espal-run.code-workspace
    # Wait for the Extension recommendations popup and Install All
    # Then wait for the Go Extensions popup and click Install All
    ```
</details>

### Linux

Tested on VirtualBox with a 20GB dynamically allocated disk and an Intel Core i7-8700K.

<details>
<summary>Ubuntu 20.04 Desktop</summary>

This will likely also work for:
- antiX
- Debian
- Deepin
- Elementary OS
- Linux Lite
- Mint
- MX
- Peppermint
- Pop!_OS
- Raspberry Pi OS
- Sparky
- Zorin OS

Instructions:
- Setup with Minimal Installation option and automatic login
- Installed basic recommended update
- (Did not install Gues Additions on purpose, to keep it as minimal as possible)
- Create `~/.gitconfig` if needed and add/update:
    ```
    [user]
        name = <your-name>
        email = <your-email>
    ```
- Open Terminal and run these commands:
    ```
    sudo dpkg --configure -a # Might be needed, just run in case
    sudo apt install -y git

    mkdir -p ~/gopath/bin
    wget https://golang.org/dl/go1.14.6.linux-amd64.tar.gz
    tar xzf go1.14.6.linux-amd64.tar.gz
    rm go1.14.6.linux-amd64.tar.gz

    echo "export PATH=~/go/bin:\$PATH" >> ~/.bashrc
    echo "export PATH=~/gopath/bin:\$PATH" >> ~/.bashrc
    echo "export GOPATH=~/gopath" >> ~/.bashrc
    source ~/.bashrc

    go get -u github.com/espal-digital-development/espal-run
    
    snap install code --classic
    mkir ~/Documents/GitHub
    cd ~/Documents/GitHub
    git clone https://github.com/espal-digital-development/espal-core # or your fork
    code espal-run/espal-core.code-workspace
    # Wait for the Extension recommendations popup and Install All
    # Then wait for the Go Extensions popup and click Install All
    ```
- If you have or make a GitHub account and have not yet generated a SSH key yet:
    ```
    ssh-keygen -t ed25591 -C "your-computer-name"
    cat ~/.ssh/id_ed25519.pub
    # Select and copy the output of the cat command.
    # Then go to github.com, login and click your Avatar
    # Go to Settings > SSH & GPG Keys > New SSH Key and paste
    # the key and save it with a name.
    ```
</details>

### BSD

TODO :: Test and describe (FreeBSD, NetBSD, OpenBSD, pkg/pkgn2g package manager?)

### Solaris

TODO :: Test and describe

### Windows

<details>
<summary>Windows 10 Pro (Click to expand)</summary>

- Install:
    - Visual Studio Core from https://code.visualstudio.com/Download
    - Go from https://golang.org/dl/
    - Git from https://git-scm.com/
    - (optional) Install GitHub Desktop from https://desktop.github.com/
- Preparation
    - Open `Command Prompt` as Administrator and execute: set `PATH=%PATH%;C:\Program Files\Git\usr\bin\`
    - Also if needed, add: `C:\Users\%user%\go\bin` && `C:\Go\bin`
- Open **Visual Studio Code**
    - In the `File menu` click `Open Workspace...`
        - Navigate to the espal-core repository and open `espal-core.code-workspace`
        - Don't hide the popups in the corner and accept the recommended extensions to be installed
    - Press `Ctrl + Shift + P`
        - Search `Go: Install/Update Tools`
        - Check the top (`all`) checkbox
        - Click `OK`
    - Press `Ctrl + Shift + P`
        - Search `Go: Lint Workspace`
        - Hit Enter (Virus scanners might make the first run a bit slower)
</details>

### Docker
TODO :: Explain

### espal-run
TODO :: Explain

### espal-store-synthesizer
TODO :: Explain

### Mocking with moq
TODO :: Explain

### Connecting to database tools

Sadly there aren't many tools that flawless integrate with CockroachDB. Most Postgres tools might work, but might give issues using the secure SSL certificate style, as use of passwords isn't the default practice for the app db users.

When using a suitable app use settings similar like so:

- Host: `localhost`
- Port: `36257`
- User: `root`
- Password: `<No password option>` (some apps might work with just leaving the field empty)
- Database (name): `app`
- SSL Mode: `Require(d)`
- SSL Key: `app/database/certs/node.key`
- SSL Cert: `app/database/certs/node.crt`
- CA Cert: `app/database/certs/ca.crt`

## Deployment

TODO :: (Guides and Tips for effective and secure deployment)

## Internals
Not everything starts from absolute scratch. Even tho the aim for Espal is to be as dependency-less as possible we can't pass on the ones that already make certain tasks very easy and robust. Espal uses the following amazing libraries:

- The easy implementable and respectively lightweight [Go PostgreSQL Driver](github.com/lib/pq)
- The handy [YAML](https://github.com/go-yaml/yaml)-to-struct library for the configuration files
- [Brian Voelker](https://github.com/brianvoe)'s [gofakeit](https://github.com/brianvoe/gofakeit) for fast data fills and variations to test certain parts of the system
- [Nathan Button](https://github.com/nbutton23)'s [zxcvbn-go](https://github.com/nbutton23/zxcvbn-go) Go port of Dropbox' zxcvbn password strength checking library
- [Taco de Wolff](https://github.com/tdewolff)'s easy-to-use [minify](https://github.com/tdewolff/minify) Swiss army knife for web formats
- [Gomail](github.com/go-gomail/gomail) dialer for sending mails without any hassle

[go-bindata](https://github.com/go-bindata/go-bindata) was used. It isn't vendored in the project, but is used for the embedding of package data into native code
- [Dave Collins](https://github.com/davecgh)' amazing and super-easy to use [go-spew](https://github.com/davecgh/go-spew) dump tool. Used a lot in development

Special thanks to Google's hard work on the [Golang](https://github.com/golang/go) itself and the [crypto](golang.org/x/crypto) library.

And of course a big thanks to all the people who contributed to all of these libraries!

## Advocate system

With the advocate system we want to create a community-driven voting pool where new feature requests can be voted on. The most popular requests will be given higher priority of being implemented.

The most important indicators will always be performance first, follow by if the feature request is implementable in a generalized way so everyone can use it without having to execute manual tasks.

If highly voted proposals still get counter voted by the developers, instead of butting heads try to think constructively together to see what other possibilities there are. There might still be ways to adjust the original request into a more viable solution. Often the 2nd or 3rd iteration of an idea turns out to be the best one.

TODO :: Make the advocate system and link it here

## Performance

Espal does not provide external linking. Popular services like Redis that are fast in scripting languages cannot keep up with the Espal core, thus cannot be linked to it. This does not mean the wheel is being reinvented; many exiting solutions are being included in the project, but they simply can't live outside the source, as it will greatly diminish performance. The database is the only external service linked to the application. Through smart data-buffering we can mostly neglect the performance-loss coming from the connection between Espal and the database.

To get a sense of the speed Espal delivers (tests run on a basic MacBook Pro 2015 15-Inch); PHP framework like Laravel and Symfony clean installations with sessions turned on and a few database queries running only achieve tops of ~80 requests per second (excluding static caching).

Espal fully deployed and on it's maximum overkill secure settings achieves about 36000 requests per second (about 450 times faster). Disabling some of the overkill security settings makes the performance amp up to slightly above 50000 requests per second. This is not to brag; but because the overall load on the Espal system is so low, it leaves a lot of room for the server to breathe, making it much easier to run heavier sites on fairly light servers.

Because of the awesome `http` response-serves are almost always guaranteed. Running `wrk` high velocity benchmarks on Apache or Nginx running any PHP framework will return timeouts and even straight-out failures up to 25%. Because all Espal's requests are running queued in separate routines 99,99% of all responses are returned instantly and only under extreme load have a slight nano- or microsecond delay (locally).

## Donate

TODO :: PayPal donation link. (also make the intro link to here)

## Questions and Answers

- Q: Are domain suffixes supported?
    - A: No and the reason is that having to check the existence of the suffix chunk would cause a massive performance impact. It's a better solution to use a wildcard SSL with subdomains than trying to cramp a performance impacting path chunk.

TODO :: More general info

## License

This package is made available under an MIT-style license. See LICENSE.txt.
