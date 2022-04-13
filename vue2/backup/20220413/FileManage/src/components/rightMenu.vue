<template>
  <div class="right-menu" v-show="visible">
    <li @click.stop="view">查看</li>
    <li >刷新</li>
    <li @click.stop="addFile">新建文件夹</li>
  </div>
</template>
<script>
  export default {
    props: {
      visible: {
        type: Boolean,
        default: function () {
          return false
        }
      }
    },
    data () {
      return {
        binded: false
      }
    },
    mounted () {
      // this.bindEvents()
    },
    methods: {
      view () {
        this.closeMenus()
      },
      addFile () {
        this.$emit('addFile')
        this.closeMenus()
      },
      // 关闭右键菜单
      closeMenus () {
        this.$emit("update:visible", false)
      }
    }
  }
</script>
<style scoped>
  .right-menu {
    border: 1px solid #eee;
    box-shadow: 0 0.5em 1em 0 rgba(0,0,0,.1);
    border-radius: 1px;
    position: fixed;
    background: #fff;
    border: 1px solid rgba(0,0,0,.2);
    border-radius: 4px;
    z-index: 999;
    padding: 4px 0;
  }
  .right-menu li {
    display: block;
    padding: 0 15px;
    line-height: 28px;
    line-height: 28px;
    color: #333;
    cursor: pointer;
    user-select: none;
    font-size: 14px;
  }
  .right-menu li:hover {
    background: #42B983;
    color: white;
  }
</style>