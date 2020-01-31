# tz-controller

apply crd
```
kubectl apply -f deployment/k8s/crd.yaml
```

create cr
```
kubectl apply -f deployment/k8s/cr.yaml
```

## feature

- TimeZone: 时区
- WhiteListNS: 不创建的podpreset的白名单
