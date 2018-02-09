# Schnelleinstieg in Go

Dies sind die Übungen für den Workshop "Schelleinstieg in Go" auf der oop 2018

## Nachtrag: Nötige Änderungen für Lösung 11 und 12

Wer Lösungen 11 und 12 schon geladen hat und einen Import Fehler bekommt, muss in den Dateien `11_solution_meta_search.go` und `12_solution.go` die Anweisung

    gitlab.innoq.com/daniel/go-workshop-utils/search/images/scraper

durch

    github.com/kesselborn/go-workshop-utils-oop2018/search/images/scraper

ersetzen und anschließend den Befehl

    go get github.com/kesselborn/go-workshop-utils-oop2018/search/images/scraper

ausführen. Kompiliert und gestartet wird das Programm anschließend mit

    go build -o metasearch
    ./metasearch -addr=":4242" -timeout=1s

Die Suchmaschine kann anschließend über http://localhost:4242 aufgerufen werden.

Die Slides vom Workshop sind in diesem Repository zu finden.

## Workshop Vorbereitungen

Teilnehmer sollten folgende Soft- und Hardware mitbringen:

* einen Laptop, auf dem Go installiert und nutzbar ist: siehe https://golang.org/doc/install für eine Anleitung
* eine funktionierende Kommandozeile (bash, zsh -- Windows Powershell funktioniert leider nicht)
* eine funktionierende Git-Installation
* einen Editor oder eine IDE zum Programmieren
* einen Kopie dieses Repositories

## Vorkenntnisse

Folgende Vorkenntnisse werden hilfreich sein, um den Workshop gut folgen zu können:

* Gute Kenntnisse in einer "Mainstream"-Programmiersprache wie Java, C# und C/C++.
* Grundkenntnisse über Nebenläufigkeitskonzepte in Programmiersprachen: Prozesse, Threads, Co-Routinen usw.
* Für die praktischen Übungen ein Notebook mit vorinstalliertem, aktuellem Go und Git sowie ein Texteditor und eine Shell.

## Nutzung der beiligenden Skripte / testen der Go-Installation

Bitte in der Kommandozeile

    ./next

aufrufen und den Anweisungen folgen. Falls eine Übung nicht beendet werden kann, hilft der Aufruf von

    ./cheat

Wenn Go korrekt installiert ist, sollte der Aufruf von `./next` in etwa folgende Ausgabe erzeugen:

     $ ./next
     M       README.md
     === RUN   TestTrue
     --- PASS: TestTrue (0.00s)
     PASS
     ok      _/Users/foobar/dev/innoq/go-workshop/go-workshop-exercise   0.006s
     [foobar 3df05d6] Commit exercise 01
      1 file changed, 3 insertions(+), 3 deletions(-)
     Fetching origin
     die nächste Übung ist noch nicht verfügbar

Viel Spaß
