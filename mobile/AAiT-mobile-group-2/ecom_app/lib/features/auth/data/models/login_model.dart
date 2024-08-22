import '../../domain/entities/login_entity.dart';

class LoginModel extends LoginEntity {
  LoginModel({required super.email, required super.password});

  Map<String, dynamic> toJson() {
    return {'email': email, 'password': password};
  }

  static LoginModel toModel(LoginEntity login_entity){
    return LoginModel(email: login_entity.email, password: login_entity.password);
  }
}
