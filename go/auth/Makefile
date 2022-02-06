##
# Makefile for development
#
# @file
# @version 0.1


# end

clean:
  rm -r .aws-sam

invoke-build:
  sam build && sam local invoke

test-manual:
  sam build && sam local start-api
