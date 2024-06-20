// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/id": {
            "get": {
                "description": "Fetches the unique identifier for a station by its name from the StationsMap. This endpoint performs a case-insensitive search and ignores spaces in the station name.\nThe Ids have format Line prefix that has the format \"SU\" followed by an abbreviation of the station name. For example \"SU-A\" for the station \"Alexanderplatz\".",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "City Data"
                ],
                "summary": "Retrieve Station ID by Name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Station name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The station id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Station not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/recent": {
            "get": {
                "description": "Fetches the most recent ticket inspector reports from the database and returns them as a JSON array.\nIf there are not enough recent reports, the endpoint will fetch additional historic reports to meet the required amount.\nThe required amount is determined by the current time of the day and the day of the week, ensuring the most relevant and timely information is provided to the user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basic Functions"
                ],
                "summary": "Retrieve information about recent ticket inspector reports",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Standard HTTP header used to make conditional requests; the response will include the requested data only if it has changed since this date and time.",
                        "name": "If-Modified-Since",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A JSON array of ticket inspector information, each entry includes details such as timestamp, station, direction, line, and historic flag.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/utils.TicketInspectorResponse"
                            }
                        }
                    },
                    "304": {
                        "description": "Returns an empty response indicating that the requested data has not changed since the time provided in the 'If-Modified-Since' header."
                    },
                    "500": {
                        "description": "Internal Server Error: An error occurred while processing the request.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/station": {
            "get": {
                "description": "Fetches the name of a station by its unique identifier from the StationsMap.\nThe Ids have format Line prefix that has the format \"SU\" followed by an abbreviation of the station name. For example \"SU-A\" for the station \"Alexanderplatz\".",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "City Data"
                ],
                "summary": "Retrieve Name by Station ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Station Id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The station id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Error getting station name",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "utils.Coordinates": {
            "type": "object",
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                }
            }
        },
        "utils.Station": {
            "type": "object",
            "properties": {
                "coordinates": {
                    "$ref": "#/definitions/utils.Coordinates"
                },
                "id": {
                    "type": "string"
                },
                "lines": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "utils.TicketInspectorResponse": {
            "type": "object",
            "properties": {
                "direction": {
                    "$ref": "#/definitions/utils.Station"
                },
                "isHistoric": {
                    "type": "boolean"
                },
                "line": {
                    "description": "String is used so that it can easily be handled by the frontend",
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "station": {
                    "$ref": "#/definitions/utils.Station"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "FreiFahren API Documentation",
	Description:      "API for the FreiFahren project, responsible for collecting and serving data about ticket inspectors on public transport.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
