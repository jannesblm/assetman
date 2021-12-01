import {createApp} from 'vue'
import {createRouter, createMemoryHistory} from 'vue-router'
import {createStore} from 'vuex'

import '@tabler/core/dist/css/tabler.css'
import '@tabler/core/dist/js/tabler'

import JQuery from 'jquery'
window.jQuery = window.$ = JQuery

import App from './App.vue'
import HelloWorld from "@/components/HelloWorld";
import AssetList from "@/components/AssetList";

const routes = [
    {path: "/", component: HelloWorld},
    {path: "/asset/list/:type", component: AssetList},
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
            assetCount: {
                hardware: 0,
                software: 0,
            },
            manufacturers: [],
        }
    },
    mutations: {
        async syncCount(state) {
            state.assetCount.hardware = await window.go.sqlite.repository.CountHardware()
            state.assetCount.software = await window.go.sqlite.repository.CountSoftware()
        },

        async updateManufacturers(state) {
            state.manufacturers = await window.go.sqlite.repository.GetAllManufacturers()
        }
    }
})

const app = createApp(App)

app.use(router)
app.use(store)

// TODO debounce

app.mount('#app')
router.replace('/asset/list')
