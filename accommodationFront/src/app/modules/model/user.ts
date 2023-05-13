export class User {
  id: String = '';
  username: String = '';
  password: String = '';
  userType: String = '';
  firstName: String = '';
  lastName: String = '';
  email: String = '';
  public constructor(obj?: any) {
    if (obj) {
      this.id = obj.id;
      this.username = obj.username;
      this.userType = obj.userType;
      this.firstName = obj.FirstName;
      this.lastName = obj.LastName;
      this.email = obj.email;
    }
  }
}
