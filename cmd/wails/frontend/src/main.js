import { createApp } from 'vue'
import {createRouter, createMemoryHistory} from 'vue-router'

import App from './App.vue'
import HelloWorld from "@/components/HelloWorld";

import '@tabler/core/dist/css/tabler.css'
import '@tabler/core/dist/js/tabler'
import AssetList from "@/components/AssetList";


const routes = [
    {path: '/', component: HelloWorld},
    {path: '/asset/list', component: AssetList,},
];

const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

createApp(App).use(router).mount('#app')
