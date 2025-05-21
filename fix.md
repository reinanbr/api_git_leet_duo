### V 0.1.2b *[30/03/25]*

***Problemas em aberto nesta versão:***

* Organização melhor das apis aqui presente.
  * Git
    * full info user
      * username
      * photoAvatar
      * calendarContriBuition
        * repo info
          * langs percent
          * name
          * date init
          * last commit date
    * streak
    * percent lang
    * total contrib
  * leet
    * user info
    * streak
    * total sub
  * duo
    * user info
    * streak
    * total xp
    * percent lang xp

Neste primeiro monento, focar apenas nas coisas que serão importantes para o painel. Mais tarde, especializar melhor as informações e criar exemplos de como usar essas api.


**06.04.25 23:17**

    Acredito que o contribuitions deve seguir o modelo de

`{
    $start = "$year-01-01T00:00:00Z";
    $end = "$year-12-31T23:59:59Z";
    return "query {
        user(login: \"$user\") {
            createdAt
            contributionsCollection(from: \"$start\", to: \"$end\") {
                contributionYears
                contributionCalendar {
                    weeks {
                        contributionDays {
                            contributionCount
                            date
                        }
                    }
                }
            }
        }
    }";
}`

Com isso, cada ponto no git, teria seu proprio query, e não apenas para todos.


**13:11 12/04/25, casa de Denise**

  Acredito que a melhor documentação para essa API, seja
    / api
      /git
      /duo
      /leet
  Em  que retorne as informações essenciais de cada site.
