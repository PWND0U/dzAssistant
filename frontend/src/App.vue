<script lang="ts" setup>
import Head from "./components/Head.vue";
import SpeechPanel from "./components/SpeechPanel.vue";
import {ref} from "vue";
import {WindowSetAlwaysOnTop} from "../wailsjs/runtime";
let toggleShow = ref(true);
function windowToggleAndTop(){
  toggleShow.value = !toggleShow.value;
  WindowSetAlwaysOnTop(!toggleShow.value)
}
</script>

<template>
  <div class="flex flex-col text-amber-50 relative h-screen">
    <Head class="flex h-[35px] w-full items-center bg-gray-900/50" v-show="toggleShow"></Head>
    <speech-panel></speech-panel>
    <el-icon
        class="text-cyan-50/50 hover:text-cyan-50 hover:cursor-pointer absolute right-1 bottom-1.5"
        size="30"
        @click="windowToggleAndTop"
    ><i-ep-HelpFilled />
    </el-icon>
  </div>
</template>

<style scoped>
</style>
