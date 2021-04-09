FROM gaul0/stratos-ground-station:latest AS build-go

ENV APP_NAME='ground-station'
ENV FRONTEND_DIR='dashboard'

COPY . .
RUN /tools/build.sh \
    'linux/amd64' \
    'linux/arm64' \
    'linux/arm' \
    'windows/amd64'

##############################################
FROM scratch
COPY --from=build-go /out/* ./