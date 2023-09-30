# AWS News Update Service for Discord

## Overview

![](https://github.com/ArthurMaverick/aws-news/blob/421f88d0f0b150ffc04ef1c88509c8ce66a15b0b/docs/Diagram.png)

The AWS News Notification Service for Discord is a tool that allows you to keep members of your Discord server informed about the latest news, announcements, and updates from Amazon Web Services (AWS). It utilizes a Discord webhook and a binary named "awsnews" to automate the update process.

## Key Features

- **Automatic Updates:** The service regularly checks AWS news sources for new announcements and updates.

- **Discord Notifications:** News updates are automatically sent to a specific channel on your Discord server via a webhook.

- **Customization:** You can configure the service to monitor specific AWS topics and send notifications only to selected channels.

- **Flexible Scheduling:** You can define how often the service checks for new updates, according to your needs.

## Requirements

- AWS account with permissions to access AWS news sources.

- A Discord server with a channel configured to receive notifications.

- The "awsnews" binary to perform configuration, updates, and destruction operations.

## Configuration

To set up the AWS News Update Service for Discord, follow these steps:

1. **Initial Configuration:**
    - Execute the "awsnews" binary with the "deploy" argument to initially configure the service. Provide necessary information, such as AWS credentials and the Discord webhook URL.

2. **Update:**
    - To update the service with new settings or Terraform code, execute the "awsnews" binary with the "update" argument.

3. **Destruction:**
    - If you wish to terminate the service and remove all associated resources, execute the "awsnews" binary with the "destroy" argument. Exercise caution when performing this operation as it will remove all resources.

## Usage

Once configured, the service will automatically check AWS news sources and send updates to the Discord channel specified via the webhook. You can access these updates at any time on the Discord server.

## Support

If you encounter issues or have questions about the AWS News Update Service for Discord, please contact our support team at [arthuracs18@gmail.com](mailto:arthuracs18@gmail.com).

## Conclusion

The AWS News Update Service for Discord is a valuable tool to keep your members informed about the latest AWS news automatically and conveniently. With the proper setup, you can ensure that everyone stays up-to-date on the latest developments in AWS.

---

Please make sure to customize this documentation to match the specific details of your service, including detailed instructions for each operation (configuration, update, and destruction). Additionally, you can add command examples for executing the "awsnews" binary and details on setting up the Discord webhook.