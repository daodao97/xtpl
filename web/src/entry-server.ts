import { renderToString } from '@vue/server-renderer'

import { createAppRouter, makeApp } from '~/main'
import type { SsrState } from '~/composables/useSsrData'

const router = createAppRouter()

async function render(url: string) {
  const initialState: SsrState = (globalThis as any).__SSR_DATA__ ?? {}
  const { app, dispose } = makeApp(initialState, { router })
  try {
    await router.replace(url)
    const context: any = {}
    ;(globalThis as any).__SSR_HEAD__ = ''
    const html = await renderToString(app, context)
    ;(globalThis as any).__SSR_HEAD__ = typeof context.teleports?.head === 'string'
      ? context.teleports.head
      : ''
    return html
  }
  finally {
    dispose()
  }
}

;(globalThis as any).ssrRender = render
