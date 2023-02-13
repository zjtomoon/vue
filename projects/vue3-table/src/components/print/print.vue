<!--https://blog.csdn.net/Ling_a_Sir/article/details/121512452-->
<!--https://juejin.cn/post/6888519415717953549-->
<template>
  <div>
    <button class="notprint" @click="printOut(printDom)">
      打印
    </button>
    <div ref="printDom">
      <hr/>
      <hr/>
      需要打印的内容
      <hr/>
    </div>
  </div>
</template>

<script>
import {reactive, ref} from 'vue'

export default {
  name: "print",
  components:{},
  setup(props) {
    const printDom = ref(null)
    const printOut = (dom) => {
      var newWindow = window.open("打印页面","_blank")
      var docStr = dom.innerHTML
      newWindow.document.close()
      newWindow.print()
      newWindow.close()
    }
    return {
      printOut,
      printDom
    }
  }
}
</script>

<style scoped>

</style>