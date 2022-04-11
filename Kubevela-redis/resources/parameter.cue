parameter: {
  redisconfig: *"""
        dir /data
        port 6379
        bind 0.0.0.0
        appendonly yes
        protected-mode no
        requirepass 123456
        pidfile /data/redis-6379.pid
  """ | string
}