# Umbrella — Junior Go Test Assignment

**Grade:** Junior Golang Developer  
**Framework:** Gin (Go)

---

## Task

Build an HTTP server with two features:

1. **Handler `GET /time`** — returns the time elapsed since **June 1, 2025** in the format `507D-7H-36M` (days, hours, minutes) with HTTP `200 OK`.
2. **Middleware** — checks the `User-Role` HTTP header. If its value equals `admin`, print `"red button user detected"` to the console using the standard `log` package.

The full task description is available in the [testtask.pdf](./testtask.pdf) file (originally specifies Echo framework, implemented with Gin).

---

## What Was Done

- Implemented a **Gin** HTTP server listening on `:8080`.
- The `/time` endpoint calculates the difference between the current UTC time and **January 1, 2025** and returns it in a human-readable format.
- A custom **middleware** (`internal/middleware/read_button_detector.go`) extracts the `User-Role` header and logs a message when an admin is detected.
- Code is structured with a separate `internal/middleware` package following Go project layout conventions.
- Request and response logging is enabled via `gin.Logger()`.

### Example

```bash
$ curl -v localhost:8080/time -H "User-Role: admin"

> GET /time HTTP/1.1
> Host: localhost:8080
> User-Role: admin

< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=utf-8
left since [Jun 1 2025]: 507D-7H-36M
```

Console output:

```
2025/01/01 ... [GIN] ... 
2025/01/01 ... red button user detected
```

---

**Time spent:** 10 minutes (without AI assistance).
