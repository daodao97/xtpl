<template>
  <header>
    <RouterLink to="/">Home</RouterLink>
  </header>
  <RouterView />
</template>

<style>
body {
  margin: 0;
  color: #172033;
  background: #f7f8fb;
  font: 16px/1.5 system-ui, sans-serif;
}

#app {
  width: min(720px, calc(100% - 32px));
  margin: 48px auto;
}

header {
  margin-bottom: 24px;
}
</style>
