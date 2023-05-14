export class UserForDelete{
    jwt: String = "";
    id: String = "";
    
    public constructor(obj?: any) {
        if (obj) {
            this.jwt = obj.jwt;
            this.id = obj.id;
        }
    }
}