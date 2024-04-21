# Golang OAuth with Fiber Framework

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

## Overview

This project is a web application built with Go (Golang) and the Fiber framework. It provides a simple and secure authentication mechanism using Google OAuth. Users can log in to the application using their Google accounts, and upon successful authentication, they are redirected to their profile page where they can view their user information fetched from Google's OAuth service.

The project aims to demonstrate how to integrate Google OAuth authentication into a Go web application using the Fiber framework. It provides a minimalistic yet functional example that can serve as a foundation for building more complex web applications with user authentication and authorization features.

## Features

- Google OAuth Integration: Seamless integration with Google OAuth for secure user authentication.
- Profile Page: Upon successful authentication, users are redirected to a profile page where they can view their user information retrieved from Google's OAuth service.
- Responsive Design: The web application utilizes responsive design principles to ensure a seamless experience across different devices and screen sizes.
- Session Management: Efficient session management to maintain user authentication state and provide a smooth browsing experience.
- Clean and Minimalistic UI: A clean and minimalistic user interface design to enhance usability and user experience.
- Error Handling: Robust error handling mechanisms to gracefully handle unexpected situations and provide informative error messages to users.

## Technologies Used

- Go (Golang)
- Fiber framework
- Google OAuth

## Installation

1. Clone the repository:

``
git clone https://github.com/anthonyq98/oauth-web-app-golang.git``

2. Navigate to the project directory:

``
  cd oauth-web-app-golang``
  
3. Install dependencies:

``
go mod tidy``

4. Configure your Google OAuth credentials:
- Go to the Google Developers Console
- Create a new project and enable the Google+ API
- Generate OAuth 2.0 credentials (client ID and client secret)
- Add authorized redirect URIs (e.g., http://localhost:8080/google_callback)

5. Create a .env file in the root of the project with your OAuth Credentials

6. Start the server with ``go run main.go``
7. Navigate to http://localhost:8080 on your browser
