import { createApp, createSSRApp } from 'vue'
import { createMemoryHistory, createRouter, createWebHistory, type Router } from 'vue-router'
import { routes } from 'vue-router/auto-routes'

import App from './App.vue'
import { createSsrDataContext, ssrDataKey, type SsrState } from '~/composables/useSsrData'

const isServer = typeof window === 'undefined'

export function createAppRouter(): Router {
  return createRouter({
    history: isServer ? createMemoryHistory() : createWebHistory('/'),
    routes,
  })
}

export function makeApp(initialState: SsrState = {}, options: { hydrate?: boolean, router?: Router } = {}) {
  const router = options.router ?? createAppRouter()
  const app = isServer || options.hydrate !== false ? createSSRApp(App) : createApp(App)
  const ssrContext = createSsrDataContext(initialState)

  app.use(router)
  app.provide(ssrDataKey, ssrContext)

  return {
    app,
    router,
    ssrContext,
    dispose() {
      app.unmount()
    },
  }
}
