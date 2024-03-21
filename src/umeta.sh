#!/bin/sh

if [ ! -f /tmp/userstyles.yml ]; then
	curl -sSL https://raw.githubusercontent.com/catppuccin/userstyles/main/scripts/userstyles.yml -o /tmp/userstyles.yml
fi

USERSTYLES_YML="cat /tmp/userstyles.yml"

choose_userstyle() {
	$USERSTYLES_YML |
		yq '.userstyles | keys[]' |
		tr -d \" |
		gum choose
}

get_details() {
	$USERSTYLES_YML | yq ".userstyles.\"$(choose_userstyle)\""
}

get_userstyle() {
	USERSTYLE=$($USERSTYLES_YML | yq ".userstyles.\"$(choose_userstyle)\"")

	echo "$USERSTYLE" | yq '{name, category, "app-link": .readme."app-link", "current-maintainers": .readme."current-maintainers"}'
}

unmaintained() {
	UNMAINTAINED=$(
		$USERSTYLES_YML |
			yq '.userstyles.[] | select(.readme."current-maintainers" == []) | .name' |
			tr -d \"
	)

	echo "$UNMAINTAINED"
	echo
	echo "Total Unmaintained: $(echo "$UNMAINTAINED" | wc -l)"
}

category() {
	CHOOSE_CATEGORY=$(
		$USERSTYLES_YML |
			yq '.userstyles.[].category' |
			tr -d \" |
			sort |
			uniq |
			gum choose
	)

	CATEGORY=$(
		$USERSTYLES_YML |
			yq '.userstyles.[] | select(.category == "'"$CHOOSE_CATEGORY"'") | .name' |
			tr -d \"
	)

	echo "$CATEGORY"
	echo
	echo "Total Userstyles in $CHOOSE_CATEGORY: $(echo "$CATEGORY" | wc -l)"
}

help() {
	echo "Syntax: umeta [option]"
	echo
	echo "options:"
	echo "h     Print this Help message."
	echo "a     Gets all details of a userstyle's metadata in json."
	echo "g     Gets basic details of a userstyle's metadata."
	echo "u     Gets a list and total of unmaintained userstyles."
	echo "c     Gets a list of userstyles by a category."
}

if [ "$#" -ne 1 ]; then
	help
	exit 1
fi

if [ "$1" = "h" ]; then
	help
elif [ "$1" = "a" ]; then
	get_details
elif [ "$1" = "g" ]; then
	get_userstyle
elif [ "$1" = "u" ]; then
	unmaintained
elif [ "$1" = "c" ]; then
	category
fi
