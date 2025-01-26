package com.mir00r.configserver.properties;

import org.springframework.boot.context.properties.ConfigurationProperties;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

@ConfigurationProperties(prefix = ActuatorSecurityProperties.PREFIX)
public class ActuatorSecurityProperties {

  public static final String PREFIX = "application.actuator-security";
  public static final String DEFAULT_AUTH = "ROLE_ACTUATOR";

  private String rootPattern = "/actuator/**";

  private String infoPattern = "/actuator/info";

  private String healthPattern = "/actuator/health";

  private User user = new User();

  public ActuatorSecurityProperties() {
    // required by client application
  }

  public String getRootPattern() {
    return rootPattern;
  }

  public ActuatorSecurityProperties setRootPattern(String rootPattern) {
    this.rootPattern = rootPattern;
    return this;
  }

  public String getInfoPattern() {
    return infoPattern;
  }

  public ActuatorSecurityProperties setInfoPattern(String infoPattern) {
    this.infoPattern = infoPattern;
    return this;
  }

  public String getHealthPattern() {
    return healthPattern;
  }

  public ActuatorSecurityProperties setHealthPattern(String healthPattern) {
    this.healthPattern = healthPattern;
    return this;
  }

  public User getUser() {
    return user;
  }

  public ActuatorSecurityProperties setUser(User user) {
    this.user = user;
    return this;
  }

  public static class User {

    private String name;

    private String password;

    private List<String> authorities = new ArrayList<>(Arrays.asList(DEFAULT_AUTH));

    public String getName() {
      return name;
    }

    public User setName(String name) {
      this.name = name;
      return this;
    }

    public String getPassword() {
      return password;
    }

    public User setPassword(String password) {
      this.password = password;
      return this;
    }

    public List<String> getAuthorities() {
      return authorities;
    }

    public User setAuthorities(List<String> authorities) {
      this.authorities = authorities;
      return this;
    }

    @Override
    @SuppressWarnings("java:S2068")
    public String toString() {
      // Sonar false alarm, Password is not hardcoded.
      return String.format("User [name=%s, password=xxxx, authorities=%s]", name, authorities);
    }
  }

  @Override
  public String toString() {
    return String.format(
      "ActuatorSecurityProperties [rootPattern=%s, infoPattern=%s, healthPattern=%s, user=%s]",
      rootPattern, infoPattern, healthPattern, user);
  }

}
