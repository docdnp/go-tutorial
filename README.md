# gotutorial

[![Publish docker images on docker hub](https://github.com/docdnp/go-tutorial/actions/workflows/main.yml/badge.svg)](https://github.com/docdnp/go-tutorial/actions/workflows/main.yml)

An interactive tutorial for go on basis of Jupyter Notebooks and gonb.

It is provided as a self-contained Docker image.

## Background

A former employer, asked me to do a tech teaser on the Go programming language before I left the company.
Earlier, rather PowerPoint-heavy tech teasers, had the big disadvantage that the learning effect was lower, but they were better suited to the character of a tech teaser.
Especially learning a new programming language is much easier when you dig your hands right into the dirt and very short feedback loops move you forward.
The idea of using a Jupyter notebook for such a format was born.
Because it was clear that the effort of such a "project" would be bigger and therefore would be done in large parts of my free time, I understandably didn't want to just "give away" this work. Fortunately it became possible to create this Jupyter notebook.

The first Jupyter kernel for Golang I used was [gophernotes](https://github.com/gopherdata/gophernotes#readme).
This was based on [gomacro](https://github.com/cosmos72/gomacro#readme), a [REPL](https://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop) for Go.
I was thrilled with how easily and quickly the tutorial's repository developed based on small examples.
But I also quickly reached the limits of how subtly different the REPL and native Go behaved in individual areas.
Then I discovered another Go kernel for Jupyter notebooks, namely [gonb](https://github.com/janpfeifer/gonb#readme).
Its concept takes advantage of the fact that the compile time of Golang is so fast that it does without a REPL and builds a Go program before an execution, which is compiled and executed immediately.
After adapting all the examples, I am sure that this approach carries better for a tutorial.
However, restarting a kernel and executing all cells takes a "relatively" long time.
One must be aware of this.

Personally, it was important to me to keep the hurdle to get started with a new programming language as low as possible.
Simply start a Docker container and get started right away.
That was my idea.
A mix of [Effective Go](https://go.dev/doc/effective_go) and the [Go Playground](https://go.dev/play/).
Supplemented by a usable shell environment and many tools that are helpful in developing Go programs.
Therefore, I envisioned that everything that belongs together would come together in a Docker image:

* The tutorial itself
* An installation of Jupyter-Lab
* The Jupyter kernel for Go
* A shell environment that makes sense
* All sorts of tools and possibilities around developing Go programs.

So this is the preliminary result, which will hopefully develop steadily.
Unfortunately, in the first version I could not finish some topics that are very important from my point of view.
But these will surely follow.
And who knows... maybe some more tutorials for other programming languages and technologies will follow.

But now have fun trying it out. :-)

## Build

```shell
build/scripts/build
```

## Execute

```shell
docker run -i --rm thednp/go-tutorial
```

## Workflow when developing

```shell
build/scripts/go-tutorial init-local
build/scripts/go-tutorial start

# [... working in Jupyter ...]
# [... saving changes ...]

build/scripts/build
build/scripts/go-tutorial start --docker    # open build result on port 8889

# [... check changes ...]
# [... working in Jupyter ...]
# [... saving changes ...]

build/scripts/build
build/scripts/go-tutorial start --docker    # open build result on port 8889

# [... check changes ...]

build/scripts/go-tutorial stop --docker
build/scripts/go-tutorial stop
build/scripts/go-tutorial rm-local
```
