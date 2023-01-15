*This is currently a proof-of-concept, but it really does work.*

Kplug
=====

A plugin platform and SDK for Kubernetes controllers.

Plugins are implemented as straightforward [gRPC](https://grpc.io/) servers running in Kubernetes
pods. Any [supported programming language](https://grpc.io/docs/languages/) can be used, e.g. Python,
Java, JavaScript, Go, etc. The Kplug [Plugin SDK](sdk/) (with eplicit per-language support) handles
gRPC setup and containerization so that plugin authors need only focus on plugin logic.

On the controller side, the Kplug [controller library](kplug/) (written in Go) can be integrated
into your Kubernetes operator code to handle the gRPC setup, APIs, and delegation to the plugins.

Rationale
---------

Writing Kubernetes controllers—especially good ones—is unfortunately very challenging. It
requires deep domain-specific knowledge of Kubernetes and ideally advanced Go programming skills.
Kplug removes that obstacle. It allows implementers to focus purely on their implementation logic
in the language of their choice.

Perhaps even more importantly, it decouples the operator semantics from the implementation
semantics. Often these are meant to be developed by different teams of engineers, with different
areas of expertise, and indeed from different organizations. Kplug allows both the controller to
be extensible (hot swappable, live upgradeable) as well as the custom resource, by using a
meta-model that attaches arbitrary "extension" custom resources to the base, agnostic one with
additional properties and status to be processed by the plugins. A plugin can thus add both
functionality and configuration properties.


Walkthrough
-----------

You can run this example and study its code [here](examples/database-table).

Let's start with a custom resource defining a generic relational (SQL) database table, with
the idea that our Kubernetes operator will create and manage the lifecycle of that table
within a database server. On the one hand we want the definition to be agnostic to any specific
database server implementation, while on the other hand we want to allow for those implementations
to expose their own unique features for configuration.

We'll make use of Kplug's supported meta-model, by which an agnostic custom resource can
"point" to extensions using a list of
[object references](https://dev-k8sref-io.web.app/docs/common-definitions/objectreference-/).
In this example our extensions will add configuration options specific to each implementation:
a choice of [storage engine for MariaDB](https://mariadb.com/kb/en/storage-engines/) and
a choice of [partitioning mode for PostgreSQL](https://www.postgresql.org/docs/current/ddl-partitioning.html).

```yaml
apiVersion: myorg.org/v1alpha1
kind: DatabaseTable
metadata:
  name: users
spec:
  serverName: authorization-cluster
  preferredImplementation: MariaDB
  columns:
  - name: id
    type: uint64
  - name: name
    type: string
  extensions:
  - apiVersion: myorg.org/v1alpha1
    kind: DatabaseTableMariaDbExtension
    name: users
  - apiVersion: myorg.org/v1alpha1
    kind: DatabaseTablePostgreSqlExtension
    name: users

---

apiVersion: myorg.org/v1alpha1
kind: DatabaseTableMariaDbExtension
metadata:
  name: users
spec:
  storageEngine: InnoDB

---

apiVersion: myorg.org/v1alpha1
kind: DatabaseTablePostgreSqlExtension
metadata:
  name: users
spec:
  partitionBy: hash
```

Note that we specified a `preferredImplementation` in this example, but the exact semantics
are domain-specific. For example, for some resource types all available plugins may be used,
rather than just one selected implementation. Also note that we do not necessarily need a custom
resource extension for a plugin to be used, as the base resource spec may be enough. Furthermore,
a plugin may support multiple custom resource extensions. In other words, the extensions are not
mapped one-to-one with plugins.

Now let's walk through the example:

1. Our Kubernetes operator starts up. Using the Kplug library it exposes a gRPC heartbeat
server via a `Service` to well-known host name for Kubernetes internal DNS. For this example
it's called `database-table-operator`.

2. Our PostreSQL plugin starts up. For this example let's assume it's written in Python. Using
the Kplug SDK for Python it exposes a gRPC server implementing the specific plugin API. It
then schedules a heartbeat to continuously make gRPC calls to the operator at that well-known
host name, `database-table-operator`, wherein it registers itself by name, version, plugin
API, and address. In this case our address is `database-table-postgresql-plugin`. The operator
now adds the plugin to its list. Should the heartbeats stop coming in after a configured time,
the plugin will be garbage-collected and removed from the list, signifying that the plugin is
no longer available (it was uninstalled or failed).

3. The operator has set up a watcher on `DatabaseTable` custom resources and detects our `users`
resource. It sees that there is a preference for MariaDB, but there is no MariaDB plugin in
its list (no heartbeat from it), so it chooses the available PostreSQL plugin instead. It updates
the `DatabaseTable` resource with that choice of plugin in its status area.

4. Finally, delegation happens: the operator, using the Kplug library, makes a gRPC call to
the PostreSQL plugin sending it our `DatabaseTable` resource manifest as well as its two linked
extensions.

5. The plugin reconciles the manifest using appropriate semantics. In this case it should check
to see if the `users` table exists on the specified PostreSQL server, create it if not there,
and if it is there ensure that it conforms to our column design, adding missing columns, deleting
superfluous columns, and modifying the type of incorrectly typed columns. The code itself is
[very straightforward](examples/database-table/postgresql-plugin/implementation.py).

6. The return value from that gRPC call minimally contains success/failure and error messages,
but may also include other status, which may go into the base `DatabaseTable` custom resource
or any of its extensions. In this case, let's say that reconciliation was successful and
that we want to add metrics about the current number of partitions of the table and their
use. These will be associated with the `DatabaseTablePostgreSqlExtension` resource. The operator
will use that return value to update the resources accordingly.

FAQ
---

### Why not use [go-plugin](https://github.com/hashicorp/go-plugin)?

go-plugin is great! But it's optimized for reliable in-process plugins. We needed a cloud native
solution for Kubernetes that expects ephemeral resources (plugins that come and go).
