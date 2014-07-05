messaged
========

a message deamon. see also messenger. only tested on linux as it requires libnotify. it accepts messages via Port 65222 (send from messenger) and displays those via a handler. now there are only 2 handlers: libnotify, which is a desktop notification, and stdout.

Running it
----------
as default it requires ssl certificates. change the type to disable ssl.
```bash
messaged &
messenger --type tcp --title Hello --message World --to localhost
```

will display a message via libnotify (hint: if you use self signed certificates use --no-ssl-verify in messenger)

Get it
-------
```bash
go get github.com/ProhtMeyhet/messaged
```

note: on my raspberry with raspbian there is a problem with the go-notify package. i have no resolution yet.

Dependencies
-------------
libnotify

https://github.com/ProhtMeyhet/libgomessage

Licence
-------
see LICENCE-AGPL
