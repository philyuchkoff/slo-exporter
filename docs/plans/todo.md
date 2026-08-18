# TODO

## Активные задачи

- [x] Фикс golangci-lint v1.61.0 → v1.64.8 в Makefile (несовместим с Go 1.23.2+)
- [x] Замена заброшенной библиотеки hpcloud/tail → nxadm/tail v1.4.11
- [x] Удаление устаревшей зависимости grafana/loki: пакет promtail positions перенесён в pkg/positions (Apache 2.0, атрибуция сохранена)
- [x] Обновление prometheus/prometheus 2021 → v0.54.1 (pkg/labels → model/labels, pkg/relabel → model/relabel, pkg/rulefmt → model/rulefmt)
- [x] Dockerfile: alpine-образ, non-root пользователь, HEALTHCHECK на /readiness
- [x] Dependabot для gomod и github-actions
- [x] Миграция CircleCI → GitHub Actions (.github/workflows/ci.yml)
- [x] Профилирование (block/mutex) вынесено за флаг --enable-pprof, по умолчанию выключено
- [x] Эндпоинты /debug/pprof/ (флаг --enable-pprof) и /logging (флаг --enable-dynamic-logging) настраиваемые
- [x] Тесты для pkg/config (7 тестов: загрузка, дефолты, ошибки, moduleConfig)
- [x] make test: добавлен -count=1 против ложного кэша

## История

- (здесь будут переноситься выполненные задачи после завершения сессии)