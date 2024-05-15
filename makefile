.PHONY: default run build test docs clean
APP_NAME=quake-iii-arena-log-decoder
VERSION=v1

help:
	@clear
	@echo "                                                                                                              "
	@echo "                                                                                                              "
	@echo "                                                                                                              "
	@echo "                                                                                                              "
	@echo "                                                                                                              "
	@echo "                                                       (                                                      "
	@echo "                                                       /                                                      "
	@echo "                                                      //                                                      "
	@echo "                                                      ///                                                     "
	@echo "                                                      ///                                                     "
	@echo "                                                     ////                                                     "
	@echo "                                                     /////                                                    "
	@echo "                                                     /////                                                    "
	@echo "                                                     /////                                                    "
	@echo "                     /(((((((                       (//////                        ((((((/                    "
	@echo "               ((((////,,                           ///////                           ,,///(((((              "
	@echo "            ((////,,                                ///////                                ,,////((           "
	@echo "          (/////,,                                 /////////                                 , /////(         "
	@echo "         ///////(/                                 /////////                                 (///////(        "
	@echo "         /////////(((/                             /////////                             ((((/////////        "
	@echo "        ,//////////////((((((((                   ///////////                  /(((((((//////////////         "
	@echo "         ,,////////////////////////(((((((((((    ///////////    (((((((((((////////////////////////,         "
	@echo "             ,////////////////////////////////    ///////////    ///////////////////////////////              "
	@echo "                  ,,//////////////////////////    ///////////    //////////////////////////,,                 "
	@echo "                          ,,,,////////////////     /////////,    ////////////////,,,                          "
	@echo "                                     ,,///////    ,/////////     ///////,,                                    "
	@echo "                                        //////     /////////     //////,                                      "
	@echo "                                       ,//////     ////////,     //////                                       "
	@echo "                                         /////     ,///////      /////,                                       "
	@echo "                                         /////      ///////      /////                                        "
	@echo "                                          ////      //////       ////,                                        "
	@echo "                                         ,////      ,/////,      ////                                         "
	@echo "                                          ////       /////       ///,                                         "
	@echo "                                          ,///       ////        ///                                          "
	@echo "                                           ///       ,///,       //,                                          "
	@echo "                                           ,//        ///        //                                           "
	@echo "                                            //        //         /,                                           "
	@echo "                                            ,/         /,        /                                            "
	@echo "                                             /         /                                                      "
	@echo "                                             ,                                                                "
	@echo ""
	@echo ""
	@echo "     - make"
	@echo "         - build: realiza o build da aplicação"
	@echo "         - run: roda aplicação via docker"
	@echo "                                                                                   By Lucas Gonçalves Bonilla"

make run-local:
	go run cmd/app/main.go

build:
	docker build -t $(APP_NAME):$(VERSION) .

run:
	docker run -it $(APP_NAME):$(VERSION)
	mkdir -p files/out
	docker cp $$(docker ps -q -l -f ancestor=$(APP_NAME):$(VERSION)):files/out/qgames.json ./files/out/qgames.json

make build-run: build run