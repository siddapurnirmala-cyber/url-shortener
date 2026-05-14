package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type TokenBucketLimiter struct {
	RDB        *redis.Client
	Ctx        context.Context
	Capacity   int64
	RefillRate int64
}

func (tb *TokenBucketLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {

		// If Redis not available → skip limiter
		if tb.RDB == nil {
			c.Next()
			return
		}

		ip := c.ClientIP()
		key := "token_bucket:" + ip

		now := time.Now().Unix()

		// Lua script (atomic token bucket)
		script := redis.NewScript(`
local key = KEYS[1]
local capacity = tonumber(ARGV[1])
local refill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])

local data = redis.call("HMGET", key, "tokens", "last_refill")
local tokens = tonumber(data[1])
local last_refill = tonumber(data[2])

if tokens == nil then
    tokens = capacity
    last_refill = now
end

-- refill tokens
local delta = math.max(0, now - last_refill)
local refill = delta * refill_rate
tokens = math.min(capacity, tokens + refill)

-- if no tokens → reject
if tokens < 1 then
    return -1
end

-- consume one token
tokens = tokens - 1

redis.call("HMSET", key, "tokens", tokens, "last_refill", now)
redis.call("EXPIRE", key, 60)

return tokens
`)

		// ✅ FIXED: Use Int64() instead of Result()
		result, err := script.Run(tb.Ctx, tb.RDB, []string{key},
			tb.Capacity,
			tb.RefillRate,
			now,
		).Int64()

		if err != nil {
			println("❌ Redis error:", err.Error())
			c.Next() // fail open
			return
		}

		// Debug log
		println("DEBUG → tokens left:", result)

		// 🚫 Rate limit hit
		if result < 0 {
			println("🚫 RATE LIMIT HIT for IP:", ip)

			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded",
			})
			c.Abort()
			return
		}

		// Allow request
		c.Next()
	}
}
