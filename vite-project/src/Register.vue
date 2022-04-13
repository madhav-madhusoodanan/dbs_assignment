<script setup>
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup
import { ref } from "vue"

const username = ref("")
const password = ref("")
const onClick = () => {
    console.log("clicked")
    console.log(username.value)
    console.log(password.value)
};

</script>

<template>
    <div class="no-scroll w-screen h-screen bg-white grid place-items-center">
        <div class="w-full h-min">
            <h1
                class="w-screen my-4 align-middle justify-center text-2xl font-semibold text-black"
            >
                Hello, Student!
            </h1>
            <div class="w-full my-4">
                <input
                    v-model="username"
                    placeholder="Username"
                    class="w-2/5 px-4 py-2 transition-all bg-gray-200 rounded-lg animation-100 border-2 border-white"
                />
            </div>
            <div class="w-full my-4">
                <input
                    v-model="password"
                    placeholder="Password"
                    class="w-2/5 px-4 py-2 transition-all bg-gray-200 rounded-lg animation-100 border-2 border-white"
                />
            </div>
            <div
                class="flex flex-row justify-between align-middle text-xl bg-slate-400 rounded w-[25vw] px-6 py-2 my-2 mx-auto cursor-pointer"
                @click="onClick"
            >
                <p class="mx-auto">Join</p>
            </div>
        </div>
    </div>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Montserrat&display=swap");
@tailwind base;
@tailwind components;
@tailwind utilities;

#app {
    font-family: "Montserrat", sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    display: flex;
    flex-direction: row;
    justify-content: space-evenly;
    align-items: center;
    height: 100vh;
    width: 100vw;
}

.no-scroll {
    -ms-overflow-style: none; /* IE and Edge */
    scrollbar-width: none; /* Firefox */
}

.no-scroll::-webkit-scrollbar {
    display: none;
}
</style>
