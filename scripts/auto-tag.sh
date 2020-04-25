#!/bin/bash

set -e

git fetch --tags

last_master_tag=$(git tag --sort=-creatordate | grep -P "^\d+.\d+.\d+$" | head -n 1)
echo -e "\nLast tag: $last_master_tag\n"

bumped_tag=$(semver bump patch "$last_master_tag")
echo -e "\nBumped tag: $bumped_tag\n"

sed -i -E "s/\"version\":.*/\"version\": \"$bumped_tag\",/" package.json

git config --global user.email "safal.pandey.sp@gmail.com"
git config --global user.name "Safal Raj Pandey"
git commit -m "$bumped_tag

[skip release]"
git tag "$bumped_tag"
git push origin HEAD --tags
