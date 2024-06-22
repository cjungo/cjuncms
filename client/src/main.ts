import { createApp } from "vue";
import { createPersistedState } from "pinia-plugin-persistedstate";
import { createPinia } from "pinia";
import VxeUITable from 'vxe-table'
import 'vxe-table/lib/style.css'
import "normalize.css";
import "./style.css";
import App from "./App.vue";
import router from "./router";

const persist = createPersistedState({
  storage: localStorage,
  key: (i) => `__store__${i}`,
});
const pinia = createPinia();
pinia.use(persist);

const app = createApp(App);
app.use(pinia);
app.use(router);
app.use(VxeUITable);
app.mount("#app");
