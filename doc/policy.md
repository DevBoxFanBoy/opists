# Policies für das Issue-Tracking-System

## Organisationen, Projekte, Einstellungen, Administratoren, Benutzer verwalten
- Ein Administrator kann alle Ressourcen verwalten.
- Ein Benutzer kann nur die Ressourcen verwalten, die in seinem Besitz sind.

## Projekte verwalten
- Ein Projektadministrator kann Ansichten, Prozesse bzw. Workflows, Issues und Mitglieder verwalten.
- Ein Projektmitglied kann nur die Issues verwalten, die ihm zugewiesen sind.

## Prozesse verwalten
- Ein Prozessadministrator kann den Prozess und seine Elemente (z.B. Wasserfall, Meilensteine, Sprints, Backlogs, Boards) verwalten.

## Ansichten verwalten
- Ein Ansichtsadministrator kann Filter einstellen.

## Workflows verwalten
- Ein Workflowadministrator kann Statusänderungen an einem Issue festlegen.

## Zugriffe verwalten
- Ein Benutzer kann die Zugriffe auf die Ressourcen, die in seinem Besitz sind, teilen.

## Mitglieder verwalten
- Ein Administrator kann die Mitgliedschaft auf Projekt-, Organisation- oder Abteilungsebene einschränken.

## Administratoren verwalten
- Ein Systemadministrator kann Administratoren ernennen und entfernen.

## Einstellungen verwalten
- Ein Systemadministrator kann das Issue-Tracking-System global konfigurieren, z.B. Datenquellen der Benutzerverwaltung, Export/Import von Daten.

## Benutzer verwalten
- Ein Systemadministrator kann Benutzer anlegen, ändern und löschen.
- Ein Benutzer kann sein eigenes Profil ändern, aber nicht sehen, aus welcher Datenquelle er stammt oder diese ändern, sowie Profile von anderen Benutzern, sowie keine Benutzer löschen oder anlegen.

## jCasbin Policies
```ini
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub_rule, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = eval(p.sub_rule) && keyMatch2(r.obj, p.obj) && r.act == p.act
```
