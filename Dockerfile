FROM golang:bullseye AS build
WORKDIR /app
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,target=.,rw=true \
    go mod tidy && \
    go build -o /webdav

FROM scratch AS final
COPY --from=build /webdav /webdav
