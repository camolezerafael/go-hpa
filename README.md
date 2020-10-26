# Curso Full Cycle

## Kubernetes & HPA | Deploy Contínuo (Atualizado 25/10/2020)

##### Autor: [Rafael Camoleze](mailto:contato@rafaelcamoleze.com)

1. Imagem no Docker Hub | [Link](https://hub.docker.com/r/camolezerafael/go-hpa)
2. Endereço do HPA no GCP | [Link](http://34.66.228.161/)
3. Triggers atualizados para regra do exercício:
    1. CI quando push/PR branch diferente de Master
    2. CI/CD quando push/merge no branch Master

Criação de pod para teste do escalonamento:

```
$ kubectl run -it loader --image=busybox /bin/sh
$ while true; do wget -q -O- http://go-hpa.default.svc.cluster.local; done;
$ watch kubectl get hpa
```
