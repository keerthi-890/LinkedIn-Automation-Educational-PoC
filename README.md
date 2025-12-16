# LinkedIn Automation – Educational PoC (Simulation Only)

> This project demonstrates automation techniques (browser control, stealth simulation) using a mock professional networking website 

Project Overview

This project is a proof-of-concept LinkedIn automation simulation
It demonstrates:

   - Browser automation with Go + Rod
   - Human-like behavioral simulation (stealth)
   - Full-stack architecture with a React + Tailwind dashboard
   - State persistence using SQLite
   - Modular Go architecture with clean separation of concerns

All actions are performed on a mock site, and connection/messaging behavior is simulated.


Tech Stack

 Backend
    - Golang (1.21+)
    - Rod (Browser Automation)
    - SQLite (state persistence)
    - REST API (/start, /stop, /stats, /logs)
    - Modular structure: auth, search, connect, message, stealth, storage, core, api, config, logger

Frontend
- React + Vite
- Tailwind CSS (dark SaaS-style dashboard)
- Stats cards, logs panel, stealth toggle, automation controls

Mock Environment
- mock-site/network.html simulates a professional networking platform
- No real LinkedIn data is used

Features

Core Automation
- Authentication Module: Simulated login using mock credentials
- Search & Targeting: Scan profiles by job title, company, location, and keywords
- Connection Requests: Click "Connect" buttons with human-like delays and logging
- Messaging System: Simulated follow-up messages with templates
- State Persistence: Tracks sent requests and messages in SQLite

Stealth & Human-like Behavior
- Mandatory:
  - Human-like mouse movement (simulated)
  - Randomized timing delays
  - Browser fingerprint abstraction
- Additional:
  - Random scrolling
  - Typing simulation (with minor typos)
  - Mouse hovering and wandering
  - Activity scheduling (business hours)
  - Rate limiting and cooldowns

Dashboard
- Live stats cards: Profiles Found, Requests Sent, Connections, Messages Sent
- Live execution logs panel
- Automation controls: Start, Pause, Stop, Emergency
- Stealth toggle switches
- Daily limit tracker



## Project Structure

