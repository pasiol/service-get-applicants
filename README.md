# service-get-applicants

Primus opiskelijarekisterin yhteydessä hyödynnettävä mikropalvelu, jonka avulla siirretään hakija opiskelijarekisteriin ja luodaan väliaikainen Wilma-tunnus, jonka avulla opiskelija voi ottaa paikan vastaan ja täyttää perustietoja. Mikropalvelu luo Wilman käyttäjätunnukset-rekisteiin opiskelijan ilmoittamaan sähköpostiosoitteeseen perustuvan tunnuksen.

Palvelu on alunperin tarkoitettu ajettavaksi konttissa esim. Kubernetes klusterissa ajastettuna. Mikropalvelu on osa laajempaa sähköisen paikanvastaanoton kokonaisuutta.

![kaavio](images/sähköinen_paikanvastaanotto.png)

## Käyttö binäärinä

Kääntäminen

    make compile

---
    HOST=palvelimen_osoite PORT=NNNN ./bin/service-get-applicants

## Primus-tuontimääritykset

main.go

	archieveApplicantConfig = ""
	insertApplicantConfig   = ""

## Suodattimet hakija- ja opiskelijarekisteriin

query.go

- täydennä filterit riveille

    pq.Search = ""


## Salaisuudet

config/secrets.go