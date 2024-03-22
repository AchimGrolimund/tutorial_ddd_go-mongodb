# Projektstruktur und Codebeschreibung

Dieses Projekt verwendet die Domain-Driven Design (DDD) Architektur, die in Go implementiert ist. Die Hauptvorteile dieser Architektur sind die klare Trennung von Verantwortlichkeiten und die Fähigkeit, den Code leicht zu warten und zu erweitern.

## Domain

Die `domain` enthält die Kerngeschäftslogik und -regeln des Systems. In diesem Projekt ist `User` das Hauptdomain-Objekt. Es repräsentiert einen Benutzer in der Anwendung.

## Application

Die `application` ist verantwortlich für die Orchestrierung der Geschäftslogik. Sie ruft Methoden aus der Domänenschicht auf und stellt sicher, dass die Geschäftsregeln eingehalten werden. In diesem Projekt enthält die `application` den `UserService`, der Methoden zum Erstellen und Abrufen von Benutzern bereitstellt.

## Infrastructure

Die `infrastructure` ist verantwortlich für die Bereitstellung von technischen Details, die für die Implementierung der Anwendung erforderlich sind. In diesem Projekt enthält die `infrastructure` den `MongoUserRepository`, der die Interaktion mit der MongoDB-Datenbank verwaltet.

## cmd/go-ddd-mongodb

Dies ist der Einstiegspunkt der Anwendung. Hier wird die Anwendung initialisiert und gestartet. Es enthält den Hauptcode, der die `UserService`-Methoden aufruft, um Benutzer zu erstellen und abzurufen.

Die Aufteilung des Codes in diese Schichten hat mehrere Vorteile:

1. **Separation of Concerns**: Jede Schicht hat eine spezifische Aufgabe. Dies macht den Code einfacher zu verstehen und zu warten.

2. **Wiederverwendbarkeit**: Da die Geschäftslogik von den technischen Details getrennt ist, kann sie in verschiedenen Kontexten wiederverwendet werden.

3. **Testbarkeit**: Da jede Schicht unabhängig ist, kann sie einzeln getestet werden. Dies erleichtert das Schreiben von Unit-Tests.

4. **Flexibilität**: Da die technischen Details in der Infrastrukturschicht gekapselt sind, können sie leicht ausgetauscht werden. Zum Beispiel könnte man von MongoDB zu einer anderen Datenbank wechseln, ohne die Geschäftslogik zu ändern.