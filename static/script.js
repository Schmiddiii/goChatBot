goWebGuiObj.responseHandler=(msg)=>{
    showMsg(msg.extras[0], false)
}

goWebGuiObj.getExtras=(className)=>{
    let arr=[]
    let value=document.getElementById("msgBox").value
    arr.push(value)
    showMsg(value, true)
    return arr
}


goWebGuiObj.init()


function showMsg(msg, fromUser){
    let node = document.createElement("p")
    node.innerHTML=msg
    if(fromUser){
        node.classList.add("msgUser")
    }else{
        node.classList.add("msgBot")
    }

    
    document.getElementById("msgList").appendChild(node)
    document.getElementById("msgList").lastChild.scrollIntoView()
}

function submit(){
    if(document.getElementById("msgBox").value!=""){    
        document.getElementById("msgBox").value=""
    }
}