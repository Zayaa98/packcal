The Pack Calculator is a Go application that calculates the number of packs required to fulfill customer orders for a product available in various pack sizes. The application adheres to three main requirements:

Only whole packs can be sent. Packs cannot be broken open.
Within the constraints of Rule 1 above, send out no more items than necessary to fulfill the order.
Within the constraints of Rules 1 & 2 above, send out as few packs as possible to fulfill each order.
This README provides an overview of the application, instructions on how to use it, and an explanation of its components.

Table of Contents
Getting Started
Prerequisites
Installation
Usage
API Endpoint
Example Input and Output
Development
Contributing
License
Getting Started
Prerequisites
Before using the Pack Calculator, ensure you have the following prerequisites:

Go: The application is written in Go, so you need to have Go installed on your system.
Git: You may want to clone the application's repository using Git.
Installation
Clone the application's repository from GitHub:
bash
Copy code
git clone https://github.com/your-username/pack-calculator.git
Change your working directory to the application's folder:
bash
Copy code
cd pack-calculator
Load the pack sizes configuration: The application simulates loading pack sizes from a configuration file or database. You can modify the loadPackSizes function in main.go to load pack sizes as needed.

Build and run the application:

bash
Copy code
go build
./pack-calculator
The application will start and listen for incoming requests on port 8080.

Usage
API Endpoint
The application provides an HTTP API endpoint that allows you to calculate the number of packs needed to fulfill customer orders.

Endpoint: /calculate-packs
Method: POST
Request Body: JSON object with an orderQuantity field representing the number of items in the order.
Example Input and Output
Here are some examples of using the API:

Input:

json
Copy code
{
  "orderQuantity": 251
}
Output:

json
Copy code
{
  "packsNeeded": {
    "500": 1
  }
}
Input:

json
Copy code
{
  "orderQuantity": 501
}
Output:

json
Copy code
{
  "packsNeeded": {
    "500": 1,
    "250": 1
  }
}
Input:

json
Copy code
{
  "orderQuantity": 12001
}
Output:

json
Copy code
{
  "packsNeeded": {
    "5000": 2,
    "2000": 1,
    "250": 1
  }
}
Development
To make changes to the Pack Calculator, you can follow these steps:

Make sure you have Go installed on your system.

Clone the application's repository from GitHub as described in the Installation section.

Modify the code in the main.go file, particularly the calculatePacks function to meet specific requirements or logic changes.

Build and run the application to test your changes.
