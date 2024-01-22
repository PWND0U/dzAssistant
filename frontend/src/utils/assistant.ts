import {GetBiliFollower} from "../../wailsjs/go/assistant/Assistant";
import {ref} from "vue"

let respStr = ref("")

function processInstructions(instruction:string){
    respStr.value = instruction
    if (instruction.indexOf("豆子同学") != -1) {
        // 分割指令
        let command = instruction.split("豆子同学");
        if (command[1].length >0){
            if (command[1].includes("粉丝数") || command[1].includes("粉丝量")) {
                GetBiliFollower().then(numFollower => {
                    respStr.value =  `当前B站粉丝量:${numFollower}`
                })
            }else {
                respStr.value = "无法理解您的指令！"  //也可以将指令交给 chatgpt 来处理，这个自行研究
            }
        }else{
            respStr.value = "你想问什么！"
        }
    }
}

export {respStr,processInstructions}

