import csv
import random
import socket

def generate_domain_ip_mapping(id):
    return f"example{id}.com.", f"172.1.0.{id}"

def main():
    # Create a new CSV file
    with open('domain_ip_mappings.csv', 'w', newline='') as file:
        writer = csv.writer(file)

        # Write the header row
        writer.writerow(['Domain', 'IP'])

        # Generate and write 100 domain-to-IP mappings
        for id in range(100):
            domain, ip = generate_domain_ip_mapping(id)
            writer.writerow([domain, ip])

    print("Domain-to-IP mappings generated and saved in domain_ip_mappings.csv")

if __name__ == '__main__':
    main()
