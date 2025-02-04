#! /bin/bash

set -x

go generate ./...

PASSING_TESTS=(
	"basic_schema_to_struct"
	"allof_schema_to_struct"
	"basic_ref_schema_to_struct"
)

CURRENT_TARGET="allof_with_refs_schema_to_struct"

# tests that should pass before continuing
if [ ${#PASSING_TESTS[@]} -gt 0 ]; then
	strs=""
	echo "----------------------------------------"
	echo "RUNNING PASSING TESTS FIRST (go-code only)"
	for test in "${PASSING_TESTS[@]}"; do
		echo "   - $test"
		if [ -z "$strs" ]; then
			strs="TestAll/$test/go-code"
		else
			strs="$strs|TestAll/$test/go-code"
		fi
	done
	echo "----------------------------------------"

	./go test -v ./pkg/generator -run "$strs"
	if [ $? -ne 0 ]; then
		echo "WHOA! PASSING TESTS NOW FAILING! GO BACK AND FIX THESE TESTS FIRST"
		exit 1
	fi
fi

echo "----------------------------------------"
echo "RUNNING CURRENT TARGET (all tests)"
echo "   - $CURRENT_TARGET"
echo "----------------------------------------"

# current target - run all tests
./go test -v ./pkg/generator -run "TestAll/$CURRENT_TARGET"

if [ $? -ne 0 ]; then
	echo "----------------------------------------"
	echo "IMPORTANT"
	echo "THE TEST COULD HAVE INCORRECT EXPECTATIONS"
	echo "TRY TO FIX THE SIMPLE PROBLEMS FIRST"
	echo "----------------------------------------"
	exit 1
fi
