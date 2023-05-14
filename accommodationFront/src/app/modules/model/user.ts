export class User {
  id: String = '';
  username: String = '';
  password: String = '';
  userType: String = '';
  FirstName: String = '';
  LastName: String = '';
  email: String = '';
  Address:String = '';
  public constructor(obj?: any) {
    if (obj) {
      this.id = obj.id;
      this.username = obj.username;
      this.userType = obj.userType;
      this.FirstName = obj.FirstName;
      this.LastName = obj.LastName;
      this.email = obj.email;
      this.Address = obj.Address;
    }
  }
}
