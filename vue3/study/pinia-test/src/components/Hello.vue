<template>
  <h2 class="">
    {{store.helloworld}}
  </h2>
  <h2 class="">
    {{store.count}}
  </h2>

  <hr/>
<!-- 下面的模板不是相应式数据 -->
  <h2 class="">{{helloworld}}</h2>  
  <h2>{{count}}</h2>  
</template>

<script lang="ts" setup>
import {mainStore} from '../store/index'
import {storeToRefs} from 'pinia'
const store = mainStore()
// 不具备响应式能力
// const {helloworld,count} = store

// 有了storeToRefs( )方法后，就可以在解构的代码中，对store使用方法了。
// 其实这时候就是把解构出来的数据作了ref响应式代理。所以数据拥有了响应式能力。
const {helloworld,count} = storeToRefs(store)
</script>

<style lang="scss" scoped></style>