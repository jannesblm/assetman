<template>
  <div class="page-body">
    <div class="container-fluid">
      <div class="row mb-4">
        <div class="col">
          <h2 class="page-title">
            Backups
          </h2>
        </div>
        <div class="col">
          <div class="col-12 d-flex justify-content-end">
            <a class="btn btn-success w-auto" href="#" @click="backup">
              <i class="ti ti-plus mr-3"></i>Create Backup
            </a>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-12">
          <DataTable :columns="columns"
                     :rows="backups"
                     :show-footer="false"
                     :show-header="false"
                     :total="backups.length"
                     @update="load">
          </DataTable>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
import DataTable from "@/components/DataTable";

export default {
  name: "Backup",

  components: {DataTable},

  data() {
    return {
      columns: [
        {
          property: 'Path',
          label: 'Path',
          noSearch: true,
          forEach: (p) => `${this.backupPath}\\${p}`
        },
        {
          property: 'Size',
          label: 'Size',
          noSearch: true,
          forEach: this.bytesToSize
        },
        {
          property: 'Modified',
          label: 'Modified',
          noSearch: true,
          forEach: (d) => this.$dayjs(d).format("DD/MM/YYYY HH:mm:ss")
        },
      ],

      backupPath: '',
      backups: []
    }
  },

  methods: {
    async load() {
      try {
        this.backups = await window.go.fs.service.GetBackupList()
      } catch (error) {
        this.$showDialog("An error occurred", error)
      }
    },

    backup: async function () {
      try {
        const name = await window.go.fs.service.CreateBackup()

        this.$showDialog("Backup created", "Backup was successfully created at \"" + name + "\"",
            true)
      } catch (error) {
        this.$showDialog("An error occurred", error)
      }

      await this.load()
    },

    bytesToSize: function bytesToSize(bytes) {
      const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
      if (bytes === 0) return '0 Byte';
      const i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)));

      return Math.round(bytes / Math.pow(1024, i), 2) + ' ' + sizes[i];
    }
  },

  async mounted() {
    this.backupPath = await window.go.fs.service.GetBackupDirectory()

    await this.load()
  }
}
</script>

<style scoped>

</style>