FROM alpine:latest

# Install Jupyter and gophernotes.
RUN apk update \
    ## -------------------------------------------------------------------------------
    ## Install Jupyter's build dependencies as virtual package
    ## (the virtual package will be removed from the image)
    && apk add \
        --repository http://dl-4.alpinelinux.org/alpine/v3.18/community \
        --repository http://dl-4.alpinelinux.org/alpine/v3.18/main \
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
    && apk \
        --repository http://dl-4.alpinelinux.org/alpine/v3.18/community \
        --repository http://dl-4.alpinelinux.org/alpine/v3.18/main \
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
        bash-completion \
        doas-sudo-shim \
        go \
    ## -------------------------------------------------------------------------------
    ## go tools
    ##    * gocomplete      bash completion for golang
    ##    * gophernotes     REPL golang kernel for jupyter
    ##    * gonb            Non-REPL golang kernel for jupyter
    ##    * dlv             Debugger for golang
    ##    * go-outline      Outline viewer for golang
    ##    * gopls           Golang Language Server
    ##    * goimports       Auto importer for Golang
    ##    * goplay          Golang Playground
    && GOPATH=/go go install github.com/posener/complete/v2/gocomplete@latest \
    && GOPATH=/go go install github.com/janpfeifer/gonb@latest \
    && GOPATH=/go go install github.com/gopherdata/gophernotes@latest \
    && GOPATH=/go go install github.com/go-delve/delve/cmd/dlv@latest \
    && GOPATH=/go go install github.com/ramya-rao-a/go-outline@latest \
    && GOPATH=/go go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && GOPATH=/go go install golang.org/x/tools/cmd/goimports@latest \
    && GOPATH=/go go install golang.org/x/tools/gopls@latest \
    ## -------------------------------------------------------------------------------
    ## jupyter notebook
    ## (to run the golang kernel)
    && ln -s /usr/include/locale.h /usr/include/xlocale.h \
    && pip3 install -U \
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
    ## -------------------------------------------------------------------------------
    ## prepare "gopher" HOME: 
    ##    * install go-jupyter-kernel from gophernotes 
    && git clone -n --depth=1 --filter=tree:0 https://github.com/gopherdata/gophernotes.git gn \
    && cd gn && git sparse-checkout set --no-cone kernel \
    && git checkout \
    && mkdir -p /home/gopher/.local/share/jupyter/kernels \
    && cp -r kernel /home/gopher/.local/share/jupyter/kernels/gophernotes \
    && cd .. && rm -rf gn \
    ## -------------------------------------------------------------------------------
    ## prepare "gopher" HOME: 
    ##    * install go-jupyter-kernel from gophernotes 
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
    && apk cache clean \
    && pip3 cache purge \
    && go clean -cache \
    && rm -rf /root/.[acpw]* ipaexg00301* /var/cache/apk/* /go/src \
               /go/pkg/mod/*

## Setup global configuration:
##     * bash profile and bash rc
##     * prvileges of user "gopher" to run commands as root
##     * jupyter LSP settings for gopls
COPY build/config/etc.profile                             /etc/profile
COPY build/config/etc.bashrc                              /etc/bash/bashrc
COPY build/config/etc.doas.conf                           /etc/doas.d/doas.conf
COPY build/scripts/goexample                              /usr/local/bin
COPY build/config/go-language-server-implementation.json  /usr/etc/jupyter/jupyter_server_config.d/go-language-server-implementation.json

## Setup user "gopher" configuration:
##      * bash rc
##      * jupyterlab settings (appearance and keyboard shortcuts)
COPY ./jupyter/home/gopher/.bashrc                          /home/gopher/.bashrc
COPY ./jupyter/home/gopher/notebooks                        /home/gopher/notebooks
COPY ./jupyter/home/gopher/.jupyter/lab                     /home/gopher/.jupyter/lab
COPY /jupyter/home/gopher/.local/share/jupyter/kernels/gonb /home/gopher/.local/share/jupyter/kernels

## Prepare environment for user "gopher"
USER gopher:gopher
ENV     GOPATH  /home/gopher/go
ENV     PATH    /go/bin:/home/gopher/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
WORKDIR         /home/gopher

## Prepare network and CLI access to this image
EXPOSE 8888
CMD [ "jupyter", "lab", "--no-browser", "--allow-root", "--ip=0.0.0.0" ]
