<template>
  <div class="container">
    <div class="content">JsBarCode</div>
    <svg id="barcode"></svg>
  </div>
</template>

<script>
import JsBarCode from "jsbarcode";
export default {
  name: "demo02",
  data() {
    return {}
  },
  mounted() {
    this.getBarCode()
  },
  methods:{
    getBarCode() {
      let options = {
        text:"123456789",
        displayValue:true,
        fontSize:12,
        height:40,
      }
      this.$nextTick(() => {
        JsBarCode("#barcode","987654321",options)
      })
    }
  }
}
</script>

<style lang="scss" scoped>
  .container{
    .content{
      background-color: aquamarine;
      font-size: 20px;
    }
  }
</style>