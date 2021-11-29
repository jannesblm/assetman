import {createApp} from 'vue'
import {createRouter, createMemoryHistory} from 'vue-router'
import {createStore} from 'vuex'

import '@tabler/core/dist/css/tabler.css'
import '@tabler/core/dist/js/tabler'

import JQuery from 'jquery'
window.jQuery = window.$ = JQuery

import lodashGet from 'lodash/get'
import VueMask from 'vue-jquery-mask';


import App from './App.vue'
import HelloWorld from "@/components/HelloWorld";
import AssetList from "@/components/AssetList";

const routes = [
    {path: '/', component: HelloWorld},
    {path: '/asset/list', component: AssetList,},
];

const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

const store = createStore({
    state() {
        return {
            user: {
                ID: 0,
                name: '',
            },
            assetCount: 0,
            manufacturers: [],
        }
    },
    mutations: {
        async updateAssetCount(state) {
            state.assetCount = await window.go.sqlite.repository.CountAll()
        },

        async updateManufacturers(state) {
            state.manufacturers = await window.go.sqlite.repository.GetAllManufacturers()
        }
    }
})

const app = createApp(App)

app.use(router)
app.use(store)
app.use(VueMask)

// TODO debounce
app.config.globalProperties.$get = lodashGet

app.mount('#app')
router.replace('/asset/list')
