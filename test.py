import requests

# Define the URL of the API endpoint
url = "http://localhost:8080/deleteurl"

# Define the form data as a dictionary
form_data = {
    "ID": "1"
}

# Send a POST request with the form data
response = requests.post(url, data=form_data)

# Print the response text
print(response.text)

