# tosutils-go

Golang library for interacting with TOS blockchain.

This library is native golang implementation of ADNL and lite protocol. It works as connection pool and can be connected to multiple lite servers in the same time, balancing is done on lib side.

It is concurrent safe and can be used from multiple goroutines under high workloads.

All main TOS protocols are implemented: ADNL, DHT, RLDP, Overlays, etc.

------

### Features
* ✅ Support cell and slice as arguments to run get method
* ✅ Reconnect on failure
* ✅ Get account state method
* ✅ Send external message
* ✅ Get transactions
* ✅ Deploy contracts
* ✅ Wallet operations
* ✅ Cell dictionaries support
* ✅ MustLoad methods
* ✅ Parse global config json
* ✅ Jettons
* ✅ DNS
* ✅ ADNL UDP Client/Server
* ✅ ADNL TCP Client/Server
* ✅ RLDP Client/Server
* ✅ TON Sites Client/Server
* ✅ DHT
* ✅ Merkle proofs validation and creation
* ✅ Overlays
* ✅ TL Parser/Serializer
* ✅ TL-B Parser/Serializer
* ✅ Payment channels
* ✅ Liteserver proofs automatic validation
* ✅ TVM (Contract execution emulation)


