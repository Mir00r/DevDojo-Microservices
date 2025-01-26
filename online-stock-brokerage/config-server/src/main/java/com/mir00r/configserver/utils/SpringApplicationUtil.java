package com.mir00r.configserver.utils;

import org.apache.commons.lang3.StringUtils;
import org.slf4j.Logger;
import org.springframework.boot.SpringApplication;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.core.env.Environment;

import java.net.InetAddress;
import java.net.UnknownHostException;
import java.util.Optional;

public class SpringApplicationUtil {

  public static final String APP_NAME_KEY = "spring.application.name";
  public static final String WEB_APP_TYPE_KEY = "spring.main.web-application-type";

  private SpringApplicationUtil() {}

  public static void logApplicationStartup(
    final Logger log,
    final Environment env) {

    final String protocol = Optional
      .of(env.getProperty("server.ssl.key-store"))
      .map(key -> "https")
      .orElse("http");
    final String serverPort = env.getProperty("server.port");
    final String contextPath = Optional
      .of(env.getProperty("server.servlet.context-path"))
      .filter(StringUtils::isNotBlank)
      .orElse("/");
    final String datasourceUrl = Optional
      .of(env.getProperty("SPRING_DATASOURCE_URL"))
      .orElse(env.getProperty("spring.datasource.url"));

    String hostAddress = "localhost";
    try {
      hostAddress = InetAddress.getLocalHost().getHostAddress();
    } catch (UnknownHostException e) {
      log.warn("The host name could not be determined, using `localhost` as fallback");
    }

    String configServerStatus = env.getProperty("configserver.status");
    if (configServerStatus == null) {
      configServerStatus = "Not found or not setup for this application";
    }

    final String projectVersion = env.getProperty("info.project.version", "-");

    final String appName = Optional
      .of(env.getProperty(APP_NAME_KEY))
      .orElse("");

    log.info("\n" +
        """
          ----------------------------------------------------------
              Application '{}' is running! with:-
              Local:         {}://localhost:{}{}
              External:      {}://{}:{}{}
              Datasource:    {}
              Profile(s):    {}
              Config Server: {}
              Version:       {}
          ----------------------------------------------------------""",
      appName,
      protocol, serverPort, contextPath,
      protocol, hostAddress, serverPort, contextPath,
      datasourceUrl,
      env.getActiveProfiles(),
      configServerStatus,
      projectVersion);
  }

  public static void exitForNoneWebApp(
    final Logger log,
    final Environment env,
    final ConfigurableApplicationContext applicationContext) {

    final String type =
      env.getRequiredProperty(WEB_APP_TYPE_KEY);
    log.info(">> {}={}", WEB_APP_TYPE_KEY, type);
    if (type.equals("none")) {
      // for CommandLineRunner or ApplicationRunner
      System.exit(SpringApplication.exit(applicationContext));
    }
  }
}
