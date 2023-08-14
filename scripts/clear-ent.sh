#!/bin/bash
cd ../ent
ls | grep -v generate.go | grep -v schema | xargs rm -rf