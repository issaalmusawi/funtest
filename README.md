# funtest
Dette skulle egentlig være en test-repo for meg selv, da filene til forrige "funtemps" repo begynte å bli mer og mer rotete. Og flere av kommentarene ble til tider mer forvirrende enn hjelpsomme. GO begynte i tillegg å slutte å gi meg output, og gi meg flere uforventede feilmeldinger (gopath/module etc), på tårds av at alle "go test-v"  ble gjennomført og godkjent. 

Resultatet ble til at jeg slettet forrige repo, slettet GO, og gjorde alt på nytt, da det å tukle med gopath og goroot gjorde ting enda mer komplisert. Jeg lagde "funTest" repo for å se om GO vile samarbeide bedre, og det funket. Nå tør jeg ikke gjøre noe som helst annerledes, da jeg har jobbet med dette i over en uke, dog med å fullføre hele funtemps(funfacts + conv). 
Forrige repo inneholdt både tester, funfacts og conv, mens denne "test"-repositoryen var kun ment til conv og main da det var den "nye" oppgaven.


Så, over til denne repo. Nå skal alle konverteringer fungere, da de ble testet fra forrige repo, og i tillegg får jeg ut riktig antall desimaler ut!!!!!! Noe som tilsynelatende virket helt umulig, da hver eneste gang jeg prøvde å bytte ut fmt.Println/fmt.Printf/fmt.Sprintf og alt av %.2f osv ikke fungerte og ga meg feil på feil.

Math.Round("funksjon"*100)/100) så ut til å gjøre magien. Så nå skal alt fungere, og igjen jeg tør ikke å gjøre noe som helst annereldes og velger å beholde denne "test" repo som min endelige innlevering. Når dere tester dette(om dere skal), så fungerer ikke kommandoen slik som "go run main.go funtemps -C -89.4 -Out K", men heller "go run main.go -C -89.4 -Out K", altså uten "funtemps", da jeg som sagt laget dette som en test, men nå som alt endelig funker så sier vi oss ferdig.



