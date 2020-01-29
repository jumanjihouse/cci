# shellcheck shell=sh
python_path="$(python3 -c "import site; print(site.USER_BASE)")"
readonly python_path

if ! printenv PATH | grep ":${python_path}/bin[:$]" >/dev/null 2>&1; then
  export PATH="${PATH}:${python_path}/bin"
fi
