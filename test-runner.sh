#! /bin/bash

PASSING_TESTS=(
	"TestNestedObjectSimple"
	"TestBasicSchemaToStruct"
	"TestStringEnumSchemaToStruct"
	"TestIntegerEnumSchemaToStruct"
	"TestAllOfSchemaToStruct"
)

CURRENT_TARGET="TestAllOfWithRefsSchemaToStruct"

# next up
# -
# - TestOneOfSchemaToStruct
# - TestAnyOfSchemaToStruct

echo "----------------------------------------"
echo "RUNNING PASSING TESTS FIRST"
echo "----------------------------------------"
# tests that should pass before continuing
./go test -v -testname ./pkg/generator/... -run $(
	IFS='|'
	echo "${PASSING_TESTS[*]}"
)
if [ $? -ne 0 ]; then
	echo "WHOA! PASSING TESTS NOW FAILING! GO BACK AND FIX THESE TESTS FIRST"
	exit 1
fi
echo "----------------------------------------"
echo "RUNNING CURRENT TARGET"
echo "----------------------------------------"

# current target
./go test -v -testname ./pkg/generator/... -run $(
	IFS='|'
	echo "$CURRENT_TARGET"
)

if [ $? -ne 0 ]; then
	echo "----------------------------------------"
	echo "IMPORTANT"
	echo "THE TEST COULD HAVE INCORRECT EXPECTATIONS"
	echo "TRY TO FIX THE SIMPLE PROBLEMS FIRST"
	echo "----------------------------------------"
	exit 1
fi
