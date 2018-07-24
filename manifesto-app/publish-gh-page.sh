#!/bin/bash

set -e

cd ..
git checkout master
git subtree split --prefix=manifesto-app/src -b gh-pages    
git push -f origin gh-pages:gh-pages
git branch -D gh-pages