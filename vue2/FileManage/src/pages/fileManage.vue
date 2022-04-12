<template>
  <div class="FIleManage">
    <div style="background: #2f2f2f;height: 60px"></div>
    <div class="content">
      <div class="nav">
        <div class="header">
          <!-- <el-button icon="el-icon-bottom" class="btnicon"></el-button>
          <div class="rbtn">
            <el-button icon="el-icon-back" class="btnicon"></el-button>
            <el-button icon="el-icon-right" class="btnicon"></el-button>
            <el-button icon="el-icon-top" class="btnicon"></el-button>
          </div> -->
          <menuTree />
        </div>
      </div>
      <div class="right">
        <div class="header">
          <el-button type="primary" class="btns"><i class="el-icon-upload2" style="font-size: 12px;"></i>上传</el-button>
          <el-button @click="addFile" class="btns btnColor"><i class="el-icon-folder-add"></i>新增文件夹</el-button>
          <el-button-group style="margin-left:10px;">
            <el-button class="btnColor"> 打包文件</el-button>
            <el-button class="btns btnColor" v-show="selection.length !== 0"><i class="el-icon-download"></i> 下载</el-button>
            <el-button class="btns btnColor" v-show="selection.length !== 0"><i class="el-icon-delete"></i> 删除</el-button>
          </el-button-group>
          <div class="rbtn">
            <div class="serchInput">
              <input type="text" class="searchInput" v-model="searchText" placeholder="搜索您的文件">
              <i class="el-icon-search"></i>
            </div>
            <el-button :icon="isMenu ? 'el-icon-menu' : 'el-icon-s-operation'" :class="{'btnicon': true, 'activeList': isMenu}" @click="isMenu = !isMenu"></el-button>
          </div>
        </div>
        <!-- 文件展示区域 -->
        <mainContent :isMenu="isMenu" :fileData="fileData" :isCreadte="isCreadte" @cancelAdd="cancelAdd"
          @addFile="addFile" @selectData="selectData" />
      </div>
    </div>
  </div>
</template>

<script>
  import mainContent from '@/pages/main'
  import menuTree from '@/components/menuTree'
  export default {
    name: 'FIleManage',
    components: {
      mainContent,
      menuTree
    },
    props: {
      msg: String
    },
    data () {
      return {
        contextMenuVisible: true,
        searchText: '',
        isMenu: true, // 是否是菜单显示
        fileData: [
          {
            img: 'Code',
            date: '2016-05-18 18:53',
            fileSize: '423M',
            fileName: 'Code的压件锦集',
            checked: false,
            id: '1'
          },
          {
            img: 'ZIP',
            fileSize: '3G',
            date: '2016-05-18 18:53',
            fileName: '超大的压缩文件锦集',
            id: '2',
            checked: false
          }
        ],
        // 是否新建文件夹
        isCreadte: true,
        selection: []
      }
    },
    methods: {
      // 新增文件夹
      addFile () {
        let flag = false
        this.fileData.map(_ => { if (_.isCreadte) { flag = true } })
        // 防止再次新增文件
        if (flag) return this.onfocus()
        this.fileData.splice(0,0, {isCreadte: true,val: '新建文件夹'})
        this.onfocus()
      },
      // 新增文件后聚焦input框
      onfocus () {
        this.$nextTick(() => {
          let e = document.getElementsByClassName('addInput')
          e[0].select()
        })
      },
      // 取消新增文件夹
      cancelAdd () {
        this.fileData.map((_, index) => {
          if (_.isCreadte) {
            this.fileData.splice(index, 1)
          }
        })
      },
      // 获取选中数据
      selectData (val) {
        console.log('val', val)
        this.selection = val
      }
    }
  }
</script>

<style>
  .FIleManage .content{
    width: 100%;
    min-width: 1100px;
    height: 100vh;
    box-sizing: border-box;
    display: flex;
  }

  .FIleManage .content .nav {
    width: 20%;
    /* padding: 0 10px; */
    background: #E8EFF2;
  }

  .FIleManage .content .nav ul {
    margin-top: 6px;
  }

  .FIleManage .content .right {
    width: 80%;
    padding: 0 10px;
  }

  .FIleManage .content .right .header .btns {
    position: relative;
    padding-left: 36px;
  }

  .FIleManage .content .right .header .btns i {
    position: absolute;
    top: 6px;
    left: 10px;
    font-size: 18px!important;
    font-weight: bold;
  }
  .FIleManage .content .right .header .btnColor {
    color: #388CFF;
    border: 1px solid #C3EAFF;
  }

  .FIleManage .content .header {
    height: 50px;
    line-height: 50px;
  }

  .FIleManage .content .rbtn{
    float: right;
  }
  .FIleManage .btnicon {
    font-size: 20px;
    padding: 6px 8px;
    vertical-align: middle;
  }
  .serchInput {
    width: 247px;
    display: inline-block;
    line-height: 30px;
    margin-right: 30px;
    background: #F1F3F5;
    border-radius: 33px;
    padding: 0 38px 0 24px;
    position: relative;
  }
  .serchInput i {
    display: block;
    width: 16px;
    height: 16px;
    border: 0;
    position: absolute;
    top: 7px;
    font-size: 16px;
    outline: 0;
    cursor: pointer;
    right: 13px;
    color: #666;
  }
  .serchInput input {
    width: 100%;
    height: 30px;
    background: 0 0;
    border: 0;
    outline: 0;
    line-height: 29px!important;
    position: relative;
    font-size: 12px;
    color: #929292;
  }
  .activeList {
    background: #dae2ea;
  }
  thead .el-table-column--selection .cell{
    padding-left: 14px;
  }
</style>
