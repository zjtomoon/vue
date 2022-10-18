<template>
    <div>
        <button @click="handleClick">点击增加</button>
    </div>
    <div>
        <button @click="handleClickPatch">
            修改数据($patch)
        </button>
    </div>

    <div>
        <button @click="handleClickMethod">
            修改数据($patch+函数)
        </button>
    </div>

    <div>
        <button @click="handleClickActions">
            修改数据(actions)
        </button>
    </div>
</template>

<script lang="ts" setup>
// 必须写'lang="ts" setup ',否则会报错
    import { mainStore } from '../store/index'
    const store = mainStore()

    const handleClick = () => {
        store.count++
        store.helloworld = store.helloworld === "Hello World"?"hello":"Hello World";
    }

    // 传递对象
    const handleClickPatch = () => {
        store.$patch({
            count:store.count + 2,
            helloworld: store.helloworld === "Hello World"?"hello":"Hello World"
        })
    }

    // 传递函数
    const handleClickMethod = () => {
        store.$patch((state) => {
            state.count++
            state.helloworld = state.helloworld === "Hello World"?"hello":"Hello World"
        })
    }

    const handleClickActions = () => {
        store.changeState()
    }
</script>

<style lang="scss" scoped>

</style>