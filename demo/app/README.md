# Demo App

The demo app extends the standard Zeebe team's benchmark application, with some code 
for the worker copied because the benchmark worker is not easily extensible.

The main goal of the worker is to add a new metric, `zeebe_client_worker_job_capacity`, which
reports a value between 0.0 and 1.0 representing the percentage of the job worker queue used
per job type and worker name.

This is added to ease integration with Horizontal Pod Autoscaling.

The percentage is computed by keeping track of the activated, non-handled job count, and
dividing it by the worker's configured capacity. 