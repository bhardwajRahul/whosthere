#!/bin/bash
RELEASE_TAG=${GITHUB_REF#refs/tags/}
mkdir assets
cd assets
for url in $(curl -s https://api.github.com/repos/${GITHUB_REPOSITORY}/releases/tags/${RELEASE_TAG} | jq -r '.assets[].browser_download_url'); do
  curl -L -O "$url"
done
sha256sum $(ls | grep -v checksums) > ../whosthere_${RELEASE_TAG}_checksums.txt
gh release upload ${RELEASE_TAG} ../whosthere_${RELEASE_TAG}_checksums.txt --clobber