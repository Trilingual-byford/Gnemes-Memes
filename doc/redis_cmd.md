```console
127.0.0.1:6379> keys *
1) "Gnemes-Authniconicocsc"
127.0.0.1:6379> type Gnemes-Authniconicocsc
hash
127.0.0.1:6379> HGETALL Gnemes-Authniconicocsc
1) "TestRedisSetFunctionality"
2) "\"This is a key from TestRedisSetFunctionality func\""
127.0.0.1:6379> HGET Gnemes-Authniconicocsc TestRedisSetFunctionality
"\"This is a key from TestRedisSetFunctionality func\""
127.0.0.1:6379> 
```