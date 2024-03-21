#!/bin/sh

if [ ! -f /tmp/userstyles.yml ]; then
	curl -sSL https://raw.githubusercontent.com/catppuccin/userstyles/main/scripts/userstyles.yml -o /tmp/userstyles.yml
	# echo "Downloaded userstyles.yml"
fi

USERSTYLES_YML="cat /tmp/userstyles.yml"

get_missing() {
	if [ -z "$2" ]; then
		$USERSTYLES_YML |
			yq ".userstyles.[] | select(.\"$1\" == null) | .name"
	else
		$USERSTYLES_YML |
			yq ".userstyles.[] | select(.\"$1\".\"$2\" == null) | .name"
	fi
}

help() {
	echo "Syntax: missing [flags] item [subitem]"
	echo
	echo "flags:"
	echo "  -c  Count missing items."
	echo "  -h  Show this help message."
}

if [ "$#" -lt 1 ]; then
	help
	exit 1
fi

if [ "$1" = "-h" ]; then
	help
elif [ "$1" = "-c" ]; then
	get_missing "$2" "$3" | wc -l
else
	get_missing "$1" "$2"
fi
