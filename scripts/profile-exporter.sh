#!/bin/sh

aws-profile-picker

selected_profile="$(cat $HOME/.aws-profile-picker)"

if [ -z "$selected_profile" ]
then
  unset AWS_PROFILE
else
  export AWS_PROFILE="$selected_profile"
fi