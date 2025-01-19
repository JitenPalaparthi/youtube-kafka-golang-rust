# Kafka Environment Variables Explained

## **Cluster and Node Identification**

- **`CLUSTER_ID: 5hDiE5vETjicqnOmlhZ9Og`**
  - Unique identifier for the Kafka cluster.
  - Required for Kafka in **KRaft mode** (without ZooKeeper).
  - Ensures all nodes belong to the same cluster.

- **`KAFKA_NODE_ID: 1`**
  - Unique ID for the Kafka broker within the cluster.
  - Each broker in a Kafka cluster must have a unique node ID.

---

## **Process Roles**

- **`KAFKA_PROCESS_ROLES: broker,controller`**
  - Defines the roles of the node:
    - **`broker`**: Handles producer/consumer client requests.
    - **`controller`**: Manages cluster metadata, leader election, etc.
  - In **KRaft mode**, a broker can act as both.

---

## **Controller Quorum**

- **`KAFKA_CONTROLLER_QUORUM_VOTERS: 1@localhost:9093`**
  - Specifies nodes eligible to vote for the Kafka controller in the quorum.
  - Format: `node_id@address:port`.
  - Example:
    - Node `1` listens at `localhost:9093` for controller communications.
  - Relevant for **KRaft mode** only.

---

## **Listeners and Networking**

- **`KAFKA_LISTENERS: PLAINTEXT://:9092,CONTROLLER://:9093`**
  - Defines endpoints for Kafka to accept connections:
    - **`PLAINTEXT://:9092`**: Client connections (producers/consumers).
    - **`CONTROLLER://:9093`**: Controller-related communications.

- **`KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://192.168.0.101:9092`**
  - Address advertised to external clients.
  - Ensures clients can connect to the broker using the host machine's IP.

- **`KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,CONTROLLER:PLAINTEXT`**
  - Maps each listener to a security protocol:
    - **PLAINTEXT**: Unsecured communication for clients.
    - **CONTROLLER**: Uses PLAINTEXT in this setup.

- **`KAFKA_INTER_BROKER_LISTENER_NAME: PLAINTEXT`**
  - Listener used for broker-to-broker communication.

- **`KAFKA_CONTROLLER_LISTENER_NAMES: CONTROLLER`**
  - Listener used for controller communication.

---

## **Logs and Storage**

- **`KAFKA_LOG_DIRS: /tmp/kraft-combined-logs`**
  - Directory where Kafka stores log data, such as:
    - Topic partitions.
    - Offsets.
    - Other metadata.
  - Logs are stored under `/tmp/kraft-combined-logs`.

---

## **Replication and High Availability**

- **`KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1`**
  - Replication factor for the **`__consumer_offsets`** topic.
  - `1` indicates no replication, suitable for single-broker setups.

- **`KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR: 1`**
  - Replication factor for the **transaction state logs** (used for transactions).
  - `1` indicates no replication.

- **`KAFKA_TRANSACTION_STATE_LOG_MIN_ISR: 1`**
  - Minimum number of in-sync replicas (ISR) required for the transaction state log.
  - `1` ensures it works in single-node setups.

- **`KAFKA_MIN_INSYNC_REPLICAS: 1`**
  - Minimum number of in-sync replicas required for a write to succeed.
  - `1` ensures that writes are acknowledged without replication.

---

## **Summary**
- **Single-node Kafka cluster** running in **KRaft mode** (without ZooKeeper).
- Node acts as both a **broker** and a **controller**.
- Networking supports internal (controller) and external (client) communication.
- Minimal replication factors, suitable for development or non-production use.