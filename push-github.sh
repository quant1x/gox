#!/usr/bin/env bash

set -e
git remote set-url origin https://github.com/quant1x/gox.git
git push --all
git push --tags
git remote -vv