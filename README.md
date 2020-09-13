#  PWNMOTHERSHIP

The place where your [pwnagotchi](https://pwnagotchi.ai/) can send their current stats.

You can run it on the **pwnahotchi** itself.



![DEMO](https://raw.githubusercontent.com/ad/pwnmothership/master/demo/demo.gif) 



## Installation

```shell
$ git clone https://github.com/ad/pwnmothership.git
$ go get github.com/rakyll/statik
$ cd pwnmothership
$ go build
$ ./pwnmothership
```

put `pwnmothership.py` into `/usr/local/share/pwnagotchi/installed-plugins` on your **pwnagotchi** and run

```shell
$ sudo pwnagotchi plugins install pwnmothership
$ sudo pwnagotchi plugins enable pwnmothership
```

It will be better with [death](https://github.com/dadav/pwnagotchi-custom-plugins) (stats) and [Experience](https://github.com/GaelicThunder/Experience-Plugin-Pwnagotchi) plugins.



## BUILD

```shell
$ go build
```

if you changed static files in `/public` directory, run

```shell
$ go generate && go build
```



## RUN

```shell
$ go get -u github.com/ad/pwnmothership

$ pwnmothership --addr=:8080
```

will start on `127.0.0.1:8080`



## CHECK

open in browser `127.0.0.1:8080/#test` 

this will show random data



## SETTINGS

to run, edit plugin settings on your **pwnagotchi**, add to your /etc/pwnagotchi/config.toml:

`main.plugins.pwnmothership.api_url = "https://your-pwnmothership-address/api/set"`

and restart

```shell
$ sudo systemctl restart pwnagotchi
```



## USING

check fingerprint of your **pwnagotchi** on page `pwnagotchi.local:8080/inbox/profile`

open in browser

`https://your-pwnmothership-address/#your-fingerprint-here`

 
