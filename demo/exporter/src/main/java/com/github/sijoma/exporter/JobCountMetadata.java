package com.github.sijoma.exporter;

import com.dslplatform.json.CompiledJson;
import com.google.common.cache.Cache;
import com.google.common.cache.CacheBuilder;
import io.micrometer.core.instrument.Gauge;
import io.micrometer.core.instrument.MeterRegistry;
import io.micrometer.core.instrument.simple.SimpleMeterRegistry;
import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;
import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;
import net.jcip.annotations.ThreadSafe;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@ThreadSafe
public final class JobCountMetadata {
  private final Jsonb serializer = JsonbBuilder.newBuilder().build();
  private final ConcurrentMap<String, Integer> counts = new ConcurrentHashMap<>();
  private final Cache<Long, String> keyToTypeCache =
      CacheBuilder.newBuilder().maximumSize(10_000).concurrencyLevel(1).build();

  private final MeterRegistry meterRegistry;
  private final Logger logger;

  public JobCountMetadata(final MeterRegistry meterRegistry, final Logger logger) {
    this.meterRegistry = meterRegistry;
    this.logger = logger;
  }

  public static void main(final String[] args) {
    final var count =
        new JobCountMetadata(
            new SimpleMeterRegistry(), LoggerFactory.getLogger(JobCountMetadata.class));
    count.increment(1, "test");
    count.increment(2, "test");
    count.increment(3, "test");
    count.increment(4, "foo");
    final var serialized = count.serialize();
    final var deser =
        new JobCountMetadata(
            new SimpleMeterRegistry(), LoggerFactory.getLogger(JobCountMetadata.class));
    deser.deserialize(serialized);

    System.out.println(count);
    System.out.println(deser);
  }

  public void increment(final long jobKey, final String jobType) {
    final var actual = counts.get(jobType);
    if (actual == null) {
      counts.put(jobType, 1);
      monitorJobType(jobType);
    } else {
      counts.put(jobType, actual + 1);
    }

    keyToTypeCache.put(jobKey, jobType);
  }

  public void incrementIncident(final long jobKey) {
    final var type = keyToTypeCache.getIfPresent(jobKey);
    if (type == null) {
      logger.debug(
          "Failed to increment job count for incident since key [{}] is not present in cache",
          jobKey);
      return;
    }

    increment(jobKey, type);
  }

  public void decrement(final String jobType, final int count) {
    final var actual = counts.get(jobType);
    if (actual != null && actual <= count) {
      counts.remove(jobType);
      return;
    }

    counts.computeIfPresent(
        jobType,
        (key, value) -> {
          final var newValue = value - count;
          if (newValue < 0) {
            return null;
          }
          return newValue <= 0 ? null : newValue;
        });
  }

  public void clear() {
    counts.clear();
  }

  public byte[] serialize() {
    final var output = new ByteArrayOutputStream();
    serializer.toJson(new State(keyToTypeCache.asMap(), counts), output);
    return output.toByteArray();
  }

  public JobCountMetadata deserialize(final byte[] bytes) {
    final var state = serializer.fromJson(new ByteArrayInputStream(bytes), State.class);

    counts.clear();
    keyToTypeCache.invalidateAll();

    counts.putAll(state.counts);
    keyToTypeCache.putAll(state.typeCache);

    return this;
  }

  @Override
  public String toString() {
    return "JobCountMetadata{"
        + "counts="
        + counts
        + ", keyToTypeCache="
        + keyToTypeCache.asMap()
        + '}';
  }

  private void monitorJobType(final String jobType) {
    Gauge.builder("zeebe.jobs.available", () -> counts.getOrDefault(jobType, 0))
        .baseUnit("job")
        .tag("jobType", jobType)
        .description("Count of available jobs per type")
        .register(meterRegistry);
  }

  @CompiledJson
  public record State(Map<Long, String> typeCache, Map<String, Integer> counts) {}
}
