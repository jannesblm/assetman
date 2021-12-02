// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
const go = {
  "backup": {
    "service": {
      /**
       * CreateBackup
       * @returns {Promise<string|Error>}  - Go Type: string
       */
      "CreateBackup": () => {
        return window.go.backup.service.CreateBackup();
      },
      /**
       * GetAppHomePath
       * @returns {Promise<string>}  - Go Type: string
       */
      "GetAppHomePath": () => {
        return window.go.backup.service.GetAppHomePath();
      },
      /**
       * GetBackupDir
       * @returns {Promise<string>}  - Go Type: string
       */
      "GetBackupDir": () => {
        return window.go.backup.service.GetBackupDir();
      },
      /**
       * GetBackupList
       * @returns {Promise<Array.<Backup>|Error>}  - Go Type: []backup.Backup
       */
      "GetBackupList": () => {
        return window.go.backup.service.GetBackupList();
      },
      /**
       * GetDatabasePath
       * @returns {Promise<string>}  - Go Type: string
       */
      "GetDatabasePath": () => {
        return window.go.backup.service.GetDatabasePath();
      },
    },
  },

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
    "assetRepository": {
      /**
       * CountAll
       * @returns {Promise<number>}  - Go Type: int64
       */
      "CountAll": () => {
        return window.go.sqlite.assetRepository.CountAll();
      },
      /**
       * CountHardware
       * @returns {Promise<number>}  - Go Type: int64
       */
      "CountHardware": () => {
        return window.go.sqlite.assetRepository.CountHardware();
      },
      /**
       * CountSoftware
       * @returns {Promise<number>}  - Go Type: int64
       */
      "CountSoftware": () => {
        return window.go.sqlite.assetRepository.CountSoftware();
      },
      /**
       * Delete
       * @param {Asset} arg1 - Go Type: storage.Asset
       * @returns {Promise<Error>}  - Go Type: error
       */
      "Delete": (arg1) => {
        return window.go.sqlite.assetRepository.Delete(arg1);
      },
      /**
       * GetById
       * @param {number} arg1 - Go Type: uint
       * @returns {Promise<Asset>}  - Go Type: storage.Asset
       */
      "GetById": (arg1) => {
        return window.go.sqlite.assetRepository.GetById(arg1);
      },
      /**
       * GetByName
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<User|Error>}  - Go Type: storage.User
       */
      "GetByName": (arg1) => {
        return window.go.sqlite.assetRepository.GetByName(arg1);
      },
      /**
       * Paginate
       * @param {string} arg1 - Go Type: string
       * @param {QueryOptions} arg2 - Go Type: storage.QueryOptions
       * @returns {Promise<Array.<Asset>|Error>}  - Go Type: []storage.Asset
       */
      "Paginate": (arg1, arg2) => {
        return window.go.sqlite.assetRepository.Paginate(arg1, arg2);
      },
      /**
       * Save
       * @param {Asset} arg1 - Go Type: storage.Asset
       * @returns {Promise<Error>}  - Go Type: error
       */
      "Save": (arg1) => {
        return window.go.sqlite.assetRepository.Save(arg1);
      },
    },
    "manufRepository": {
      /**
       * CountAll
       * @returns {Promise<number>}  - Go Type: int64
       */
      "CountAll": () => {
        return window.go.sqlite.manufRepository.CountAll();
      },
      /**
       * GetAll
       * @returns {Promise<Array.<Manufacturer>|Error>}  - Go Type: []storage.Manufacturer
       */
      "GetAll": () => {
        return window.go.sqlite.manufRepository.GetAll();
      },
      /**
       * GetById
       * @param {number} arg1 - Go Type: uint
       * @returns {Promise<Manufacturer>}  - Go Type: storage.Manufacturer
       */
      "GetById": (arg1) => {
        return window.go.sqlite.manufRepository.GetById(arg1);
      },
      /**
       * GetByName
       * @param {string} arg1 - Go Type: string
       * @returns {Promise<User|Error>}  - Go Type: storage.User
       */
      "GetByName": (arg1) => {
        return window.go.sqlite.manufRepository.GetByName(arg1);
      },
      /**
       * Paginate
       * @param {QueryOptions} arg1 - Go Type: storage.QueryOptions
       * @returns {Promise<Array.<Manufacturer>|Error>}  - Go Type: []storage.Manufacturer
       */
      "Paginate": (arg1) => {
        return window.go.sqlite.manufRepository.Paginate(arg1);
      },
    },
  },

};
export default go;
