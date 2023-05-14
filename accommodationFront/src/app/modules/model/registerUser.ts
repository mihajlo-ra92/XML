export class RegisterUser {
  userType: String = '';
    username: String = '';
    password: String = '';
    email: String = '';
    FirstName: String = '';
    LastName: String = '';
    Address:String = '';
    public constructor(obj?: any) {
      if (obj) {
        this.userType = obj.userType;
        this.username = obj.username;
        this.password = obj.password;
        this.email = obj.email;
        this.FirstName = obj.FirstName;
        this.LastName = obj.LastName;
        this.Address = obj.Address;
      }
    }
  }
  