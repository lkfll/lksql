"use strict";



// 基础类型（int string）数组的渲染
// 使用json生成ul元素,RespJson必须是，基础类型的数组 
// []int || []string
Object.prototype.ToUl = function (element) {
    this.RespPromise
        .then(f => {
            let ul = document.createElement('ul');
            for (let i = 0; i < this.RespJson.length; i++) {
                let li = document.createElement('li');
                li.innerHTML = this.RespJson[i]
                ul.appendChild(li);
            }
            element.appendChild(ul)
        })
}

// 多个对象的渲染
// 使用json生成ul元素,RespJson必须是，struct类型的数组 
// []struct
Object.prototype.ToTable = function (element) {
    this.RespPromise
        .then(f => {
            let table = document.createElement('table');
            let keys = Object.keys(this.RespJson[0]) // 获取单个对象的keys
            let tr = document.createElement('tr');
            for (let i = 0; i < keys.length; i++) {
                let th = document.createElement("th")
                th.innerHTML = keys[i]
                tr.appendChild(th)
            }
            table.appendChild(tr)
            for (let i = 0; i < this.RespJson.length; i++) {
                let tr = document.createElement('tr');
                for (let j = 0; j < keys.length; j++) {
                    let td = document.createElement("td")
                    td.innerHTML = this.RespJson[i][keys[j]]
                    tr.appendChild(td)
                }
                table.appendChild(tr);
            }
            element.appendChild(table)
        })
} 