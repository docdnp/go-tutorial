FROM golang:1.23.0-alpine3.20

# Install Jupyter and gophernotes.
RUN apk update \
    ## -------------------------------------------------------------------------------
    ## Install Jupyter's build dependencies as virtual package
    ## (the virtual package will be removed from the image)
    && ALPINE_VERSION=$(cat /etc/alpine-release | sed -re s/\.[0-9]$//) \
    && ALPINE_REPO=http://dl-4.alpinelinux.org/alpine/v$ALPINE_VERSION \
    && apk add \
    --repository $ALPINE_REPO/community\
    --repository $ALPINE_REPO/main \
    --arch=x86_64 \
    --virtual jupyter.builddeps  \
    g++ \
    gcc \
    git \   
    libffi-dev \
    pkgconfig \
    python3-dev \
    mercurial \
    mesa-dev \
    musl-dev \
    g++ \
    gcc \
    git \
    libffi-dev \
    pkgconfig \
    python3-dev \
    mercurial \
    mesa-dev \
    musl-dev \
    zeromq-dev \
    cargo \
    ## -------------------------------------------------------------------------------
    ## Install all permanent tools and runtime environments
    ## (these will remain in the image)
    && ALPINE_VERSION=$(cat /etc/alpine-release | sed -re s/\.[0-9]$//) \
    && ALPINE_REPO=http://dl-4.alpinelinux.org/alpine/v$ALPINE_VERSION \
    && apk \
    --repository $ALPINE_REPO/community\
    --repository $ALPINE_REPO/main \
    --arch=x86_64 add \
    python3 \
    ca-certificates \
    py3-pip \
    py3-pyzmq \
    su-exec \
    nano \
    zeromq \
    bash \
    curl \
    perl \
    bash-completion \
    doas-sudo-shim \
    ## -------------------------------------------------------------------------------
    ## go tools
    ##    * gocomplete      bash completion for golang
    ##    * gonb            Non-REPL golang kernel for jupyter
    ##    * dlv             Debugger for golang
    ##    * go-outline      Outline viewer for golang
    ##    * gopls           Golang Language Server
    ##    * goimports       Auto importer for Golang
    ##    * goplay          Golang Playground
    && GOPATH=/go go install github.com/posener/complete/v2/gocomplete@latest \
    && GOPATH=/go go install github.com/janpfeifer/gonb@latest \
    && GOPATH=/go go install github.com/go-delve/delve/cmd/dlv@latest \
    && GOPATH=/go go install github.com/ramya-rao-a/go-outline@latest \
    && GOPATH=/go go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && GOPATH=/go go install golang.org/x/tools/cmd/goimports@latest \
    && GOPATH=/go go install golang.org/x/tools/gopls@latest \
    ## -------------------------------------------------------------------------------
    ## jupyter notebook
    ## (to run the golang kernel)
    && ln -s /usr/include/locale.h /usr/include/xlocale.h \
    && pip3 install -U --break-system-packages \
    pip \
    jupyter \
    jupyterlab-lsp \
    notebook \
    pyzmq \
    tornado \
    ipykernel \
    ## -------------------------------------------------------------------------------
    ## add user: gopher:gopher
    && adduser \
    --disabled-password \
    --gecos "" \
    --home "/home/gopher" \
    --shell "/bin/bash" \
    --uid "1000" \
    gopher \
    && mkdir -p /home/gopher/.local/share/jupyter/kernels \
    ## -------------------------------------------------------------------------------
    ## prepare "gopher" HOME: 
    ##    * setup bash only environment
    ##    * setup inputrc
    ##    * set password: "gopher" 
    && rm /bin/ash /bin/sh \
    && echo -e "#!/bin/bash\nexec /bin/bash \"\$@\"" > /bin/sh \
    && chmod +x /bin/sh \
    && chown -R gopher.gopher /home/gopher \
    && su - gopher -c "curl https://raw.githubusercontent.com/docdnp/inputrc/main/shell.inputrc -o ~/.inputrc" \
    && echo -e "gopher\ngopher" | passwd gopher \
    #! clean
    && apk del jupyter.builddeps \
    && pip3 cache purge \
    && go clean -cache \
    && rm -rf /root/.[acpw]* ipaexg00301* /var/cache/apk/* /go/src \
    /go/pkg/mod/*

## Setup global configuration:
##     * bash profile and bash rc
##     * prvileges of user "gopher" to run commands as root
##     * jupyter LSP settings for gopls
ADD build/config/etc.profile                             /etc/profile
ADD build/config/etc.bashrc                              /etc/bash/bashrc
ADD build/config/etc.doas.conf                           /etc/doas.d/doas.conf
ADD build/scripts/change-home-owner                      /usr/local/bin
ADD build/scripts/goexample                              /usr/local/bin
ADD build/scripts/go-tutorial                            /usr/local/bin
ADD build/scripts/entry                                  /usr/local/bin
ADD build/scripts/url                                    /usr/local/bin
ADD build/config/go-language-server-implementation.json  /usr/etc/jupyter/jupyter_server_config.d/go-language-server-implementation.json

## Setup user "gopher" configuration:
##      * bash rc
##      * jupyterlab settings (appearance and keyboard shortcuts)
ADD notebooks                                            /home/gopher/notebooks
ADD build/config/home.bashrc                             /home/gopher/.bashrc
ADD build/config/home.jupyter.lab                        /home/gopher/.jupyter/lab
ADD build/config/home.local.share.jupyter.kernels.gonb   /home/gopher/.local/share/jupyter/kernels/gonb

## Prepare environment for user "gopher"
ENV     GOPATH  /home/gopher/go
ENV     PATH    /go/bin:/home/gopher/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
WORKDIR         /home/gopher
ADD VERSION /etc/VERSION
RUN chown -R gopher.gopher /home/gopher; chmod og-w /etc/doas.d/doas.conf
USER gopher:gopher

## Prepare network and CLI access to this image
EXPOSE 8888/tcp
ENTRYPOINT [ "entry" ]
