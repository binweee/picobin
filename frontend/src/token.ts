class TokenServer{
    private static TOKEN = 'token'

    static saveToken(token : string):void{
        localStorage.setItem(this.TOKEN,token)
    }

    static getToken():string | null{
        return localStorage.getItem(this.TOKEN)
    }

    static removeToken():void{
        localStorage.removeItem(this.TOKEN)
    }
}

export default TokenServer