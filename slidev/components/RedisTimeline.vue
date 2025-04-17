<template>
    <div class="flex content-center justify-center justify-center mt-12">
        <Timeline :value="data" align="alternate" class="w-full md:w-80">
            <template #content="slotProps">
                <div>
                    <p class="text-xs">{{ slotProps.item.date }}</p>
                    <p class="text-2xl">{{ slotProps.item.title }}</p>
                    <p v-if="slotProps.item.description">{{ slotProps.item.description }}</p>
                </div>
            </template>
        </Timeline>
    </div>
</template>

<script setup>
import { defineProps } from 'vue'

defineProps({
    data: Array
})
</script>