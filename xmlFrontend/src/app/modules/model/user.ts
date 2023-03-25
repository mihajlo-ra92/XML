export class User {
  id: String = '';
  username: String = '';
  password: String = '';
  userType: String = '';
  firstName: String = '';
  lastName: String = '';
  gender: String = '';
  birthDate!: Date;
  email: String = '';
  public constructor(obj?: any) {
    if (obj) {
      this.id = obj.id;
      this.username = obj.username;
      this.userType = obj.userType;
      this.firstName = obj.firstName;
      this.lastName = obj.lastName;
      this.gender = obj.gender;
      this.birthDate = obj.birthDate;
      this.email = obj.email;
    }
  }
}
