# Groupie Tracker

![Groupie Tracker](static/img/branding.png)

## Overview

Groupie Tracker is a web application that allows users to explore and discover information about various artists, their albums, concert locations, and more. The application provides a modern, responsive interface with search functionality to help users find specific artists or information.

## Features

- **Artist Discovery**: Browse through a collection of artists with their images and basic information
- **Search Functionality**: Search for artists by name, band members, first album, creation date, or concert locations
- **Artist Details**: View detailed information about each artist, including:
  - Band members
  - First album release date
  - Creation date
  - Concert locations
  - Concert dates
- **Interactive UI**: Modern interface with hover effects and animations

## Technologies Used

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS, JavaScript
- **API**: Groupie Tracker API (https://groupietrackers.herokuapp.com/api)

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/bytebabee/groupie-tracker.git
   ```

2. Navigate to the project directory:
   ```
   cd groupie-tracker
   ```

3. Run the application:
   ```
   go run main.go
   ```

4. Open your browser and navigate to the URL displayed in the terminal (http://localhost:PORT)

## Project Structure

- `main.go`: Entry point of the application
- `handler/`: Contains HTTP handlers for different routes
- `models/`: Data structures for the application
- `templates/`: HTML templates
- `static/`: Static assets (CSS, JavaScript, images)
- `utilities/`: Helper functions and utilities

## API Endpoints

The application uses the following API endpoints:

- Artists: `https://groupietrackers.herokuapp.com/api/artists`
- Locations: `https://groupietrackers.herokuapp.com/api/locations`
- Dates: `https://groupietrackers.herokuapp.com/api/dates`
- Relations: `https://groupietrackers.herokuapp.com/api/relation`

## Authors

- Created by Irem ([@iremnurc](https://github.com/iremnurc)) and Taha ([@tcavuslu](https://github.com/tcavuslu)) 
