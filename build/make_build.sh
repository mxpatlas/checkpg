#!/bin/bash
PATH="$PATH:-/bin:/sbin:/usr/bin:/usr/sbin}"

BUILDROOT="${BUILDROOT:-buildroot}"
if [ -z $"BUILDROOT" ]; then
	echo "BUILDROOT required"
  exit 1
fi

if [ ! -d "$BUILDROOT" ]; then 
	echo "Create BUILDROOT='$BUILDROOT'"
  mkdir -p "$BUILDROOT" || exit 1
fi


BINDIR="$BUILDROOT/bin"
if [ ! -d "$BINDIR" ]; then
	echo "Create BINDIR='$BINDIR'"
  mkdir -p "$BINDIR" || exit 1
fi

go build -o "$BINDIR/checkpg" cmd/checkpg/main.go

