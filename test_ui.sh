#!/bin/bash
# Test UI script - regenerates, rebuilds, and restarts the UI

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
FORGEXML_DIR="/home/mmcnew/repos/forgexml"
SCAP_DIR="/home/mmcnew/repos/forgexml-scap"

echo "=== ForgeXML-SCAP UI Test Script ==="
echo ""

# Step 1: Kill any existing scap-ui processes
echo "[1/5] Stopping existing processes..."
pkill -f "scap-ui" 2>/dev/null || true
sleep 1

# Step 2: Regenerate code from forgexml
echo "[2/5] Regenerating code..."
cd "$FORGEXML_DIR"
go run ./cmd/forgexml/main.go -config "$SCAP_DIR/forgexml.json" -generate-ui

# Step 3: Clean and rebuild
echo "[3/5] Rebuilding application..."
cd "$SCAP_DIR"
rm -f bin/scap-ui
go build -o bin/scap-ui ./cmd/ui

# Step 4: Start the UI
echo "[4/5] Starting UI server..."
cd "$SCAP_DIR"
nohup ./bin/scap-ui > logs/ui.log 2>&1 &
UI_PID=$!
disown $UI_PID
sleep 2

# Step 5: Verify it's running
echo "[5/5] Verifying..."
if curl -s http://localhost:8080/api/types/root > /dev/null 2>&1; then
    echo ""
    echo "=== SUCCESS ==="
    echo "UI is running at http://localhost:8080"
    echo "PID: $UI_PID"
    echo ""
    echo "Root types available:"
    curl -s http://localhost:8080/api/types/root | jq -r '.[]' 2>/dev/null || curl -s http://localhost:8080/api/types/root
    echo ""
    echo "To stop: kill $UI_PID"
else
    echo ""
    echo "=== FAILED ==="
    echo "UI failed to start"
    exit 1
fi
