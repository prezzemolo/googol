#!/usr/bin/fish
alias commit='bash -c \'git add --all && git commit -m "end of $(git diff --name-only --cached | grep .go | cut -d\\/ -f1) section $(git diff --name-only --cached | grep .go | rev | cut -d/ -f1 | cut -c 4- | rev | sed -e \\\'s/\\\\./ (/\\\' -e \\\'s/-/ /g\\\'))"\''
