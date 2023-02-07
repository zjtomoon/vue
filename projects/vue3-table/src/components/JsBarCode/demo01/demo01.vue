<template>
  <div class="js_barcode">
    <div v-for="(item,index) in jsBarCodeList" :key="index">
      <img :id="'jsbarcodeImg' + index" style="width: 260px"/>
    </div>
  </div>
</template>

<script setup>
import {onMounted, ref} from 'vue'

import JsBarCode from 'jsbarcode'

const jsBarCodeList = ref(['ETC6987', 'VIP6666', 'TXSH7845'])

onMounted(() => {
  setTimeout(() => {
    generateJsBarCodeImg()
  }, 500)
})

function generateJsBarCodeImg() {
  jsBarCodeList.forEach((v, index) => {
    JsBarCode('#jsbarcodeImg' + index, v, {
      format: 'CODE39',
      width: 2,
      height: 40
    })
  })
}

</script>

<style lang="scss" scoped>

</style>