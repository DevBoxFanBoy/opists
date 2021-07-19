/*!
    * Start Bootstrap - SB Admin v6.0.2 (https://startbootstrap.com/template/sb-admin)
    * Copyright 2013-2020 Start Bootstrap
    * Licensed under MIT (https://github.com/StartBootstrap/startbootstrap-sb-admin/blob/master/LICENSE)
    */
(function ($) {
    "use strict";

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
            dataType: "json",
            contentType: "application/json"
        });
    })
})(jQuery);
