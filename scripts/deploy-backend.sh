#!/bin/bash

echo "Starting backend deployment..."

aws autoscaling start-instance-refresh \
  --auto-scaling-group-name starttech-asg

echo "Deployment triggered."
