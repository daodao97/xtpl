import { createAdmin, setCmp, regViews } from '@okiss/oms'
import '@okiss/oms/style.css'
import { defineAsyncComponent } from 'vue'
import app from './app'
import { RegComponents, regCustomFormComps } from '@okiss/vbtf'


// register dashboard
setCmp('dashboard', defineAsyncComponent(() => import('./views/dashboard/index.vue')))

// RegComponents({
//   'ChannelTest': defineAsyncComponent(() => import('./components/ChannelTest.vue')),
// })

// // register custom form components
// regCustomFormComps({
//   'UserTokenChannelBinding': defineAsyncComponent(() => import('./components/UserTokenChannelBinding.vue')),
// })

// register custom page views
regViews(import.meta.glob('./views/**/**.vue', { eager: true }))

const env = import.meta.env
const isProdMode = env.PROD

const options = {
  // mock: true,
  settings: {
    formMutex: false,
    title: 'XgoAdmin',
    showPageJsonSchema: false,
    logo: 'https://dow.chatbee.cc/chatgpt.jpeg',
    defaultAvatar: 'https://dow.chatbee.cc/chatgpt.jpeg',
    captcha: false
  },
  plugins: [app],
  axios: {
    baseURL: (window.BASE_URL || env.VITE_BASE_API) + ''
  },
  form: {
    vsPath: isProdMode ? location.pathname + 'assets/monaco-editor/vs' : 'node_modules/monaco-editor/min/vs'
  }
}

createAdmin(options)
