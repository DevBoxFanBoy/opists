Feature: get one or more issues from projects or by id
  Show all my issues and errors whenever project or issue not found.

  Scenario: get all project's issues
    When I send "GET" request to "/rest/v1/projects/DF/issues" with body ''
    Then the response code should be 200
    And the response should match json:
      """
      {
        "issues": [
          {
            "id": 0,
            "name": "New Bug",
            "description": "An error raise when...",
            "status": "open",
            "priority": "Highest",
            "projectKey": "DF",
            "components": [
              "DrinkOwnChampagne",
              "EatMyOwnApplication"
            ],
            "sprints": [
              "Sprint2"
            ],
            "estimatedPoints": 0,
            "estimatedTime": "0h",
            "affectedVersion": "1.2.3",
            "fixedVersion": "1.2.4"
          }
        ]
      }
      """

  Scenario: not found project key for issue
    When I send "GET" request to "/rest/v1/projects/Z0ZZZZZZZZZZZZZZZZZ0/issues" with body ''
    Then the response code should be 404
    And the response should match json:
      """
      {
        "code": 404,
        "message": "Project with Key Z0ZZZZZZZZZZZZZZZZZ0 not found!"
      }
      """

  Scenario: get issues by ID
    When I send "GET" request to "/rest/v1/projects/DF/issues/0" with body ''
    Then the response code should be 200
    And the response should match json:
      """
      {
        "id": 0,
        "name": "New Bug",
        "description": "An error raise when...",
        "status": "open",
        "priority": "Highest",
        "projectKey": "DF",
        "components": [
          "DrinkOwnChampagne",
          "EatMyOwnApplication"
        ],
        "sprints": [
          "Sprint2"
        ],
        "estimatedPoints": 0,
        "estimatedTime": "0h",
        "affectedVersion": "1.2.3",
        "fixedVersion": "1.2.4"
      }
      """

  Scenario: not found issue ID
    When I send "GET" request to "/rest/v1/projects/DF/issues/9223372036854775807" with body ''
    Then the response code should be 404
    And the response should match json:
       """
      {
        "code": 404,
        "message": "Issue with ID 9223372036854775807 not found!"
      }
      """

  Scenario: responses that A is invalid for ID
    When I send "GET" request to "/rest/v1/projects/DF/issues/A" with body ''
    Then the response code should be 400
    And the response should match json:
       """
      {
        "code": 400,
        "message": "ID A is invalid!"
      }
      """

  Scenario: responses that -1 is invalid for ID
    When I send "GET" request to "/rest/v1/projects/DF/issues/-1" with body ''
    Then the response code should be 400
    And the response should match json:
       """
      {
        "code": 400,
        "message": "ID -1 is invalid!"
      }
      """
#create issue
  Scenario: creates a new issue
    When I send "POST" request to "/rest/v1/projects/DF/issues" with body '{"name":"New Bug2","description":"2 description","status":"open","priority":"Highest","projectKey":"DF","components":["DrinkOwnChampagne","EatMyOwnApplication"],"sprints":["Sprint2"],"estimatedPoints":0,"estimatedTime":"0h","affectedVersion":"1.2.3","fixedVersion":"1.2.4"}'
    Then the response code should be 201
    And the response header "Location" match value "/DF/1"

#update issue

  Scenario: updates a existing issue
    When I send "PUT" request to "/rest/v1/projects/DF/issues" with body '{"id":0,"name":"New Bug","description":"Other description","status":"open","priority":"Highest","projectKey":"DF","components":["DrinkOwnChampagne","EatMyOwnApplication"],"sprints":["Sprint2"],"estimatedPoints":0,"estimatedTime":"0h","affectedVersion":"1.2.3","fixedVersion":"1.2.4"}'
    Then the response code should be 204

  Scenario: responses that project not equal to given project key
    When I send "PUT" request to "/rest/v1/projects/DF/issues" with body '{"id":0,"projectKey":"BB","name":"New Bug","description":"An error raise when...","status":"open","priority":"Highest","components":["DrinkOwnChampagne","EatMyOwnApplication"],"sprints":["Sprint2"],"estimatedPoints":0,"estimatedTime":"0h","affectedVersion":"1.2.3","fixedVersion":"1.2.4"}'
    Then the response code should be 400
    And the response should match json:
       """
      {
        "code": 400,
        "message": "Issue's ProjectKey BB is not equal to DF!"
      }
      """

  Scenario: not found project key for project
    When I send "PUT" request to "/rest/v1/projects/Z0ZZZZZZZZZZZZZZZZZ0/issues" with body '{"id":0,"projectKey":"Z0ZZZZZZZZZZZZZZZZZ0","name":"New Bug","description":"An error raise when...","status":"open","priority":"Highest","components":["DrinkOwnChampagne","EatMyOwnApplication"],"sprints":["Sprint2"],"estimatedPoints":0,"estimatedTime":"0h","affectedVersion":"1.2.3","fixedVersion":"1.2.4"}'
    Then the response code should be 404
    And the response should match json:
       """
      {
        "code": 404,
        "message": "Project with Key Z0ZZZZZZZZZZZZZZZZZ0 not found!"
      }
      """

  Scenario: not found issue ID for update the issue
    When I send "PUT" request to "/rest/v1/projects/DF/issues" with body '{"id":9223372036854775807,"projectKey":"DF","name":"New Bug","description":"An error raise when...","status":"open","priority":"Highest","components":["DrinkOwnChampagne","EatMyOwnApplication"],"sprints":["Sprint2"],"estimatedPoints":0,"estimatedTime":"0h","affectedVersion":"1.2.3","fixedVersion":"1.2.4"}'
    Then the response code should be 404
    And the response should match json:
       """
      {
        "code": 404,
        "message": "Issue with ID 9223372036854775807 not found!"
      }
      """

  Scenario: issue ID required for update
    When I send "PUT" request to "/rest/v1/projects/DF/issues" with body '{"projectKey":"DF","name":"New Bug","description":"An error raise when...","status":"open","priority":"Highest","components":["DrinkOwnChampagne","EatMyOwnApplication"],"sprints":["Sprint2"],"estimatedPoints":0,"estimatedTime":"0h","affectedVersion":"1.2.3","fixedVersion":"1.2.4"}'
    Then the response code should be 400
    And the response should match json:
       """
      {
        "code": 400,
        "message": "Issue's ID is required!"
      }
      """
