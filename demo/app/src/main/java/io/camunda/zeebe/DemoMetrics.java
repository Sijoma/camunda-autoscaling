package io.camunda.zeebe;

import io.camunda.zeebe.client.api.worker.JobWorkerMetrics;
import io.camunda.zeebe.config.WorkerCfg;
import io.micrometer.core.instrument.MeterRegistry;
import io.micrometer.core.instrument.Tags;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.concurrent.atomic.AtomicReference;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class DemoMetrics implements JobWorkerMetrics {
  private static final Logger LOGGER = LoggerFactory.getLogger(DemoMetrics.class);
  private final AtomicReference<Double> capacityPercentageRef = new AtomicReference<>((double) 0);
  private final AtomicInteger activatedJobCount = new AtomicInteger();
  private final JobWorkerMetrics delegate;
  private final int capacity;

  public DemoMetrics(final MeterRegistry registry, final WorkerCfg config) {
    final var tags = Tags.of("workerName", config.getWorkerName(), "jobType", config.getJobType());
    delegate = JobWorkerMetrics.micrometer().withMeterRegistry(registry).withTags(tags).build();
    capacity = config.getCapacity();

    registry.gauge(
        "zeebe.client.worker.job.capacity", tags, capacityPercentageRef, AtomicReference::get);
  }

  @Override
  public void jobActivated(int count) {
    delegate.jobActivated(count);
    LOGGER.trace("Activated {} jobs", count);
    updateCapacityPercentage(count);
  }

  @Override
  public void jobHandled(int count) {
    delegate.jobHandled(count);
    LOGGER.trace("Handled {} jobs", count);
    updateCapacityPercentage(-count);
  }

  private void updateCapacityPercentage(final int count) {
    final var activatedCount = this.activatedJobCount.addAndGet(count);
    final double capacityPercentage = (double) activatedCount / capacity;
    LOGGER.trace("Updating capacity to {} (count: {})", capacityPercentage * 100, activatedCount);
    capacityPercentageRef.set(capacityPercentage);
  }
}
