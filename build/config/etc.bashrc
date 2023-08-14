[ -z "$PS1" ] && return

shopt -s checkwinsize

PS1='\[\033[01;32m\]\u@gonotebook\[\033[00m\]:\[\033[01;34m\]\w\[\033[00m\]$ '

. /etc/bash/bash_completion.sh

alias grep='grep --color=auto'
alias l='ls -CF'
alias la='ls -A'
alias ll='ls -alF'
alias ls='ls --color=auto'

cd "$HOME"
