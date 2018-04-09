#!/bin/bash
ignores=macos,emacs,vim,go

curl -L -s https://www.gitignore.io/api/$ignores > .gitignore
echo vendor >> .gitignore
