export class LoginUser{
    username: String = "";
    password: String = "";
    
    public constructor(obj?: any) {
        if (obj) {
            this.username = obj.username;
            this.password = obj.password;
        }
    }
}