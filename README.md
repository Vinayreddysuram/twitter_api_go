# Twitter API Project

## Overview

This project uses Go to interact with the Twitter API. It provides functionality to post tweets and delete them based on their IDs. The project demonstrates how to authenticate with the Twitter API and perform basic operations like posting and deleting tweets.

## Features

- **Post a Tweet**: Allows users to post a new tweet with a specified text.
- **Delete a Tweet**: Allows users to delete a tweet by providing its ID.
- **Error Handling**: Includes robust error handling for various API responses, such as unauthorized access, rate limits, and tweet not found errors.

## Prerequisites

- Go 1.16 or higher
- Twitter Developer account and API keys
- Basic knowledge of Go and working with APIs

## Setting Up a Twitter Developer Account

To use the Twitter API, you'll need to create a Twitter Developer account and obtain your API credentials. Follow these steps to set up your account:

1. **Sign Up for a Twitter Developer Account**:
   - Go to the [Twitter Developer Portal](https://developer.twitter.com/).
   - Click on the **Sign In** button in the top right corner and log in with your Twitter account.
   - If you don't have a Twitter account, you'll need to create one first.

2. **Apply for Developer Access**:
   - After signing in, click on the **Apply** button for a developer account.
   - Complete the application form. You will need to provide details such as:
     - Your intended use case for the Twitter API.
     - How you plan to use the data obtained from the API.
     - Any specific features you intend to access.
   - Review and accept the Developer Agreement and Twitter’s policies.

3. **Wait for Approval**:
   - Once you've submitted your application, Twitter will review it. This process may take some time.
   - You will receive an email notification about the status of your application. If approved, you'll gain access to the Developer Portal.

4. **Create a Project and App**:
   - After approval, log in to the [Twitter Developer Portal](https://developer.twitter.com/en/portal/dashboard).
   - Click on the **Projects & Apps** tab.
   - Click on the **New Project** button to create a new project.
   - Fill in the project details and click **Next**.
   - Create a new App under the project by following the prompts.

5. **Get Your API Credentials**:
   - Once your app is created, go to the **Keys and tokens** tab.
   - Here you will find your **API Key**, **API Key Secret**, **Access Token**, and **Access Token Secret**.
   - You can regenerate tokens if needed, but ensure to keep them secure.

## Project Details and Executions

1. **Posting the Tweet**:
   - The program allows users to input text for a new tweet and posts it to Twitter.
  
<img width="625" alt="image" src="https://github.com/user-attachments/assets/27bb0c1c-91e8-4a66-a558-9a67ed7626a9">

   - **Verification**: After posting, the user can verify the tweet appears on their Twitter feed.
  
<img width="625" alt="image" src="https://github.com/user-attachments/assets/78a953f5-2c3b-4b18-bce2-fc1fb7f37490">


2. **Delete the Tweet**:
   - Users can enter the tweet ID of the tweet they want to delete.
   - The program sends a request to delete the specified tweet.
  
<img width="625" alt="image" src="https://github.com/user-attachments/assets/60b79c9d-bd34-45af-8362-e6804e23ec01">

   - **Verification**: Users can check their Twitter feed to confirm the tweet has been deleted.

<img width="625" alt="image" src="https://github.com/user-attachments/assets/640edc1d-5432-46dc-bb68-7f5a9457bf83">


4. **Error Handling**:
   - The program includes robust error handling for various scenarios, including unauthorized access, rate limits, and tweet not found errors.
   - Users receive informative error messages based on the API responses.
  
<img width="625" alt="image" src="https://github.com/user-attachments/assets/e89301a5-93ed-4897-aee8-e14299f6e1ec">



<img width="625" alt="image" src="https://github.com/user-attachments/assets/63b36b99-f7b6-4cb7-b24d-ade60444c82c">
