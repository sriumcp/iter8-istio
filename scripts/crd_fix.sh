# !/bin/bash
#
# A script to fix the bug in the controller

SCRIPTDIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
FILE_PATH=$SCRIPTDIR"/../config/crds/iter8_v1alpha1_canary.yaml"
suffix=".original"
line=$(grep 'lastTransitionTime' -n config/crds/iter8_v1alpha1_canary.yaml | sed 's/:.*//')
line=$(( $line + 1 ))

sed -i$suffix  "${line}s/object/string/" $FILE_PATH
rm $FILE_PATH$suffix
