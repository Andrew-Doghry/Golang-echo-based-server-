// let pro = {
//     type : "food and pavarage"
// }
// let data = JSON.stringify(pro)
// fetch('http://localhost:2000/product',{method:"POST",body:data}).then(data=>data.json()).then(console.log)


const btn = document.querySelector('.btn')
btn.onclick = (e)=>{
    let pro = [
    {
        TYPE : "food and pavarage"
    },
    {
        TYPE : "tic tac"
    },
    {
        TYPE : "dental instruments"
    }
]
    let data = 
    fetch('http://localhost:2000/product',{method:"POST",body:JSON.stringify(pro)}).then(data=>data.json()).then(console.log)

}