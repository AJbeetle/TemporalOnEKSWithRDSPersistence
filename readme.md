OBJECTIVE :
Deploy a fully self-hosted Temporal cluster on Amazon EKS persisted by Amazon RDS (PostgreSQL), and validate durable workflow execution through a fault-tolerant order processing demo that demonstrates:
Automatic retry on activity failure
Stateful workflow persistence across restarts
Multi-step orchestration with heartbeating and backoff policies
DEMO APPLICATION TO CREATE : 
START
  │
  ▼
Activity: ValidateOrder
  • Check item availability (mock)
  • Retry: max 3 attempts, backoff 2s
  │
  ▼
Activity: ReserveInventory
  • Deduct stock (mock DB write)
  • Heartbeat every 2s
  • Retry: max 5 attempts, backoff 5s
  │
  ▼
Activity: ChargePayment
  • Simulate payment gateway (INTENTIONALLY FAIL 50% of calls)
  • Retry: max 3 attempts, backoff 10s  ← demonstrates automatic retry
  │
  ▼
Activity: SendConfirmationEmail
  • Log "email sent" to stdout
  │
  ▼
COMPLETE -- return OrderID + status

