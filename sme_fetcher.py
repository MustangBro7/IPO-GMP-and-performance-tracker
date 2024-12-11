import requests
import json
import sys
def fetch_nse_data(symbol):
    url = "https://www.nseindia.com/api/live-analysis-emerge"

    # Headers to mimic a browser
    headers = {
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
        "Accept": "application/json, text/plain, */*",
        "Accept-Language": "en-US,en;q=0.9",
        "Referer": "https://www.nseindia.com/market-data/sme-market",
        "Connection": "keep-alive",
    }

    # Make a GET request
    session = requests.Session()
    try:
        # Perform an initial request to get cookies
        session.get("https://www.nseindia.com", headers=headers)
        
        # Now fetch the actual data
        response = session.get(url, headers=headers)
        response.raise_for_status()  # Raise an exception for HTTP errors

        # Parse the JSON data
        data = response.json()

        # Filter data for the provided symbol
        for item in data.get("data", []):
            if item.get("symbol").upper() == symbol.upper():

                print(item.get('lastPrice'))
                return
        
        print("Symbol not found.")
    
    except requests.exceptions.RequestException as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print(json.dumps({"error": "Stock symbol is required"}))
        sys.exit(1)
    symbol = sys.argv[1]
    fetch_nse_data(symbol)
