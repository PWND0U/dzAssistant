<script setup lang="ts">
import {Quit, WindowHide, WindowMaximise, WindowMinimise, WindowToggleMaximise} from "../../wailsjs/runtime";
import {ref} from  "vue";
import {GetAppName} from "../../wailsjs/go/mainApp/MainApp";
let appName =  ref("");
GetAppName().then((name) =>{
  appName.value = name;
})
</script>

<template>
  <div class="flex px-2">
    <div>{{appName}}</div>
    <div class="flex-grow h-full" style="--wails-draggable:drag"></div>
    <div class="space-x-2 flex">
      <span class="cursor-pointer flex items-center" @click="WindowHide()"><el-icon size="20px"><i-ep-ArrowDown /></el-icon></span>
      <span class="cursor-pointer flex items-center" @click="WindowMinimise()"><el-icon size="20px"><i-ep-RemoveFilled /></el-icon></span>
      <span class="cursor-pointer flex items-center" @click="WindowToggleMaximise()"><el-icon size="20px"><i-ep-Rank /></el-icon></span>
      <span class="cursor-pointer flex items-center" @click="Quit()"><el-icon size="20px"><i-ep-CloseBold /></el-icon></span>

    </div>
  </div>
</template>

<style scoped>

</style>