"use strict";

// 配置公共url
Object.prototype.Url = ""
// 占位置，标志要使用
Object.prototype.RespJson = null
// 占位置，标志要使用 Promise<Response>
Object.prototype.RespPromise = null

// 响应处理函数, 必须是箭头函数，不然this会指向windows
Object.prototype.SetRespJson = json => {
    if (json.Code === 0) {
        this.RespJson = json.Data
        Printf("Code:", json.Code, "Data:", json.Data)
    } else {
        //this.RespJson = json.Data
        Printf("Code:", json.Code, "Data:", json.Data)
    }
}

// 发送当前对象的json数据
Object.prototype.Fetch = function (urlStr) {
    this.RespPromise = fetch(this.Url + urlStr, {
        method: "POST",
        body: JSON.stringify(this)
    })
        .then(response => response.json())
        .then(eval(this.SetRespJson.toString()))// 先转为字符串在转为函数，曲线救国确保this指向
    return this
}

// 控制台打印
function Printf(...obj) {
    console.log(...obj)
}

// 控制台打印响应json
Object.prototype.PrintRespJson = function () {
    this.RespPromise
        .then(f => Printf(this.RespJson))
    return this
}

