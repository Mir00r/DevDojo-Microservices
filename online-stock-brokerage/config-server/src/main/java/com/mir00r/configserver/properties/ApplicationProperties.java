package com.mir00r.configserver.properties;


import org.springframework.boot.context.properties.ConfigurationProperties;

@ConfigurationProperties(prefix = "application", ignoreUnknownFields = false)
public class ApplicationProperties {

  private RegistryProperties registry = new RegistryProperties();
  private final ActuatorSecurityProperties actuatorSecurity = new ActuatorSecurityProperties();
  private final InternalSecurityProperties internalSecurity = new InternalSecurityProperties();

  public RegistryProperties getRegistry() {
    return registry;
  }

  public void setRegistry(RegistryProperties registry) {
    this.registry = registry;
  }

  public ActuatorSecurityProperties getActuatorSecurity() {
    return actuatorSecurity;
  }

  public InternalSecurityProperties getInternalSecurity() {
    return internalSecurity;
  }
}
