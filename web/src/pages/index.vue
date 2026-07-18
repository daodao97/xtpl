<script setup lang="ts">
import { useSsrData } from '~/composables/useSsrData'

interface HomePayload {
  message?: string
  path?: string
}

const payload = useSsrData<HomePayload>()
</script>

<template>
  <main>
    <h1>gossr is ready</h1>
    <p>{{ payload.message ?? 'Add a payload handler in the Go host.' }}</p>
    <small>Path: {{ payload.path ?? '/' }}</small>
  </main>
</template>
