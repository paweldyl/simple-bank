#!/bin/sh

set -e

echo "🚀 Starting the API entrypoint..."

echo "📦 Running DB migration..."
if /app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up; then
    echo "✅ Migration completed."
else
    echo "❌ Migration failed!"
    exit 1
fi

echo "🚀 Launching the API..."
exec "$@"