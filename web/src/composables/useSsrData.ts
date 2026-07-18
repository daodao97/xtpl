import { inject, shallowRef, type ShallowRef } from 'vue'

export type SsrState = Record<string, unknown>

export interface SsrDataContext {
  state: ShallowRef<SsrState>
  setState: (value: SsrState) => void
}

export const ssrDataKey = Symbol('ssr-data')
const emptySsrState = shallowRef<SsrState>({})

export function createSsrDataContext(initialState: SsrState): SsrDataContext {
  const state = shallowRef<SsrState>(initialState)
  return {
    state,
    setState(value) {
      state.value = value && typeof value === 'object' ? value : {}
    },
  }
}

export function useSsrData<T extends object = SsrState>(): ShallowRef<T> {
  const context = inject<SsrDataContext | null>(ssrDataKey, null)
  return (context?.state ?? emptySsrState) as ShallowRef<T>
}
