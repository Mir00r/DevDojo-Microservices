package com.mir00r.configserver.properties;

public class RegistryProperties {

  private final User user = new User();

  private String host = "localhost";

  private int port = 8761;

  public RegistryProperties() {
    // required by client application
  }

  public String getHost() {
    return host;
  }

  public void setHost(String host) {
    this.host = host;
  }

  public int getPort() {
    return port;
  }

  public void setPort(int port) {
    this.port = port;
  }

  public User getUser() {
    return user;
  }

  public static class User {

    private String name;

    private String password;

    public String getName() {
      return name;
    }

    public void setName(String name) {
      this.name = name;
    }

    public String getPassword() {
      return password;
    }

    public void setPassword(String password) {
      this.password = password;
    }
  }

}
