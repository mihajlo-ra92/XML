export class Jwt {
    jwt: String = '';
    
    public constructor(obj?: any) {
      if (obj) {
        this.jwt = obj.jwt;
      }
    }
  }