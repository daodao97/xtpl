import type { RouteLocationNormalized } from 'vue-router'

import { makeApp } from '~/main'
import type { SsrState } from '~/composables/useSsrData'

declare global {
  interface Window {
    __SSR_DATA__?: SsrState
  }
}

const hasInitialPayload = window.__SSR_DATA__ !== undefined
const initialState = window.__SSR_DATA__ ?? {}
const appElement = document.querySelector('#app')
const shouldHydrate = appElement !== null
  && appElement.innerHTML.trim() !== ''
  && !appElement.innerHTML.includes('<!--app-html-->')
const { app, router, ssrContext } = makeApp(initialState, { hydrate: shouldHydrate })
const initialPath = window.location.pathname + window.location.search + window.location.hash
let firstNavigation = true
let latestRequest = 0

router.beforeResolve((to, from) => {
  if (firstNavigation) {
    firstNavigation = false
    return true
  }
  if (to.fullPath === from.fullPath)
    return true
  if (to.meta.ssrData === false) {
    ssrContext.setState({})
    return true
  }

  const request = ++latestRequest
  void fetchSsrData(to)
    .then((data) => {
      if (request === latestRequest)
        ssrContext.setState(data)
    })
    .catch(error => console.error('Failed to fetch SSR data', error))
  return true
})

void router.replace(initialPath)
router.isReady().then(() => {
  app.mount('#app')
  delete window.__SSR_DATA__
  if (!hasInitialPayload && router.currentRoute.value.meta.ssrData !== false) {
    void fetchSsrData(router.currentRoute.value)
      .then(data => ssrContext.setState(data))
      .catch(error => console.error('Failed to fetch initial SSR data', error))
  }
})

async function fetchSsrData(route: RouteLocationNormalized): Promise<SsrState> {
  const url = new URL(route.fullPath, window.location.origin)
  const response = await fetch(`/_ssr/data${url.pathname}${url.search}`, {
    credentials: 'same-origin',
    headers: { Accept: 'application/json' },
  })
  if (!response.ok)
    throw new Error(`SSR data request failed with status ${response.status}`)
  const data: unknown = await response.json()
  return data && typeof data === 'object' ? data as SsrState : {}
}
