#!/bin/bash

echo "Rollback initiated"

aws autoscaling cancel-instance-refresh \
  --auto-scaling-group-name starttech-asg

echo "Rollback completed"