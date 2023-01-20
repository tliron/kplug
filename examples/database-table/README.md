Kplug Example: Database Table
=============================

This functions as a "hello world" for Kplug.

Start by installing the operator and its custom resource definition:

```
kubectl apply -f assets/kubernetes/operator.yaml
```

Now install the plugins and their own custom resource definitions:

```
kubectl apply -f assets/kubernetes/mariadb-plugin.yaml
kubectl apply -f assets/kubernetes/postgresql-plugin.yaml
```

Actually, while we are adding a MariaDB custom resource definition, we are
not actually including a plugin to support it. This is deliberate in order to
demonstrate what happens if a certain plugin is not available.

To check the logs of the operator and PostreSQL plugin:

```
kubectl logs deployment/database-table-operator -n kplug-example -f
kubectl logs deployment/database-table-postgresql-plugin -n kplug-example -f
```

Now apply the example custom resources:

```
kubectl apply -f assets/kubernetes/example.yaml
```

We have specified that the `preferredImplementation` is MariaDB, but because that
plugin doesn't exist, PostgreSQL will be selected instead.

Check the logs again to see what the operator and plugin are doing.

Check to see that the custom resource and its PostreSQL custom resource have had
their statuses updated:

```
kubectl get mariadbdatabasetable/users -n kplug-example -o yaml
kubectl get PostgreSqlDatabaseTable/users -n kplug-example -o yaml
```
