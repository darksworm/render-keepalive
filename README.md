# Render Keepalive

A minimal Docker application that pings a specified URL once. Designed to be as small and simple as possible.

## Features

- Reads URL from environment variables
- Sends a request to the URL once without waiting for a response (fire and forget)
- Extremely small Docker image (built on `scratch`)
- Simple, clean Go code with proper error handling

## Environment Variables

- `PING_URL` (required): The URL to ping

## Building the Docker Image

```bash
docker build -t render-keepalive .
```

## Running the Docker Image

```bash
docker run --rm -e PING_URL=https://example.com render-keepalive
```

## Example Output

```
Starting to ping URL: https://example.com
Sending request to https://example.com without waiting for response...
Request sent successfully
```

## Why Go?

Go was chosen for this project because:
1. It compiles to a single binary with no runtime dependencies
2. It has a small memory footprint
3. It has built-in HTTP client capabilities
4. It can be compiled to a static binary that can run in a minimal Docker image

## Docker Image Size

The resulting Docker image is extremely small (typically a few MB) because:
1. It uses a multi-stage build
2. The final image is based on `scratch` (the most minimal base image)
3. It contains only the compiled Go binary with no OS components
