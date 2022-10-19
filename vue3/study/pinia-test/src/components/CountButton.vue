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

    <div>
        <button @click="handleClickChangePhone">
            改变电话号码
        </button>
    </div>

    <div>
        <button @click="getList">
            显示List
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

    // 调用该方法后，phoneHidden方法也随着调用了一次
    const handleClickChangePhone = () => {
        store.phone = "15139333999"
    }

    const getList = () => {
        store.getList();
    }
</script>

<style lang="scss" scoped>

</style>