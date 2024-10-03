FROM gcr.io/zeebe-io/worker:SNAPSHOT
LABEL org.opencontainers.image.source=https://github.com/Sijoma/camunda-autoscaling-hackday

# Copy over custom payload JSON
COPY ./worker-payload.json /app/resources/bpmn/demo/worker-payload.json
