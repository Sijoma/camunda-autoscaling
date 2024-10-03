package io.camunda.zeebe;

import io.camunda.zeebe.App;
import io.camunda.zeebe.ResponseChecker;
import io.camunda.zeebe.Worker;
import io.camunda.zeebe.client.ZeebeClient;
import io.camunda.zeebe.client.ZeebeClientBuilder;
import io.camunda.zeebe.client.api.worker.JobHandler;
import io.camunda.zeebe.client.api.worker.JobWorker;
import io.camunda.zeebe.client.api.worker.JobWorkerMetrics;
import io.camunda.zeebe.config.AppCfg;
import io.camunda.zeebe.config.WorkerCfg;
import io.camunda.zeebe.util.logging.ThrottledLogger;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;

public final class DemoWorker extends Worker {
  public static final Logger LOGGER = LoggerFactory.getLogger(DemoWorker.class);
  private static final Logger THROTTLED_LOGGER = new ThrottledLogger(LOGGER, Duration.ofSeconds(5));

  private final AppCfg appCfg;
  private final WorkerCfg workerCfg;

  DemoWorker(final AppCfg appCfg) {
    super(appCfg);
    this.appCfg = appCfg;
    this.workerCfg = appCfg.getWorker();
  }

  public static void main(final String[] args) {
    createApp(DemoWorker::new);
  }

  @Override
  public void run() {
    final String jobType = workerCfg.getJobType();
    final long completionDelay = workerCfg.getCompletionDelay().toMillis();
    final boolean isStreamEnabled = workerCfg.isStreamEnabled();
    final var variables = readVariables(workerCfg.getPayloadPath());
    final BlockingQueue<Future<?>> requestFutures = new ArrayBlockingQueue<>(10_000);
    final ZeebeClient client = createZeebeClient();
    final JobWorkerMetrics metrics = new DemoMetrics(prometheusRegistry, workerCfg);
    printTopology(client);

    final JobWorker worker =
        client
            .newWorker()
            .jobType(jobType)
            .handler(handleJob(client, variables, completionDelay, requestFutures))
            .streamEnabled(isStreamEnabled)
            .metrics(metrics)
            .open();

    final ResponseChecker responseChecker = new ResponseChecker(requestFutures);
    responseChecker.start();

    Runtime.getRuntime()
        .addShutdownHook(
            new Thread(
                () -> {
                  worker.close();
                  client.close();
                  responseChecker.close();
                }));
  }

  private JobHandler handleJob(
      final ZeebeClient client,
      final String variables,
      final long completionDelay,
      final BlockingQueue<Future<?>> requestFutures) {
    return (jobClient, job) -> {
      final long startHandlingTime = System.currentTimeMillis();
      if (workerCfg.isSendMessage()) {
        final var correlationKey =
            job.getVariable(workerCfg.getCorrelationKeyVariableName()).toString();

        final boolean messagePublishedSuccessfully = publishMessage(client, correlationKey);
        if (!messagePublishedSuccessfully) {
          return;
        }
      }

      final var command = jobClient.newCompleteCommand(job.getKey()).variables(variables);
      addDelayToCompletion(completionDelay, startHandlingTime);
      requestFutures.add(command.send());
    };
  }

  private ZeebeClient createZeebeClient() {
    final var timeout =
        workerCfg.getTimeout() != Duration.ZERO
            ? workerCfg.getTimeout()
            : workerCfg.getCompletionDelay().multipliedBy(6);
    final ZeebeClientBuilder builder =
        ZeebeClient.newClientBuilder()
            .gatewayAddress(appCfg.getBrokerUrl())
            .numJobWorkerExecutionThreads(workerCfg.getThreads())
            .defaultJobWorkerName(workerCfg.getWorkerName())
            .defaultJobTimeout(timeout)
            .defaultJobWorkerMaxJobsActive(workerCfg.getCapacity())
            .defaultJobPollInterval(workerCfg.getPollingDelay())
            .withProperties(System.getProperties())
            .withInterceptors(monitoringInterceptor);

    if (!appCfg.isTls()) {
      builder.usePlaintext();
    }

    return builder.build();
  }

  private boolean publishMessage(final ZeebeClient client, final String correlationKey) {
    final var messageName = workerCfg.getMessageName();

    LOGGER.debug("Publish message '{}' with correlation key '{}'", messageName, correlationKey);
    final var messageSendFuture =
        client
            .newPublishMessageCommand()
            .messageName(messageName)
            .correlationKey(correlationKey)
            .send();

    try {
      messageSendFuture.get(10, TimeUnit.SECONDS);
      return true;
    } catch (final Exception ex) {
      THROTTLED_LOGGER.error(
          "Exception on publishing a message with name {} and correlationKey {}",
          messageName,
          correlationKey,
          ex);
      return false;
    }
  }

  private void addDelayToCompletion(
      final long completionDelay, final long startHandlingTime) {
    try {
      final var elapsedTime = System.currentTimeMillis() - startHandlingTime;
      if (elapsedTime < completionDelay) {
        final long sleepTime = completionDelay - elapsedTime;
        LOGGER.debug("Sleep for {} ms", sleepTime);
        Thread.sleep(sleepTime);
      } else {
        LOGGER.debug(
            "Skip sleep. Elapsed time {} is larger then {} completion delay.",
            elapsedTime,
            completionDelay);
      }
    } catch (final Exception e) {
      THROTTLED_LOGGER.error("Exception on sleep with completion delay {}", completionDelay, e);
    }
  }
}