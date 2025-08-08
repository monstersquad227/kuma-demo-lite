import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import {
	Button, Table
} from 'ant-design-vue';

const app = createApp(App)

app.use(Button)
app.use(Table)
app.use(router)
app.mount('#app')
