#!/usr/bin/env bash
set -euo pipefail

# if first argument is "test", use gotestsum
if [ "${1:-}" == "test" ]; then
	shift

	cc=0
	ff=0
	testname=0
	pnametfails=0
	real_args=()
	extra_args=""
	format="pkgname"
	# Handle each argument
	for arg in "$@"; do
		if [ "$arg" = "-custom-coverage" ]; then
			cc=1
		elif [ "$arg" = "-force" ]; then
			ff=1
		elif [ "$arg" = "-testname" ]; then
			testname=1
		elif [ "$arg" = "-pnametfails" ]; then
			pnametfails=1
		else
			real_args+=("$arg")
		fi
	done

	if [[ "$cc" == "1" ]]; then
		tmpcoverdir=$(mktemp -d)
		function print_coverage() {
			echo "================================================"
			echo "Function Coverage"
			echo "------------------------------------------------"
			go tool cover -func=$tmpcoverdir/coverage.out
			echo "================================================"

		}
		extra_args=" -coverprofile=$tmpcoverdir/coverage.out -covermode=atomic "
		trap "print_coverage" EXIT
	fi

	if [[ "$ff" == "1" ]]; then
		extra_args="$extra_args -count=1 "
	fi

	if [[ "$testname" == "1" ]]; then
		format="testname"
	fi

	if [[ "$pnametfails" == "1" ]]; then
		format="pkgname-and-test-fails"
	fi

	./scripts/run-tool.sh gotestsum \
		--format $format \
		--format-icons hivis \
		--hide-summary=skipped \
		--raw-command -- go test -vet=all -json -cover $extra_args "${real_args[@]}"
	exit $?
fi

# otherwise run go directly with all arguments
exec go "$@"
