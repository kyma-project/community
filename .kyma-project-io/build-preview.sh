#!/usr/bin/env bash
# Script for build preview of this repo like in https://kyma-project.io/community/ on every PR.
# For more information, please contact with: @michal-hudy @m00g3n @aerfio @magicmatatjahu

set -e
set -o pipefail

pushd "$(pwd)" > /dev/null

on_error() {
  echo -e "${RED}✗ Failed${NC}"
  exit 1
}
trap on_error ERR

on_exit() {
  popd > /dev/null
}
trap on_exit EXIT

readonly KYMA_PROJECT_IO_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

readonly WEBSITE_DIR="website"
readonly WEBSITE_REPO="https://github.com/magicmatatjahu/website"

readonly BUILD_DIR="${KYMA_PROJECT_IO_DIR}/${WEBSITE_DIR}"

#"https://github.com/kyma-project/website"

# Colors
readonly RED='\033[0;31m'
readonly GREEN='\033[0;32m'
readonly YELLOW='\033[0;33m'
readonly NC='\033[0m' # No Color

pass() {
  local message="$1"
  echo -e "${GREEN}√ ${message}${NC}"
}

step() {
  local message="$1"
  echo -e "\\n${YELLOW}${message}${NC}"
}

copy-website-repo() {
  git clone -b "docs-community-preview" --single-branch "${WEBSITE_REPO}" "${WEBSITE_DIR}"
}

build-preview() {
  make -C "${BUILD_DIR}" netlify-community-preview
}

main() {
  step "Copying kyma/website repo"
  copy-website-repo
  pass "Copied"

  step "Building preview"
  build-preview
  pass "Builded"
}
main