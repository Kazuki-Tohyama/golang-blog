#!/bin/sh

set -e

host="$1"
shift
user="$1"
shift
password="$1"
shift
cmd="$@"

echo "wait.sh started"
until mysqladmin ping -h"$host" -P 3306 -u"$user" -p"$password" --silent;do
	echo "waiting for mysql..."
  sleep 1
done

echo "mysql started!"
exec $cmd
