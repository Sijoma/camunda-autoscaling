package com.github.sijoma.exporter;

import io.camunda.zeebe.exporter.api.Exporter;
import io.camunda.zeebe.exporter.api.context.Context;
import io.camunda.zeebe.exporter.api.context.Controller;
import io.camunda.zeebe.protocol.record.Record;
import io.camunda.zeebe.protocol.record.RecordType;
import io.camunda.zeebe.protocol.record.ValueType;
import io.camunda.zeebe.protocol.record.intent.IncidentIntent;
import io.camunda.zeebe.protocol.record.intent.JobBatchIntent;
import io.camunda.zeebe.protocol.record.intent.JobIntent;
import io.camunda.zeebe.protocol.record.value.IncidentRecordValue;
import io.camunda.zeebe.protocol.record.value.JobBatchRecordValue;
import io.camunda.zeebe.protocol.record.value.JobRecordValue;
import io.micrometer.core.instrument.composite.CompositeMeterRegistry;
import java.util.EnumSet;
import java.util.Set;

public final class JobMetricExporter implements Exporter {
  private static final Set<ValueType> VALUE_TYPES =
      EnumSet.of(ValueType.JOB, ValueType.JOB_BATCH, ValueType.INCIDENT);

  private final CompositeMeterRegistry meterRegistry = new CompositeMeterRegistry();

  private JobCountMetadata jobCounters;
  private Controller controller;

  @Override
  public void configure(final Context context) throws Exception {
    meterRegistry
        .add(context.getMeterRegistry())
        .config()
        .commonTags("partitionId", String.valueOf(context.getPartitionId()));
    context.setFilter(
        new Context.RecordFilter() {
          @Override
          public boolean acceptType(RecordType recordType) {
            return recordType == RecordType.EVENT;
          }

          @Override
          public boolean acceptValue(ValueType valueType) {
            return VALUE_TYPES.contains(valueType);
          }
        });

    jobCounters = new JobCountMetadata(meterRegistry, context.getLogger());
  }

  @Override
  public void open(final Controller controller) {
    this.controller = controller;
    controller.readMetadata().ifPresent(jobCounters::deserialize);
  }

  @Override
  public void close() {
    meterRegistry.close();
    jobCounters.clear();
  }

  @SuppressWarnings("unchecked")
  @Override
  public void export(final Record<?> record) {
    if (record.getIntent() instanceof JobBatchIntent intent) {
      final var jobRecord = (Record<JobBatchRecordValue>) record;
      if (intent == JobBatchIntent.ACTIVATED) {
        jobCounters.decrement(
            jobRecord.getValue().getType(), jobRecord.getValue().getJobs().size());
      }
    } else if (record.getIntent() instanceof JobIntent intent) {
      final var jobRecord = (Record<JobRecordValue>) record;
      final var type = (jobRecord).getValue().getType();
      switch (intent) {
        case CREATED:
        case TIMED_OUT:
        case YIELDED:
        case RECURRED_AFTER_BACKOFF:
          jobCounters.increment(jobRecord.getKey(), type);
          break;
        case FAILED:
          // on failed, only increment if we would retry immediately; otherwise wait for the recur
          // event
          if (jobRecord.getValue().getRetries() > 0
              && jobRecord.getValue().getRetryBackoff() <= 0) {
            jobCounters.increment(jobRecord.getKey(), type);
          }
          break;
        default:
          break;
      }
    } else if (record.getIntent() instanceof IncidentIntent intent) {
      final var incidentRecord = (Record<IncidentRecordValue>) record;
      final var jobKey = incidentRecord.getValue().getJobKey();
      if (intent == IncidentIntent.RESOLVED && jobKey > 0) {
        jobCounters.incrementIncident(jobKey);
      }
    }

    controller.updateLastExportedRecordPosition(record.getPosition(), jobCounters.serialize());
  }
}
