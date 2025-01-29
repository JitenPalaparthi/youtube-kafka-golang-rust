# Kafka Topic Configuration

Kafka topic configuration controls how topics behave, including replication, partitioning, and retention policies. Below are key configuration options with explanations.

---

## **1. Partitioning and Replication**
- **`num.partitions`**  
  - The number of partitions for the topic.  
  - Determines parallelism and scalability.  
  - **Example**: `num.partitions=3`  
  - More partitions = better throughput but more coordination overhead.

- **`replication.factor`**  
  - The number of replicas for each partition.  
  - Ensures durability by replicating data across brokers.  
  - **Example**: `replication.factor=3` (requires at least 3 brokers).

---

## **2. Message Retention**
- **`retention.ms`**  
  - Time to retain messages (in milliseconds).  
  - **Example**: `retention.ms=604800000` (7 days).  
  - **Default**: Unlimited (`-1`).  
  - Useful for purging old data automatically.

- **`retention.bytes`**  
  - Maximum size of data to retain (per partition).  
  - **Example**: `retention.bytes=1073741824` (1 GB).  
  - If both `retention.ms` and `retention.bytes` are set, messages are deleted when either condition is met.

- **`delete.retention.ms`**  
  - Time to retain tombstone (delete) markers for log compaction.  
  - **Example**: `delete.retention.ms=86400000` (1 day).

---

## **3. Log Segments**
- **`segment.bytes`**  
  - Maximum size of a log segment file.  
  - **Example**: `segment.bytes=1073741824` (1 GB).  
  - Smaller segments = quicker log cleanup but increased overhead.

- **`segment.ms`**  
  - Time after which a log segment is rolled over.  
  - **Example**: `segment.ms=604800000` (7 days).

---

## **4. Cleanup Policies**
- **`cleanup.policy`**  
  - Specifies how Kafka handles old messages.  
  - **Options**:  
    - `delete`: Deletes old messages based on `retention.ms` or `retention.bytes`.  
    - `compact`: Retains the latest message with the same key (log compaction).  
  - **Example**: `cleanup.policy=delete,compact`.

- **`min.cleanable.dirty.ratio`**  
  - Minimum proportion of data that must be "dirty" (changed) before log compaction is triggered.  
  - **Example**: `min.cleanable.dirty.ratio=0.5`.

---

## **5. Message Size**
- **`max.message.bytes`**  
  - Maximum size of a single message (in bytes).  
  - **Example**: `max.message.bytes=1048576` (1 MB).  
  - Producer messages larger than this will fail.

---

## **6. Replication Settings**
- **`min.insync.replicas`**  
  - Minimum number of replicas that must acknowledge a write for it to be considered successful.  
  - **Example**: `min.insync.replicas=2`.  
  - Works with `acks=all` for high durability.

- **`unclean.leader.election.enable`**  
  - If set to `true`, allows partitions to be assigned to out-of-sync replicas when all in-sync replicas fail.  
  - **Example**: `unclean.leader.election.enable=false` (safer, avoids data loss).

---

## **7. Compression**
- **`compression.type`**  
  - Compression used for messages at the topic level.  
  - **Options**: `gzip`, `snappy`, `lz4`, `zstd`, or `producer` (uses the producer's configuration).  
  - **Example**: `compression.type=snappy`.

---

## **8. Time-based Indexing**
- **`message.timestamp.type`**  
  - Determines the timestamp type for messages.  
  - **Options**:  
    - `CreateTime`: Timestamp set by the producer.  
    - `LogAppendTime`: Timestamp set by the broker when the message is appended.  
  - **Example**: `message.timestamp.type=CreateTime`.

- **`message.timestamp.difference.max.ms`**  
  - Maximum allowed difference between a broker and a producer’s timestamp.  
  - **Example**: `message.timestamp.difference.max.ms=86400000` (1 day).

---

## **Topic Creation Example (Kafka CLI)**

You can configure topics using the Kafka CLI when creating them:

```bash
kafka-topics.sh --create \
    --bootstrap-server localhost:9092 \
    --replication-factor 3 \
    --partitions 3 \
    --topic example-topic \
    --config retention.ms=604800000 \
    --config cleanup.policy=delete \
    --config compression.type=snappy