#! /bin/bash

go generate ./...

PASSING_TESTS=(
	# "TestBasicSchemaToStruct"

)

# Default target if none provided
CURRENT_TARGET=${CURRENT_TARGET:-"TestBasicSchemaToStruct"}

# next up
# "TestNestedObjectSimple"
# "TestStringEnumSchemaToStruct"
# "TestIntegerEnumSchemaToStruct"
# "TestAllOfSchemaToStruct"
# "TestBasicRefSchemaToStruct"
# - TestOneOfSchemaToStruct
# - TestAnyOfSchemaToStruct

echo "----------------------------------------"
echo "RUNNING PASSING TESTS FIRST"
for test in "${PASSING_TESTS[@]}"; do
	echo "   - $test"
done
echo "----------------------------------------"

# tests that should pass before continuing
if [ ${#PASSING_TESTS[@]} -gt 0 ]; then
	./go test -v ./pkg/generator/... -run $(
		IFS='|'
		echo "${PASSING_TESTS[*]}"
	)
	if [ $? -ne 0 ]; then
		echo "WHOA! PASSING TESTS NOW FAILING! GO BACK AND FIX THESE TESTS FIRST"
		exit 1
	fi
fi

echo "----------------------------------------"
echo "RUNNING CURRENT TARGET"
echo "   - $CURRENT_TARGET"
echo "----------------------------------------"

# current target
./go test -v ./pkg/generator/... -run "$CURRENT_TARGET"

if [ $? -ne 0 ]; then
	echo "----------------------------------------"
	echo "IMPORTANT"
	echo "THE TEST COULD HAVE INCORRECT EXPECTATIONS"
	echo "TRY TO FIX THE SIMPLE PROBLEMS FIRST"
	echo "----------------------------------------"
	exit 1
fi
