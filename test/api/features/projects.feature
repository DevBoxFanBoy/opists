Feature: get one or more projects
  Show all my projects and errors whenever project not found.

  Scenario: get all projects
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
    When I send "GET" request to "/rest/v1/projects/Z0ZZZZZZZZZZZZZZZZZ0" with body ''
    Then the response code should be 404
    And the response should match json:
      """
      {
        "code": 404,
        "message": "Project with Key Z0ZZZZZZZZZZZZZZZZZ0 not found!"
      }
      """
