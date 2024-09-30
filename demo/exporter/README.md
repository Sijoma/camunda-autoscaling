# Job Metric Exporter

This exporter tracks the count of available jobs per type as a gauge,
providing an absolute count per job type.

## Usage

Build the project:

```shell
mvn clean install -DskipTests -DskipChecks -T1C
```

Copy over the resulting uber JAR:

```shell
cp target/job-metric-exporter-uber.jar /usr/local/zeebe/exporters/job-metric-exporter.jar
```

Configure your broker with:

```yaml
zeebe:
  broker:
    exporters:
      job-metrics:
        className: "com.github.sijoma.exporter.JobMetricExporter"
        jarPath: "/usr/local/zeebe/exporters/job-metric-exporter.jar" 
```

The count is measured via the metric `zeebe.jobs.available` (or in
Prometheus, `zeebe_jobs_available`), with a `jobType` label and
`partitionId` label.

> [!Note]
> During transition, count may fluctuate because on restarting from snapshot,
> it will restart from whatever the count was at the time of the snapshot. On
> replaying events, it will then come back to the normal count.

## How it works

> [!Note]
> The count cannot be negative, and it is simply not tracked anymore once it hits 0.

Essentially, for every job type, there is a gauge. The gauge is incremented
whenever a new record with one of the following intents is exported:

- `JobIntent.CREATED`: this means a job was created, and is therefore
  immediately available.
- `JobIntent.TIMED_OUT`: this means a timed out, and is now available to be
  activated again.
- `JobIntent.YIELDED`: this means a job failed to be pushed out, so is now
  available to be activated.
- `JobIntent.RECURRED_AFTER_BACKOFF`: this means a job had failed, but was not
  made immediately available as it must wait for a back off before. This event
  is written once that backoff has expired.
- `JobIntent.FAILED`: only increment if the job has no recur-backoff property
  (see above), and there are some retries left, as it means the job is immediately
  available for activation.
- `IncidentIntent.RESOLVED`: only if this is a job related incident.

The gauge is decremented only when a job is activated, i.e. being worked on,
so on `JobBatchIntent.ACTIVATED`.

> [!Note]
> There are some inefficiencies, since with job push, you can have a sequence of
> `JobIntent.CREATED`, `JobBatchIntent.ACTIVATED`, `JobIntent.YIELDED`, for example,
> in rapid succession. This leads to unnecessary work, but that's acceptable for a demo.

### Incident problem

Since incident records only track the job key, we need to correlate the key to the job
type. For this, we have a LRU cache (job key => job type), capped at 10,000.

This means there is a small chance that an incident resolve intent does not increment
the count, which could lead to missed jobs depending on your scaling policy.

### Transitions/restarts

To cope with role transitions/restarts, the known counts are cached in memory and
persisted as the exporter metadata. The incident key-type cache is also serialized
and kept in the exporter state. Serialization format is JSONB.

## Limitations

As mentioned, there are some limitations to this approach via an exporter:

- The metric is potentially high cardinality, as the domain space of `jobType` is
  theoretically infinite.
    - This means we create one gauge per job type. Acceptable if you have a finite amount.
- The behavior on `IncidentIntent.RESOLVED` may not always succeed, e.g. the cache was
  not large enough.
- Serializing the counts and type cache in the exporter metadata is not a great pattern, because that
  contains user data (job types), which could be sent over unencrypted traffic (e.g. UDP). Additionally,
  that could become quite large, slowing down exporting.
- Serialization of the type cache is sub-optimal. We serialize each key - type pair,
  but it's very likely many keys share the same type, so we're writing out the same type
  over and over and over. Compression helps somewhat, but we could likely improve this.
- The exporter depends on the semantics of the engine, i.e. when a job is made activatable.
  Since new intents/records can be introduced, it's possible that newer versions would break
  the logic here.
 