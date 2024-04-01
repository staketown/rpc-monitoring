# RPC Monitoring Tool

RPC Monitor is a simple tool what call you rpcs and check if it's crashed with AppHash or any panic and notify you
over provided discord webhook


# How To Start?

1. Edit prometheus file with adding you rpc `./rpc-monitoring/prometheus/prometheus.yml` and add/edit the
`static_configs` section with your rpc.
2. Edit `./rpc-monitoring/prometheus/alert_manager/alertmanager.yml` file to add your discord webhook.
3. From root directory execute the following commands:
```bash
docker compose build --no-cache
docker compose up -d
```

# Useful commands

Start containers
```bash
docker compose up -d
```

Stop containers
```bash
docker compose down
```