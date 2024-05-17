# Quake III Arena log decoder

	                                               (
	                                               /
	                                              //
	                                              ///
	                                              ///
	                                             ////
	                                             /////
	                                             /////
	                                             /////
	             /(((((((                       (//////                        ((((((/
	       ((((////,,                           ///////                           ,,///(((((
	    ((////,,                                ///////                                ,,////((
	  (/////,,                                 /////////                                 , /////(
	 ///////(/                                 /////////                                 (///////(
	 /////////(((/                             /////////                             ((((/////////
	,//////////////((((((((                   ///////////                  /(((((((//////////////
	 ,,////////////////////////(((((((((((    ///////////    (((((((((((////////////////////////,
	     ,////////////////////////////////    ///////////    ///////////////////////////////
	          ,,//////////////////////////    ///////////    //////////////////////////,,
	                  ,,,,////////////////     /////////,    ////////////////,,,
	                             ,,///////    ,/////////     ///////,,
	                                //////     /////////     //////,
	                               ,//////     ////////,     //////
	                                 /////     ,///////      /////,
	                                 /////      ///////      /////
	                                  ////      //////       ////,
	                                 ,////      ,/////,      ////
	                                  ////       /////       ///,
	                                  ,///       ////        ///
	                                   ///       ,///,       //,
	                                   ,//        ///        //
	                                    //        //         /,
	                                    ,/         /,        /
	                                     /         /
	                                     ,

O projeto em questão opera sobre a leitura de arquivos no formato .log armazenados em um diretório local, gerando uma saída detalhada por partida. Tal saída inclui totais de mortes de cada jogador e os tipos de mortes registrados em cada jogo. Os resultados são apresentados em uma tabela no terminal, bem como em um arquivo .json em uma pasta designada.

Para viabilizar o desenvolvimento, optou-se pela adoção da arquitetura hexagonal, dada sua propensão à fácil separação de responsabilidades, reutilização de código, testabilidade e escalabilidade. A coleta de dados a partir do arquivo .log foi implementada por meio de uma goroutine, enquanto múltiplas goroutines foram empregadas para o consumo desses dados. A escalabilidade do sistema é assegurada por uma variável de ambiente presente no arquivo config.toml.

O algoritmo de leitura e geração dos dados de partida e morte foi concebido sob a ótica de um modelo produtor/consumidor, onde uma goroutine é responsável por gerar os dados enquanto outras consomem esses dados por meio de um canal de comunicação. Tal abordagem confere ao sistema uma eficiente operação paralela na geração e no consumo dos dados.

Em pontos críticos do projeto, a utilização de mutexes se mostrou necessária para prevenir condições de corrida (race conditions) entre os recursos do sistema.

Para garantir a robustez do sistema, foram desenvolvidos testes abrangentes, concentrando-se em áreas de maior criticidade. Atualmente, o projeto ostenta uma cobertura total de testes de 78.6%.

## Como rodar o projeto

O projeto foi implementado pensando na utilização de um make file. Para tanto:

### Comandos
``` make help```: Mostra o cabeçalho do projeto com os possíveis comandos

``` make run-local```: Roda a aplicação via golang local. (É necessário que a variável de ambiente type esteja configurada para **dev**)

``` make build```: Realiza o build da aplicação via docker. (É necessário que a variável de ambiente type esteja configurada para **prd**)

``` make run```: Roda a aplicação do build

``` make build-run```: Realiza o build da aplicação via docker e o run no mesmo comando. (É necessário que a variável de ambiente type esteja configurada para **prd**)

```make test```: Realiza os testes da aplicação via docker

```test-local```: Realiza testes locais. Necessita do pacote gocov instalado

### Saídas

As saídas do algoritmo se dão em duas partes. No terminal através de uma tabela separada para cada partida no log e um arquivo qgames.json na pasta *src/files/out/*

#### Exemplo de saída via terminal

```
+---------+-------------+---------+-------+--------+
| GAME ID | TOTAL KILLS | PLAYERS | KILLS | DEATHS |
+---------+-------------+---------+-------+--------+
| game_1  |           0 |         |       |        |
+---------+-------------+---------+-------+--------+

+---------+-------------+---------------------+----------------------------+--------------------------------+
| GAME ID | TOTAL KILLS |       PLAYERS       |           KILLS            |             DEATHS             |
+---------+-------------+---------------------+----------------------------+--------------------------------+
| game_2  |          11 | Isgalamido, Mocinha | Isgalamido: -5, Mocinha: 0 | MOD_TRIGGER_HURT: 7,           |
|         |             |                     |                            | MOD_ROCKET_SPLASH: 3,          |
|         |             |                     |                            | MOD_FALLING: 1                 |
+---------+-------------+---------------------+----------------------------+--------------------------------+

+---------+-------------+--------------------------------+--------------------------------+--------------------------------+
| GAME ID | TOTAL KILLS |            PLAYERS             |             KILLS              |             DEATHS             |
+---------+-------------+--------------------------------+--------------------------------+--------------------------------+
| game_3  |           4 | Mocinha, Isgalamido, Zeh, Dono | Mocinha: 0, Isgalamido: 1,     | MOD_ROCKET: 1,                 |
|         |             | da Bola                        | Zeh: -2, Dono da Bola: -1      | MOD_TRIGGER_HURT: 2,           |
|         |             |                                |                                | MOD_FALLING: 1                 |
+---------+-------------+--------------------------------+--------------------------------+--------------------------------+

...
```

#### Exemplo de saída via arquiv .json

```
{
    "game_1": {
        "total_kills": 0,
        "players": [],
        "kills": {},
        "deaths": {}
    },
	    "game_2": {
        "total_kills": 11,
        "players": [
            "Isgalamido",
            "Mocinha"
        ],
        "kills": {
            "Isgalamido": -5,
            "Mocinha": 0
        },
        "deaths": {
            "MOD_FALLING": 1,
            "MOD_ROCKET_SPLASH": 3,
            "MOD_TRIGGER_HURT": 7
        }
    },
	"game_3": {
        "total_kills": 4,
        "players": [
            "Mocinha",
            "Isgalamido",
            "Zeh",
            "Dono da Bola"
        ],
        "kills": {
            "Dono da Bola": -1,
            "Isgalamido": 1,
            "Mocinha": 0,
            "Zeh": -2
        },
        "deaths": {
            "MOD_FALLING": 1,
            "MOD_ROCKET": 1,
            "MOD_TRIGGER_HURT": 2
        }
    },
	...
}
```