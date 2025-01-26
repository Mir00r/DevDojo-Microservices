package com.mir00r.configserver;

import com.mir00r.configserver.properties.ApplicationProperties;
import com.mir00r.configserver.utils.SpringApplicationUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.cloud.config.server.EnableConfigServer;
import org.springframework.context.ConfigurableApplicationContext;
import org.springframework.core.env.Environment;

@SpringBootApplication
@EnableConfigServer  // Enable Config Server
@EnableConfigurationProperties({ApplicationProperties.class})
public class ConfigServerApplication {

  private static final Logger LOG = LoggerFactory.getLogger(ConfigServerApplication.class);

	public static void main(String[] args) {
    final SpringApplication app = new SpringApplication(ConfigServerApplication.class);
    final ConfigurableApplicationContext applicationContext = app.run(args);
    final Environment env = applicationContext.getEnvironment();
    SpringApplicationUtil.logApplicationStartup(LOG, env);
	}

}
