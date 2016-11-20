# watcher [![CircleCI](https://circleci.com/gh/roboll/watcher.svg?style=svg)](https://circleci.com/gh/roboll/watcher)

Execute handlers in response to filesystem events.

[![Docker Repository on Quay](https://quay.io/repository/roboll/watcher/status "Docker Repository on Quay")](https://quay.io/repository/roboll/watcher)

## about

`watch.Handler` is a function that gets executed in response to filesystem events. The default handler, `exec`, performs a shell exec. The container ships with `kubectl` and `curl`, which should address most needs of hitting webhooks, sending signals, etc. The primary use case is sending signals in response to secret or configmap updates, maintaining flexibility to define alternate handlers in go for more complex scenarios.

## example usage

```
        - name: watcher
          image: quay.io/roboll/watcher:v0.1.0
          args:
            - -watch=/etc/secrets/my-secret/..data,/etc/secrets/my-other-secret/..data
            - -exec-command=kubectl -n $(POD_NAMESPACE) exec $(POD_NAME) -c my-app-container -- pkill -HUP my-app
          volumeMounts:
            - name: my-secret
              mountPath: /etc/secrets/my-secret/
              readOnly: true
            - name: my-other-secret
              mountPath: /etc/secrets/my-other-secret/
              readOnly: true
          env:
            - name: POD_NAME
              valueFrom:
                fieldRef:
                  fieldPath: metadata.name
            - name: POD_NAMESPACE
              valueFrom:
                fieldRef:
                  fieldPath: metadata.namespace
```
