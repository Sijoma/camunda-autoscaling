FROM gcr.io/zeebe-io/starter:SNAPSHOT
LABEL org.opencontainers.image.source=https://github.com/Sijoma/camunda-autoscaling-hackday

# Copy over custom BPMN and start input
COPY sign-up-process.bpmn /app/resources/bpmn/demo/sign-up-process.bpmn
COPY sign-up-input.json /app/resources/bpmn/demo/sign-up-input.json
