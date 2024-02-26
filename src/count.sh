#!/bin/sh

if [ ! -f /tmp/userstyles.yml ]; then
	curl -sSL https://raw.githubusercontent.com/catppuccin/userstyles/main/scripts/userstyles.yml -o /tmp/userstyles.yml
	# echo "Downloaded userstyles.yml"
fi

USERSTYLES_YML="cat /tmp/userstyles.yml"

TOTAL_USERSTYLES=$(
	$USERSTYLES_YML |
		yq '.userstyles | keys[]' |
		wc -l |
		tr -d " "
)

TOTAL_MAINTAINERS=$(
	$USERSTYLES_YML |
		yq '.collaborators | length'
)

CURRENT_MAINTAINERS=$(
	$USERSTYLES_YML |
		yq '.userstyles.[].readme."current-maintainers" | add | select( . != null ) | .url' |
		sort |
		uniq -c |
		wc -l |
		tr -d " "
)

PAST_MAINTAINERS=$((TOTAL_MAINTAINERS - CURRENT_MAINTAINERS))

echo "Userstyles: $TOTAL_USERSTYLES"
echo ""
echo "Total Maintainers: $TOTAL_MAINTAINERS"
echo "Current Maintainers: $CURRENT_MAINTAINERS"
echo "Past Maintainers: $PAST_MAINTAINERS"
