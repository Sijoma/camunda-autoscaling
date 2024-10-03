FROM gcr.io/zeebe-io/worker:SNAPSHOT

# Copy over custom payload JSON
COPY ./worker-payload.json /app/resources/bpmn/demo/worker-payload.json