package com.mir00r.configserver;

import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.security.access.AccessDeniedException;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.web.AuthenticationEntryPoint;
import org.springframework.security.web.access.AccessDeniedHandler;
import org.springframework.web.servlet.HandlerExceptionResolver;

public class SecurityProblemSupport
  implements AuthenticationEntryPoint, AccessDeniedHandler {

  private final HandlerExceptionResolver resolver;

  @Autowired
  public SecurityProblemSupport(
    @Qualifier("handlerExceptionResolver") final HandlerExceptionResolver resolver) {
    this.resolver = resolver;
  }

  @Override
  public void commence(
    final HttpServletRequest request,
    final HttpServletResponse response,
    final AuthenticationException exception) {
    resolver.resolveException(request, response, null, exception);
  }

  @Override
  public void handle(
    final HttpServletRequest request,
    final HttpServletResponse response,
    final AccessDeniedException exception) {
    resolver.resolveException(request, response, null, exception);
  }
}
