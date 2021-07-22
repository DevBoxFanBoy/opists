/*!
    * Start Bootstrap - SB Admin v6.0.2 (https://startbootstrap.com/template/sb-admin)
    * Copyright 2013-2020 Start Bootstrap
    * Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-sb-admin/blob/master/LICENSE)
    */
(function ($) {
    "use strict";

    // https://coderwall.com/p/ostduq/escape-html-with-javascript
    // List of HTML entities for escaping.
    var htmlEscapes = {
        '&': '&amp;',
        '<': '&lt;',
        '>': '&gt;',
        '"': '&quot;',
        "'": '&#x27;',
        '/': '&#x2F;'
    };

    // Regex containing the keys listed immediately above.
    var htmlEscaper = /[&<>"'\/]/g;

    // Escape a string for HTML interpolation.
    function escapeHtml(string) {
        return ('' + string).replace(htmlEscaper, function (match) {
            return htmlEscapes[match];
        });
    };

    // Add active state to sidbar nav links
    var path = window.location.href; // because the 'href' property of the DOM element is the absolute path
    $("#layoutSidenav_nav .sb-sidenav a.nav-link").each(function () {
        if (this.href === path) {
            $(this).addClass("active");
        }
    });

    // Toggle the side navigation
    $("#sidebarToggle").on("click", function (e) {
        e.preventDefault();
        $("body").toggleClass("sb-sidenav-toggled");
    });

    // Create new Project
    $("#createProject").on("click", function (e) {
        e.preventDefault()
        let newProject = JSON.stringify({
            key: $("#prjKeyField").val(),
            name: $("#prjNameField").val(),
            description: $("#prjDescriptionField").val()
        });
        $.ajax({
            type: "POST",
            url: "/rest/v1/projects",
            data: newProject,
            success: function () {
                location.reload();
            },
            error: function (xhr) {
                //TODO validate xhr, xhr.responseJSON, ...
                //console.info(xhr);
                //console.info(xhr.responseJSON);
                //console.info(xhr.responseJSON.code);
                //console.info(xhr.responseJSON.message);
                //console.info(xhr.status);
                $("#createProjectModalErrors").append(
                    "<div class=\"alert alert-danger alert-dismissible fade show\" role=\"alert\">\n" +
                    escapeHtml(xhr.responseJSON.message) + "\n" +
                    "<button type=\"button\" class=\"close\" data-dismiss=\"alert\" aria-label=\"Close\">\n" +
                    "<span aria-hidden=\"true\">&times;</span>\n" +
                    "</button>\n" +
                    "</div>"
                )
            },
            dataType: "json",
            contentType: "application/json"
        });
    })
    // Create new Issue
    $("#createIssue").on("click", function (e) {
        e.preventDefault()
        let projectKey = $("#issuePrjKeyField").val()
        let newIssue = JSON.stringify({
            name: $("#issueNameField").val(),
            description: $("#issueDescriptionField").val()
        });
        $.ajax({
            type: "POST",
            url: "/rest/v1/projects/" + projectKey + "/issues",
            data: newIssue,
            success: function () {
                location.reload();
            },
            error: function (xhr) {
                //TODO validate xhr, xhr.responseJSON, ...
                //console.info(xhr);
                //console.info(xhr.responseJSON);
                //console.info(xhr.responseJSON.code);
                //console.info(xhr.responseJSON.message);
                //console.info(xhr.status);
                $("#createIssueModalErrors").append(
                    "<div class=\"alert alert-danger alert-dismissible fade show\" role=\"alert\">\n" +
                    escapeHtml(xhr.responseJSON.message) + "\n" +
                    "<button type=\"button\" class=\"close\" data-dismiss=\"alert\" aria-label=\"Close\">\n" +
                    "<span aria-hidden=\"true\">&times;</span>\n" +
                    "</button>\n" +
                    "</div>"
                )
            },
            dataType: "json",
            contentType: "application/json"
        });
    })
})(jQuery);
