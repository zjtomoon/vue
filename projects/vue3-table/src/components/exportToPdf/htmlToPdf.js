import html2canvas from "html2canvas";
import JsPDF from 'jspdf'

export default {
    install(Vue,options) {
        Vue.prototype.htmlToPdf = (id,title) => {
            const element = document.getElementById(`${id}`)
            const opts = {
                scale:12,
                useCORS:true,
                allowTaint:false,
                taintTest:true,
                logging:true
            }
            html2canvas(element,opts).then((canvas)=> {
                console.log(canvas)
                const contentWidth = canvas.width
                const contentHeight = canvas.height
                const pageHeight = (contentWidth / 592.28) * 841.89
                let leftHeight = contentHeight
                let position = 0
                const imgWidth = 595.28
                const imgHeight = (592.28 / contentWidth) * contentHeight
                const pageData = canvas.toDataURL('./image/jpeg',1.0)
                console.log(pageData)
                const PDF = new JsPDF('','pt','a4')
                if(leftHeight < pageHeight) {
                    PDF.addImage(pageData,'JPEG',0,position,imgWidth,imgHeight)
                } else {
                    while (leftHeight > 0) {
                        PDF.addImage(pageData,'JPEG',0,position,imgWidth,imgHeight)
                        leftHeight -= 841.89
                        if (leftHeight > 0) {
                            PDF.addPage()
                        }
                    }
                }
                PDF.save(title + '.pdf')
            }).catch((error)=>{
                console.log('打印失败',error)
            })
        }
    }
}