package com.mir00r.configserver.configs;


import com.mir00r.configserver.SecurityProblemSupport;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Import;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configurers.AbstractHttpConfigurer;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.context.DelegatingSecurityContextRepository;
import org.springframework.security.web.context.NullSecurityContextRepository;
import org.springframework.security.web.context.RequestAttributeSecurityContextRepository;
import org.springframework.security.web.context.SecurityContextRepository;

@Configuration
@EnableWebSecurity
@Import(SecurityProblemSupport.class)
public class WebSecurityConfig {

  WebSecurityConfig() {}

  @Configuration
  public static class DefaultSecurityConfig {

    private final SecurityProblemSupport problemSupport;

    public DefaultSecurityConfig(
      final SecurityProblemSupport problemSupport) {

      this.problemSupport = problemSupport;
    }

    @Bean
    SecurityFilterChain defaultFilterChain(
      HttpSecurity http)
      throws Exception {

      http
        .csrf(AbstractHttpConfigurer::disable)
        .exceptionHandling(config -> config
          .authenticationEntryPoint(problemSupport)
          .accessDeniedHandler(problemSupport))
        .sessionManagement(spec -> spec
          .sessionCreationPolicy(SessionCreationPolicy.STATELESS))
        .securityContext(config -> config
          .requireExplicitSave(true)
          .securityContextRepository(securityContextRepository()))
        .authorizeHttpRequests(authorize -> authorize
          .anyRequest().authenticated())
        .httpBasic(spec -> spec
          .authenticationEntryPoint(problemSupport));

      return http.build();
    }

    private SecurityContextRepository securityContextRepository() {

      return new DelegatingSecurityContextRepository(
        new RequestAttributeSecurityContextRepository(),
        new NullSecurityContextRepository());
    }
  }
}
