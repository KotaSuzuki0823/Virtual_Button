new Vue({
    el:'#app',
    data:{

    },
    // メソッド定義
    methods: {
        // サーバを叩く(指定したルーティンを実行させる)
        runAction() {
            axios.get('/action')
            .then(response => {
                if (response.status != 200) {
                    throw new Error('レスポンスエラー')
                } else {
                    throw new Error('成功')
                }
            })
        },
    }
})

var buttonhandler = new Vue({
    el : '#buttonhandler',
    data: {
        counter: 0
    }
})