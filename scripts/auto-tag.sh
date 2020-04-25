#!/bin/bash

last_master_tag=$(git tag --sort=-creatordate | grep -P "^\d+.\d+.\d+$" | head -n 1)
bumped_tag=$(semver bump patch "$last_master_tag")

git tag "$bumped_tag"
git push origin HEAD --tags
