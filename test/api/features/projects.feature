Feature: get one or more projects
  Show all my projects and errors whenever project not found.

  Scenario: get all projects
    Given As "admin" User
    When I send "GET" request to "/rest/v1/projects" with body ''
    Then the response code should be 200
    And the response should match json:
      """
      {
        "projects": [
          {
            "key": "DF",
            "name": "DogFooding",
            "description": "The Project used intern for Development.",
            "versions": [
              "1.2.3",
              "1.2.4"
            ],
            "components": [
              {
                "name": "DrinkOwnChampagne",
                "description": "Used intern for Development.",
                "versions": [
                  "DOC 1.0.0",
                  "DOC 1.0.1"
                ]
              }
            ],
            "sprints": [
              {
                "key": "Sprint2",
                "name": "Sprint 2 - Consume DogFooding",
                "start": "2020-11-12T07:00:34.7Z",
                "end": "2020-11-26T15:18:36.33Z"
              }
            ]
          }
        ]
      }
      """

  Scenario: get one project
    Given As "admin" User
    When I send "GET" request to "/rest/v1/projects/DF" with body ''
    Then the response code should be 200
    And the response should match json:
      """
      {
        "key": "DF",
        "name": "DogFooding",
        "description": "The Project used intern for Development.",
        "versions": [
          "1.2.3",
          "1.2.4"
        ],
        "components": [
          {
            "name": "DrinkOwnChampagne",
            "description": "Used intern for Development.",
            "versions": [
              "DOC 1.0.0",
              "DOC 1.0.1"
            ]
          }
        ],
        "sprints": [
          {
            "key": "Sprint2",
            "name": "Sprint 2 - Consume DogFooding",
            "start": "2020-11-12T07:00:34.7Z",
            "end": "2020-11-26T15:18:36.33Z"
          }
        ]
      }
      """

  Scenario: not found project key for project
    Given As "admin" User
    When I send "GET" request to "/rest/v1/projects/Z0ZZZZZZZZZZZZZZZZZ0" with body ''
    Then the response code should be 404
    And the response should match json:
      """
      {
        "code": 404,
        "message": "Project with Key Z0ZZZZZZZZZZZZZZZZZ0 not found!"
      }
      """

  #create project
  Scenario: create a new project via UI
    Given As "admin" User
    When I send "POST" request to "/rest/v1/projects" with body '{"key":"AAA","name":"Triple A Game","description":"The fancy 3D Game."}'
    Then the response code should be 201
    And the response header "Location" match value "/AAA"
    And the response should match json:
       """
       {
         "key": "AAA",
         "name": "Triple A Game",
         "description": "The fancy 3D Game."
       }
       """

  #update project
  Scenario: update a existing project
    Given As "admin" User
    When I send "PUT" request to "/rest/v1/projects/DF" with body '{"key": "DF","name": "DogFooding","description": "The Project used intern for Development. Updated"}'
    Then the response code should be 204

  Scenario: Bob want to get all projects
    Given As "bob" User
    When I send "GET" request to "/rest/v1/projects" with body ''
    Then the response code should be 200

  Scenario: Bob want to get project DF
    Given As "bob" User
    When I send "GET" request to "/rest/v1/projects/DF" with body ''
    Then the response code should be 200

  Scenario: Bob want to create a new project but has no permission
    Given As "bob" User
    When I send "POST" request to "/rest/v1/projects" with body '{"key":"AAA","name":"Triple A Game","description":"The fancy 3D Game."}'
    Then the response code should be 403

  Scenario: Bob want to update project DF but has no permission
    Given As "bob" User
    When I send "PUT" request to "/rest/v1/projects/DF" with body '{"key": "DF","name": "DogFooding","description": "The Project used intern for Development. Updated"}'
    Then the response code should be 403

  Scenario: Alice want to get all projects
    Given As "alice" User
    When I send "GET" request to "/rest/v1/projects" with body ''
    Then the response code should be 200

  Scenario: Alice want to get project DF
    Given As "alice" User
    When I send "GET" request to "/rest/v1/projects/DF" with body ''
    Then the response code should be 200

  Scenario: Alice want to create a new project
    Given As "alice" User
    When I send "POST" request to "/rest/v1/projects" with body '{"key":"BBB","name":"Triple A Game","description":"The fancy 3D Game."}'
    Then the response code should be 201

  Scenario: Alice want to update project DF
    Given As "alice" User
    When I send "PUT" request to "/rest/v1/projects/DF" with body '{"key": "DF","name": "DogFooding","description": "The Project used intern for Development. Updated"}'
    Then the response code should be 204