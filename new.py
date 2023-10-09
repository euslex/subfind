import requests

# Function to get the HTTP status code of a subdomain
def get_status_code(subdomain):
    url = f"http://{subdomain}"  # You can modify this to use HTTPS if needed
    try:
        response = requests.get(url)
        return response.status_code
    except requests.exceptions.RequestException:
        return "Error"

# Read subdomains from the file
with open("sub_domain.txt", "r") as file:
    subdomains = file.read().splitlines()

# Iterate through the subdomains and get their status codes
for subdomain in subdomains:
    status_code = get_status_code(subdomain)
    print(f"{subdomain}: {status_code}")
