[![Build Status](https://travis-ci.org/Khaledgarbaya/contentful.go.svg?branch=master)](https://travis-ci.org/Khaledgarbaya/contentful.go)
# contentful_go
contetnful sdk client built on top of golang

## This is not an official SDK, I built it to learn the GO programming Language

## DO NOT USE THIS FOR PRODUCTION 

## SUPER EXPERIMENTAL 

# About

[Contentful](https://www.contentful.com) is a content management platform for web applications, mobile apps and connected devices. It allows you to create, edit & manage content in the cloud and publish it anywhere via a powerful API. Contentful offers tools for managing editorial teams and enabling cooperation between organizations.

## Pre-requiset

To get content from Contentful, an app should authenticate with an with an OAuth bearer token.

You can create API keys using [Contentful's web interface](https://app.contentful.com). Go to the app, open the space that you want to access (top left corner lists all the spaces), and navigate to the APIs area. Open the API Keys section and create your first token. Done.

Don't forget to also get your Space ID.

For more information, check the Contentful's REST API reference on [Authentication](https://www.contentful.com/developers/docs/references/authentication/).

## Example Usage

```go
import(
  "fmt"
  
  "github.com/Khaledgarbaya/contentful_go"
)
// first create a contentful client
// token and space id you can get once you create a space 
contentful := contentful_go.New("SPACE_ID", "DELIVERY_TOKEN")

// get an entry 
entry, _ := contentful.GetEntry("ENTRY_ID")
fmt.Printf("found entry with the name %s", entry.Name)
```
