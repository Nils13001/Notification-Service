
# Notification Service — E-Commerce Microservices Architecture

**Status:** Foundations (Setup Only)  
**Language:** Go (Chi Router)  
**OS:** Windows (No Docker in Sprint 1)  
**Epic:** MSA E01 — Polyglot Event Driven E-Commerce Microservices  

---

## 📌 Overview
`notification-service` is part of a polyglot microservices architecture for an e-commerce platform.  
This skeleton provides:
- Health endpoint (`GET /api/notification/health`)
- Payment authorization endpoint (`POST /api/notificaion/notify`) — returns **501 Not Implemented** (placeholder)
- Basic Chi router setup with middleware support (logging planned)

---

## ✅ Quickstart (Windows)
**Prerequisites:**
- Go 1.25+
- Git installed

**Run locally:**
```bash
go run ./cmd/notification-service
