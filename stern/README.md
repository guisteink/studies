# Comandos Stern - Monitoramento de Logs

Aqui estão alguns comandos Stern para monitorar logs em diferentes contextos do Kubernetes.

## Monitorar Logs por Número de Telefone no Contexto 'wa'

```
stern -l phone=XXXXXXXXXXXX --context {{}} --color never -s 1s > XXXXXXXXXXXX.log
stern -l phone=551140029000 --context wa --color never -s 1s > 551140029000.log
```

## Filtrar Logs por Número no Contexto 'prod'
```
stern -l app=connector-whatsapp-facebook --context prod --color always -s 1m | cut -d ' ' -f3- | fgrep 5491152460410 | jq .
```
