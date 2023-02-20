# funtest
Dette skulle egentlig være en test-repo for meg selv, da filene til forrige "funtemps" repo begynte å bli mer og mer rotete. GO begynte i tillegg å slutte å gi meg output, og gi meg flere uforventede feilmeldinger (gopath/module etc), på tårds av at alle "go test-v"  ble gjennomført og godkjent. 

Resultatet ble til at jeg slettet forrige repo, slettet GO, og gjorde alt på nytt, da det å tukle med gopath og goroot gjorde ting enda mer komplisert. Jeg lagde "funTest" repo for å se om GO vile samarbeide bedre, og det funket. Nå tør jeg ikke gjøre noe som helst annerledes, da jeg har jobbet med dette i over en uke, dog med å fullføre hele funtemps(funfacts + conv). 
Forrige repo inneholdt både tester, funfacts og conv, mens denne "test"-repositoryen var kun ment til conv og main da det var den "nye" oppgaven.


Så, over til denne repo. Nå skal alle konverteringer fungere, da de ble testet fra forrige repo, og i tillegg får jeg ut riktig antall desimaler ut!!!!!! Noe som tilsynelatende virket helt umulig, da hver eneste gang jeg prøvde å bytte ut fmt.Println/fmt.Printf/fmt.Sprintf og alt av %.2f osv ikke fungerte og ga meg feil på feil.

Math.Round("funksjon"*100)/100) så ut til å gjøre magien. Så nå skal alt fungere, og igjen jeg tør ikke å gjøre noe som helst annereldes og velger å beholde denne "test" repo som min endelige innlevering. 

Testing: 
kommandolinje instruksjon for testing av konverteringer av temperatur, eksemepl: "./funtest -C 25 -out F" (returnerer 25 C i F)

//For funfacts(ikke helt ferdig) : "./funtest -funfacts Terra" (returnerer fakta til jorden)

NY for funfacts: "./funtest -t (temptype, foreløpig kun C) -funfact (Sun,Luna,Terra)