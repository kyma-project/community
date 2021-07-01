#!/usr/bin/env bash
# Script for build preview of this repo like in https://kyma-project.io/community/ on every PR.
# For more information, please contact with: @m00g3n @aerfio @magicmatatjahu

set -eo pipefail

on_error() {
    echo -e "${RED}✗ Failed${NC}"
    exit 1
}
trap on_error ERR

readonly KYMA_PROJECT_IO_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

readonly WEBSITE_DIR="website"
readonly WEBSITE_REPO="https://github.com/dbadura/website"

readonly BUILD_DIR="${KYMA_PROJECT_IO_DIR}/${WEBSITE_DIR}"

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


merge() {
  git config --global user.email "CI@kyma.project"
  git config --global user.name "CI"
  git checkout -B pull-request

  git checkout main
  step "Last commit from main"
  git log --max-count=1

  git merge pull-request
}
remove-cached-content() {
  ( rm -rf "${BUILD_DIR}" ) || true
}

copy-website-repo() {
  git clone -b "new-navigation-tree" --single-branch "${WEBSITE_REPO}" "${WEBSITE_DIR}"
}

build-preview() {
  export APP_COMMUNITY_SOURCE_DIR="${KYMA_PROJECT_IO_DIR}"/../docs
  export APP_PREVIEW_SOURCE_DIR="${BUILD_DIR}/content/community"
  make -C "${BUILD_DIR}" netlify-community-preview
}

main() {
  step "Remove website cached content"
  remove-cached-content
  pass "Removed"

  step "Merge with main branch"
  merge
  step "Merge done"

  step "Copying kyma/website repo"
  copy-website-repo
  pass "Copied"

  step "Delete old Content directory from website"
  rm -rf "${BUILD_DIR}"/content/community
  step "Deleted"

  step "Building preview"
  build-preview
  pass "Builded"

  step "Add Redirect rule"
  echo "/ /community/" > "${BUILD_DIR}"/public/_redirects
  step "Added"

  tree "${BUILD_DIR}"/content/community
  ls -l "${KYMA_PROJECT_IO_DIR}/.."
}
main
