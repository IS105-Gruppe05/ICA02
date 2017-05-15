# ICA02
  
  Navnliste: Abdikani Gureye, Brede Knutsen Meli, Eirik Aanestad Fintland, Jan Kevin Henriksen, Mats Skjærvik, Mikael Kimerud, Morten Johan Mygland, Nils Fredrik Iselvmo Bjerk.

I denne oppgaven har alle bidratt. Ikke alle jobber på hver sin maskin, vi jobber av og til bare på noen av pcene og det er da de som laster opp det som er gjort. Derfor kan det vise at det er bare 5/8 som har lastet opp noe. 


Oppgave 1:

Specs ligger inne i docs (linken under).
https://docs.google.com/spreadsheets/d/1TUlQXl-SlkGzoc0ZrM6NQQLayWV9iOdv5WyQhnWOXKE/edit#gid=0 

Oppstart: POST →  BIOS/EFI
Prosesser: CPU, primærminne, dvs. ram og sekundærminne, dvs. HDD/SDD.

Maskinen starter ved å gjennomføre en power-on self test (POST), og deretter går videre til oppstart programmer i firmware på hovedkortet. Basic Input/Output System (BIOS) er en fellesbetegnelse for slike programmer. Det er vanlig å inkludere Extended Firmware Interface (EFI) etter 2010.

En prosess må ha tilgang til prosessor, primær- og sekundær minne(for å kunne overleve uten strøm) for å kjøre. Hver plattform har sin egen arkitektur for administrasjon av prosesser. OSX har /sbin/launchd prosess med ID 1, Linux /sbin/init (kan variere for ulike distros) med ID 1. OSX og Linux (fordi de er baserte på unix) har begge et hierarkisk system for prosesser. NT Windows har ntoskrnl.exe, og bruker ikke et direkte hierarkisk system på samme måte som OSX og Linux. 

I Windows kan man bruke taskmgr for å ha oversikt over hvilke prosesser som kjører og hvilke som ligger klar til bruk (idle). Når en prosess befinner seg “idle” kjøres en software med lavt aktivitet nivå, slik at prosessen er klar til å startes når den trengs. Man kan lukke ikke-kritiske prosesser via taskmgr. I Linux og OSX kan man bruke top, ps og  for å ha oversikt for prosesser. 

Prosesser:
Både linux serveren, windows og mac varierer prosessene som kjøres. Når du starter et program i windows vil nødvendige prosesser bli tatt i bruk. Under oppstart av et program vil flere prosesser bli brukt, for så å gå i idle tilstand igjen. Det varierer dog hvor mange prosesser hvert system benytter for programmer. Når det gjelder linux serveren virker den å være mer stabil enn windows, men varierer fremdeles litt.


Oppgave 2:

![Alt Text](https://github.com/IS105-Gruppe05/ICA02/blob/master/bilder/Oppgave%202/Skjermbilde%202017-05-12%20kl.%2013.35.36.png?raw=trueg)
