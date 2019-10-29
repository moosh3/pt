[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

# pt
`pt` is a robust cli tool that interacts with [pomotodo](). 

# Overview

# Installation
Make sure you have a working Go environment (go 1.1+ is *required*). [See the install instructions](http://golang.org/doc/install.html).

To get `pt`, run:
```
$ go get github.com/aleccunningham/pt
```

To install it in your path so that `pt`can be easily used:

```
$ cd $GOPATH/src/github.com/aleccunningham/pt
$ GOBIN="/usr/local/bin" go install
```

### Docker

You can also use the official docker image hosted at `r.kubernetes.lol`.

```
$ make docker
$ docker run -t -v $(pwd)/config.yaml:/config.yaml r.kubernetes.lol/pt:$VERSION <command>
```

## Getting Started

The **start** command will start tracking time for a given taskname. **Note that a taskname cannot have white spaces**, they serve as identifiers.

```
$ golog start {taskname}
```

To stop tracking use the **stop** command, if you want to **resume** a stopped task just golog start {taskname} again.

```
$ golog stop {taskname}
```

With the **list** command you can see all your tasks and see which of them are active.

```
$ golog list
0h:1m:10s    create-readme (running)
0h:0m:44s    do-some-task
```

If you only want to check the status of one task, use the **status** command.

```
$ golog status create-readme
0h:3m:55s    create-readme (running)
```

The lifetime of the info I usually need is very short (actually is just a workday), in the next day it's unlikely that i'll need previous info. This is one case where **clear** command is handy.

```
$ golog clear
All tasks were deleted.
```

# Contribution Guidelines
@TODO
If you have any questions feel free to link @aleccunningham to the issue in question and we can review it together.
