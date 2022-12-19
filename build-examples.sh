#!/bin/sh

pushd examples >/dev/null

for dir in *; do
  echo build $dir

  pushd $dir >/dev/null

  go build

  popd >/dev/null
done

popd >/dev/null