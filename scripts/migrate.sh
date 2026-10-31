#!/usr/bin/env bash
# Runs goose migrations using the database URL from ~/.gatorconfig.json.
# Usage: ./migrate.sh up | down | status | reset
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
action="${1:?usage: $0 up|down|status|reset}"

# -e makes jq fail if db_url is missing or null
url="$(jq -er '.db_url' "$HOME/.gatorconfig.json")"

while true; do
  read -r -p "Run 'goose $action' on the gator database? (yes/no) " answer
  case "$answer" in
  yes) break ;;
  no)
    echo "Aborted."
    exit 0
    ;;
  *) echo "Please answer yes or no in full." ;;
  esac
done

goose -dir "$root/sql/schema" postgres "$url" "$action"
