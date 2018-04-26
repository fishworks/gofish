#!/usr/bin/env bash

# Copyright (C) 2016, Matt Butcher

# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:

# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.

# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.
#
# Ripped from github.com/technosophos/helm-template's get-binary.sh script, with a few tweaks to fetch gofish.

PROJECT_NAME="gofish"
PROJECT_GH="fishworks/gofish"

: ${INSTALL_PREFIX:="/usr/local/bin"}
: ${VERSION:="v0.5.0"}

if [[ $SKIP_BIN_INSTALL == "1" ]]; then
  echo "Skipping binary install"
  exit
fi

# initArch discovers the architecture for this system.
initArch() {
  ARCH=$(uname -m)
  case $ARCH in
    armv5*) ARCH="armv5";;
    armv6*) ARCH="armv6";;
    armv7*) ARCH="armv7";;
    aarch64) ARCH="arm64";;
    x86) ARCH="386";;
    x86_64) ARCH="amd64";;
    i686) ARCH="386";;
    i386) ARCH="386";;
  esac
}

# initOS discovers the operating system for this system.
initOS() {
  OS=$(echo `uname`|tr '[:upper:]' '[:lower:]')

  case "$OS" in
    # Msys support
    msys*) OS='windows';;
    # Minimalist GNU for Windows
    mingw*) OS='windows';;
  esac
}

# verifySupported checks that the os/arch combination is supported for
# binary builds.
verifySupported() {
  local supported="linux-386\nlinux-amd64\nlinux-arm\nlinux-arm64\nlinux-ppc64le\ndarwin-amd64\nwindows-amd64"
  if ! echo "${supported}" | grep -q "${OS}-${ARCH}"; then
    echo "No prebuilt binary for ${OS}-${ARCH}."
    exit 1
  fi

  if ! type "curl" > /dev/null && ! type "wget" > /dev/null; then
    echo "Either curl or wget is required"
    exit 1
  fi
}

# getDownloadURL checks the latest available version.
getDownloadURL() {
  DOWNLOAD_URL="https://gofi.sh/releases/$PROJECT_NAME-$VERSION-$OS-$ARCH.tar.gz"
}

# downloadFile downloads the latest binary package and also the checksum
# for that binary.
downloadFile() {
  TMP_CACHE_FILE="/tmp/${PROJECT_NAME}.tgz"
  echo "Downloading $DOWNLOAD_URL"
  if type "curl" > /dev/null; then
    curl -L "$DOWNLOAD_URL" -o "$TMP_CACHE_FILE"
  elif type "wget" > /dev/null; then
    wget -q -O "$TMP_CACHE_FILE" "$DOWNLOAD_URL"
  fi
}

# installFile verifies the SHA256 for the file, then unpacks and
# installs it.
installFile() {
  TMPDIR="/tmp/$PROJECT_NAME"
  mkdir -p "$TMPDIR"
  tar xf "$TMP_CACHE_FILE" -C "$TMPDIR"
  TMPDIR_BIN="$TMPDIR/$OS-$ARCH/$NAME"
  echo "Preparing to install into ${INSTALL_PREFIX}"
  # Use * to also copy the file withe the exe suffix on Windows
  mkdir -p "$INSTALL_PREFIX"
  sudo cp "$TMPDIR_BIN/$PROJECT_NAME" "$INSTALL_PREFIX"
}

# fail_trap is executed if an error occurs.
fail_trap() {
  result=$?
  if [ "$result" != "0" ]; then
    echo -e "!!!\tFailed to install $PROJECT_NAME"
    echo -e "!!!\tFor support, go to https://github.com/$PROJECT_GH."
  fi
  exit $result
}

# testVersion tests the installed client to make sure it is working.
testVersion() {
  set +e
  echo "$PROJECT_NAME installed into $INSTALL_PREFIX/$PROJECT_NAME"
  echo "Run '$PROJECT_NAME init' to get started!"
  # To avoid to keep track of the Windows suffix,
  # call the plugin assuming it is in the PATH
  PATH=$PATH:$INSTALL_PREFIX
  gofish -h > /dev/null
  set -e
}

# Execution

#Stop execution on any error
trap "fail_trap" EXIT
set -e
initArch
initOS
verifySupported
getDownloadURL
downloadFile
installFile
testVersion


# -----BEGIN PGP PUBLIC KEY BLOCK-----

# mQINBFp8op4BEADDk9QQYbTaq+nXTvxxTyjBqmTS3CsN94y8TfxslVLbQym3wuT5
# 9zWUv+JRlIZoqZiXJvrXFuJnUgTuzniYkrLtxvUWRKY8cISqcuA26d40YuBeQXWl
# TvHAeDiJ3vyLRbS38/tzcEYThojuc0CMIklqDwrwI2J3DAegkfc4jAB60Q9oYo/M
# IlsRxS1jaLMoe3xUFJV8Lq9BQOehqNSpL+L0lCHsXmLJjczuAE+pYReIuAlyeke5
# o5/t4fdEqc5hziTN7XJGF5qAHK4+iEuzYx1M1UHuHqCOzZX8P/KOXT90Iru8HZ11
# r/NueLUHIK/PALFBd8tC1zu03pyEKBgSd/qdsyFvIlJHQIEUOe66RaBOLhhBm9o7
# 97dWceYrktl0xIJxLRlcwQAb8/+xivDsQAdERU3xwVbWk9UQJx8Jy9fjI8bvJF/n
# 3XKvGIX/FZVzgicZExJVbN/braP9obHswvUe5mKDwsllQPz7FuTjoaxAj+Px9iPt
# l5jLzk++pjEXSmbljeW6lDBh4T997szZTs8Vp1LdJprpZZcyOmA5ct/uilNsXJaY
# 8K1agfHNTb5n+wDwrVnepwPo7bpwONgktzgnvxa6Vhi4UEvE5JXx6R9YFM8cpS61
# TDVAg82i0PGZa98RKxRs/NH2ynUsplhkzIDbbDWZZPYCwp5EMAk0i/NyuwARAQAB
# tCpNYXR0aGV3IEZpc2hlciA8bWF0dC5maXNoZXJAbWljcm9zb2Z0LmNvbT6JAjgE
# EwECACIFAlp8op4CGwMGCwkIBwMCBhUIAgkKCwQWAgMBAh4BAheAAAoJEM3sZ2h+
# +qNOo50P/040uUwZ4FE7iHkwn7Da2Nqbe1hvadcvWUfVkXxCHTKXnAd3RLeEmPBE
# a+vkhNPA3vd57EGqDfBw2zdMVx+jbKgrUxwUrn506TgwnHFq+gaqm+JVQQrEz0/Z
# ZeO7xqlbSrwHP5jwMKFPCUF6QdBuyCqH4GnUiCMa69Vp3hEx1bGqEPeSAMWoBq1k
# i5H8A/oB/9498J8SSNGOhCN0igG50IgQw1bHoHvWMmsbNslFHnQ8s+M6q0LZzC6w
# 6gUnntvcNouWZqggMyBDh9QT8APH2c1iRuwPZU0t821O7RX9kzjDfDPpC/9QLW9c
# 2nEF5QU0EGF7HCd+7feqQHiDZT7B/VEXoWt89j7Fz8lDLs+jXl+G8MG+0F+tNcq6
# YkdNjTDw7HG4TQnreUnb+pKbRSb7By3gQA59ovwtRhpLCwIXLjO5YyYvyzjlKrQt
# 0jYtkwBCrulKq2EnF6lsg1eowTNQ2FaGBZZCCT46IDT1z39bSfKxHA0lZyE5FApS
# QKFqTkNnspIMLBpblYAyYeF4wLwyHmKXUZioTq09fd7QXWKUwIj5q5xJHBDtwlly
# HcGTjb6mJACsK46hBL5LIlYqCiJmaLilWQ02K+drVBAqFnQrGI7stH/8S/yP6+PQ
# EUjjmbDAxuplLhlo094qqmpPDJxBoUnEOdbERppKXrlUZf6HeOkluQINBFp8op4B
# EADdlBwTZ8XKGJUf0oINAQ7/4GZJCinBGWKHRmM+HrWsB+O13FEeWvGHNJQ5hswa
# aKaIylq4+uF8nnpf3OV4zEfZEguECIzqHo78JUAGRbVREnqzBN7eLQHeAnyvt/yd
# veEza1Uu/G/Gyi/QfL+nGfInjPcsqOY6kWfhlSYGLWMmi86cBwAHSxxCgRjLGZtB
# lID1qLXPeCZyqyvPLMuENIt01I6qnFQL+RLaXFpCjeXCF2XY5lT6VZ9DPJxmb/dE
# ZHFCZeplCPBIJd7ihMPO4j8gen7nuh1n6egU6h+aWJ4APtwkAQi8ESOM8yGTgPnW
# AFqEAIkVdR9erf894LZ6WKfuPKhtQULoVIBMM0idmmccqlejG6DlFsCua1i1Mce7
# xgfqRd1GxW6U0NzY4BPdyU91ggc+hhkUmpfB4F/ggZOs8N4ckiDFnXkN6ShVpXNt
# eMcaltsv5XpHqsBBy6vixn9vMdWavViBGLDapiloba97CbiHLajGuUDeyOHZSdCK
# 38ecMOfVHhWGI5QOb0ayDwI4pSsBmxdGc/mNsMLNS5ScTJBIsZKMKMnPIUzAFgZp
# /c4dCk/+cqAssejoX24UElkmoucrFJtxIt0blDlODybm8kTSHf0gVOYuK41OvbkJ
# cz5tLIUfZMRdlMfz7wSGVXamITjoEwSQ9uzD2fXA8ui1pQARAQABiQIfBBgBAgAJ
# BQJafKKeAhsMAAoJEM3sZ2h++qNOj/gQAI1MN167Gz9Zb4uzep+GzYSB2Bp6AGnw
# VHAIVFTTuUfz9yUa2Cv+nVuUdh4W0CHfwWUbaUqlP9NiYiMi/tcwgSiCdeCK1G/x
# PmJyyp17bfOiFhw042kbpX2yp2N2e3IUfNjOOyxh+X44rVcLLAuwIqVpDG8RCn6b
# /wBfXc2YCq4wWX+QYqkRrb3Qd0cOVfGDyyaSiLrj20oi9KO0PjwZ1qKRlk7xBC6R
# /0IvZwwFRVJfFbLlmCHmGkaSwpwvMLumN/wlV3T8Px/XGIYuK7RxhslZhB2b75Sj
# xnF8TfQPyTRpYVscujad0HdPd9uNrsQxrOsXwzMIyyJ+Nsn5IEt8jyMFS7mw3orW
# V9nmxBrdBUlUzxiCuwwzHrxdUXRwA6gJVNS5rk8m1tpIl+JNao66i3fGbYSSvGV4
# wqOwwpl0gLhnx5Xy3lwfrdLIDSgRDZ0i815Pcr4Q+j4O/PJXWt6J83aH8FiCTEuM
# BAjo8CuHR37jaO7ZliGsJW/kWVntM8uW7C1r+46sqG7nBlHmNW/E79/C9XqtT0Rp
# NOkfRaUyYKkjSgckeM7VzNCT4iLIUPUHGm+YmwHuh83IMRNuoELlMpoBr7DgqUtY
# N6lsOkvlvsLIL+TzuWb0sUKbtEUh1jiQUDtzGig1LGsimRBevH2IvHgBWxzPVpw6
# Qz0L9+I2ai68
# =dBiU
# -----END PGP PUBLIC KEY BLOCK-----
