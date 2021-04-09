ARG OUTPUT_DIR='/out'

FROM gaul0/stratos-ground-station:latest AS build-go
# Use --build-arg to set ARGs when building the image
ARG BUILD_TARGETS='linux/amd64,linux/arm64,linux/arm/7,windows/amd64'
ARG APP_NAME='ground-station'
ARG FRONTEND_DIR='dashboard'
ARG SKIP_FRONTEND
ARG OUTPUT_DIR
ARG USE_ANSI=1

ENV APP_NAME=$APP_NAME \
    FRONTEND_DIR=$FRONTEND_DIR \
    SKIP_FRONTEND=$SKIP_FRONTEND \
    OUTPUT_DIR=$OUTPUT_DIR \
    USE_ANSI=$NO_ANSI

COPY . .

RUN echo "$BUILD_TARGETS" | xargs -d, /tools/build.sh

##################################################
FROM scratch
ARG OUTPUT_DIR
COPY --from=build-go "$OUTPUT_DIR" .