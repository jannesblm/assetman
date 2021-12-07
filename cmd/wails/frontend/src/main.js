import {createApp} from "vue"
import {createMemoryHistory, createRouter} from "vue-router"
import {createStore} from "vuex"
import {extend, get, has, isObjectLike, isPlainObject, uniq, unset} from "lodash"

import dayjs from "dayjs";

import "@tabler/core/dist/css/tabler.css"
import "@tabler/core/dist/js/tabler"

import App from "./App.vue"
import AssetList from "@/components/AssetList";
import Login from "@/components/Login";
import ManufacturerList from "@/components/ManufacturerList";
import ErrorModal from "@/components/ErrorModal";
import AssetEdit from "@/components/AssetEdit";
import Backup from "@/components/Backup";
import VulnerabilityModal from "@/components/VulnerabilityModal";

window._ = {
    has,
    get,
    isObjectLike,
    isPlainObject,
    unset,
    extend,
    uniq
}

const routes = [
    {path: "/login", name: "login", component: Login},
    {path: "/asset/list/:type", name: "assets", component: AssetList},
    {path: "/manufacturers", name: "manufacturers", component: ManufacturerList},
    {path: "/backup", name: "backup", component: Backup},
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
                Name: '',
                Group: '',
            },

            count: {
                hardware: 0,
                software: 0,
                manufacturers: 0,
            },

            manufacturers: [],

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

    getters: {
        loggedIn: (state) => {
            return state.user.ID > 0
        },

        isAdmin: (state) => {
            return state.user.Group === "admin"
        },
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

        setAssetCounts(state, count) {
            state.count = count
        },

        setManufacturers(state, list) {
            state.manufacturers = list
        },

        setUser(state, user) {
            state.user = user
        }
    },

    actions: {
        setUser({commit}, user) {
            commit("setUser", user)
        },

        showModal({commit}, opts) {
            commit("showModal", opts)
        },

        hideModal({commit, state}, status) {
            commit("hideModal")
            _.has(state.modal.on, "close") && state.modal.on.close(status)
        },

        handleModalMessage({state}, message) {
            _.has(state.modal.on, message.type)
            && typeof (state.modal.on[message.type]) === "function"
            && state.modal.on[message.type](...message.args)
        },

        async syncAssetCounts({commit}) {
            commit("setAssetCounts", {
                hardware: await window.go.sqlite.assetRepository.CountHardware(),
                software: await window.go.sqlite.assetRepository.CountSoftware(),
                manufacturers: await window.go.sqlite.manufRepository.CountAll(),
            })
        },

        async syncManufacturers({commit}) {
            commit("setManufacturers", await window.go.sqlite.manufRepository.GetAll())
        },
    }
})

const app = createApp(App)

app.component("ErrorModal", ErrorModal)
app.component("AssetEditModal", AssetEdit)
app.component("VulnerabilityModal", VulnerabilityModal)

app.use(router)
app.use(store)

app.config.globalProperties.$dayjs = dayjs

app.config.globalProperties.$showDialog = function (title, description, success = false) {
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
                    success: success,
                    showCancel: false,
                    confirmText: "OK",
                    title: title,
                    description: description,
                })
            },
            on: {
                close: () => {
                }
            }
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

window.go.auth.service.GetUser()
    .then(user => {
        store.dispatch("setUser", user)
        router.replace('/asset/list/hardware')
    })
    .catch(() => {
        router.replace({name: 'login'})
    })
