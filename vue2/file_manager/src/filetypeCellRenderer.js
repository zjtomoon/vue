import Vue from 'vue'

export default Vue.extend({
  template: `
          <span style="margin-left: 10px;">

              <font-awesome-icon :icon="imgIcon" size="2x" style="color: #9A2530;" />

          </span>
          `,
  data () {
    return {
      imgIcon: null,
      value: ''
    }
  },
  beforeMount () {
    switch (this.params.value) {
      case 'pdf':
        this.imgIcon = 'file-pdf'
        break

      case 'jpg':

      // eslint-disable-next-line no-fallthrough
      case 'jpeg':

      // eslint-disable-next-line no-fallthrough
      case 'png':

      // eslint-disable-next-line no-fallthrough
      case 'gif':
        this.imgIcon = 'file-image'
        break

      case 'doc':

      // eslint-disable-next-line no-fallthrough
      case 'docx':
        this.imgIcon = 'file-word'
        break

      case 'ppt':

      // eslint-disable-next-line no-fallthrough
      case 'pptx':
        this.imgIcon = 'file-powerpoint'
        break

      case 'xls':

      // eslint-disable-next-line no-fallthrough
      case 'xlsx':
        this.imgIcon = 'file-excel'
        break

      case 'mp4':

      // eslint-disable-next-line no-fallthrough
      case 'mkv':

      // eslint-disable-next-line no-fallthrough
      case 'avi':

      // eslint-disable-next-line no-fallthrough
      case 'mov':

      // eslint-disable-next-line no-fallthrough
      case 'wmv':

      // eslint-disable-next-line no-fallthrough
      case 'mpeg':

      // eslint-disable-next-line no-fallthrough
      case 'mpg':

      // eslint-disable-next-line no-fallthrough
      case 'm4v':

      // eslint-disable-next-line no-fallthrough
      case 'flv':
        this.imgIcon = 'file-video'
        break

      case 'zip':

      // eslint-disable-next-line no-fallthrough
      case 'rar':
        this.imgIcon = 'file-archive'
        break

      default:
        this.imgIcon = 'file-alt'
        break
    }
    this.value = this.params.value
  }
})
