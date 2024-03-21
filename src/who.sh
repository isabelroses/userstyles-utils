#!/bin/sh

if [ ! -f /tmp/userstyles.yml ]; then
	curl -sSL https://raw.githubusercontent.com/catppuccin/userstyles/main/scripts/userstyles.yml -o /tmp/userstyles.yml
	# echo "Downloaded userstyles.yml"
fi

USERSTYLES_YML="cat /tmp/userstyles.yml"

CHOOSEN_COLLABORATOR=$(
	$USERSTYLES_YML |
		yq '.collaborators.[] | if .name == null then .url else .name end' |
		sed "s/https:\/\/github.com\///g" |
		tr -d \" |
		gum choose
)

GET_USERSTYLES_MAINTAINED=".userstyles.[] | select(.readme.\"current-maintainers\".[] | if (.name) then .name == \"$CHOOSEN_COLLABORATOR\" else .url == \"https://github.com/$CHOOSEN_COLLABORATOR\" end) | if (.name | type==\"array\") then .name.[0] else .name end"

USERSTYLES_MAINTAINED=$(
	$USERSTYLES_YML |
		yq "$GET_USERSTYLES_MAINTAINED" |
		tr -d \"
)

TOTAL_USERSTYLES_MAINTAINED=$(
	echo "$USERSTYLES_MAINTAINED" |
		wc -l |
		tr -d " "
)

echo "Userstyles Maintained by $CHOOSEN_COLLABORATOR:"
echo
echo "$USERSTYLES_MAINTAINED"
echo
echo "Total Userstyles Maintained: $TOTAL_USERSTYLES_MAINTAINED"
