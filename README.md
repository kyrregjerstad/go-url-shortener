# URL Shortener

A simple URL shortener built while learning Go. The project consists of a Go backend API and a SvelteKit frontend, using turborepo.

## Features

- Shorten long URLs into easy-to-share links
- Track visit analytics for each shortened URL
- View browser and referrer statistics
- Modern, responsive UI built with SvelteKit and Tailwind CSS

## Project Structure

```
.
├── apps
│   └── frontend      # SvelteKit frontend
└── packages
    └── api          # Go backend API
```

## Development

1. start all services:

```bash
turbo dev
```

## Learning Goals

This project was built to learn:

- Go backend development
- RESTful API design
- Database interactions in Go
