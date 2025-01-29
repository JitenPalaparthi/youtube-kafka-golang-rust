# Kafka Producer Configuration

Kafka producers are responsible for sending records (messages) to Kafka topics. The producer configuration parameters control aspects like message serialization, retries, acknowledgments, and compression.

---

## **1. Broker Connection**
- **`bootstrap.servers`**  
  - A list of host:port pairs of Kafka brokers the producer connects to.  
  - **Example**: `bootstrap.servers=localhost:9092,broker2:9092`.  

---

## **2. Message Serialization**
- **`key.serializer`**  
  - Serializer class for the message key.  
  - Must implement `org.apache.kafka.common.serialization.Serializer`.  
  - **Example**: `key.serializer=org.apache.kafka.common.serialization.StringSerializer`.

- **`value.serializer`**  
  - Serializer class for the message value.  
  - Must implement `org.apache.kafka.common.serialization.Serializer`.  
  - **Example**: `value.serializer=org.apache.kafka.common.serialization.StringSerializer`.

---

## **3. Acknowledgments (Durability)**
- **`acks`**  
  - Controls when the producer considers a write successful.  
  - **Options**:  
    - `0`: Producer does not wait for acknowledgment from the broker. (Fast but risks data loss).  
    - `1`: Producer waits for acknowledgment from the leader broker only.  
    - `all` (or `-1`): Producer waits for acknowledgment from all in-sync replicas. (Highest durability).  
  - **Example**: `acks=all`.

---

## **4. Retries and Idempotence**
- **`retries`**  
  - Number of retry attempts for failed send requests.  
  - **Example**: `retries=5`.  
  - Use in conjunction with `acks=all` for reliability.

- **`enable.idempotence`**  
  - Ensures exactly-once delivery by guaranteeing no duplicate messages.  
  - **Example**: `enable.idempotence=true`.  
  - Automatically sets `acks=all` and `retries` to a high value.

---

## **5. Compression**
- **`compression.type`**  
  - Compression type for producer messages.  
  - **Options**: `none`, `gzip`, `snappy`, `lz4`, `zstd`.  
  - **Example**: `compression.type=snappy`.  

---

## **6. Batching**
- **`batch.size`**  
  - Maximum size (in bytes) of a batch of messages sent to a partition.  
  - **Example**: `batch.size=16384` (16 KB).  

- **`linger.ms`**  
  - Time to wait before sending a batch of messages.  
  - Adds latency but can improve throughput by allowing more messages to be batched together.  
  - **Example**: `linger.ms=5`.

---

## **7. Memory and Buffering**
- **`buffer.memory`**  
  - Total memory available for buffering unsent messages.  
  - **Example**: `buffer.memory=33554432` (32 MB).  

- **`max.block.ms`**  
  - Maximum time the `send()` method will block if the buffer is full.  
  - **Example**: `max.block.ms=60000` (1 minute).  

---

## **8. Timeouts**
- **`request.timeout.ms`**  
  - Time for the broker to respond to a produce request before the producer retries or fails.  
  - **Example**: `request.timeout.ms=30000` (30 seconds).  

- **`delivery.timeout.ms`**  
  - Maximum time to attempt delivering a record, including retries.  
  - **Example**: `delivery.timeout.ms=120000` (2 minutes).  

---

## **9. Record Ordering**
- **`max.in.flight.requests.per.connection`**  
  - Maximum number of unacknowledged requests per connection.  
  - Lower values (e.g., 1) ensure strict ordering but may reduce throughput.  
  - **Example**: `max.in.flight.requests.per.connection=5`.

---

## **10. Security**
- **`security.protocol`**  
  - Protocol used to communicate with brokers.  
  - **Options**: `PLAINTEXT`, `SSL`, `SASL_PLAINTEXT`, `SASL_SSL`.  
  - **Example**: `security.protocol=SSL`.

- **`sasl.mechanism`**  
  - Authentication mechanism for SASL.  
  - **Example**: `sasl.mechanism=PLAIN`.

---

## **11. Additional Settings**
- **`client.id`**  
  - Identifier for the producer instance. Useful for tracking and logging.  
  - **Example**: `client.id=my-producer`.

- **`enable.idempotence`**  
  - Ensures exactly-once delivery.  
  - **Example**: `enable.idempotence=true`.

- **`transactional.id`**  
  - Used to enable transactional messaging for exactly-once semantics across multiple partitions.  
  - **Example**: `transactional.id=transaction-1`.

---

## **Producer Configuration Example **

```properties
bootstrap.servers=localhost:9092
key.serializer=org.apache.kafka.common.serialization.StringSerializer
value.serializer=org.apache.kafka.common.serialization.StringSerializer
acks=all
retries=5
compression.type=snappy
linger.ms=5
batch.size=16384
buffer.memory=33554432
enable.idempotence=true
client.id=my-producer
```

### Delivery Guarantees

    - At most once: acks=0, low retry, and no idempotence.
	- At least once: acks=1, retries > 0.
	- Exactly once: acks=all, enable.idempotence=true.
