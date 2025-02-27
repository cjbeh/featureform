# Snowflake

Featureform supports [Snowflake](https://www.snowflake.com/) as an Offline Store.

## Implementation <a href="#implementation" id="implementation"></a>

### Primary Sources

#### Tables

Table sources are used directly via a view. Featureform will never write to a primary source.

#### Files

Files are copied into a Snowflake table via a Kubernetes Job kicked off by our coordinator. If scheduling is set, the table will atomically be re-copied over.

### Transformation Sources

SQL transformations are used to create a view. By default, those views are materialized and updated according to the schedule parameter. Deprecated transformations are converted to un-materialized views to save storage space.

### Offline to Inference Store Materialization

When a feature is registered, Featureform creates an internal transformation to get the newest value of every feature and its associated entity. A Kubernetes job is then kicked off to sync this up with the Inference store.

### Training Set Generation

Every registered feature and label is associated with a view table. That view contains three columns, the entity, value, and timestamp. When a training set is registered, it is created as a materialized view via a JOIN on the corresponding label and feature views.

## Configuration <a href="#configuration" id="configuration"></a>

First we have to add a declarative Snowflake configuration in Python.

### Credentials with Account and Organization

{% code title="snowflake_config.py" %}

```python
import featureform as ff

ff.register_snowflake(
    name = "snowflake_docs",
    description = "Example inference store",
    team = "Featureform",
    username = snowflake_username,
    password = snowflake_password,
    account = snowflake_account,
    organization = snowflake_org,
    database = snowflake_database,
    schema = snowflake_schema,
)
```

{% endcode %}

### Legacy Credentials

[Older Snowflake accounts](https://docs.snowflake.com/en/user-guide/admin-account-identifier.html#using-an-account-locator-as-an-identifier) may have credentials that use an `Account Locator` rather than
an `account` and `organization` to connect. Featureform provides a separate registration
function to support these credentials.

{% code title="snowflake_config.py" %}

```python
import featureform as ff

ff.register_snowflake_legacy(
    name = "snowflake_docs",
    description = "Example inference store",
    team = "Featureform",
    username = snowflake_username,
    password =  snowflake_password,
    account_locator = snowflake_account_locator,
    database = snowflake_database,
    schema = snowflake_schema,
)
```

{% endcode %}

Once our config file is complete, we can apply it to our Featureform deployment

```bash
featureform apply snowflake_config.py --host $FEATUREFORM_HOST
```

We can re-verify that the provider is created by checking the [Providers tab of the Feature Registry](../getting-started/exploring-the-feature-registry.md).

### Mutable Configuration Fields

* `description`
* `username`
* `password`
* `role`
