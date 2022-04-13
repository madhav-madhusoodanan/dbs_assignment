<script setup>
defineProps({
    course: String,
    units: Number,
    id: String,
    replace: Function,
    selected: Boolean,
})

import Replace from "./replace.vue"
</script>

<template>
    <div v-if="selected" class="flex flex-row">
        <div
            class="flex flex-row justify-between align-middle text-xl bg-slate-400 rounded min-w-[25vw] px-6 py-2 my-2 cursor-pointer"
        >
            <div class="text-black">
                <h1>{{ course }}</h1>
            </div>
            <div class="text-black">
                <h2>{{ units }}</h2>
            </div>
        </div>
        <div class="h-full flex flex-row align-middle my-auto">
            <div
                class="mx-1 my-auto text-green-400 rounded cursor-pointer"
                @click="replace(id)"
            >
                <Replace />
            </div>
        </div>
    </div>
    <div v-else class="flex flex-row">
        <div
            class="flex flex-row justify-between align-middle text-xl bg-red-400 rounded min-w-[25vw] px-6 py-2 my-2 cursor-pointer"
        >
            <div class="text-white">
                <h1>{{ course }}</h1>
            </div>
            <div class="text-white">
                <h2>{{ units }}</h2>
            </div>
        </div>
    </div>
</template>
