#!/bin/bash

set -e

last_master_tag=`git tag --sort=-creatordate | grep -P "^\d+.\d+.\d+$" | head -n 1`
echo -e "\nLast tag: $last_master_tag\n"

bumped_tag=`semver bump patch "$last_master_tag"`
echo -e "\nBumped tag: $bumped_tag\n"

git tag "$bumped_tag"
git push origin HEAD --tags
