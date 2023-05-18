Docker volume plugin for GlusterFS

This is a managed Docker volume plugin to allow Docker containers to access
GlusterFS volumes. The GlusterFS client does not need to be installed on the
host and everything is managed within the plugin.

## Usage

1 - Install the plugin

```
docker plugin install --alias glusterfs jorussek/glusterfs:latest

# optional you can set a default server list and/or volume
docker plugin install --alias glusterfs jorussek/glusterfs SERVERS=<server1,server2,...,serverN> VOLNAME=<volname>

# or to enable debug
docker plugin install --alias glusterfs jorussek/glusterfs DEBUG=1
```

2 - Create a volume

> Make sure the **_gluster volume exists_**.
>
> Or the mounting of the volume will fail.

```
$ docker volume create -d glusterfs -o servers=<server1,server2,...,serverN> -o volname=<volname> -o subdir=<subdir> glustervolume
glustervolume
$ docker volume ls
DRIVER           VOLUME NAME
glusterfs:next   glustervolume
```

or if you set the defaults for the plugin, you can create a volume without any options:

```
$ docker volume create -d glusterfs glustervolume
glustervolume
$ docker volume ls
DRIVER           VOLUME NAME
glusterfs:next   glustervolume
```

3 - Use the volume

```
$ docker run -it -v glustervolume:<path> bash ls <path>
```

## Options

- servers [required, if no default set]: A comma-separated list of servers e.g.: 192.168.2.1,192.168.1.1
- volname [required, if no default set]: The name of the glusterfs volume e.g.: gv0. Needs to be defined on the glusterfs cluster.
- subdir [optional, default: volume name]: The name of the subdir. Will be created, if not found.

For additional options see [man mount.glusterfs](https://github.com/gluster/glusterfs/blob/release-6/doc/mount.glusterfs.8).

## Supported tags and respective `Dockerfile` links

- jorussek/glusterfs:latest -> jorussek/glusterfs:9
- [jorussek/glusterfs:9](https://github.com/jorussek/docker-volume-glusterfs/blob/glusterfs-9/Dockerfile) (Ubuntu 20.04)
- [jorussek/glusterfs:8](https://github.com/jorussek/docker-volume-glusterfs/blob/glusterfs-8/Dockerfile) (Ubuntu 20.04)
- [jorussek/glusterfs:7](https://github.com/jorussek/docker-volume-glusterfs/blob/glusterfs-7/Dockerfile) (Ubuntu 20.04)
- [jorussek/glusterfs:6](https://github.com/jorussek/docker-volume-glusterfs/blob/glusterfs-6/Dockerfile) (Ubuntu 20.04)
- [jorussek/glusterfs:1.1.0](https://github.com/jorussek/docker-volume-glusterfs/blob/4af73f9ba63e816958f25a2bddf5665f6c859fe9/Dockerfile) (Ubuntu 18.04)
- [jorussek/glusterfs:1.0.0](https://github.com/jorussek/docker-volume-glusterfs/blob/4af73f9ba63e816958f25a2bddf5665f6c859fe9/Dockerfile) (Ubuntu 18.04)

## TODO

- write integration tests

## LICENSE

MIT
