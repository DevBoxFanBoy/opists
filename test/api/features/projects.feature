Feature: get one or more projects
  Show all my projects.

  Scenario: should get all projects
    When I send "GET" request to "/v1/projects"
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

  Scenario: should get one project
    When I send "GET" request to "/v1/projects/DF"
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