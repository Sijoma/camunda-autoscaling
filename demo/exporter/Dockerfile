# An init container which will copy the exporter JAR into /exporters/job-metric-exporter.jar
FROM ubuntu:noble
COPY ./target/job-metric-exporter-uber.jar /app/job-metric-exporter.jar

RUN groupadd --gid 1001 camunda && useradd --system --gid 1001 --uid 1001 camunda
USER 1001:1001

CMD ["bash", "-c", "cp /app/job-metric-exporter.jar /exporters/job-metric-exporter.jar"]