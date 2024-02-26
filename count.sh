#!/usr/bin/env bash

USERSTYLES=$(
	curl -sSL https://raw.githubusercontent.com/catppuccin/userstyles/main/scripts/userstyles.yml |
		yq '.userstyles | keys[]' |
		wc -l
)

MAINTAINERS=$(
	curl -sSL https://raw.githubusercontent.com/catppuccin/userstyles/main/scripts/userstyles.yml |
		yq '.collaborators | length'
)

echo "Userstyles: $USERSTYLES"
echo "Maintainers: $MAINTAINERS"
