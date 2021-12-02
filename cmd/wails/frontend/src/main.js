import {createApp} from "vue"
import {createMemoryHistory, createRouter} from "vue-router"
import {createStore} from "vuex"
import {get, has, isObjectLike, unset} from "lodash"
import dayjs from "dayjs";
import JQuery from 'jquery'

import '@tabler/core/dist/css/tabler.css'
import '@tabler/core/dist/js/tabler'

import App from './App.vue'
import HelloWorld from "@/components/HelloWorld";
import AssetList from "@/components/AssetList";
import Login from "@/components/Login";
import ManufacturerList from "@/components/ManufacturerList";
import ErrorModal from "@/components/ErrorModal";
import AssetEdit from "@/components/AssetEdit";

window.jQuery = window.$ = JQuery

window._ = {
    has,
    get,
    isObjectLike,
    unset,
}

const routes = [
    {path: "/", name: "home", component: HelloWorld},
    {path: "/login", name: "login", component: Login},
    {path: "/asset/list/:type", name: "assets", component: AssetList},
    {path: "/manufacturers", name: "manufacturers", component: ManufacturerList},
]

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
                isLoggedIn: true,
            },

            count: {
                hardware: 0,
                software: 0,
                manufacturers: 0,
            },

            modal: {
                show: false,
                component: "ErrorModal",
                props: {opts: ErrorModal.props.opts.default()},
                classes: [],
                on: {
                    close: () => {
                    },
                },
            }
        }
    },

    mutations: {
        showModal(state, opts) {
            state.modal.show = true

            if (_.has(opts, "component")) {
                state.modal.component = opts.component
            }

            if (_.has(opts, "props")) {
                state.modal.props = opts.props
            }

            if (_.has(opts, "classes")) {
                state.modal.classes = opts.classes
            }

            if (_.has(opts, "on")) {
                state.modal.on = opts.on
            }
        },

        hideModal(state) {
            state.modal.show = false
        },

        async setAssetCounts(state, count) {
            state.count = count
        },
    },

    actions: {
        showModal({commit}, opts) {
            commit("showModal", opts)
        },

        hideModal({commit, state}, status) {
            commit("hideModal")
            _.has(state.modal.on, "close") && state.modal.on.close(status)
        },

        async syncAssetCounts({ commit }) {
            commit("setAssetCounts", {
                hardware: await window.go.sqlite.assetRepository.CountHardware(),
                software: await window.go.sqlite.assetRepository.CountSoftware(),
                manufacturers: await window.go.sqlite.manufRepository.CountAll(),
            })
        },

        handleModalMessage({state}, message) {
            console.log(_.has(state.modal.on, message.type))
            _.has(state.modal.on, message.type) && state.modal.on[message.type](...message.args)
        }
    }
})

const app = createApp(App)

app.component("ErrorModal", ErrorModal)
app.component("AssetEditModal", AssetEdit)

app.use(router)
app.use(store)

app.config.globalProperties.$dayjs = dayjs

app.config.globalProperties.$showError = function (title, description) {
    let timeout = 0

    if (store.state.modal.show) {
        store.dispatch("hideModal").then()
        timeout = 500
    }

    setTimeout(() => {
        store.dispatch("showModal", {
            classes: ['modal-sm', 'modal-dialog-centered'],
            component: "ErrorModal",
            props: {
                opts: Object.assign(ErrorModal.props.opts.default(), {
                    success: false,
                    showCancel: false,
                    confirmText: "OK",
                    title: title,
                    description: description,
                })
            },
        }).then()
    }, timeout)
}

app.config.globalProperties.$confirm = function (title, description, callback) {
    let timeout = 0

    if (store.state.modal.show) {
        store.dispatch("hideModal").then()
        timeout = 500
    }

    setTimeout(() => {
        store.dispatch("showModal", {
            classes: ['modal-sm', 'modal-dialog-centered'],
            component: "ErrorModal",
            props: {
                opts: Object.assign(ErrorModal.props.opts.default(), {
                    success: false,
                    confirmText: "Yes",
                    title: title,
                    description: description,
                })
            },
            on: {
                close: callback
            }
        }).then()
    }, timeout)
}

app.mount('#app')
router.replace('/asset/list/software')
