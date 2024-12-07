import sys
import yfinance as yf
import json

def fetch_stock(symbol):
    try:
        # Add ".NS" for NSE-listed stocks
        stock = yf.Ticker(symbol + ".NS")
        # Fetch the most recent day's data
        data = stock.history(period="1d")
        # Get the closing price for the latest day
        current_price = data["Close"][-1]
        print(current_price)
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print(json.dumps({"error": "Stock symbol is required"}))
        sys.exit(1)
    symbol = sys.argv[1]
    fetch_stock(symbol)
