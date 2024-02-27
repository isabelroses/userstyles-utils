#!/bin/sh

if [ ! -f /tmp/userstyles.yml ]; then
	curl -sSL https://raw.githubusercontent.com/catppuccin/userstyles/main/scripts/userstyles.yml -o /tmp/userstyles.yml
fi

USERSTYLES_YML="cat /tmp/userstyles.yml"

ChooseUserstyle() {
	$USERSTYLES_YML |
		yq '.userstyles | keys[]' |
		tr -d \" |
		gum choose
}

GetDetails() {
	USERSTYLE=$($USERSTYLES_YML | yq ".userstyles.\"$(ChooseUserstyle)\"")

	echo "$USERSTYLE"
}

GetUserstyle() {
	USERSTYLE=$($USERSTYLES_YML | yq ".userstyles.\"$(ChooseUserstyle)\"")

	echo "$USERSTYLE" | yq '{name, category, "app-link": .readme."app-link", "current-maintainers": .readme."current-maintainers"}'
}

Unmaintained() {
	UNMAINTAINED=$(
		$USERSTYLES_YML |
			yq '.userstyles.[] | select(.readme."current-maintainers" == []) | .name' |
			tr -d \"
	)

	echo "$UNMAINTAINED"
	echo
	echo "Total Unmaintained: $(echo "$UNMAINTAINED" | wc -l)"
}

Category() {
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

Help() {
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
	Help
	exit 1
fi

if [ "$1" = "h" ]; then
	Help
elif [ "$1" = "a" ]; then
	GetDetails
elif [ "$1" = "g" ]; then
	GetUserstyle
elif [ "$1" = "u" ]; then
	Unmaintained
elif [ "$1" = "c" ]; then
	Category
fi
