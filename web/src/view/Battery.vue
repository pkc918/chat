<script lang="ts" setup>
import { useBattery } from '@vueuse/core'

const { charging, chargingTime, dischargingTime, level } = useBattery()

console.log(charging, chargingTime, dischargingTime, level)

</script>

<template>
    <div>
      charging: {{charging}}
      chargingTime: {{chargingTime}}
      dischargingTime: {{dischargingTime}}
      level: {{level}}
    </div>
</template>

<style scoped></style>
