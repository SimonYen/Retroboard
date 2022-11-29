function setClock(){
    c=document.getElementById("banner")
    d=new Date()
    str=d.getFullYear()+'/'+(d.getMonth()+1)+'/'+d.getDay()+' '+d.getHours()+':'+d.getMinutes()
    c.innerHTML=str
}
function resetBanner(){
    b=document.getElementById("banner")
    b.innerHTML="RETRO DASHBOARD"
}