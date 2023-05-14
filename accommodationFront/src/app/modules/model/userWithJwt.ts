import{ User } from './user'

export class UserWithJwt{
    jwt: string = "";
    user: User = new User();
    public constructor(obj?: any) {
        if (obj) {
            this.jwt = obj.jwt;
            this.user = obj.user;
        }
    }
}