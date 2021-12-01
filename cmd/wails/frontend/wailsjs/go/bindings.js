// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
const go = {
  "main": {
    "App": {
      /**
       * Greet
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<string>}  - Go Type: string
       */
      "Greet": (arg1) => {
        return window.go.main.App.Greet(arg1);
      },
    },
  },

  "sqlite": {
    "repository": {
      /**
       * CountAll
       * @returns {Promise<number>}  - Go Type: int64
       */
      "CountAll": () => {
        return window.go.sqlite.repository.CountAll();
      },
      /**
       * CountHardware
       * @returns {Promise<number>}  - Go Type: int64
       */
      "CountHardware": () => {
        return window.go.sqlite.repository.CountHardware();
      },
      /**
       * CountSoftware
       * @returns {Promise<number>}  - Go Type: int64
       */
      "CountSoftware": () => {
        return window.go.sqlite.repository.CountSoftware();
      },
      /**
       * GetAll
       * @returns {Promise<Array.<Asset>|Error>}  - Go Type: []storage.Asset
       */
      "GetAll": () => {
        return window.go.sqlite.repository.GetAll();
      },
      /**
       * GetAllManufacturers
       * @returns {Promise<Array.<Manufacturer>|Error>}  - Go Type: []storage.Manufacturer
       */
      "GetAllManufacturers": () => {
        return window.go.sqlite.repository.GetAllManufacturers();
      },
      /**
       * GetById
       * @param {number} arg1 - Go Type: uint
       * @returns {Promise<Asset>}  - Go Type: storage.Asset
       */
      "GetById": (arg1) => {
        return window.go.sqlite.repository.GetById(arg1);
      },
      /**
       * GetByName
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<User|Error>}  - Go Type: storage.User
       */
      "GetByName": (arg1) => {
        return window.go.sqlite.repository.GetByName(arg1);
      },
      /**
       * PaginateByName
       * @param {string} arg1 - Go Type: string
       * @param {QueryOptions} arg2 - Go Type: storage.QueryOptions
       * @returns {Promise<Array.<Asset>|Error>}  - Go Type: []storage.Asset
       */
      "PaginateByName": (arg1, arg2) => {
        return window.go.sqlite.repository.PaginateByName(arg1, arg2);
      },
      /**
       * PaginateByTypeAndName
       * @param {string} arg1 - Go Type: string
       * @param {string} arg2 - Go Type: string
       * @param {QueryOptions} arg3 - Go Type: storage.QueryOptions
       * @returns {Promise<Array.<Asset>|Error>}  - Go Type: []storage.Asset
       */
      "PaginateByTypeAndName": (arg1, arg2, arg3) => {
        return window.go.sqlite.repository.PaginateByTypeAndName(arg1, arg2, arg3);
      },
      /**
       * Save
       * @param {Asset} arg1 - Go Type: storage.Asset
       * @returns {Promise<Error>}  - Go Type: error
       */
      "Save": (arg1) => {
        return window.go.sqlite.repository.Save(arg1);
      },
    },
  },

};
export default go;