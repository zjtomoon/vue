<template>
  <div class="contents">
    <div>
      <el-dialog
          title="档案保送单"
          :close-on-click-modal="false"
          :visible.sync="dialogFormVisible"
          width="60%"
      >
        <div id="printTest1" v-if="listinfo.length>0">
          <table
              v-for="(item,index) in listinfo"
              :key="index"
              border="0"
              width="70%"
              style="margin: auto!important;">
            <thead>
            <tr>
              <th colspan="8" id="title">新任中管干部档案保送单</th>
            </tr>
            </thead>
            <tbody>
            <tr style="height: 40px">
              <td>姓名</td>
              <td colspan="3" v-if="name!=''">{{ item.name }}</td>
              <td colspan="3">档案数</td>
              <td colspan="1">1</td>
            </tr>
            <tr style="height: 40px">
              <td>工作单位及职务</td>
              <td colspan="7">{{ item.jobTitle }}</td>
            </tr>
            <tr style="height: 40px">
              <td>任职时间</td>
              <td colspan="7">{{ item.joinWorkTime | dateFmt('YYYY-MM-DD') }}</td>
            </tr>
            <tr style="height: 60px">
              <td>档案整理人</td>
              <td colspan="3" v-if="name!=''">{{ item.name }}</td>
              <td colspan="3">档案审核人</td>
              <td colspan="1" v-if="name!=''">{{ item.name }}</td>
            </tr>
            <tr style="height: 180px">
              <td>
                <div style="width: 70px;margin: auto!important;">报送单位预览</div>
              </td>
              <td colspan="7">
                <div class="right">
                  <div class="top">
                    <span class="leader" v-if="name!=''">领导签字:{{ item.name }}</span>
                    <span class="official_seal">(公章)</span>
                  </div>
                  <div class="bottom">
                    <span class="year">{{ item.workingTime | dateFmt('YYYY') }}年</span>
                    <span class="month">{{ item.workingTime | dateFmt('MM') }}月</span>
                    <span class="day">{{ item.workingTime | dateFmt('DD') }}日</span>
                  </div>
                </div>
              </td>
            </tr>
            </tbody>
          </table>
        </div>
      </el-dialog>
    </div>
  </div>
</template>

<script>
export default {
  name: "SubmissionForm",
  props: {
    dialogFormVisible: {
      type: Boolean,
      default: false
    },
  },
  watch: {
    dialogFromVisible() {
      this.chushi();
      if (this.dialogFormVisible == false) {
        this.$emit('dialog')
      }
    }
  },
  data() {
    return {
      listinfo: [],
      sids: [],
    }
  },
  created() {
    this.chushi()
  },
  methods: {
    init(sids) {
      this.sids = sids
    },
    chishi() {
      request({
        url: '/submit/getPrintList',
        method: 'get',
        params: {
          ids: this.sids + ''
        }
      }).then(res => {
        this.listinfo = res.data.data
      })
    },
    closeDialog() {
      this.listinfo = [];
    },
    //打印
    confirm() {

    },
    cancel() {
      this.$emit('dialog')
    }
  }
}
</script>

<style scoped lang="less">
//表格样式
#printTest1 .right .top .leader {
  float: left;
  margin-left: 30px !important;
}

#printTest1 .right .top {
  margin-top: 100px !important;
}

#printTest1 .right .top .official_seal {
  margin-right: 30px !important;
}

#printTest1 .bottom {
  margin-top: 10px !important;
}

#printTest1 .bottom .year, #printTest1 .bottom .month {
  margin-right: 40px !important;
}

#printTest1 #title {
  padding-top: 50px !important;
  padding-bottom: 30px !important;
}

#printTest1 {
  width: 70%;
  margin: auto;
  border: 2px solid black;
  padding-bottom: 30px !important;
}

#printTest1 table {
  border-collapse: collapse;
}

#printTest1 table thead th {
  font-size: 20px;
  padding: 10px;
}

#printTest1 table tbody tr {
  height: 30px;
  font-size: 14px;
}

#printTest1 table tbody td {
  width: 25%;
  border: 1px solid black;
  text-align: center !important;
}

#printTest1 table tbody td span {
  margin-right: 20px;
}

//截至
//下面的样式是 弹框组件样式 如:标题等等
.contents /deep/ .el-dialog__title {
  font-size: 15px !important;
}

.contents /deep/ .el-form {
  /*width: 410px!important;*/
  margin: auto !important;
}

.contents /deep/ .el-dialog__body {
  padding: 20px 10px !important;
}

.contents .ml {
  margin-bottom: 15px !important;
  border-left: 3px solid #447FC1;
  padding-left: 10px !important;
  font-size: 18px !important;
  font-weight: 500;
  font-size: 15px !important;
  margin-left: 15px !important;
}

.style /deep/ .el-input__inner {
  margin-bottom: 15px !important;
}
</style>