# Релизный образ: минимальный, без прав root, со встроенной проверкой готовности.
# UID 65534 соответствует стандартному пользователю nobody в alpine.
FROM alpine:3.21

RUN apk add --no-cache ca-certificates \
    && mkdir -p /slo_exporter \
    && chown -R 65534:65534 /slo_exporter

COPY --chown=65534:65534 slo_exporter /slo_exporter/slo_exporter

USER 65534:65534
WORKDIR /slo_exporter

EXPOSE 8080

# Проверка готовности: readiness возвращает 503 во время остановки, что позволяет оркестратору вывести под из ротации.
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget -qO- http://127.0.0.1:8080/readiness || exit 1

ENTRYPOINT ["/slo_exporter/slo_exporter"]

CMD ["--help"]
