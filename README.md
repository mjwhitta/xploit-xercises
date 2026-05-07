# Xploit-Xercises

## Usage

You can modify and use the provided [compose.yaml] or you can download
and run the newest release:

```
$ docker compose up -d
$ # OR
$ docker run \
    --detach \
    --hostname="xx" \
    --name="xx" \
    --publish="2222:22" \
    --rm \
    ghcr.io/mjwhitta/xploit-xercises:rolling
```

You can access the container via SSH like so:

```
$ ssh -l nebulous -p 2222 localhost
$ ssh -l protostellar -p 2222 localhost
$ ssh -l fission -p 2222 localhost
```

The passwords all match the associated usernames. These accounts all
have password-less `sudo` access, in case you need to install any
tools. NO CHEATING!

Alternatively, you can choose not to publish a port and access the
container like so:

```
$ docker exec -i -t -u nebulous xx bash -l
$ docker exec -i -t -u protostellar xx bash -l
$ docker exec -i -t -u fission xx bash -l
```

## Supported challenges

### Nebulous

This is a modern re-implementation of [Nebula]. Instead of 32-bit C
binaries, this uses a 64-bit Go binaries. Instead of Ubuntu 11.10,
this uses Alpine (edge), Debian (stable and stable-slim), or Ubuntu
(latest and rolling). The default is Ubuntu (rolling). See the
[Dockerfile] to build an alternative. Instead of being distributed as
an ISO, this is distributed as a Docker container that you can pull or
build yourself. As such, some of the vulnerabilities may be simulated,
but the spirit of Nebula has been preserved.

After you login as `nebulous` there is a folder for each level. Inside
is a `README.md` with some details. There is also a `hints.md` if you
feel stuck. The ultimate goal of each level is to run `getflag` as the
associated `flagXX` account.

### Protostellar

[TODO][Protostar]

### Fission

[TODO][Fusion]

[compose.yaml]: ./compose.yaml
[Dockerfile]: ./Dockerfile
[Fusion]: https://exploit.education/fusion
[Nebula]: https://exploit.education/nebula
[Protostar]: https://exploit.education/protostar
