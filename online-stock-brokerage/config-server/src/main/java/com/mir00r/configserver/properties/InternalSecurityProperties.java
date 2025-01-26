package com.mir00r.configserver.properties;

import org.springframework.boot.context.properties.ConfigurationProperties;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

@ConfigurationProperties(prefix = InternalSecurityProperties.PREFIX)
public class InternalSecurityProperties {

  public static final String PREFIX = "application.internal-security";
  public static final String DEFAULT_AUTH = "ROLE_INTERNAL";

  private String rootPattern = "/internal/**";

  private List<User> user = new ArrayList<>();

  public InternalSecurityProperties() {
    // required by client application
  }

  public String getRootPattern() {
    return rootPattern;
  }

  public InternalSecurityProperties setRootPattern(String rootPattern) {
    this.rootPattern = rootPattern;
    return this;
  }

  public List<User> getUser() {
    return user;
  }

  public InternalSecurityProperties setUser(List<User> user) {
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
    return String.format("InternalSecurityProperties [rootPattern=%s, user=%s]", rootPattern, user);
  }

}
