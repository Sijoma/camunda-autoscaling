package com.github.sijoma.exporter;

import io.micrometer.core.instrument.Gauge;
import io.micrometer.core.instrument.MeterRegistry;
import java.nio.ByteOrder;
import java.nio.charset.Charset;
import java.nio.charset.StandardCharsets;
import java.util.Objects;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;
import net.jcip.annotations.ThreadSafe;
import org.agrona.BufferUtil;
import org.agrona.collections.MutableInteger;
import org.agrona.concurrent.UnsafeBuffer;

@ThreadSafe
final class JobCountMetadata {
  private static final ByteOrder ENDIANNESS = ByteOrder.LITTLE_ENDIAN;
  private static final Charset CHARSET = StandardCharsets.UTF_8;
  private final ConcurrentMap<String, Integer> counts = new ConcurrentHashMap<>();
  private final MeterRegistry meterRegistry;

  JobCountMetadata(MeterRegistry meterRegistry) {
    this.meterRegistry = meterRegistry;
  }

  public void increment(final String jobType) {
    final var actual = counts.get(jobType);
    if (actual == null) {
      counts.put(jobType, 1);
      monitorJobType(jobType);
    } else {
      counts.put(jobType, actual + 1);
    }
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
    final var length = length();

    if (length == 0) {
      return BufferUtil.NULL_BYTES;
    }

    final var bytes = new byte[length];
    final var buffer = new UnsafeBuffer(bytes);
    final var offset = new MutableInteger(0);
    counts.forEach(
        (type, count) -> offset.value += serializeCount(buffer, offset.value, type, count));

    assert offset.value == length : "expected total bytes serialized to equal predicted length";
    return bytes;
  }

  public void deserialize(final byte[] bytes) {
    final var buffer = new UnsafeBuffer(bytes);
    counts.clear();

    if (bytes.length == 0) {
      return;
    }

    int offset = 0;
    while (offset < bytes.length) {
      final var type = buffer.getStringUtf8(offset, ENDIANNESS);
      offset += Integer.BYTES + getStringLength(type);

      final var count = buffer.getInt(offset, ENDIANNESS);
      offset += Integer.BYTES;

      counts.put(type, count);
      monitorJobType(type);
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) return true;
    if (o == null || getClass() != o.getClass()) return false;
    JobCountMetadata that = (JobCountMetadata) o;
    return Objects.equals(counts, that.counts);
  }

  @Override
  public int hashCode() {
    return Objects.hashCode(counts);
  }

  @Override
  public String toString() {
    return "JobCountMetadata{" + "counts=" + counts + '}';
  }

  private int serializeCount(
      final UnsafeBuffer buffer, final int offset, final String type, final int count) {
    int bytesWritten = buffer.putStringUtf8(offset, type, ENDIANNESS);
    buffer.putInt(offset + bytesWritten, count, ENDIANNESS);
    return bytesWritten + Integer.BYTES;
  }

  private int length() {
    final var length = new MutableInteger(0);
    counts.forEach((type, ignored) -> length.value += 2 * Integer.BYTES + getStringLength(type));
    return length.value;
  }

  private int getStringLength(final String type) {
    return type.getBytes(CHARSET).length;
  }

  private void monitorJobType(final String jobType) {
    Gauge.builder("zeebe.jobs.available", () -> counts.getOrDefault(jobType, 0))
        .baseUnit("job")
        .tag("jobType", jobType)
        .description("Count of available jobs per type")
        .register(meterRegistry);
  }
}
