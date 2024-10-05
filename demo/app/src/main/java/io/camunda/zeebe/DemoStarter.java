package io.camunda.zeebe;

import io.camunda.zeebe.config.AppCfg;

public class DemoStarter extends Starter {
  DemoStarter(final AppCfg appCfg) {
    super(appCfg);
  }

  public static void main(final String[] args) {
    createApp(DemoStarter::new);
  }
}
