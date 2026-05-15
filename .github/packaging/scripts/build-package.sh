#!/usr/bin/env bash

# Build and sign a Linux package with nfpm inside the container.
#
# Required env vars:
#   NFPM_PACKAGER  - nfpm packager: "rpm" or "deb"
#   PACKAGE        - "avalanchego" or "subnet-evm"
#   VERSION        - Semantic version without "v" prefix (e.g., "1.14.1")
#   TAG            - Git tag (e.g., "v1.14.1")
#   PACKAGE_ARCH   - Architecture (x86_64/aarch64 for RPM, amd64/arm64 for DEB)
#   OUTPUT_DIR     - Directory for the output package (bind-mounted from host)
#
# Optional env vars:
#   GPG_KEY_FILE                              - Path to GPG private key
#   NFPM_RPM_PASSPHRASE / NFPM_DEB_PASSPHRASE  - GPG passphrase
#   AVALANCHEGO_COMMIT                         - Git commit hash (auto-detected if not set)

set -euo pipefail

: "${PACKAGE:?PACKAGE must be set (avalanchego or subnet-evm)}"
: "${VERSION:?VERSION must be set}"
: "${TAG:?TAG must be set}"
: "${PACKAGE_ARCH:?PACKAGE_ARCH must be set}"
: "${OUTPUT_DIR:?OUTPUT_DIR must be set}"
: "${NFPM_PACKAGER:?NFPM_PACKAGER must be set (rpm or deb)}"

NFPM_PACKAGER="${NFPM_PACKAGER,,}"
pkg_format_upper="${NFPM_PACKAGER^^}"

REPO_ROOT="/build"
PACKAGING_DIR="${REPO_ROOT}/.github/packaging"
NFPM_CONFIG_TEMPLATE="${PACKAGING_DIR}/nfpm/${PACKAGE}-${NFPM_PACKAGER}.yml"
NFPM_CONFIG_RESOLVED="${REPO_ROOT}/build/${PACKAGE}-${NFPM_PACKAGER}-resolved.yml"
if [[ ! -f "${NFPM_CONFIG_TEMPLATE}" ]]; then
    echo "Unknown nfpm packager or package: ${NFPM_PACKAGER} / ${PACKAGE}" >&2
    exit 1
fi

# shellcheck disable=SC1091
source "${PACKAGING_DIR}/scripts/lib-build-common.sh"

# Well-known paths referenced by nfpm configs
export NFPM_CHANGELOG="${REPO_ROOT}/build/nfpm-changelog.yml"
export NFPM_SIGNING_KEY="${REPO_ROOT}/build/gpg/signing-key.asc"
export NFPM_RPM_PASSPHRASE="${NFPM_RPM_PASSPHRASE:-}"
export NFPM_DEB_PASSPHRASE="${NFPM_DEB_PASSPHRASE:-}"

echo "=== Building ${PACKAGE} ${pkg_format_upper} for ${PACKAGE_ARCH} (tag: ${TAG}) ==="

init_build_env
build_binary "${PACKAGE}"
generate_changelog "${VERSION}"

# ── GPG signing ───────────────────────────────────────────────────

GPG_KEY_FILE="${GPG_KEY_FILE:-}"
case "${NFPM_PACKAGER}" in
    rpm)
        GPG_PASSPHRASE_ENV="NFPM_RPM_PASSPHRASE"
        ;;
    deb)
        GPG_PASSPHRASE_ENV="NFPM_DEB_PASSPHRASE"
        ;;
    *)
        echo "Unsupported nfpm packager: ${NFPM_PACKAGER}" >&2
        exit 1
        ;;
esac

GPG_PUBLIC_KEY="${OUTPUT_DIR}/${pkg_format_upper}-GPG-KEY-avalanchego"

# Ephemeral keys use a known throwaway passphrase so local and CI builds
# exercise passphrase handling without release credentials.
if [[ -z "${GPG_KEY_FILE}" ]]; then
    use_ephemeral_gpg_passphrase "${GPG_PASSPHRASE_ENV}"
fi

setup_gpg "${GPG_KEY_FILE}" "${GPG_PUBLIC_KEY}" "${pkg_format_upper}"

# ── Package with nfpm ─────────────────────────────────────────────

export VERSION PACKAGE_ARCH BINARY_PATH

PKG_FILENAME="${PACKAGE}-${TAG}-${PACKAGE_ARCH}.${NFPM_PACKAGER}"
PKG_PATH="${OUTPUT_DIR}/${PKG_FILENAME}"

run_nfpm_package \
    "${NFPM_CONFIG_TEMPLATE}" \
    "${NFPM_CONFIG_RESOLVED}" \
    "${NFPM_PACKAGER}" \
    "${PKG_PATH}"

echo "${pkg_format_upper} built: ${PKG_PATH}"
