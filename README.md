# BookTracker 📚✍️


## Features
-Track your writing per day
-How many words your wrote per day
-How many words you are currently at
---

## Tech Stack

- **Frontend:** React-ts
- **Backend:** Go
- **Desktop:** Wails

---

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) ≥ 1.21
- [Node.js](https://nodejs.org/) ≥ 18
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### Steps

```bash
# Clone the repository
git clone https://github.com/yourusername/booktracker.git
cd booktracker

# Install frontend dependencies
cd frontend
npm install

# Build frontend
npm run build

# Build the Wails app
cd ..
wails build

# Run the app
./build/bin/BookTracker
