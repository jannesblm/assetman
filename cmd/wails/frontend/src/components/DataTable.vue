<template>
  <div class="card">
    <div v-if="showHeader" class="card-body border-bottom py-3">
      <div class="d-flex flex-column flex-lg-row">
        <div class="ms-auto mb-3 mb-lg-0 ms-lg-0 text-muted">
          Show
          <div class="mx-2 d-inline-block">
            <input :value="limit" class="form-control form-control-sm" size="3"
                   type="text" @keyup.enter="limit = parseInt($event.target.value); emitUpdate()">
          </div>
          entries
        </div>
        <div v-if="searchableColumns.length > 0" class="ms-auto text-muted d-flex">
          Search:
          <div class="ms-2 input-group">
            <select id="query-field" v-model="queryFieldIndex" class="form-select form-control-sm ms-2 w-25">
              <option v-for="(column, index) in searchableColumns" :key="index" :value="index">
                {{ column.label }}
              </option>
            </select>
            <input v-model="query" class="form-control form-control-sm w-50" type="text">
          </div>
        </div>
      </div>
    </div>
    <div class="table-responsive">
      <table class="table table-vcenter card-table">
        <thead>
        <tr>
          <th v-if="checkboxes"></th>
          <th v-for="(column, index) in columns" :key="column" :class="column.class">
            <span
                :class="[{'text-green font-weight-extrabold': sortIndex === index, 'cursor-pointer': isSortable(column)}, 'd-flex']"
                @click="sort = index">
              {{ column.label }}
              <span v-if="isSortable(column)">
                <i v-if="sortIndex === index"
                   :class="[{'ti-chevron-up': sortDir === 'asc', 'ti-chevron-down': sortDir === 'desc'}, 'ti']"></i>
                <i v-else class="ti ti-chevron-down"></i>
              </span>
            </span>
          </th>
        </tr>
        </thead>
        <tbody v-if="rows.length > 0">
        <tr v-for="(row, index) in rows" :key="index">
          <td v-if="checkboxes && has(row, 'checked')">
            <input class="form-check-input" type="checkbox">
          </td>
          <td v-for="(column, index) in modelColumns" :key="index">
            <div v-if="!isTrusted(column)" class="column-content">{{ col(row, column) }}</div>
            <div v-else class="column-content" v-html="col(row, column)"></div>
          </td>
          <slot :row="row"></slot>
        </tr>
        </tbody>
        <tbody v-else>
        <tr>
          <td :colspan="columns.length" class="text-center">
            <small>No results found</small>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
    <div v-if="showFooter" class="card-footer d-flex align-items-center">
      <p class="m-0 text-muted">Showing {{ offset > 0 ? offset : 1 }} to {{ offset + rows.length }} of {{ total }}
        entries</p>
      <ul class="pagination m-0 ms-auto">
        <li :class="[{disabled: page - 1 === 0}, 'page-item']">
          <a class="page-link" @click="(page - 1) >= 0 ? page-- : () => {}">
            <i class="ti ti-chevron-left"></i>
            prev
          </a>
        </li>
        <li v-for="i in availablePages" :key="i" :class="[{active: page === i}, 'page-item', 'cursor-pointer']">
          <a :class="[{active: page === i}, 'page-link']" @click="page = i">{{ i }}</a>
        </li>
        <li :class="[{disabled: page >= availablePages}, 'page-item']">
          <a class="page-link" @click="(page + 1) > 0 ? page++ : () => {}">
            next
            <i class="ti ti-chevron-right"></i>
          </a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: "DataTable",

  emits: ["update"],

  props: {
    checkboxes: {
      type: Boolean,
      required: false,
      default: false,
    },

    checkboxProp: {
      type: String,
      required: false,
      default: "",
    },

    columns: {
      type: Array,
      required: true,
      default() {
        return [
          {
            property: 'name of property',
            label: 'label of column',
          },
        ];
      }
    },

    rows: {
      type: Array,
      required: true,
      default() {
        return [];
      }
    },

    total: {
      type: Number,
      required: true,
      default() {
        return 0;
      }
    },

    showHeader: {
      type: Boolean,
      required: false,
      default: true,
    },

    showFooter: {
      type: Boolean,
      required: false,
      default: true,
    }
  },

  data() {
    return {
      query: "",
      queryFieldIndex: -1,
      sortIndex: -1,
      sortDir: "desc",
      limit: 10,
      page: 1,
    }
  },

  methods: {
    has: _.has,

    toSQLName: function (str) {
      return str && str
          // GORM will alias the joined tables like StructName.snake_cased_property, this matches that
          .match(/[A-za-z0-9]{2,}?\.|[A-za-z0-9]{2,}?(?=[A-Z]|$)/g)
          .reduce((c, i) => c + (i.endsWith('.') ? i : `${i.toLowerCase()}_`), '')
          .slice(0, -1)
    },

    isSortable: function (col) {
      return _.has(col, 'sortable')
          && _.has(col, 'property')
          && col.sortable
    },

    isSearchable: function (col) {
      return _.has(col, 'property')
          && (!_.has(col, 'noSearch') || !col.noSearch)
    },

    isTrusted: function (col) {
      return _.has(col, 'trusted') && col.trusted
    },

    col: function (row, col) {
      return this.forEach(col)(_.get(row, col.property))
    },

    init: function () {
      let firstSortable = this.columns.findIndex(col => this.isSortable(col)
          && _.has(col, 'defaultSort')
          && col.defaultSort)

      if (firstSortable === -1) {
        firstSortable = this.columns.findIndex(this.isSortable)
      }

      if (firstSortable >= 0) {
        this.sortIndex = firstSortable
      }

      let firstSearchable = this.searchableColumns.findIndex(col => _.has(col, 'defaultSearch')
          && col.defaultSearch)

      if (firstSearchable === -1) {
        firstSearchable = this.searchableColumns.findIndex(this.isSearchable)
      }

      if (firstSearchable >= 0) {
        this.queryFieldIndex = firstSearchable
      }
    },

    toggleSortDir: function () {
      this.sortDir = this.sortDir === "desc" ? "asc" : "desc"
    },

    forEach(column) {
      if (typeof column.forEach === "function") {
        return column.forEach;
      }

      return (o) => o
    },

    emitUpdate() {
      this.$emit("update", {
        query: this.query,
        queryField: this.queryField,
        offset: this.offset,
        limit: this.limit,
        sort: this.sort,
      })
    },
  },

  watch: {
    columns: function () {
      this.init()
    },

    query: function () {
      this.emitUpdate()
    },

    sort: function () {
      this.emitUpdate()
    },

    page: function () {
      this.emitUpdate()
    }
  },

  computed: {
    modelColumns: function () {
      return this.columns.filter(c => _.has(c, 'property'))
    },

    searchableColumns: function () {
      return this.columns.filter(this.isSearchable)
    },

    availablePages: function () {
      return Math.ceil(this.total / this.limit)
    },

    offset: function () {
      return (this.page - 1) * this.limit
    },

    sort: {
      get: function () {
        if (this.sortIndex >= 0) {
          return (_.has(this.columns[this.sortIndex], "prefix")
                  ? `${this.columns[this.sortIndex].prefix}.` : '')
                  + this.toSQLName(this.columns[this.sortIndex].property)
                  + " " + this.sortDir


        }

        return ""
      },

      set: function (newVal) {
        if (!_.has(this.columns, `[${newVal}]`) || !this.isSortable(this.columns[newVal])) {
          return
        }

        if (this.sortIndex === newVal) {
          this.toggleSortDir()
        } else {
          this.sortDir = "desc"
        }

        this.sortIndex = newVal
      }
    },

    queryFieldLabel: function () {
      if (this.queryFieldIndex >= 0) {
        return this.searchableColumns[this.queryFieldIndex].label
      }

      return ""
    },

    queryField: function () {
      if (this.queryFieldIndex >= 0) {
        return (_.has(this.columns[this.queryFieldIndex], "prefix")
                ? `${this.columns[this.queryFieldIndex].prefix}.` : '')
                + this.toSQLName(this.searchableColumns[this.queryFieldIndex].property)
      }

      return ""
    }
  },

  mounted() {
    this.init()
  }
}
</script>

<style scoped>
.font-weight-extrabold {
  font-weight: 900;
}

.table-responsive {
  overflow-x: visible;
}

#query-field {
  min-height: calc(1.3333333em + 0.25rem + 2px);
  padding: 0.125rem 0.5rem;
  font-size: 0.75rem;
  border-radius: 2px
}
</style>